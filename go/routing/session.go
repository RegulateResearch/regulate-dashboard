package routing

import (
	"frascati/routing/grouping"
	"frascati/setup"
)

func setupSessionRouting(routers grouping.Routes, handlers setup.Handlers) {
	sessionHandler := handlers.Session
	routers.Login.GET("/session", sessionHandler.CheckSession)
}
