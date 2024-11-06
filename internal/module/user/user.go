package user

import (
	"context"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/platform/logger"
)

type UserModule struct {
	log         logger.Logger
	persistence storage.User
}

func Init(persistence storage.User, log logger.Logger) module.User {
	return UserModule{
		log:         log,
		persistence: persistence,
	}
}

// Login implements module.User.
func (u UserModule) Login(ctx context.Context, email string, password string) {
	panic("unimplemented")
}

// SignUp implements module.User.
func (u UserModule) SignUp(ctx context.Context, email string, password string) {
	panic("unimplemented")
}
