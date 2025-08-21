package handler

import (
	"frascati/converter"
	"frascati/dto"
	"frascati/response"
	"frascati/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	handler := AuthHandler{
		authService: authService,
	}

	return handler
}

func (h AuthHandler) Register(ctx *gin.Context) {
	var userRegister dto.UserRegister
	err := ctx.ShouldBindBodyWithJSON(&userRegister)
	if err != nil {
		ctx.Error(err)
		return
	}

	userWrite := converter.UserRegisterToEntity(userRegister)
	userReturn, exc := h.authService.Register(ctx, userWrite)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusCreated, converter.UserEntityToDTO(userReturn))
}

func (h AuthHandler) Login(ctx *gin.Context) {
	var userLogin dto.UserLogin
	err := ctx.ShouldBindBodyWithJSON(&userLogin)
	if err != nil {
		ctx.Error(err)
		return
	}

	userWrite := converter.UserLoginToEntity(userLogin)
	token, exc := h.authService.Login(ctx, userWrite)

	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(token, "login success"))
}
