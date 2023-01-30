package models

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int64 `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	AccountNumber int64     `json:"accountNumber"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	// constructor
	return &Account{
		// ID:            rand.Intn(1000000),
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(100000000)),
		CreatedAt:     time.Now().Local().UTC(),
	}
}
