package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	// if r.Method == "GET" {
	// 	return s.handleGetAccount(w, r)
	// }
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("%v request method not allowed", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	// account := models.NewAccount("Ishan", "Shrestha")
	return writeJSON(w, http.StatusOK, id)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil

}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
