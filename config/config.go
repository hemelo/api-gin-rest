package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TestMode bool
}

var Settings Config

func DatabaseConnector() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	var dbName, hostName, port, user, password string
	if Settings.TestMode {
		dbName = os.Getenv("DB_NAME_TEST")
		hostName = os.Getenv("DB_HOST_TEST")
		port = os.Getenv("DB_PORT_TEST")
		user = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASSWORD_TEST")
	} else {
		dbName = os.Getenv("DB_NAME")
		hostName = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
	}

	return dbName, hostName, port, user, password
}