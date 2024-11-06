package persistencedb

import (
	"go-clean-architecture/internal/constant/model/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PersistenceDB struct {
	*db.Queries
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) PersistenceDB {
	return PersistenceDB{
		pool:    pool,
		Queries: db.New(pool),
	}
}
