package handler

import (
	"frascati/response"
	"frascati/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TryHandler struct {
	tryService service.TryService
}

func NewTryHandler(tryService service.TryService) TryHandler {
	return TryHandler{
		tryService: tryService,
	}
}

func (h TryHandler) TryLongOp(ctx *gin.Context) {
	time.Sleep(20 * time.Second)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse("success", "success"))
}

func (h TryHandler) TryBackgroundSuccess(ctx *gin.Context) {
	res := h.tryService.TryBackgroundSuccess(ctx.Param("task_name"))
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(res, "success"))
}

func (h TryHandler) TryBackgroundFail(ctx *gin.Context) {
	res := h.tryService.TryBackgroundFail(ctx.Param("task_name"))
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(res, "success"))
}
