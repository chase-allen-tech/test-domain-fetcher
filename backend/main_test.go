package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/v1/:status", StatusHandler)

	tests := []struct {
		status       string
		expectedCode int
	}{
		{STATUS_AMAZON, 200},
		{STATUS_GOOGLE, 200},
		{STATUS_ALL, 200},
		{"invalid-status", 400},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "/v1/"+tt.status, nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, tt.expectedCode, resp.StatusCode)
	}
}
