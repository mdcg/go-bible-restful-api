package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllBooks() (*[]db.Books, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var books []db.Books
	was_found := true

	if conn.Model(&db.Books{}).Find(&books).RecordNotFound() {
		was_found = false
	}
	return &books, was_found
}

func FindBookByAbbrev(abbr string) (*db.Books, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var book db.Books
	was_found := true

	if conn.Model(&db.Books{}).Where("abbrev = ?", abbr).First(&book).RecordNotFound() {
		was_found = false
	}
	return &book, was_found
}
