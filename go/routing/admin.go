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

func SetupAdminRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware, db *sql.DB) {
	group := router.Group("/admin")
	group.Use(authMiddleware.Authenticate)
	group.Use(middleware.Authorize(func(userData dto.UserTokenReturn) bool {
		return userData.Role == constant.ROLE_ADMIN
	}))

	adminHandler := handler.NewAdminHandler(service.NewUserService(repository.NewUserRepository(db)))
	group.GET("", adminHandler.GetAll)
}
