package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"player/pkg/cache"
	"player/pkg/database"
	"player/pkg/handler"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db, err := database.NewPostgresDB(initDBConfig())

	if err != nil {
		logrus.Fatalf("Error connecting to Database: %s", err.Error())
	}
	defer db.Close()

	repository := database.NewInvoiceRepository(db)
	caching := cache.NewCache()

	h := handler.NewHandler(repository, caching)
	if err := h.Init(); err != nil {
		logrus.Printf("Error occurred while running HTTP server: %s", err.Error())
	}
}

func initDBConfig() database.Config {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "qwerty"
	}

	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "postgres"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}

	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	return database.Config{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
}
