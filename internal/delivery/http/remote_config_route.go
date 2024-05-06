package http

import (
    "github.com/gofiber/fiber/v2"
    "cms-config/internal/app/controller"
    "cms-config/internal/app/usecase"
    "cms-config/internal/app/repository"
    "google.golang.org/api/firebaseremoteconfig/v1"
)

func SetupRemoteConfigRoutes(app *fiber.App, firebaseClient *firebaseremoteconfig.Service) {
    repo := repository.NewFirebaseRepository(firebaseClient)
    uc := usecase.NewUsecase(repo)
    controller := controller.NewRemoteConfigController(uc)

    app.Get("/api/v1/remote-config", controller.GetTemplate)
}
