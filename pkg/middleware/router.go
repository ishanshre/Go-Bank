package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/Go-Bank/pkg/storage"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	// convert into JSON as we are making JSON api
	w.Header().Set("Content-Type", "application/json") // header attributes must be set before w.WriteHeader
	w.WriteHeader(status)                              // set the status of the response
	// set the format of the response
	return json.NewEncoder(w).Encode(v) // Response writer encoded into json and retured as a response
}

type apiFunc func(http.ResponseWriter, *http.Request) error // our handle func signature we are using to handle our api

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"success"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	// this func decorate all the router.handler functions we create into HandleFunc
	// it is a wrapper function that wraps our handler into http.HandleFunc()
	return func(w http.ResponseWriter, r *http.Request) {
		// handles the error from our router handle func and return it as a http.HandleFunc
		if err := f(w, r); err != nil {
			//handle the error by responding the request with error with datatype string
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount)) // wrap the function
	router.HandleFunc("/account/{id}", withJWTAuth(makeHTTPHandleFunc(s.handleGetAccountById), s.store))
	router.HandleFunc("/transfer/{id}", makeHTTPHandleFunc(s.handleTransfer))
	log.Println("JSON API web server starting at port number: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func NewApiServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}
