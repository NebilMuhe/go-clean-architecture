package initiator

import (
	"go-clean-architecture/internal/handler"
	"go-clean-architecture/internal/handler/user"
	"go-clean-architecture/platform/logger"
)

type Handler struct {
	user handler.UserHandler
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		user: user.Init(module.user, log),
	}
}
