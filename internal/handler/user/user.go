package user

import (
	"go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/platform/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userModule module.User
	log        logger.Logger
}

func Init(module module.User, log logger.Logger) handler.UserHandler {
	return UserHandler{
		userModule: module,
		log:        log,
	}
}

// Login implements handler.UserHandler.
func (u UserHandler) Login(ctx *gin.Context) {
	panic("unimplemented")
}

// SignUp implements handler.UserHandler.
func (u UserHandler) SignUp(ctx *gin.Context) {
	panic("unimplemented")
}
