package controller

import (
	"cms-config/internal/app/usecase"
	"cms-config/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type RemoteConfigController struct {
	usecase *usecase.RemoteConfigUseCase
	client  *http.Client
}

var url string

func init() {
	viper := config.NewViper()
	url = viper.GetString("REMOTE_CONFIG_URL")
}

func NewRemoteConfigController(usecase *usecase.RemoteConfigUseCase, client *http.Client) *RemoteConfigController {
	return &RemoteConfigController{usecase: usecase, client: client}
}

func (rc *RemoteConfigController) GetTemplate(c *fiber.Ctx) error {
	token, err := rc.usecase.GetOauthToken()
	if err != nil {
		log.Fatalf("error getting token: %v", err)
	}

	viper := config.NewViper()

	url := viper.GetString("REMOTE_CONFIG_URL")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}
	// set authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	resp, err := rc.client.Do(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	if resp.StatusCode != 200 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": resp.Status,
		})
	}

	var remoteConfigBody map[string]interface{}

	// decode body
	err = json.NewDecoder(resp.Body).Decode(&remoteConfigBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// overwrite remote-config.json file
	file, err := os.Create("remote-config.json")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// Marshal the decoded data to a byte slice
	jsonData, err := json.Marshal(remoteConfigBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// Open the file for writing with error handling
	file, err = os.OpenFile("remote-config.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// Write the JSON data to the file
	defer file.Close() // Ensure file is closed even on errors
	_, err = file.Write(jsonData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	log.Infoln(remoteConfigBody)

	return c.Status(fiber.StatusOK).JSON(remoteConfigBody)
}

func (rc *RemoteConfigController) GetDefaultValue(c *fiber.Ctx) error {
	token, err := rc.usecase.GetOauthToken()
	if err != nil {
		log.Fatalf("error getting token: %v", err)
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}
	// get path parameter
	pathParam := c.Params("param_name")
	log.Info(pathParam)

	// set authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	resp, err := rc.client.Do(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	if resp.StatusCode != 200 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": resp.Status,
		})
	}

	var remoteConfigBody map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&remoteConfigBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	// Check if "parameters" key exists and is a map
	parameters, ok := remoteConfigBody["parameters"].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Parameters not found",
		})
	}

	// Check if pathParam key exists in parameters map and is a map
	parameterName, ok := parameters[pathParam].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Parameter not found",
		})
	}

	// Check if "defaultValue" key exists in param map and is a map
	defaultValueMap, ok := parameterName["defaultValue"].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Default value not found",
		})
	}

	// Check if "value" key exists in defaultValueMap and is a string
	defaultValue, ok := defaultValueMap["value"].(string)
	if !ok || defaultValue == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Value not found or empty",
		})
	}

	// if not found, return error
	if defaultValue == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Parameter not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(defaultValue)
}
