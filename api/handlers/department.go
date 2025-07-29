package handlers

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleDepartmentRoutes(router *mux.Router, service services.DepartmentService) {
	router.HandleFunc("", CreateDepartment(service)).Methods("POST")
}

func CreateDepartment(service services.DepartmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Department
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			RespondError(w, http.StatusBadRequest, "error decoding request body", err)
		}
		defer r.Body.Close()

		err = service.CreateDepartment(r.Context(), req)
		if err != nil {
			RespondError(w, http.StatusInternalServerError, "error creating department", err)
		}

		RespondJSON(w, http.StatusCreated, nil)
	}
}
