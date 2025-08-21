package grouping

import (
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func loginRoute(router gin.IRouter, middlewares setup.Middlewares) Route {
	return newRoute(router, "", middlewares.Auth.Authenticate)
}
