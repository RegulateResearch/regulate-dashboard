package setup

import (
	"database/sql"
	"frascati/comp/logger"
	"frascati/config"
	"frascati/exception"
	"frascati/prep"
	"os"
)

type App struct {
	db          *sql.DB
	warnFile    *os.File
	errFile     *os.File
	Logger      logger.EnhancedLogger
	Handlers    Handlers
	Middlewares Middlewares
}

func SetupApp() (App, exception.Exception) {
	config.InitEnv()
	db, exc := setupDb()

	if exc != nil {
		return App{}, exc
	}

	warnFile, errFile := prep.PrepFile()
	logger := setupEnhanceLogger(warnFile, errFile)
	jwtService, bcryptService := setupAuthUtils()

	repos := setupRepositories(db)
	services := setupServices(repos, jwtService, bcryptService)
	middlewares := setupMiddlewares(jwtService, logger)
	handlers := setupHandlers(services)

	app := App{
		db:          db,
		warnFile:    warnFile,
		errFile:     errFile,
		Logger:      logger,
		Middlewares: middlewares,
		Handlers:    handlers,
	}

	return app, nil
}

func (a App) Close() exception.Exception {
	errs := make([]error, 0)

	dbCloseErr := a.db.Close()
	if dbCloseErr != nil {
		errs = append(errs, dbCloseErr)
	}

	warnCloseErr := a.warnFile.Close()
	if warnCloseErr != nil {
		errs = append(errs, warnCloseErr)
	}

	errCloseErr := a.errFile.Close()
	if errCloseErr != nil {
		errs = append(errs, errCloseErr)
	}

	if len(errs) > 0 {
		return exception.NewMultipleException(exception.CAUSE_INTERNAL, "app shutdown", "app's components fail to shut down", errs...)
	}

	return nil
}
