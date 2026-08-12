package handler

import (
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"frascati/typing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	baseHandler
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
	userReturn, exc := h.authService.Register(h.extractCtx(ctx), userWrite)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	userReturnDto := converter.UserEntityToDTO(userReturn)
	ctx.JSON(http.StatusCreated, response.NewSuccessResponse(userReturnDto, "success create"))
}

func (h AuthHandler) Login(ctx *gin.Context) {
	var userLogin dto.UserLogin
	err := ctx.ShouldBindBodyWithJSON(&userLogin)
	if err != nil {
		ctx.Error(err)
		return
	}

	userWrite := converter.UserLoginToEntity(userLogin)
	deepCtx := typing.NewDictionaryContext(ctx.Request.Context())
	token, exc := h.authService.Login(deepCtx, userWrite)

	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(token, "login success"))
}
