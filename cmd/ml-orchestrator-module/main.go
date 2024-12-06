package main

import (
	"github.com/gorilla/mux"
	"log"
	"ml-orchestrator-module/internal/config"
	"ml-orchestrator-module/internal/handlers"
	"ml-orchestrator-module/internal/middleware"
	"net/http"
)

var cfg *config.Config

func main() {

	cfg := config.MustLoad()

	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(next, cfg)
	})

	r.HandleFunc("/resume/beta", func(w http.ResponseWriter, r *http.Request) {
		handlers.ResumeDaddy(cfg, w, r)
	}).Methods(http.MethodPost)

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
