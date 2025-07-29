package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleUserRoutes(router *mux.Router, service services.UserService) {
	router.HandleFunc("", CreateUser(service)).Methods(http.MethodPost)
}

func CreateUser(service services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.UserPassport
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			RespondError(w, http.StatusBadRequest, "error decoding request body", err)
			return
		}
		defer r.Body.Close()

		user, err := service.CreateUser(r.Context(), req)
		if err != nil {
			RespondError(w, http.StatusBadRequest, "Error creating user", err)
			return
		}

		RespondJSON(w, http.StatusOK, user)
	}
}
