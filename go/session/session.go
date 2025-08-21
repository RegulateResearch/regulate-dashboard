package session

import (
	"frascati/entity"
	"frascati/exception"

	"github.com/gin-gonic/gin"
)

func PassAuthValue(ctx *gin.Context) (entity.Session, exception.Exception) {
	userDataRaw, ok := ctx.Get("user_data")
	emptyData := entity.Session{}
	if !ok {
		return emptyData, exception.NewBaseException(
			exception.CAUSE_INTERNAL,
			"auth/session",
			"something is wrong in our end",
			ErrSessionValueNotPassed,
		)
	}

	userData, ok := userDataRaw.(entity.Session)
	if !ok {
		return emptyData, exception.NewBaseException(
			exception.CAUSE_INTERNAL,
			"auth/session",
			"something is wrong in our end",
			ErrSessionDataParsing,
		)
	}

	return userData, nil
}
