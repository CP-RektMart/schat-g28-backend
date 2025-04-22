package server

import (
	"github.com/CP-RektMart/schat-g28-backend/doc"
	"github.com/gofiber/fiber/v2"
	// fiberSwagger "github.com/gofiber/swagger"
	"github.com/swaggo/swag"
	"github.com/yokeTH/go-pkg/scalar"
)

func (s *Server) RegisterDocs() {
	swag.Register(doc.SwaggerInfo.InfoInstanceName, doc.SwaggerInfo)
	s.app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})
	s.app.Get("/swagger/*", scalar.Handler("./doc/swagger.json"))
	s.app.Get("/openapi", func(c *fiber.Ctx) error {
		return c.SendFile("doc/swagger.yaml")
	})
}
