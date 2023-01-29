package main

import (
	"log"

	"github.com/ishanshre/Go-Bank/pkg/middleware"
	"github.com/ishanshre/Go-Bank/pkg/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatalf("Error in main.go line 13: %v", err)
	}
	if err := store.Init(); err != nil {
		log.Fatalf("Error in line 16: %v", err)
	}
	server := middleware.NewApiServer(":8000", store)
	server.Run()
}
