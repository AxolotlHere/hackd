package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type app struct {
	server_config config
}

type config struct {
	Addr string
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("STATUS: 200 OK"))
}

func (application_instance *app) mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 10))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Route("/v1", func(router chi.Router) {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
		router.Get("/health", checkHealth)
		router.Get("/getUser", func(w http.ResponseWriter, r *http.Request) {
			email := r.URL.Query().Get("email")
			if email == "" {
				http.Error(w, "missing parameter", http.StatusBadRequest)
				return
			}
			user := getUserData(email)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		})

	})
	return router
}

func (application_instance *app) initialiseServer(mux http.Handler) error {
	server := &http.Server{
		Addr:         application_instance.server_config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}
	return server.ListenAndServe()
}
