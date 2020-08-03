package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVerses() (*[]db.Verses, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	not_found := false

	if conn.Model(&db.Verses{}).Limit(100).Find(&verses).RecordNotFound() {
		not_found = true
	}
	return &verses, not_found
}
