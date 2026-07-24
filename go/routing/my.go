package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupMyRouter(routers grouping.Routes, handlers setup.Handlers) {
	myHandler := handlers.My
	userGroup := routers.User.Group("/my")
	generalGroup := routers.General.Group("/my")

	generalGroup.GET("/profile", myHandler.MyProfile)
	userGroup.GET("/courses", myHandler.MyCourses)
}
