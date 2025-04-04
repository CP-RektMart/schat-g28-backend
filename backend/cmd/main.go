package main

import (
	"github.com/CP-RektMart/computer-network-g28/backend/internal/config"
	db "github.com/CP-RektMart/computer-network-g28/backend/internal/db"
	"github.com/CP-RektMart/computer-network-g28/backend/internal/server"
)

func main() {
	config := config.Load()
	store := db.New(db.Config{
		PostgresURL: config.PostgresURL,
	})
	server := server.New(server.Config{
		ServerAddr: config.ServerAddr,
	}, store)

	server.Start()
}