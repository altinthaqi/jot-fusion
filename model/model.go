package model

import (
	"math/rand"
	"time"
)

type (
	CreateAccountRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	Account struct {
		ID        int       `json:"id"`
		FitstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Number    int64     `json:"number"`
		Balance   int64     `json:"balance"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FitstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Int63n(10000)),
		CreatedAt: time.Now().UTC(),
	}
}
