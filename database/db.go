package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectStringDB := "host=localhost user=root password=root dbname=personalities port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectStringDB))

	if err != nil {
		log.Panic("Erro ao se conectar com banco de dados")
	}
	fmt.Println("Conectado no banco de dados..")
}
