package grouping

import (
	"frascati/constants"
	"frascati/obj/entity"
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func userRoute(router gin.IRouter, middlewares setup.Middlewares) Route {
	return newRoute(router, "/user", middlewares.Auth.Authenticate, middlewares.Auth.Authorize(func(userData entity.Session) bool {
		return userData.Role == constants.ROLE_USER
	}))
}
