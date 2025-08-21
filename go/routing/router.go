package routing

import (
	"frascati/middleware"
	"frascati/routing/grouping"
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handlers setup.Handlers, middlewares setup.Middlewares) *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Logger.LogActivities)
	r.Use(gin.Recovery())
	r.Use(middleware.HandleError)

	routes := grouping.AllRoutes(r, middlewares)

	setupEndpoints(routes, handlers)
	return r
}

func setupEndpoints(routes grouping.Routes, handlers setup.Handlers) {
	setupAuthRouting(routes, handlers)
	setupSessionRouting(routes, handlers)
	setupTryGetUser(routes, handlers)
}
