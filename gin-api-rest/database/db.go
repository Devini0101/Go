package database

import (
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	// creates a table based on the modal plus the id, created, updated and deleted at timestamps
	DB.AutoMigrate(&models.Student{})
}
