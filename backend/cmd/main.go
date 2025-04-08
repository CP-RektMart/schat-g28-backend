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
		Host:     config.Postgres.Host,
		User:     config.Postgres.User,
		Password: config.Postgres.Password,
		DBName:   config.Postgres.DBName,
		Port:     config.Postgres.Port,
		SSLMode:  config.Postgres.SSLMode,
	})

	server := server.New(server.Config{
		Name:         config.Server.Name,
		Port:         config.Server.Port,
		MaxBodyLimit: config.Server.MaxBodyLimit,
	}, store)

	server.Start(ctx, stop)
}
