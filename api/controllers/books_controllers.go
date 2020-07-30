package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllBooks() *[]db.Books {
	conn := db.GetDB()
	defer conn.Close()

	var books []db.Books
	conn.Model(&db.Books{}).Find(&books)
	return &books
}
