package initiator

import (
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
	_ = InitLogger()
	sampleLog.Info("logger initialized")
}
