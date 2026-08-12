package setup

import (
	"frascati/comp/auth"
	"frascati/comp/graceful"
	"frascati/comp/logging"
	"frascati/middleware"
)

type Middlewares struct {
	Auth         middleware.AuthMiddleware
	Logger       middleware.LoggerMiddleware
	Gatekeeper   middleware.GatekeeperMiddleware
	CourseAccess middleware.CourseAccessMiddleware
}

func setupMiddlewares(jwt auth.JwtService, logger logging.ExceptionSupportLogger, gatekeeper graceful.Gatekeeper, serv services) Middlewares {
	return Middlewares{
		Auth:         middleware.NewAuthMiddleware(jwt),
		Logger:       middleware.NewLoggerMiddleware(logger),
		Gatekeeper:   middleware.NewGatekeeperMiddleware(gatekeeper, logger),
		CourseAccess: middleware.NewCourseAccessMiddleware(serv.courseMember),
	}
}
