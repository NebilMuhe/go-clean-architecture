package storage

import "context"

type User interface {
	SignUp(ctx context.Context, email, password string)
	Login(ctx context.Context, email, password string)
}
