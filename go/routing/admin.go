package routing

import (
	"database/sql"
	"frascati/constants"
	"frascati/entity"
	"frascati/handler"
	"frascati/middleware"
	"frascati/repository"
	"frascati/service"

	"github.com/gin-gonic/gin"
)

func SetupAdminRouting(router *gin.Engine, authMiddleware middleware.AuthMiddleware, db *sql.DB) {
	group := router.Group("/admin")
	group.Use(authMiddleware.Authenticate)
	group.Use(middleware.Authorize(func(userData entity.SessionData) bool {
		return userData.Role == constants.ROLE_ADMIN
	}))

	adminHandler := handler.NewAdminHandler(service.NewUserService(repository.NewUserRepository(db)))
	group.GET("", adminHandler.GetAll)
}
