package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
}
