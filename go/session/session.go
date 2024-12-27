package session

import (
	"frascati/dto"
	"frascati/exception"

	"github.com/gin-gonic/gin"
)

func PassAuthValue(ctx *gin.Context) (dto.UserTokenReturn, exception.Exception) {
	userDataRaw, ok := ctx.Get("user_data")
	if !ok {
		return dto.UserTokenReturn{}, exception.NewBaseException(
			exception.CAUSE_INTERNAL,
			"auth/session",
			"something is wrong in our end",
			ErrSessionValueNotPassed,
		)
	}

	userData, ok := userDataRaw.(dto.UserTokenReturn)
	if !ok {
		return dto.UserTokenReturn{}, exception.NewBaseException(
			exception.CAUSE_INTERNAL,
			"auth/session",
			"something is wrong in our end",
			ErrSessionDataParsing,
		)
	}

	return userData, nil
}
