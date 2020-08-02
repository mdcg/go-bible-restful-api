package routes

import (
	"net/http"

	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllVerses(w http.ResponseWriter, r *http.Request) {
	verses, was_found := controllers.FindAllVerses()
	if was_found {
		response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.FailResponse("Nenhum verso foi encontrado.")
	utils.JSONResponse(w, http.StatusFound, response)
}
