package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupAuthRouting(routers grouping.Routes, handlers setup.Handlers) {
	authHandler := handlers.Auth
	sessionHandler := handlers.Session
	authGroup := routers.NoLogin.Group("/auth")
	loginGroup := routers.Login.Group("/auth")

	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	loginGroup.GET("/session", sessionHandler.CheckSession)
}
