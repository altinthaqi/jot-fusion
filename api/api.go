package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/altinthaqi/jot-fusion/db"
	"github.com/altinthaqi/jot-fusion/helper"
	"github.com/altinthaqi/jot-fusion/model"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type (
	apiFunc func(http.ResponseWriter, *http.Request) error

	ApiError struct {
		Error string `json:"error"`
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

	router.HandleFunc("/login", makeHTTPHandlerFunc(s.handleLogin))
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandlerFunc(s.handleAccountByID), s.store))

	router.HandleFunc("/transfer", makeHTTPHandlerFunc(s.handleTransfer))

	log.Print("JSON API server running on port: ", s.listenAddr, "\n")

	http.ListenAndServe(s.listenAddr, router)
}

func withJWTAuth(handlerFunc http.HandlerFunc, s db.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("jwt middleware")

		token, err := validateJWT(r.Header.Get("x-jwt-token"))

		if err != nil || !token.Valid {
			permissionDenied(w)
			return
		}

		id, err := helper.GetQueryInt(r, "id")
		if err != nil {
			permissionDenied(w)
			return
		}

		account, err := s.GetAccountByID(id)
		if err != nil {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if account.Number != int64(claims["accountNumber"].(float64)) {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, ApiError{Error: "permission denied"})
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil

	})

}

func createJWT(account *model.Account) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     15000,
		"accountNumber": account.Number,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
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
