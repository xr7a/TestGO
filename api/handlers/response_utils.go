package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Printf("Failed to encode JSON response: %v\n", err)
		}
	}
}

func RespondError(w http.ResponseWriter, status int, message string, internalErr error) {
	errorRes := ErrorResponse{
		Error:   internalErr.Error(),
		Message: message,
	}
	RespondJSON(w, status, errorRes)
}
