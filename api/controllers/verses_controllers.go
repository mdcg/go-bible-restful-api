package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVerses() (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	was_found := true

	if conn.Model(&db.Verses{}).Limit(100).Find(&verses).RecordNotFound() {
		was_found = false
	}
	return &verses, was_found
}
