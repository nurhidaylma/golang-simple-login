package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	connection, err := gorm.Open(postgres.Open(Database_url), &gorm.Config{})
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb, _ := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")
	return connection
}
