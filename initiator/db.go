package initiator

import (
	"context"
	"go-clean-architecture/platform/logger"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitDB(ctx context.Context, url string, log logger.Logger) *pgxpool.Pool{
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(ctx, "failed to parse config", zap.Error(err))
	}

	idleTimeout := viper.GetDuration("db.idle_time")
	if idleTimeout == 0 {
		idleTimeout = time.Minute * 4
	}

	config.MaxConnIdleTime = idleTimeout
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(ctx, "failed to create pool", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(ctx, "failed to ping database", zap.Error(err))
	}

	return pool
}
