package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for _, user := range users {
		if id == fmt.Sprint(user.ID) {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				log.Error(err)
				return
			}
			return
		}
	}

	log.Error("User not found")
	http.Error(w, "User not found", http.StatusNotFound)
}
