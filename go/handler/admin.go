package handler

import (
	"frascati/response"
	"frascati/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userService service.UserService
}

func NewAdminHandler(userService service.UserService) AdminHandler {
	return AdminHandler{
		userService: userService,
	}
}

func (h AdminHandler) GetAll(ctx *gin.Context) {
	res, exc := h.userService.FindAll(ctx)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(res, "success admin"))
}
