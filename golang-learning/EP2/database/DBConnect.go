package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

var err error

func DD() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=Demo port=5432 sslmode=disable TimeZone=Asia/Shanghai search_path=Demo"
	DBConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
