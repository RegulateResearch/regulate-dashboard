package setup

import (
	"database/sql"
	"frascati/routing"

	"github.com/gin-gonic/gin"
)

func SetupApplication(r *gin.Engine, db *sql.DB) {
	authService, _ := SetupAuthFunctionalities(db)

	routing.SetupRouterHandler(r, authService)
}
