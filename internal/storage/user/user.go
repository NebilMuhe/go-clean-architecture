package user

import (
	"context"
	"go-clean-architecture/internal/constant/model/db"
	"go-clean-architecture/internal/constant/model/dto"
	persistencedb "go-clean-architecture/internal/constant/model/persistenceDB"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/platform/logger"

	"go.uber.org/zap"
)

type UserPersistence struct {
	log logger.Logger
	db  persistencedb.PersistenceDB
}

func Init(db persistencedb.PersistenceDB, log logger.Logger) storage.User {
	return &UserPersistence{
		log: log,
		db:  db,
	}
}

func (u *UserPersistence) SignUp(ctx context.Context, user dto.User) (dto.UserResponse, error) {
	userResponse,err := u.db.AddUser(ctx, db.AddUserParams{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil{
		u.log.Error(ctx,"failed to save user",zap.Error(err))
		return dto.UserResponse{},err
	}

	return dto.UserResponse{
		ID: userResponse.ID.String(),
		Email: userResponse.Email,
	},nil
}

// Login implements storage.User.
func (u *UserPersistence) Login(ctx context.Context, email string, password string) {
	panic("unimplemented")
}
