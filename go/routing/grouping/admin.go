package grouping

import (
	"frascati/constants"
	"frascati/obj/entity"
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func adminRoute(router gin.IRouter, middlewares setup.Middlewares) Route {
	return newRoute(router, "/admin", middlewares.Auth.Authenticate, middlewares.Auth.Authorize(func(userData entity.Session) bool {
		return userData.Role == constants.ROLE_ADMIN
	}))
}
