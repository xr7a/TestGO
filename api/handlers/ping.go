package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandlePingRoutes(router *mux.Router) {
	router.HandleFunc("", Ping).Methods(http.MethodGet)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
