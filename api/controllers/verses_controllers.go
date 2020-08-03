package controllers

import (
	"fmt"

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

func FindVersesByChapterVersionAndBook(chapter int, version string, book_id int) (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	not_found := false

	if conn.Model(&db.Verses{}).Where("chapter = ? AND version = ? AND book = ?", chapter, version, book_id).Find(&verses).RecordNotFound() {
		not_found = true
	}
	return &verses, not_found
}

func FindVerseByChapterVersionAndBook(verse int, chapter int, version string, book_id int) (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	not_found := false

	if conn.Model(&db.Verses{}).Where("verse = ? AND chapter = ? AND version = ? AND book = ?", verse, chapter, version, book_id).Find(&verses).RecordNotFound() {
		not_found = true
	}
	return &verses, not_found
}

func SearchByVerseText(version string, word string) (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	not_found := false

	word_to_be_searched := fmt.Sprintf("%%%v%%", word)
	if conn.Model(&db.Verses{}).Where("version = ? AND text LIKE ?", version, word_to_be_searched).Find(&verses).RecordNotFound() {
		not_found = true
	}
	return &verses, not_found
}
