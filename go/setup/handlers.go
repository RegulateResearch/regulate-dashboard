package setup

import "frascati/handler"

type Handlers struct {
	Auth         handler.AuthHandler
	Course       handler.CourseHandler
	CourseMember handler.CourseMemberHandler
	Session      handler.SessionHandler
	Admin        handler.AdminHandler
	User         handler.UserHandler
	Sso          handler.SsoHandler
	My           handler.MyHandler
	Record       handler.RecordHandler
	Try          handler.TryHandler
}

func setupHandlers(services services) Handlers {
	return Handlers{
		Auth:         handler.NewAuthHandler(services.auth),
		Course:       handler.NewCourseHandler(services.course),
		CourseMember: handler.NewCourseMemberHandler(services.courseMember),
		Session:      handler.NewSessionHandler(),
		Admin:        handler.NewAdminHandler(services.user),
		User:         handler.NewUserHandler(services.user),
		Sso:          handler.NewSsoHandler(services.sso),
		My:           handler.NewMyHandler(services.my),
		Record:       handler.NewRecordHandler(services.record),
		Try:          handler.NewTryHandler(services.try),
	}
}
