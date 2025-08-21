package grouping

import (
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	NoLogin Route
	Login   Route
	User    Route
	Admin   Route
}

func AllRoutes(router gin.IRouter, middlewares setup.Middlewares) Routes {
	return Routes{
		NoLogin: nologinRoute(router),
		Login:   loginRoute(router, middlewares),
		User:    userRoute(router, middlewares),
		Admin:   adminRoute(router, middlewares),
	}
}
