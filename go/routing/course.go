package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupCourseRouter(routers grouping.Routes, handlers setup.Handlers) {
	courseHandler := handlers.Course
	adminCourseGroup := routers.Admin.Group("/courses")

	adminCourseGroup.GET("", courseHandler.AllCourse)
	adminCourseGroup.POST("", courseHandler.NewCourse)
}
