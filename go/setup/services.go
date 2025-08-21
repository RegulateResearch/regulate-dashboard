package setup

import (
	"frascati/service"
	"frascati/service/auth"
)

type services struct {
	auth service.AuthService
	user service.UserService
}

func setupServices(repos repositories, jwt auth.JwtService, bcrypt auth.BcryptService) services {
	return services{
		auth: service.NewAuthService(repos.auth, bcrypt, jwt),
		user: service.NewUserService(repos.user),
	}
}
