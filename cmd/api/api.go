package main

import (
	"encoding/json"
	"fmt"
	"net"

	"net/http"
	"time"

	"github.com/AxolotlHere/hackd/internal/env"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
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

var store = sessions.NewCookieStore([]byte(env.GetSessionCreds()))

func (application_instance *app) mount() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Second * 10))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Route("/v1", func(router chi.Router) {

		router.Get("/health", checkHealth)
		router.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("called")
			email := r.URL.Query().Get("email")
			passwd := r.URL.Query().Get("passwd")
			if email == "" || passwd == "" {
				http.Error(w, "Missing parameters", http.StatusBadRequest)
				return
			}
			if validateUser(email, passwd) != true {
				http.Error(w, "Invalid username/password", http.StatusBadRequest)
				return
			}
			session, err := store.Get(r, "session-auth")
			if err != nil {
				http.Error(w, "Session Error", http.StatusInternalServerError)
			}
			session.Values["email"] = email
			session.Values["authenticated"] = true
			session.Save(r, w)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		})
		router.Get("/getUserData", func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "session-auth")
			auth, ok := session.Values["authenticated"].(bool)
			email := r.URL.Query().Get("email")
			fmt.Println("[DEBUG] Session Data:", session.Values)
			fmt.Println("[DEBUG] Auth OK:", ok, "Auth Bool:", auth)
			if !ok || !auth {
				http.Error(w, "Forbidden request", http.StatusForbidden)
				return
			}

			user_creds := getUserData(email)
			fmt.Println(user_creds)
			if user_creds.Username != "" {
				http.Error(w, "Bad status: User not found", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(user_creds)
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
