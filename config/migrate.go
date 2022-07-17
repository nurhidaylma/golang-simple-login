package config

import (
	"github.com/google/uuid"
	"github.com/nurhidaylma/golang-simple-login/repository"
	"gorm.io/gorm"
)

func InitialMigration() {
	userID, _ := uuid.NewRandom()

	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(repository.User{
		ID: userID,
	})
}

func Closedatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}
