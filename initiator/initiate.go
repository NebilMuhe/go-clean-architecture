package initiator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
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

	log.Info(context.Background(), "initializing migration")
	m := InitMigration(context.Background(), viper.GetString("db.url"), log)
	UpMigration(context.Background(), m, log)
	log.Info(context.Background(), "initialized migration")

	log.Info(context.Background(), "initializing persistence")
	InitPersistence(context.Background(), log)
	log.Info(context.Background(), "initialized persistence")

	log.Info(context.Background(), "initializing module")
	InitModule(context.Background(), log)
	log.Info(context.Background(), "module initialized")

	log.Info(context.Background(), "initializing handler")
	InitHandler(context.Background(), log)
	log.Info(context.Background(), "handler initialized")

	log.Info(context.Background(), "initializing server")
	server := gin.New()
	server.Use(ginzap.RecoveryWithZap(log.GetZapLogger().Named("gin"), true))
	log.Info(context.Background(),"server initialized")

	log.Info(context.Background(),"initializing route")
	server.Group("/v1")
	InitRoute()
	log.Info(context.Background(),"route initialized")

	log.Info(context.Background(),"initializing http server")
	srv := &http.Server{
		Addr: viper.GetString("server.host")+":" + viper.GetString("server.port"),
		Handler: server,
	}

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt)
	signal.Notify(quit,syscall.SIGTERM)

	go func ()  {
		log.Info(context.Background(),"server started",
			zap.String("host",viper.GetString("server.host")),
			zap.String("port",viper.GetString("server.port")),
		)	
		log.Info(context.Background(),fmt.Sprintf("server stopped with error %v",srv.ListenAndServe()))
	}()
	sig := <-quit
	log.Info(context.Background(),fmt.Sprintf("server shutting down with signal %v",sig))
	ctx,cancel := context.WithTimeout(context.Background(),viper.GetDuration("server.timeout"))
	defer cancel()

	if err := srv.Shutdown(ctx); err!= nil{
		log.Fatal(context.Background(),"error while shutting down server",zap.Error(err))
	}
	log.Info(context.Background(),"server shutdown successfully")
}
