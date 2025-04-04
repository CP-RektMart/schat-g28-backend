package main

import (
	"github.com/CP-RektMart/computer-network-g28/internal/config"
	"github.com/CP-RektMart/computer-network-g28/internal/database"
	"github.com/CP-RektMart/computer-network-g28/internal/server"
)

func main() {
	config := config.Load()
	store := database.New(database.Config{
		PostgresURL: config.PostgresURL,
	})
	server := server.New(server.Config{
		ServerAddr: config.ServerAddr,
	}, store)

	server.Start()
}