package handler

import (
	"frascati/obj/converter"
	"frascati/response"
	"frascati/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct{}

func NewSessionHandler() SessionHandler {
	return SessionHandler{}
}

func (h SessionHandler) CheckSession(ctx *gin.Context) {
	sessionData, err := session.PassAuthValue(ctx)
	sessionDto := converter.SessionDataToDto(sessionData)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(sessionDto, "token is valid"))
}
