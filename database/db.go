package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"api-go-gin/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}