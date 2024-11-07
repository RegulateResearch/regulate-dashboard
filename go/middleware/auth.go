package middleware

import (
	"frascati/constant"
	"frascati/service/auth"

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

}

func Authorize(role constant.Role) func(*gin.Context) {
	return func(ctx *gin.Context) {

	}
}
