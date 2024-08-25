package integration

import (
	"bytes"
	"encoding/json"
	"golang/backend/database"
	"golang/backend/routers"
	"golang/backend/utils"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestPaginationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	} `json:"data"`
}

func setupApp() *fiber.App {
	app := fiber.New()
	utils.LoadENVTest()
	database.ConnectDb()
	routers.SetupRouter(app)
	return app
}

func TestGetAllProductsIntegration(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var response TestPaginationResponse
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, "Product succesfully retrieve", response.Message)
	assert.Equal(t, 1, response.Data.Page)
	assert.Equal(t, 10, response.Data.Limit)
}

func TestAddProductIntegration(t *testing.T) {
	app := setupApp()

	// Define the input payload for the POST request
	input := `{"category_id":"c84fe982-4ac6-46fa-b6f5-c0dbccb179b5","name":"Test Product","description":"Test Description"}`

	// Create a new POST request with the input payload
	req := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBufferString(input))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := app.Test(req)
	body, _ := io.ReadAll(resp.Body)
	expectedResponse := `{
    "success": true,
    "message": "Product successfully created",
    "data": null,
    "error": null
}`

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.JSONEq(t, expectedResponse, string(body))
}
