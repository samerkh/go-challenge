package handlers

import (
	"encoding/json"
	"fmt"
	"go_challenge/api"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	userChan := make(chan api.User)
	errChan := make(chan error)

	go fetchUser(id, userChan, errChan)

	select {
	case user := <-userChan:
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Printf("error encoding user: %v", err)
			http.Error(w, "Error encoding user", http.StatusInternalServerError)
			return
		}
	case err := <-errChan:
		log.Printf("error fetching user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

// represents a long running operation
func fetchUser(id string, userChan chan<- api.User, errChan chan<- error) {
	for _, user := range users {
		if id == fmt.Sprint(user.ID) {
			userChan <- user
			return
		}
	}
	errChan <- fmt.Errorf("user not found")
}
