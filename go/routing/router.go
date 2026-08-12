package routing

import (
	"frascati/middleware"
	"frascati/routing/grouping"
	"frascati/setup"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app setup.App) *gin.Engine {
	r := gin.New()

	middlewares := app.Middlewares()
	r.Use(middlewares.Gatekeeper.Process)
	r.Use(middlewares.Logger.LogActivities)
	r.Use(gin.Recovery())
	r.Use(middleware.HandleError)

	routes := grouping.AllRoutes(r, middlewares)

	setupEndpoints(routes, app.Handlers(), middlewares)
	return r
}

func setupEndpoints(routes grouping.Routes, handlers setup.Handlers, middlewares setup.Middlewares) {
	setupAuthRouting(routes, handlers)
	setupCourseRouter(routes, handlers, middlewares)
	setupUsersRouting(routes, handlers)
	setupMyRouter(routes, handlers)
	setupRecordRoutes(routes, handlers)
	setupTryRoutes(routes, handlers)
}
