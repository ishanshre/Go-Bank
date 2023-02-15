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

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		validate, err := validateLogin(req, s.store)
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, validate)
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}

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
	if r.Method == "GET" {
		id, err := getId(r)
		if err != nil {
			return err

		}
		account, err := s.store.GetAccountById(id)
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, account)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createNewAccount := new(models.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createNewAccount); err != nil {
		return nil
	}
	account, err := models.NewAccount(createNewAccount.FirstName, createNewAccount.LastName, createNewAccount.Username, createNewAccount.Password)
	if err != nil {
		return err
	}
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	tokenString, err := createJWT(account)
	if err != nil {
		return err
	}
	fmt.Println("JWT Token: ", tokenString)
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, map[string]int{"deleted_id": id})
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}
	accountUpdated := new(models.Account)
	if err := json.NewDecoder(r.Body).Decode(&accountUpdated); err != nil {
		return nil
	}
	if err := s.store.UpdateAccount(id, accountUpdated); err != nil {
		return err
	}
	log.Printf("Account with id %v successfully updated", id)
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "account updated successfull"})

}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		transferRequest := new(models.TransferRequest)
		if err := json.NewDecoder(r.Body).Decode(&transferRequest); err != nil {
			return err
		}
		defer r.Body.Close()
		return writeJSON(w, http.StatusOK, transferRequest)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func getId(r *http.Request) (int, error) {
	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return id, fmt.Errorf("invalid id or pasing id error: %s", idstr)
	}
	return id, nil
}
