package routing

import (
	"frascati/middleware"
	"frascati/prep/logger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(logger logger.EnhancedLogger) *gin.Engine {
	r := gin.New()
	r.Use(middleware.NewLoggerMiddleware(logger).LogActivities)
	r.Use(gin.Recovery())
	r.Use(middleware.HandleError)
	return r
}
