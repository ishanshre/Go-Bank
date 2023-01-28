package models

import "math/rand"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	AccountNumber int64  `json:"accountNumber"`
	Balance       int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	// constructor
	return &Account{
		ID:            rand.Intn(1000000),
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(100000000)),
	}
}
