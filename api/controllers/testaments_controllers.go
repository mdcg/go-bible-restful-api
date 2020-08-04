package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

// Search and return all testaments from the bible.
func FindAllTestaments() (*[]db.Testament, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var testaments []db.Testament
	not_found := false

	if conn.Model(&db.Testament{}).Find(&testaments).RecordNotFound() {
		not_found = true
	}
	return &testaments, not_found
}

// Search and return the data for a specific bible testament from your part.
func FindTestamentByPart(part string) (*db.Testament, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var testament db.Testament
	not_found := false

	if conn.Model(&db.Testament{}).Where("id = ?", part).Find(&testament).RecordNotFound() {
		not_found = true
	}
	return &testament, not_found
}
