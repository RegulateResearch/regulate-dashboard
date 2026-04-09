package setup

import "frascati/handler"

type Handlers struct {
	Auth    handler.AuthHandler
	Course  handler.CourseHandler
	Session handler.SessionHandler
	Admin   handler.AdminHandler
	User    handler.UserHandler
	Sso     handler.SsoHandler
	Try     handler.TryHandler
}

func setupHandlers(services services) Handlers {
	return Handlers{
		Auth:    handler.NewAuthHandler(services.auth),
		Course:  handler.NewCourseHandler(services.course),
		Session: handler.NewSessionHandler(),
		Admin:   handler.NewAdminHandler(services.user),
		User:    handler.NewUserHandler(services.user),
		Sso:     handler.NewSsoHandler(services.sso),
		Try:     handler.NewTryHandler(services.try),
	}
}
