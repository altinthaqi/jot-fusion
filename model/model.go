package model

import (
	"math/rand"
	"time"
)

type (
	Account struct {
		ID        int       `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Number    int64     `json:"number"`
		Balance   int64     `json:"balance"`
		CreatedAt time.Time `json:"createdAt"`
	}

	CreateAccountRequest struct {
		FirstName string `json:"firstName" validate:"required|min_len:1"`
		LastName  string `json:"lastName" validate:"required|min_len:1"`
	}

	DeleteAccountRequest struct {
		ID int `json:"id" validate:"required|min:1"`
	}
)

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Int63n(10000)),
		CreatedAt: time.Now().UTC(),
	}
}
