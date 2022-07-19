package main

import (
	"github.com/basterrus/http_rest_api_service_golang"
	"github.com/basterrus/http_rest_api_service_golang/pkg/handler"
	"github.com/basterrus/http_rest_api_service_golang/pkg/repository"
	"github.com/basterrus/http_rest_api_service_golang/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	// logrus json formatter
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// initConfig -- read variables from config.yaml
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error load configuration file %s", err.Error())
	}
	//Load password from .env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Connect from postgres db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка подключения к базе: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(http_rest_api_service_golang.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error server run %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
