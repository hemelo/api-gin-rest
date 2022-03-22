package database

import (
	"fmt"
	"log"

	"api-go-gin/config"
	"api-go-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	var dbName, hostName, port, user, password = config.DatabaseConnector()
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", hostName, port, user, dbName, password)

	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	
	DB.AutoMigrate(&models.Aluno{})
}
