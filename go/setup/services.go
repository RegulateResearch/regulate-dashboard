package setup

import (
	"frascati/comp/auth"
	"frascati/service"
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
