package handler

import (
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SsoHandler struct {
	serv service.SsoUiService
}

func NewSsoHandler(ssoService service.SsoUiService) SsoHandler {
	return SsoHandler{
		serv: ssoService,
	}
}

func (h SsoHandler) Validate(ctx *gin.Context) {
	var validateData dto.SsoValidateData
	err := ctx.ShouldBindBodyWithJSON(&validateData)
	if err != nil {
		ctx.Error(err)
		return
	}

	res, exc := h.serv.Validate(ctx, validateData.Ticket, validateData.Service)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(res, "success"))
}
