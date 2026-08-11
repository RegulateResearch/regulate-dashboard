package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupCourseRouter(routers grouping.Routes, handlers setup.Handlers) {
	courseHandler := handlers.Course

	adminCourseGroup := routers.Admin.Group("/courses")
	userCourseGroup := routers.User.Group("/courses")

	userCourseGroup.GET("", courseHandler.AllCourse)

	adminCourseGroup.GET("", courseHandler.AllCourse)
	adminCourseGroup.POST("", courseHandler.NewCourse)

	adminCourseGroup.GET("/:course_id", courseHandler.CourseById)
	adminCourseGroup.PUT("/:course_id", courseHandler.UpdateById)
	adminCourseGroup.DELETE("/:course_id", courseHandler.DeleteById)

	setupCourseMemberRouter(routers, handlers)
	setupCourseItemRouter(routers, handlers)
}

func setupCourseMemberRouter(routers grouping.Routes, handlers setup.Handlers) {
	courseMemberHandler := handlers.CourseMember
	adminCourseGroup := routers.Admin.Group("/courses")

	adminCourseGroup.GET("/:course_id/members", courseMemberHandler.FindByCourse)
	adminCourseGroup.POST("/:course_id/members", courseMemberHandler.AddNewMember)
	adminCourseGroup.PUT("/:course_id/members", courseMemberHandler.UpdateMember)
	adminCourseGroup.DELETE("/:course_id/members", courseMemberHandler.DeleteMember)
}

func setupCourseItemRouter(routers grouping.Routes, handlers setup.Handlers) {
	courseItemHandler := handlers.CourseItem
	adminCourseGroup := routers.Admin.Group("/courses")

	adminCourseGroup.GET("/:course_id/items", courseItemHandler.GetByCourse)
	adminCourseGroup.POST("/:course_id/items", courseItemHandler.AddBulk)
	adminCourseGroup.PUT("/:course_id/items/:item_id", courseItemHandler.UpdateSingular)
	adminCourseGroup.DELETE("/:course_id/items/:item_id", courseItemHandler.DeleteSingular)
}
