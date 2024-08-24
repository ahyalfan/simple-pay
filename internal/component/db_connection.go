package component

import (
	"ahyalfan/golang_e_money/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(conf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Databases.Host, conf.Databases.User, conf.Databases.Password, conf.Databases.Name, conf.Databases.Port, "Asia/Jakarta")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err.Error())
	}
	return db
}
