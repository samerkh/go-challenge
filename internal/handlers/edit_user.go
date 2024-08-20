package handlers

import (
	"encoding/json"
	"fmt"
	"go_challenge/api"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for i, user := range users {
		if id == fmt.Sprint(user.ID) {
			var user api.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Error(err)
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			users[i] = user

			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Error(err)
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			return
		}
	}

	log.Error("User not found")
	http.Error(w, "User not found", http.StatusNotFound)
}
