package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVersions() (*[]db.Versions, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var versions []db.Versions
	was_found := true

	if conn.Model(&db.Versions{}).Find(&versions).RecordNotFound() {
		was_found = false
	}
	return &versions, was_found
}

func FindVersionByAbbrev(abbr string) (*db.Versions, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var version db.Versions
	was_found := true

	if conn.Model(&db.Versions{}).Where("abbrev = ?", abbr).First(&version).RecordNotFound() {
		was_found = false
	}
	return &version, was_found
}
