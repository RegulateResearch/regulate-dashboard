package setup

import (
	"frascati/comp/logger"
	"frascati/middleware"
	"frascati/util/auth"
)

type Middlewares struct {
	Auth   middleware.AuthMiddleware
	Logger middleware.LoggerMiddleware
}

func setupMiddlewares(jwt auth.JwtService, logger logger.EnhancedLogger) Middlewares {
	return Middlewares{
		Auth:   middleware.NewAuthMiddleware(jwt),
		Logger: middleware.NewLoggerMiddleware(logger),
	}
}
