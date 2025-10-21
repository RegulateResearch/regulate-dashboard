package routing

import (
	"frascati/middleware"
	"frascati/routing/grouping"
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app setup.App) *gin.Engine {
	r := gin.New()
	r.Use(app.Middlewares.Logger.LogActivities)
	r.Use(gin.Recovery())
	r.Use(middleware.HandleError)

	routes := grouping.AllRoutes(r, app.Middlewares)

	setupEndpoints(routes, app.Handlers)
	return r
}

func setupEndpoints(routes grouping.Routes, handlers setup.Handlers) {
	setupAuthRouting(routes, handlers)
	setupSessionRouting(routes, handlers)
	setupTryGetUser(routes, handlers)
}
