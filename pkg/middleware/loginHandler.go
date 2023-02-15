package middleware

import (
	"github.com/ishanshre/Go-Bank/pkg/models"
	"github.com/ishanshre/Go-Bank/pkg/storage"
	"golang.org/x/crypto/bcrypt"
)

func validateLogin(req models.LoginRequest, s storage.Storage) (*models.LoginResponse, error) {
	check, err := s.GetAccountByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(check.EncryptedPassword), []byte(req.Password)); err != nil {
		return nil, err
	}
	tokenString, err := createJWT(check)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{
		Username:    check.Username,
		AccessToken: tokenString,
	}, nil
}
