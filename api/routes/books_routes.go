package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

// Route to search all the books available in the bible.
func FindAllBooks(w http.ResponseWriter, r *http.Request) {
	books, not_found := controllers.FindAllBooks()
	if not_found {
		response := utils.FailResponse("Nenhum livro foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"books": books})
	utils.JSONResponse(w, http.StatusFound, response)
}

// Route to list the books of a particular testament of the bible.
func FindBooksByTestament(w http.ResponseWriter, r *http.Request) {
	part := mux.Vars(r)["part"]
	testament, not_found := controllers.FindTestamentByPart(part)
	if not_found {
		response := utils.FailResponse("Nenhum testamento foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

	books, not_found := controllers.FindBooksByTestament(testament.Id)
	if not_found {
		response := utils.FailResponse("Nenhum livro foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"books": books})
	utils.JSONResponse(w, http.StatusFound, response)
}

// Route to search the data for a particular book from its abbreviation.
func FindBookByAbbrev(w http.ResponseWriter, r *http.Request) {
	abbrev := mux.Vars(r)["abbrev"]
	book, not_found := controllers.FindBookByAbbrev(abbrev)
	if not_found {
		response := utils.FailResponse("Nenhum livro foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"book": book})
	utils.JSONResponse(w, http.StatusFound, response)
}
