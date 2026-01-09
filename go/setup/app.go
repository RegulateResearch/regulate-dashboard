package setup

import (
	"database/sql"
	"frascati/comp/graceful"
	"frascati/comp/logger"
	"frascati/config"
	"frascati/exception"
	"frascati/lambda"
	"frascati/prep"
	"os"
)

type App interface {
	Open()
	Close(appCloseSig chan struct{}, serverCloseSig chan struct{}, gateClosedSig chan struct{}) exception.Exception
	CloseComp() exception.Exception
	Handlers() Handlers
	Middlewares() Middlewares
	Logger() logger.EnhancedLogger
}

type app struct {
	db          *sql.DB
	warnFile    *os.File
	errFile     *os.File
	logger      logger.EnhancedLogger
	gatekeeper  graceful.Gatekeeper
	handlers    Handlers
	middlewares Middlewares
	isClosed    bool
}

func SetupApp() (App, exception.Exception) {
	config.InitEnv()
	db, exc := setupDb()

	if exc != nil {
		return nil, exc
	}

	warnFile, errFile := prep.PrepFile()
	logger := setupEnhanceLogger(warnFile, errFile)
	gatekeeper := graceful.NewGateKeeper()
	jwtService, bcryptService := setupAuthUtils()

	repos := setupRepositories(db)
	services := setupServices(repos, jwtService, bcryptService)
	middlewares := setupMiddlewares(jwtService, logger, gatekeeper)
	handlers := setupHandlers(services)

	app := &app{
		db:          db,
		warnFile:    warnFile,
		errFile:     errFile,
		logger:      logger,
		gatekeeper:  gatekeeper,
		middlewares: middlewares,
		handlers:    handlers,
		isClosed:    false,
	}

	return app, nil
}

func (a *app) Handlers() Handlers {
	return a.handlers
}

func (a *app) Middlewares() Middlewares {
	return a.middlewares
}

func (a *app) Logger() logger.EnhancedLogger {
	return a.logger
}

func (a *app) Open() {
	a.gatekeeper.Open()
}

func (a *app) Close(appCloseSig chan struct{}, serverCloseSig chan struct{}, gateClosedSig chan struct{}) exception.Exception {
	defer func() { appCloseSig <- struct{}{} }()
	a.gatekeeper.Close()
	a.gatekeeper.Wait()

	gateClosedSig <- struct{}{}

	<-serverCloseSig

	return a.CloseComp()
}

func (a *app) CloseComp() exception.Exception {
	if a.isClosed {
		return nil
	}

	type closable interface {
		Close() error
	}

	closables := []closable{a.db, a.warnFile, a.errFile}

	errs := lambda.FilterList(lambda.MapList(closables, func(c closable) error {
		return c.Close()
	}), func(err error) bool {
		return err != nil
	})

	if len(errs) > 0 {
		return exception.NewMultipleException(exception.CAUSE_INTERNAL, "app shutdown", "app's components fail to shut down", errs...)
	}

	a.isClosed = true

	return nil
}
