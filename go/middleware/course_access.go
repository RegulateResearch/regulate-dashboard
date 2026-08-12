package middleware

import (
	"errors"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/service"
	"frascati/session"
	"frascati/typing"

	"github.com/gin-gonic/gin"
)

type CourseAccessMiddleware struct {
	baseMiddleware
	memberService service.CourseMemberService
}

func NewCourseAccessMiddleware(memberService service.CourseMemberService) CourseAccessMiddleware {
	return CourseAccessMiddleware{
		memberService: memberService,
	}
}

func (m CourseAccessMiddleware) GetViewAccess(ctx *gin.Context) {
	m.GetAccessData(func(cm entity.CourseMember) bool {
		return cm.Role.HaveViewAccess()
	})(ctx)
}

func (m CourseAccessMiddleware) GetWriteAccess(ctx *gin.Context) {
	m.GetAccessData(func(cm entity.CourseMember) bool {
		return cm.Role.HaveWriteAccess()
	})(ctx)
}

func (m CourseAccessMiddleware) GetAccessData(accessFn func(entity.CourseMember) bool) func(*gin.Context) {
	return func(ctx *gin.Context) {
		userData, exc := session.PassAuthValue(ctx)
		if exc != nil {
			ctx.Error(exc)
			ctx.Abort()
			return
		}

		courseId := typing.IDFromString(ctx.Param("course_id"))
		userId := userData.ID

		memberData, err := m.memberService.GetAccessData(m.extractCtx(ctx), courseId, userId)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		if !accessFn(memberData) {
			newErr := errors.New("this user is not granted access to this particular endpoint")
			exc := exception.NewBaseException(exception.CAUSE_FORBIDDEN, "course_access/middleware", exception.NOT_FOUND, newErr)
			ctx.Error(exc)
			ctx.Abort()
			return
		}
	}
}
