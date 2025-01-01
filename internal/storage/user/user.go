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
	userResponse, err := u.db.AddUser(ctx, db.AddUserParams{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		u.log.Error(ctx, "failed to save user", zap.Error(err))
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:    userResponse.ID.String(),
		Email: userResponse.Email,
	}, nil
}

func (u *UserPersistence) UserExist(ctx context.Context, email string) (bool, error) {
	exists, err := u.db.Queries.UserExists(ctx, email)
	if err != nil {
		u.log.Error(ctx, "failed to check user exist", zap.Error(err))
		return false, err
	}
	return exists, nil
}

func (u *UserPersistence) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	user, err := u.db.GetUserByEmail(ctx, email)
	if err != nil {
		u.log.Error(ctx, "failed to get user by email", zap.Error(err))
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Password: user.Password,
	}, nil
}


