package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		response := utils.FailResponse("Nenhum versículo foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindVersesByChapterVersionAndBook(w http.ResponseWriter, r *http.Request) {
	version_abbrev := mux.Vars(r)["version_abbrev"]
	book_abbrev := mux.Vars(r)["book_abbrev"]
	chapter := mux.Vars(r)["chapter"]

	converted_chapter, err := strconv.Atoi(chapter)
	if err != nil {
		response := utils.FailResponse("Capítulo inválido.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

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

	verses, not_found := controllers.FindVersesByChapterVersionAndBook(converted_chapter, version.Abbrev, book.Id)
	if not_found {
		response := utils.FailResponse("Nenhum versículo foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
	utils.JSONResponse(w, http.StatusFound, response)
}

func FindVerseByChapterVersionAndBook(w http.ResponseWriter, r *http.Request) {
	version_abbrev := mux.Vars(r)["version_abbrev"]
	book_abbrev := mux.Vars(r)["book_abbrev"]
	chapter := mux.Vars(r)["chapter"]
	verse := mux.Vars(r)["verse"]

	converted_chapter, err := strconv.Atoi(chapter)
	if err != nil {
		response := utils.FailResponse("Capítulo inválido.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

	converted_verse, err := strconv.Atoi(verse)
	if err != nil {
		response := utils.FailResponse("Versículo inválido.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

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

	verses, not_found := controllers.FindVerseByChapterVersionAndBook(converted_verse, converted_chapter, version.Abbrev, book.Id)
	if not_found {
		response := utils.FailResponse("Nenhum versículo foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
	utils.JSONResponse(w, http.StatusFound, response)
}

type searchInfo struct {
	Phrase string
}

func SearchByVerseText(w http.ResponseWriter, r *http.Request) {
	version_abbrev := mux.Vars(r)["version_abbrev"]
	var searchObj searchInfo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&searchObj); err != nil {
		response := utils.FailResponse("Utilize o parâmetro 'phrase' para realizar a pesquisa.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	defer r.Body.Close()

	version, not_found := controllers.FindVersionByAbbrev(version_abbrev)
	if not_found {
		response := utils.FailResponse("Nenhuma versão encontrada.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}

	verses, not_found := controllers.SearchByVerseText(version.Abbrev, searchObj.Phrase)
	if not_found {
		response := utils.FailResponse("Nenhum versículo foi encontrado.")
		utils.JSONResponse(w, http.StatusFound, response)
		return
	}
	response := utils.SuccessResponse(map[string]interface{}{"verses": verses})
	utils.JSONResponse(w, http.StatusFound, response)
}
