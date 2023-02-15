package models

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

type TransferRequest struct {
	ToAccount int64 `json:"toAccount"`
	Amount    int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Username          string    `json:"username"`
	EncryptedPassword string    `json:"-"`
	AccountNumber     int64     `json:"accountNumber"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName, username, password string) (*Account, error) {
	// constructor
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		// ID:            rand.Intn(1000000),
		FirstName:         firstName,
		LastName:          lastName,
		Username:          username,
		EncryptedPassword: string(encPass),
		AccountNumber:     int64(rand.Intn(100000000)),
		CreatedAt:         time.Now().Local().UTC(),
	}, nil
}
