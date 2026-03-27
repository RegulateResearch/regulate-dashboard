package setup

import (
	"frascati/comp/auth"
	"frascati/comp/background"
	"frascati/service"
)

type services struct {
	auth   service.AuthService
	course service.CourseService
	user   service.UserService
	record service.RecordService
	try    service.TryService
}

func setupServices(repos repositories, jwt auth.JwtService, bcrypt auth.BcryptService, backgroundProcessor background.Processor) services {
	return services{
		auth:   service.NewAuthService(repos.auth, bcrypt, jwt, repos.transactor),
		course: service.NewCourseService(repos.course),
		user:   service.NewUserService(repos.user),
		record: service.NewRecordService(repos.record),
		try:    service.NewTryService(backgroundProcessor),
	}
}
