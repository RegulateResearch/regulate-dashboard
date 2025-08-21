package middleware

import (
	"frascati/entity"
	"frascati/exception"
	middleware_exception "frascati/middleware/exception"
	"frascati/service/auth"
	"frascati/session"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService auth.JwtService
}

func NewAuthMiddleware(jwtService auth.JwtService) AuthMiddleware {
	return AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m AuthMiddleware) Authenticate(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	splitStrs := strings.Split(header, " ")
	splitLen := len(splitStrs)

	if splitLen <= 0 || splitLen != 2 || splitStrs[0] != "Bearer" {
		exc := exception.NewBaseException(exception.CAUSE_USER, "auth/middleware", "invalid token", middleware_exception.ErrAuthInvalidBearerTokenFormat)
		ctx.Error(exc)
		ctx.Abort()

		return
	}

	tokenStr := splitStrs[1]
	userData, exc := m.jwtService.ParseToken(tokenStr)
	if exc != nil {
		ctx.Error(exc)
		ctx.Abort()

		return
	}

	ctx.Set("user_data", userData)
}

func (m AuthMiddleware) Authorize(authorizeFn func(entity.Session) bool) func(*gin.Context) {
	return func(ctx *gin.Context) {
		userData, exc := session.PassAuthValue(ctx)
		if exc != nil {
			ctx.Error(exc)
			ctx.Abort()
			return
		}

		if !authorizeFn(userData) {
			exc = exception.NewBaseException(exception.CAUSE_FORBIDDEN, "auth/middleware", "resource not found", middleware_exception.ErrAuthorizationFailsForbidden)
			ctx.Error(exc)
			ctx.Abort()
			return
		}
	}
}
