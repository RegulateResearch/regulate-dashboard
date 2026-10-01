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

	userGroup.GET("/tasks", myHandler.MyTasks)
	userGroup.GET("/tasks/:task_id", myHandler.MyTaskById)
	userGroup.PUT("/tasks/:task_id", myHandler.UpdateTask)
}
