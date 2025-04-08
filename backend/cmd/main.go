package main

import (
	"context"

	"github.com/CP-RektMart/computer-network-g28/backend/internal/config"
	db "github.com/CP-RektMart/computer-network-g28/backend/internal/db"
	"github.com/CP-RektMart/computer-network-g28/backend/internal/server"
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	config := config.Load()

	store := db.New(db.Config{
		PgConfig: config.Postgres,
	})
	server := server.New(server.Config{
		Name:         config.Server.Name,
		Port:         config.Server.Port,
		MaxBodyLimit: config.Server.MaxBodyLimit,
	}, store)

	server.Start(ctx, stop)
}
