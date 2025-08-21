package setup

import (
	"database/sql"
	"frascati/prep/logger"
)

func SetupApplication(db *sql.DB) (logger.EnhancedLogger, Handlers, Middlewares) {
	logger := setupLogger()

	jwtService, bcryptService := setupAuthUtils()
	repos := setupRepositories(db)
	services := setupServices(repos, jwtService, bcryptService)

	handlers := setupHandlers(services)
	middlewares := setupMiddlewares(jwtService, logger)

	return logger, handlers, middlewares
}
