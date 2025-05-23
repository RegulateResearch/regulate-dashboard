package routing

import (
	"frascati/handler"
	"frascati/middleware"

	"github.com/gin-gonic/gin"
)

func SetupSessionRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware) {
	group := router.Group("")
	group.Use(authMiddleware.Authenticate)

	group.GET("/session", handler.CheckSession)
}
