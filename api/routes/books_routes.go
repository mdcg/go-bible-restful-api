package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books, was_found := controllers.FindAllBooks()
	if was_found {
		response := utils.SuccessResponse(map[string]interface{}{"books": books})
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.FailResponse("Nenhum livro foi encontrado.")
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindBookByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	book, was_found := controllers.FindBookByAbbrev(abbrev)
	if was_found {
		response := utils.SuccessResponse(map[string]interface{}{"book": book})
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.FailResponse("Nenhum livro foi encontrado.")
	utils.JSONResponse(w, http.StatusFound, response)
}
