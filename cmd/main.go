package main

import (
	"cms-config/internal/config"
	"cms-config/internal/delivery/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/firebaseremoteconfig/v1"
)

func main() {
	app := fiber.New()
	viper := config.NewViper()
	webPort := viper.GetString("WEB_PORT")

	// Setup routes
	http.SetupRemoteConfigRoutes(app, &firebaseremoteconfig.Service{})

	// Start the server
	err := app.Listen(fmt.Sprintf(":%s", webPort))
	if err != nil {
		panic(err)
	}
}
