package handlers

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/parsing"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ElementInfo struct {
	Symbol string
	Details []string
}

type ElementDiscription struct {
	Formula string
	Total string
	Elements []ElementInfo 
}

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
	
	total := 0.0
	for key, count := range parsed {
		mass, err := parsing.GetMolarMass(key)
		if err != nil {
			return c.String(400, "pizdec")
		}
		elements[key] = mass * float64(count)
		total += elements[key]
	}
	
	var disc ElementDiscription
	disc.Total = fmt.Sprintf("%.3f g/mol", total)
	disc.Formula = formula
	for key, count := range parsed {
		russian, _ := parsing.GetRussianName(key)
		percent := (elements[key] / total) * 100;
		disc.Elements = append(disc.Elements,
			ElementInfo{key, []string{"Название: " + russian,
									  "Масса в соединении: " + fmt.Sprintf("%.3f g/mol", elements[key]),
									  "Количество атомов: " + strconv.Itoa(count),
									  "Процент от общей массы: " + fmt.Sprintf("%.3f %%", percent)}})
	}
	return c.Render(200, "molar", disc)
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
