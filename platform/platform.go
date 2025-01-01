package platform

import "context"

type Token interface {
	GenerateAccessToken(ctx context.Context,id string) (string,error)
	GenerateRefreshToken(ctx context.Context,id string) (string,error)
}