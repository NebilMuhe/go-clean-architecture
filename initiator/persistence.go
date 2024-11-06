package initiator

import (
	persistencedb "go-clean-architecture/internal/constant/model/persistenceDB"
	"go-clean-architecture/internal/storage"
	"go-clean-architecture/internal/storage/user"
	"go-clean-architecture/platform/logger"
)

type Persistence struct {
	user storage.User
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{
		user: user.Init(db,log),
	}
}
