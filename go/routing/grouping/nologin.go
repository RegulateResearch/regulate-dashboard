package grouping

import "github.com/gin-gonic/gin"

func nologinRoute(router gin.IRouter) Route {
	return newRoute(router, "")
}
