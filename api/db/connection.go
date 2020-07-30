package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	// Command to set env variables: set -o allexport; source .env; set +o allexport
	database := os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	passwd := os.Getenv("MYSQL_PASSWORD")

	uri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, host, database)
	db, err := gorm.Open("mysql", uri)

	if err != nil {
		log.Fatal(err)
	}
	return db
}
