package initiator

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	sampleLog, err := zap.NewProduction()
	if err != nil {
		sampleLog.Fatal("failed to start sample logger", zap.Error(err))
	}
	sampleLog.Info("initializing config")
	InitConfig("config", "config", "yaml", sampleLog)
	sampleLog.Info("config initialized")

	sampleLog.Info("initializing logger")
	log := InitLogger()
	sampleLog.Info("logger initialized")

	log.Info(context.Background(), "initializing database")
	InitDB(context.Background(), viper.GetString("db.url"), log)
	log.Info(context.Background(), "database initialized")
}
