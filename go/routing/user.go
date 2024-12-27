package routing

import (
	"database/sql"
	"frascati/constant"
	"frascati/dto"
	"frascati/handler"
	"frascati/middleware"
	"frascati/repository"
	"frascati/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware, db *sql.DB) {
	group := router.Group("/user")
	group.Use(authMiddleware.Authenticate)
	group.Use(middleware.Authorize(func(userData dto.UserTokenReturn) bool {
		return userData.Role == constant.ROLE_USER
	}))

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))
	group.GET("", userHandler.GetAll)
}
