package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllTestaments() *[]db.Testament {
	conn := db.GetDB()
	defer conn.Close()

	var testaments []db.Testament
	conn.Model(&db.Testament{}).Find(&testaments)
	return &testaments
}
