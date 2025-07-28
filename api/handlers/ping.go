package handlers

import (
	"awesomeProject/services"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func PingConfig(service services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service.GetConfig()
	}
}
