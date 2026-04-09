package setup

import (
	"frascati/comp/auth"
	"frascati/comp/background"
	"frascati/comp/ssoui"
	"frascati/service"
)

type services struct {
	auth   service.AuthService
	course service.CourseService
	user   service.UserService
	sso    service.SsoUiService
	try    service.TryService
}

func setupServices(repos repositories, jwt auth.JwtService, bcrypt auth.BcryptService, backgroundProcessor background.Processor, ssoClient ssoui.Client) services {
	return services{
		auth:   service.NewAuthService(repos.auth, bcrypt, jwt),
		course: service.NewCourseService(repos.course),
		user:   service.NewUserService(repos.user),
		try:    service.NewTryService(backgroundProcessor),
		sso:    service.NewSsoUiService(ssoClient),
	}
}
