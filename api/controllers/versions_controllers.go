package controllers

import (
	"github.com/mdcg/go-bible-restful-api/api/db"
)

func FindAllVersions() *[]db.Versions {
	conn := db.GetDB()
	defer conn.Close()

	var versions []db.Versions
	conn.Model(&db.Versions{}).Find(&versions)
	return &versions
}
