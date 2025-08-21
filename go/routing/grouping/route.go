package grouping

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	groups *gin.RouterGroup
}

func newRoute(router gin.IRouter, path string, handlers ...gin.HandlerFunc) Route {
	groups := router.Group(path, handlers...)
	return Route{
		groups: groups,
	}
}

func (r Route) Group(path string, handlers ...gin.HandlerFunc) Route {
	return newRoute(r.groups, path, handlers...)
}

func (r Route) GET(path string, handlers ...gin.HandlerFunc) {
	r.groups.GET(path, handlers...)
}

func (r Route) POST(path string, handlers ...gin.HandlerFunc) {
	r.groups.POST(path, handlers...)
}

func (r Route) PUT(path string, handlers ...gin.HandlerFunc) {
	r.groups.PUT(path, handlers...)
}

func (r Route) DELETE(path string, handlers ...gin.HandlerFunc) {
	r.groups.DELETE(path, handlers...)
}
