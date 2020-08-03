package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdcg/go-bible-restful-api/api/utils"

	"github.com/mdcg/go-bible-restful-api/api/controllers"
)

func FindVersesByVersionAndBook(w http.ResponseWriter, r *http.Request) {
	version_abbrev := mux.Vars(r)["version_abbrev"]
	book_abbrev := mux.Vars(r)["book_abbrev"]

	version, not_found := controllers.FindVersionByAbbrev(version_abbrev)
	if not_found {
		response := utils.FailResponse("Nenhuma versão encontrada.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

	book, not_found := controllers.FindBookByAbbrev(book_abbrev)
	if not_found {
		response := utils.FailResponse("Nenhum livro foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

	verses, not_found := controllers.FindVersesByVersionAndBook(version.Abbrev, book.Id)
	if not_found {
		response := utils.FailResponse("Nenhum verso foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
	utils.JSONResponse(w, http.StatusFound, response)
}
