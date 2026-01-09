package handler

import (
	"frascati/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TryHandler struct{}

func NewTryHandler() TryHandler {
	return TryHandler{}
}

func (h TryHandler) TryLongOp(ctx *gin.Context) {
	time.Sleep(20 * time.Second)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse("success", "success"))
}
