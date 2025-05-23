package session

import (
	"fmt"
	"frascati/entity"
	"frascati/exception"

	"github.com/gin-gonic/gin"
)

func PassAuthValue(ctx *gin.Context) (entity.SessionData, exception.Exception) {
	userDataRaw, ok := ctx.Get("user_data")
	emptyData := entity.SessionData{}
	if !ok {
		return emptyData, exception.NewBaseException(
			exception.CAUSE_INTERNAL,
			"auth/session",
			"something is wrong in our end",
			ErrSessionValueNotPassed,
		)
	}

	userData, ok := userDataRaw.(entity.SessionData)
	fmt.Printf("id: %d", userData.ID)
	fmt.Printf("role: %s", userData.Role.ToString())
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
