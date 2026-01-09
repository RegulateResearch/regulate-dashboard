package setup

import (
	"frascati/comp/auth"
	"frascati/comp/graceful"
	"frascati/comp/logger"
	"frascati/middleware"
)

type Middlewares struct {
	Auth       middleware.AuthMiddleware
	Logger     middleware.LoggerMiddleware
	Gatekeeper middleware.GatekeeperMiddleware
}

func setupMiddlewares(jwt auth.JwtService, logger logger.EnhancedLogger, gatekeeper graceful.Gatekeeper) Middlewares {
	return Middlewares{
		Auth:       middleware.NewAuthMiddleware(jwt),
		Logger:     middleware.NewLoggerMiddleware(logger),
		Gatekeeper: middleware.NewGatekeeperMiddleware(gatekeeper, logger),
	}
}
