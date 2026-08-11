package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"frascati/typing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	baseHandler
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (h UserHandler) GetAll(ctx *gin.Context) {
	res, exc := h.userService.FindAll(h.extractCtx(ctx))
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.UserEntityToDTO)

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success user"))
}

func (h UserHandler) GetById(ctx *gin.Context) {
	id := typing.IDFromString(ctx.Param("id"))
	res, err := h.userService.FindById(h.extractCtx(ctx), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := converter.UserEntityToDTO(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h UserHandler) UpdateAccess(ctx *gin.Context) {
	var updateData []dto.UserAccess
	err := ctx.ShouldBindBodyWithJSON(&updateData)
	if err != nil {
		ctx.Error(err)
		return
	}

	users := lambda.MapList(updateData, converter.UserAccessToEntity)
	res, exc := h.userService.UpdateAccessBulk(h.extractCtx(ctx), users)
	if exc != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.UserEntityToDTO)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
