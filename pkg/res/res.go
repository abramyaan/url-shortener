package res

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Если передана строка, упаковываем её в ErrorResponse
	if message, ok := data.(string); ok {
		data = ErrorResponse{Message: message}
	}

	json.NewEncoder(w).Encode(data)
}