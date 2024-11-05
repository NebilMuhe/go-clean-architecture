package initiator

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(path, name, configType string, logger *zap.Logger) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("failed to initialize config", zap.Error(err))
	}

	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("config file changed to", zap.String("file name:", in.Name))
	})
}
