package routing

import (
	"frascati/handler"
	"frascati/service/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouterHandler(router *gin.Engine, authService auth.AuthService) {
	authHandler := handler.NewAuthHandler(authService)
	group := router.Group("/auth")
	group.POST("/register", authHandler.Register)
	group.POST("/login", authHandler.Login)
}
