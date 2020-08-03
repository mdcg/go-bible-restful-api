package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllTestaments(w http.ResponseWriter, r *http.Request) {
	testaments, not_found := controllers.FindAllTestaments()
	if not_found {
		response := utils.FailResponse("Nenhum testamento foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"testaments": testaments})
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindTestamentsByPart(w http.ResponseWriter, r *http.Request) {
	part := mux.Vars(r)["part"]
	testament, not_found := controllers.FindTestamentByPart(part)
	if not_found {
		response := utils.FailResponse("Nenhum testamento foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"testament": testament})
	utils.JSONResponse(w, http.StatusFound, response)
}
