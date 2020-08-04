package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

// Search and return all the books of the bible.
func FindAllBooks() (*[]db.Books, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var books []db.Books
	not_found := false

	if conn.Model(&db.Books{}).Find(&books).RecordNotFound() {
		not_found = true
	}
	return &books, not_found
}

// Searches and returns the books of the bible of a given testament.
func FindBooksByTestament(part int) (*[]db.Books, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var books []db.Books
	not_found := false

	if conn.Model(&db.Books{}).Where("testament = ?", part).Find(&books).RecordNotFound() {
		not_found = true
	}
	return &books, not_found
}

// Searches and returns the data for a particular book of the Bible through its abbreviation.
func FindBookByAbbrev(abbr string) (*db.Books, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var book db.Books
	not_found := false

	if conn.Model(&db.Books{}).Where("abbrev = ?", abbr).First(&book).RecordNotFound() {
		not_found = true
	}
	return &book, not_found
}
