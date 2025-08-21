package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

// func SetupSessionRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware) {
// 	group := router.Group("")
// 	group.Use(authMiddleware.Authenticate)

// 	group.GET("/session", handler.CheckSession)
// }

func setupSessionRouting(routers grouping.Routes, handlers setup.Handlers) {
	sessionHandler := handlers.Session
	routers.Login.GET("/session", sessionHandler.CheckSession)
}
