package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/altinthaqi/jot-fusion/db"
	"github.com/gorilla/mux"
)

type (
	apiFunc func(http.ResponseWriter, *http.Request) error

	ApiError struct {
		Error string
	}

	APIServer struct {
		listenAddr string
		store      db.Storage
	}
)

func NewAPIServer(listenAddr string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccountByID))

	log.Print("JSON API server running on port: ", s.listenAddr, "\n")

	http.ListenAndServe(s.listenAddr, router)
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
