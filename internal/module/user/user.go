package user

import (
	"context"
	"go-clean-architecture/internal/constant/model/dto"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/platform/logger"

	"go.uber.org/zap"
)

type UserModule struct {
	log         logger.Logger
	persistence storage.User
}

func Init(persistence storage.User, log logger.Logger) module.User {
	return &UserModule{
		log:         log,
		persistence: persistence,
	}
}



func (u *UserModule) SignUp(ctx context.Context, user dto.User) (dto.UserResponse, error) {
	if err := user.Validate(); err!= nil{
		u.log.Error(ctx,"invalid user input",zap.Error(err))
		return dto.UserResponse{},err
	}
	userResponse,err := u.persistence.SignUp(ctx,user)
	if err != nil{
		return dto.UserResponse{},err
	}

	return userResponse,nil
}

func (u *UserModule) Login(ctx context.Context, email string, password string) {
	panic("unimplemented")
}
