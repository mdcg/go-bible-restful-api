package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVersions() (*[]db.Versions, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var versions []db.Versions
	not_found := false

	if conn.Model(&db.Versions{}).Find(&versions).RecordNotFound() {
		not_found = true
	}
	return &versions, not_found
}

func FindVersionByAbbrev(abbr string) (*db.Versions, bool) {
	conn := db.GetDB()
	defer conn.Close()

	var version db.Versions
	not_found := false

	if conn.Model(&db.Versions{}).Where("abbrev = ?", abbr).First(&version).RecordNotFound() {
		not_found = true
	}
	return &version, not_found
}
