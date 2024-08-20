package handlers

import (
	"go_challenge/api"

	"github.com/go-chi/chi"
)

var users []api.User = []api.User{}

func Handler(r *chi.Mux) {

	r.Route("/users", func(router chi.Router) {
		router.Get("/", GetUsers)
		router.Get("/{id}", GetUser)
		router.Post("/", CreateUser)
		router.Put("/{id}", UpdateUser)
		router.Delete("/{id}", DeleteUser)
	})

}
