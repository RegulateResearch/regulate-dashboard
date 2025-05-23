package setup

import (
	"database/sql"
	"frascati/middleware"
	"frascati/routing"

	"github.com/gin-gonic/gin"
)

func SetupApplication(r *gin.Engine, db *sql.DB) {
	authService, jwtService := SetupAuthFunctionalities(db)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	routing.SetupRouterHandler(r, authService)
	routing.SetupAdminRouting(r, authMiddleware, db)
	routing.SetupUserRouting(r, authMiddleware, db)
	routing.SetupSessionRouting(r, authMiddleware)
}
