package routing

import (
	"database/sql"
	"frascati/constants"
	"frascati/handler"
	"frascati/middleware"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware, db *sql.DB) {
	group := router.Group("/user")
	group.Use(authMiddleware.Authenticate)
	group.Use(authMiddleware.Authorize(func(userData entity.Session) bool {
		return userData.Role == constants.ROLE_USER
	}))

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))
	group.GET("", userHandler.GetAll)
}
