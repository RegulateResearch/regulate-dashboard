package handler

import (
	"frascati/response"
	"frascati/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSession(ctx *gin.Context) {
	sessionData, err := session.PassAuthValue(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(sessionData, "token is valid"))
}
