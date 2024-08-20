package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID   int
	Name string
}

var users []User = []User{}

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	Handler(r)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func Handler(r *chi.Mux) {

	r.Route("/users", func(router chi.Router) {
		router.Get("/", GetUsers)
		router.Get("/{id}", GetUser)
		router.Post("/", CreateUser)
		router.Put("/{id}", UpdateUser)
		router.Delete("/{id}", DeleteUser)
	})

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Error(err)
		http.Error(w, "User not found", http.StatusNotFound)

		return
	}
}

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

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user User
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for i, user := range users {
		if id == fmt.Sprint(user.ID) {
			var user User
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for i, user := range users {
		if id == fmt.Sprint(user.ID) {
			users = append(users[:i], users[i+1:]...)

			err := json.NewEncoder(w).Encode(user)
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
