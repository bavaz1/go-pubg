package main

import (
	"net/http"

	"github.com/bavaz1/go-pubg/database"
	"github.com/bavaz1/go-pubg/server"
)

func main() {
	storage := database.New()

	server := server.Server{
		&http.Client{},
		":8080",
		storage,
	}

	server.Listen()
}
