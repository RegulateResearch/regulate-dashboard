package middleware

import (
	"frascati/typing"

	"github.com/gin-gonic/gin"
)

type baseMiddleware struct{}

func (m baseMiddleware) extractCtx(ctx *gin.Context) typing.Context {
	return typing.NewDictionaryContext(ctx.Request.Context())
}
