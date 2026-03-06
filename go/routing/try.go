package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupTryRoutes(routers grouping.Routes, handlers setup.Handlers) {
	group := routers.NoLogin.Group("/try")

	group.GET("/long", handlers.Try.TryLongOp)
	group.GET("/background/:task_name/success", handlers.Try.TryBackgroundSuccess)
	group.GET("/background/:task_name/fail", handlers.Try.TryBackgroundFail)
}
