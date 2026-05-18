package api

import (
	handlers "GolangProjects/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
}
