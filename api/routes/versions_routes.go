package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/db"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVersions(w http.ResponseWriter, r *http.Request) {
	versions := controllers.FindAllVersions()
	utils.JSONResponse(w, http.StatusFound, versions)
}

func FindVersionByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	version := controllers.FindVersionByAbbrev(abbrev)

	if utils.WasAnyDataFound(version, db.Versions{}) {
		utils.JSONResponse(w, http.StatusFound, version)
	}
	utils.JSONResponse(w, http.StatusNotFound, map[string]string{})
}
