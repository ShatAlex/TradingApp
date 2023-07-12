package main

import (
	"os"

	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/handler"
	"github.com/ShatAlex/trading-app/pkg/repository"
	"github.com/ShatAlex/trading-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables:%s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initilized db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	handlers := handler.NewHandler(service)

	server := new(trade.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occured while running server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
