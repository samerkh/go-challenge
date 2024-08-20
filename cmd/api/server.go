package main

import (
	"fmt"
	"go_challenge/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
