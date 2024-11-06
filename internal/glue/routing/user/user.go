package user

import (
	"go-clean-architecture/internal/glue/routing"
	"go-clean-architecture/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)



func InitRoute(group *gin.RouterGroup,handler handler.UserHandler) {
	userRoutes := []routing.Router{
		{
			Method: http.MethodPost,
			Path: "/signup",
			Handler: handler.SignUp,
			Middleware: []gin.HandlerFunc{},
		},
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: handler.Login,
			Middleware: []gin.HandlerFunc{},
		},
	}

	routing.RegisterRoutes(group,userRoutes)
}