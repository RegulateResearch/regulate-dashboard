package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupCourseRouter(routers grouping.Routes, handlers setup.Handlers) {
	courseHandler := handlers.Course
	courseMemberHandler := handlers.CourseMember
	adminCourseGroup := routers.Admin.Group("/courses")

	adminCourseGroup.GET("", courseHandler.AllCourse)
	adminCourseGroup.POST("", courseHandler.NewCourse)

	adminCourseGroup.GET("/:id", courseHandler.CourseById)
	adminCourseGroup.PUT("/:id", courseHandler.UpdateById)
	adminCourseGroup.DELETE("/:id", courseHandler.DeleteById)

	adminCourseGroup.GET("/:id/members", courseMemberHandler.FindByCourse)
	adminCourseGroup.POST("/:id/members", courseMemberHandler.AddNewMember)
	adminCourseGroup.PUT("/:id/members", courseMemberHandler.UpdateMember)
	adminCourseGroup.DELETE("/:id/members", courseMemberHandler.DeleteMember)
}
