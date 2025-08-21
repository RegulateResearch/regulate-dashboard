package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupAuthRouting(routers grouping.Routes, handlers setup.Handlers) {
	authHandler := handlers.Auth
	authGroup := routers.NoLogin.Group("/auth")

	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
}
