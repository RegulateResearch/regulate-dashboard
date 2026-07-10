package grouping

import (
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func generalLoginRoute(router gin.IRouter, middlewares setup.Middlewares) Route {
	return newRoute(router, "/common", middlewares.Auth.Authenticate)
}
