package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupRecordRoutes(routes grouping.Routes, handlers setup.Handlers) {
	group := routes.NoLogin.Group("/records")
	recordHandler := handlers.Record

	group.GET("", recordHandler.FindAll)
	group.POST("", recordHandler.AddBulk)
}
