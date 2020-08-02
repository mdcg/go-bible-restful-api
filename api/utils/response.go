package utils

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	status string
	data   interface{}
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	parsed_data, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(parsed_data))
}

func SuccessResponse(payload interface{}) interface{} {
	return map[string]interface{}{
		"status": "success",
		"data":   payload,
	}
}

func FailResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "fail",
		"message": message,
	}
}
