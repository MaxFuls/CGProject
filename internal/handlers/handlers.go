package handlers

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/parsing"
	"fmt"
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

func MolarGetHandler(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/molar.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func SicretPage(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/fortune.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func MolarPostHandler(c echo.Context) error {
	formula := c.FormValue("formula")
	parsed := parsing.ParseFormula(formula)
	elements := make(map[string]float64)
	general_mass := 0.0
	for key, count := range parsed {
		mass, err := parsing.GetMolarMass(key)
		if err != nil {
			c.String(400, "You are debil")
		}
		elements[key] = mass * float64(count)
		general_mass += elements[key]
	}
	var str string
	for key, value := range elements {
		str += fmt.Sprintf("%s - %f %% \n", key, (value/general_mass)*100.0)
	}
	str += fmt.Sprintf("General mass is %f", general_mass)
	return c.String(200, str)
}

func BalanceGetHandler(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/balance.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func BalancePostHandler(c echo.Context) error {
	return nil
}
