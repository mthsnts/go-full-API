package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	fmt.Println(dsn)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Failed to connect to db")
	}

}
