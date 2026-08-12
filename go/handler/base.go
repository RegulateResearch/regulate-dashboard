package handler

import (
	"frascati/typing"

	"github.com/gin-gonic/gin"
)

type baseHandler struct{}

func (h baseHandler) extractCtx(ctx *gin.Context) typing.Context {
	return typing.NewDictionaryContext(ctx.Request.Context())
}
