package storage

import (
	"context"
	"go-clean-architecture/internal/constant/model/dto"
)

type User interface {
	SignUp(ctx context.Context, user dto.User) (dto.UserResponse,error)
	UserExist(ctx context.Context, email string) (bool,error)
	GetUserByEmail(ctx context.Context, email string) (dto.UserResponse,error)
}
