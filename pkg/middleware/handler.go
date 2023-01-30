package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-Bank/pkg/models"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	// if r.Method == "GET" {
	// 	return s.handleGetAccount(w, r)
	// }
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("%v request method not allowed", r.Method)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("Error parsing the id to int: %v", err)
		return err
	}
	account, err := s.store.GetAccountById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createNewAccount := new(models.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createNewAccount); err != nil {
		return nil
	}
	account := models.NewAccount(createNewAccount.FirstName, createNewAccount.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
