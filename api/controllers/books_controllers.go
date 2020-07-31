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

func FindBookByAbbrev(abbr string) *db.Books {
	conn := db.GetDB()
	defer conn.Close()

	var book db.Books
	conn.Model(&db.Books{}).Where("abbrev = ?", abbr).First(&book)
	return &book
}
