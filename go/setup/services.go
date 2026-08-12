package setup

import (
	"frascati/comp/auth"
	"frascati/comp/background"
	"frascati/comp/ssoui"
	"frascati/service"
)

type services struct {
	auth         service.AuthService
	course       service.CourseService
	courseItem   service.CourseItemService
	courseMember service.CourseMemberService
	user         service.UserService
	sso          service.SsoUiService
	my           service.MyService
	record       service.RecordService
	try          service.TryService
}

func setupServices(repos repositories, jwt auth.JwtService, bcrypt auth.BcryptService, backgroundProcessor background.Processor, ssoClient ssoui.Client) services {
	return services{
		auth:         service.NewAuthService(repos.auth, bcrypt, jwt, repos.transactor),
		course:       service.NewCourseService(repos.course),
		courseItem:   service.NewCourseItemService(repos.courseItem, repos.course),
		courseMember: service.NewCourseMemberService(repos.courseMember, repos.course, repos.user, repos.transactor),
		user:         service.NewUserService(repos.user),
		sso:          service.NewSsoUiService(ssoClient, repos.auth, jwt),
		my:           service.NewMyService(repos.user, repos.course),
		record:       service.NewRecordService(repos.record),
		try:          service.NewTryService(backgroundProcessor),
	}
}
