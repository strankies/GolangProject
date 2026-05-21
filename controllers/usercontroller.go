package handlers

import (
	"GolangProjects/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var users = []models.User{
	{ID: "1", Username: "user1", Email: "user1@gmail.com"},
	{ID: "2", Username: "user2", Email: "user2@gmail.com"},
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
}
