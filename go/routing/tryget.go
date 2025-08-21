package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupTryGetUser(routers grouping.Routes, handlers setup.Handlers) {
	routers.Admin.GET("/try/users", handlers.Admin.GetAll)
	routers.User.GET("/try/users", handlers.User.GetAll)
}
