package model

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type (
	Account struct {
		ID                int       `json:"id"`
		FirstName         string    `json:"firstName"`
		LastName          string    `json:"lastName"`
		Number            int64     `json:"number"`
		Balance           int64     `json:"balance"`
		EncryptedPassword string    `json:"-"`
		CreatedAt         time.Time `json:"createdAt"`
	}

	LoginRequest struct {
		Number   int64  `json:"number"`
		Password string `json:"password"`
	}

	TransferRequest struct {
		ToAccount int `json:"toAccount"`
		Amount    int `json:"amount"`
	}

	CreateAccountRequest struct {
		FirstName string `json:"firstName" validate:"required|min_len:1"`
		LastName  string `json:"lastName" validate:"required|min_len:1"`
		Password  string `json:"password"`
	}
)

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		EncryptedPassword: string(encPw),
		Number:            int64(rand.Int63n(10000)),
		CreatedAt:         time.Now().UTC(),
	}, nil
}
