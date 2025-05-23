package handler

import (
	"frascati/converter"
	"frascati/dto"
	"frascati/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) AuthHandler {
	handler := AuthHandler{
		authService: authService,
	}

	return handler
}

func (h AuthHandler) Register(ctx *gin.Context) {
	var userRegister dto.UserWrite
	err := ctx.ShouldBindBodyWithJSON(&userRegister)
	if err != nil {
		ctx.Error(err)
		return
	}

	userWrite := converter.ConvertUserWriteToEntity(userRegister)
	userReturn, exc := h.authService.Register(ctx, userWrite)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusCreated, converter.ConvertUserEntityToDTO(userReturn))
}

func (h AuthHandler) Login(ctx *gin.Context) {
	var userLogin dto.UserWrite
	err := ctx.ShouldBindBodyWithJSON(&userLogin)
	if err != nil {
		ctx.Error(err)
		return
	}

	userWrite := converter.ConvertUserWriteToEntity(userLogin)
	token, exc := h.authService.Login(ctx, userWrite)

	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
