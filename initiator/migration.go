package initiator

import (
	"context"
	"fmt"
	"go-clean-architecture/platform/logger"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func InitMigration(ctx context.Context, url string, log logger.Logger) *migrate.Migrate {
	path := fmt.Sprintf("cockroachdb://%s", strings.Split(url, "://")[1])
	file := "file://db/migrations"
	m, err := migrate.New(file, path)

	if err != nil {
		log.Fatal(ctx, "failed to initialize migration", zap.Error(err))
	}

	return m
}

func UpMigration(ctx context.Context, m *migrate.Migrate, log logger.Logger) {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(ctx, "failed to migrate", zap.Error(err))
	}
}
