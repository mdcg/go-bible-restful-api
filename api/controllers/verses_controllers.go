package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindVersesByVersionAndBook(version string, book_id int) (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	not_found := false

	if conn.Model(&db.Verses{}).Where("version = ? AND book = ?", version, book_id).Find(&verses).RecordNotFound() {
		not_found = true
	}
	return &verses, not_found
}
