package server

import (
	"context"
	"fmt"
	"log"

	database "github.com/CP-RektMart/computer-network-g28/backend/internal/db"
	"github.com/CP-RektMart/computer-network-g28/backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Name         string `env:"NAME"`
	Port         int    `env:"PORT"`
	MaxBodyLimit int    `env:"MAX_BODY_LIMIT"`
}

type CorsConfig struct {
	AllowedOrigins   string `env:"ALLOWED_ORIGINS"`
	AllowedMethods   string `env:"ALLOWED_METHODS"`
	AllowedHeaders   string `env:"ALLOWED_HEADERS"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS"`
}

type Server struct {
	config Config
	App    *fiber.App
	DB     *database.Store
}

func New(config Config, DB *database.Store) *Server {
	return &Server{
		config: config,
		App:    fiber.New(),
		DB:     DB,
	}
}

func (s *Server) Start(ctx context.Context, stop context.CancelFunc) {
	s.App.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(dto.HttpResponse[string]{
			Result: "ok",
		})
	})
	go func() {
		if err := s.App.Listen(fmt.Sprintf("localhost:%d", s.config.Port)); err != nil {
			log.Fatalf("failed to start server: %v", err)
			stop()
		}
	}()

	defer func() {
		if err := s.App.Shutdown(); err != nil {
			log.Printf("failed to shutdown server: %v.", err)
		}
	}()

	<-ctx.Done()

	log.Println("shutting down server...")
}
