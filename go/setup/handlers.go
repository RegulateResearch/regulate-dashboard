package setup

import "frascati/handler"

type Handlers struct {
	Auth    handler.AuthHandler
	Session handler.SessionHandler
	Admin   handler.AdminHandler
	User    handler.UserHandler
	Try     handler.TryHandler
}

func setupHandlers(services services) Handlers {
	return Handlers{
		Auth:    handler.NewAuthHandler(services.auth),
		Session: handler.NewSessionHandler(),
		Admin:   handler.NewAdminHandler(services.user),
		User:    handler.NewUserHandler(services.user),
		Try:     handler.NewTryHandler(),
	}
}
