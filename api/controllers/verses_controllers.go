package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVerses() *[]db.Verses {
	conn := db.GetDB()
	defer conn.Close()

	var verses []db.Verses
	conn.Model(&db.Verses{}).Limit(100).Find(&verses)
	return &verses
}
