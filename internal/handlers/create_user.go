package handlers

import (
	"encoding/json"
	"go_challenge/api"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user api.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Error(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	users = append(users, user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Error(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
}
