package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(c echo.Context) error {
	response := HealthResponse{
		Status: "ok",
	}
	return c.JSON(http.StatusOK, response)
}
