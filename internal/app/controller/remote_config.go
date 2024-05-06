package controller

import (
    "github.com/gofiber/fiber/v2"
    "cms-config/internal/app/usecase"
)

type RemoteConfigController struct {
    usecase usecase.RemoteConfigUsecase
}

func NewRemoteConfigController(usecase usecase.RemoteConfigUsecase) *RemoteConfigController {
    return &RemoteConfigController{
        usecase: usecase,
    }
}

func (rc *RemoteConfigController) GetTemplate(c *fiber.Ctx) error {
    // Call the use case to get the template
    template, err := rc.usecase.GetTemplate()
    if err != nil {
        // Return an error response
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    // Return the template as a JSON response
    return c.JSON(fiber.Map{"template": template})
}
