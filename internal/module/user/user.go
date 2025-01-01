package user

import (
	"context"
	"fmt"
	"go-clean-architecture/internal/constant/model/dto"
	"go-clean-architecture/internal/module"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/platform"
	"go-clean-architecture/platform/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserModule struct {
	log         logger.Logger
	persistence storage.User
	platform    platform.Token
}

func Init(persistence storage.User, platform platform.Token, log logger.Logger) module.User {
	return &UserModule{
		log:         log,
		persistence: persistence,
		platform:    platform,
	}
}

func (u *UserModule) SignUp(ctx context.Context, user dto.User) (dto.UserResponse, error) {
	if err := user.Validate(); err != nil {
		u.log.Error(ctx, "invalid user input", zap.Error(err))
		return dto.UserResponse{}, err
	}
	hashedPassword, err := u.hashPassword(user.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user.Password = hashedPassword
	userResponse, err := u.persistence.SignUp(ctx, user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserModule) Login(ctx context.Context, user dto.User) (dto.Token, error) {
	if err := user.Validate(); err != nil {
		u.log.Error(ctx, "invalid user input", zap.Error(err))
		return dto.Token{}, err
	}

	exist, err := u.persistence.UserExist(ctx, user.Email)
	if err != nil {
		return dto.Token{}, err
	}

	if !exist {
		return dto.Token{}, fmt.Errorf("user not found")
	}

	userResponse, err := u.persistence.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return dto.Token{}, err
	}

	if err := u.comparePassword(userResponse.Password, user.Password); err != nil {
		return dto.Token{}, err
	}

	accessToken, err := u.platform.GenerateAccessToken(ctx, userResponse.ID)
	if err != nil {
		return dto.Token{}, err
	}

	refreshToken, err := u.platform.GenerateRefreshToken(ctx, userResponse.ID)
	if err != nil {
		return dto.Token{}, err
	}

	return dto.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserModule) hashPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error(context.Background(), "failed to hash password", zap.Error(err))
		return "", err
	}
	return string(hashedByte), nil
}

func (u *UserModule) comparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		u.log.Error(context.Background(), "failed to compare password", zap.Error(err))
		return err
	}
	return nil
}
