package main

import "github.com/ishanshre/Go-Bank/pkg/middleware"

func main() {
	server := middleware.NewApiServer(":8000")
	server.Run()
}
