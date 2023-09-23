package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	web_programing_susu "web-programing-susu"
	handlers2 "web-programing-susu/pkg/handlers"
	"web-programing-susu/pkg/repository"
	"web-programing-susu/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) // Формат логов делаем в json виде
	// Достаем книфиги
	if err := initConfig(); err != nil {
		log.Fatalf(err.Error())
	}

	// Загрузим переменнные окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable %s", err.Error())
	}

	// Инициализируем бд
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslMode"),
	})
	if err != nil {
		logrus.Fatalf("faled to initialize db, %s", err.Error())
	}

	// Создаем репозиторий и сервис, затем инициализируем бд, а далее создаем хендлеры
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers2.NewHandler(services)

	// Создаем сервер и присваем ему хендлеры
	fmt.Println("Server listening...")
	srv := new(web_programing_susu.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
func initConfig() error {
	// Достаем содержимое конфиг файла
	viper.AddConfigPath("config") // Директория конфиг файла
	viper.SetConfigName("config") // Название конфиг файла
	return viper.ReadInConfig()

}
