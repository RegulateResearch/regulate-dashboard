package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupMyRouter(routers grouping.Routes, handlers setup.Handlers) {
	myHandler := handlers.My
	route := routers.User
	group := route.Group("/my")

	group.GET("/profile", myHandler.MyProfile)
}
