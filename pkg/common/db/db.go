package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



func Init(url string) *gorm.DB{
	Database, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}
	
	return Database
}
