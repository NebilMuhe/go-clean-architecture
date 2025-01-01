package initiator

import (
	"context"
	"go-clean-architecture/platform"
	"go-clean-architecture/platform/logger"
	"go-clean-architecture/platform/token"
)

type Platform struct{
	token platform.Token
}


func InitPlatform(ctx context.Context,log logger.Logger,secretKey string) Platform{
	return Platform{
		token: token.InitJWT(log,secretKey),
	}
}