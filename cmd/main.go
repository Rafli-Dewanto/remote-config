package main

import (
	"github.com/gofiber/fiber/v2"
	"cms-config/internal/delivery/http"
	"google.golang.org/api/firebaseremoteconfig/v1"
)

func main() {
	app := fiber.New()

	// Setup routes
	http.SetupRemoteConfigRoutes(app, &firebaseremoteconfig.Service{})

	// Start the server
	err := app.Listen(":3002")
	if err != nil {
		panic(err)
	}
}
