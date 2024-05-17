package http

import (
	"cms-config/internal/app/controller"
	"cms-config/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/firebaseremoteconfig/v1"
	"net/http"
)

func SetupRemoteConfigRoutes(app *fiber.App, firebaseClient *firebaseremoteconfig.Service) {
    useCase := usecase.NewUsecase(firebaseClient)
    remoteConfigController := controller.NewRemoteConfigController(useCase, &http.Client{})

    app.Get("/api/v1/remote-config", remoteConfigController.GetTemplate)
    app.Get("/api/v1/remote-config/:param_name", remoteConfigController.GetDefaultValue)
}
