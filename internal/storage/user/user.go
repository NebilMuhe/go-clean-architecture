package user

import (
	"context"
	persistencedb "go-clean-architecture/internal/constant/model/persistenceDB"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/platform/logger"
)

type UserPersistence struct {
	log logger.Logger
	db  persistencedb.PersistenceDB
}

func Init(db persistencedb.PersistenceDB,log logger.Logger) storage.User {
	return &UserPersistence{
		log: log,
		db:  db,
	}
}

// Login implements storage.User.
func (u *UserPersistence) Login(ctx context.Context, email string, password string) {
	panic("unimplemented")
}

// SignUp implements storage.User.
func (u *UserPersistence) SignUp(ctx context.Context, email string, password string) {
	panic("unimplemented")
}
