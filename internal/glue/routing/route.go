package routing

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Method     string
	Path       string
	Handler    gin.HandlerFunc
	Middleware []gin.HandlerFunc
}

func RegisterRoutes(group *gin.RouterGroup, router []Router) {

	for _, route := range router {
		var handler []gin.HandlerFunc

		handler = append(handler, route.Middleware...)
		handler = append(handler, route.Handler)

		group.Handle(route.Method, route.Path, handler...)
	}
}
