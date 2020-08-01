package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVersions(w http.ResponseWriter, r *http.Request) {
	versions := controllers.FindAllVersions()
	utils.JSONResponse(w, 200, versions)
}

func FindVersionByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	versions := controllers.FindVersionByAbbrev(abbrev)
	utils.JSONResponse(w, 200, versions)
}
