package initiator

import (
	"go-clean-architecture/platform/logger"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func InitLogger() logger.Logger{
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(viper.GetInt("logger.level")))

	zapLogger,err := config.Build(zap.AddCallerSkip(1))
	if err!=nil{
		log.Fatalf("failed to initiatlize logger %s",err)
	}

	l := logger.NewLogger(zapLogger)
	return l
}