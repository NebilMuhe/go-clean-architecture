package token

import (
	"context"
	"go-clean-architecture/platform"
	"go-clean-architecture/platform/logger"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type Token struct {
	log       logger.Logger
	secretKey string
}

func InitJWT(log logger.Logger, secretKey string) platform.Token {
	return &Token{
		log:       log,
		secretKey: secretKey,
	}
}

func (t *Token) GenerateAccessToken(ctx context.Context, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString(t.secretKey)
	if err != nil {
		t.log.Error(ctx, "failed to generate token", zap.Error(err))
		return "", err
	}
	return tokenString, nil
}

func (t *Token) GenerateRefreshToken(ctx context.Context,id string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":id,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString,err := token.SignedString(t.secretKey)
	if err != nil{
		t.log.Error(ctx,"failed to generate token",zap.Error(err))
		return "",err
	}

	return tokenString,nil
}