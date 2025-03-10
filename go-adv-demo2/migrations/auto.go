package main

import (
	"github.com/joho/godotenv"
	"go/adv-demo2/internal/link"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("Dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{})

}
