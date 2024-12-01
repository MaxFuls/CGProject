package handlers

import (
	"ChemistryPR/internal/config"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func RootHandlerFunc(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/index.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func MolarHandlerFunc(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/molar.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func BalanceHandlerFunc(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/balance.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}
