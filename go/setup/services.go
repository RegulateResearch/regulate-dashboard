package setup

import (
	"frascati/comp/auth"
	"frascati/comp/background"
	"frascati/service"
)

type services struct {
	auth service.AuthService
	user service.UserService
	try  service.TryService
}

func setupServices(repos repositories, jwt auth.JwtService, bcrypt auth.BcryptService, backgroundProcessor background.Processor) services {
	return services{
		auth: service.NewAuthService(repos.auth, bcrypt, jwt),
		user: service.NewUserService(repos.user),
		try:  service.NewTryService(backgroundProcessor),
	}
}
