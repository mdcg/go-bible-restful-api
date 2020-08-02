package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVersions(w http.ResponseWriter, r *http.Request) {
	versions, was_found := controllers.FindAllVersions()
	if was_found {
		response := utils.SuccessResponse(map[string]interface{}{"versions": versions})
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.FailResponse("Nenhuma versão foi encontrada.")
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindVersionByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	version, was_found := controllers.FindVersionByAbbrev(abbrev)
	if was_found {
		response := utils.SuccessResponse(map[string]interface{}{"version": version})
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.FailResponse("Nenhuma versão foi encontrada.")
	utils.JSONResponse(w, http.StatusFound, response)
}
