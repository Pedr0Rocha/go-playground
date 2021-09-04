package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	VERSION = "v1.0.0"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK!"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server running status"))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Running, version " + VERSION))
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Handler)

	subrouter := router.PathPrefix("/api").Subrouter()

	subrouter.HandleFunc("/health", healthHandler)
	subrouter.HandleFunc("/version", versionHandler)

	http.ListenAndServe(":8755", router)
}
