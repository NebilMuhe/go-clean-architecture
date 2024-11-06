package initiator

import (
	"go-clean-architecture/internal/module"
	"go-clean-architecture/internal/module/user"
	"go-clean-architecture/platform/logger"
)

type Module struct {
	user module.User
}

func InitModule(persistence Persistence, log logger.Logger) Module {
	return Module{
		user: user.Init(persistence.user, log),
	}
}
