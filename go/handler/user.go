package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/response"
	"frascati/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h UserHandler) GetAll(ctx *gin.Context) {
	res, exc := h.userService.FindAll(ctx)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.UserEntityToDTO)

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success user"))
}
