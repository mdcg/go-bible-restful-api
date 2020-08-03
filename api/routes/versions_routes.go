package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVersions(w http.ResponseWriter, r *http.Request) {
	versions, not_found := controllers.FindAllVersions()
	if not_found {
		response := utils.FailResponse("Nenhuma versão foi encontrada.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"versions": versions})
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindVersionByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	version, not_found := controllers.FindVersionByAbbrev(abbrev)
	if not_found {
		response := utils.FailResponse("Nenhuma versão foi encontrada.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"version": version})
	utils.JSONResponse(w, http.StatusFound, response)
}
