package storage

import (
	"context"
	"go-clean-architecture/internal/constant/model/dto"
)

type User interface {
	SignUp(ctx context.Context, user dto.User) (dto.UserResponse,error)
	Login(ctx context.Context, email, password string)
}
