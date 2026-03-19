package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupUsersRouting(routers grouping.Routes, handlers setup.Handlers) {
	adminGroup := routers.Admin
	userHandler := handlers.User

	adminGroup.GET("/users", userHandler.GetAll)
	adminGroup.GET("/users/:id", userHandler.GetById)
}
