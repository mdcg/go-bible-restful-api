package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WasAnyDataFound(data interface{}, _type interface{}) bool {
	// Não está funcionando -- Investigar
	return reflect.DeepEqual(data, _type)
}
