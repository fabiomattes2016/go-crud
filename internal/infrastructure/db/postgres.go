package db

import (
	"log"

	"github.com/fabiomattes2016/go-crud/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host, user, password, dbname, port string) {
	var err error

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=America/Sao_Paulo"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&domain.User{})

	log.Println("Database connection established")
}
