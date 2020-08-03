package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllTestaments() (*[]db.Testament, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var testaments []db.Testament
	was_found := true

	if conn.Model(&db.Testament{}).Find(&testaments).RecordNotFound() {
		was_found = false
	}
	return &testaments, was_found
}

func FindTestamentByPart(part string) (*db.Testament, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var testament db.Testament
	was_found := true

	if conn.Model(&db.Testament{}).Where("id = ?", part).Find(&testament).RecordNotFound() {
		was_found = false
	}
	return &testament, was_found
}
