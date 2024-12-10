package handlers

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/parsing"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

// type SubstanceDiscription struct {
// 	Name string
// 	Details []string
// }

type BalanceResponse struct {
	Reaction string
	Result string
	Reagents []string//SubstanceDiscription
	Products []string//SubstanceDiscription
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
	reaction := c.FormValue("reaction")
	reaction = strings.ReplaceAll(reaction, " ", "")
	var response BalanceResponse
	response.Reaction = reaction
	var reagents []string
	var products []string
	is_reagents := true  
	previous := 0
	for i, c := range reaction {
		if c == '+' && is_reagents {
			reagents = append(reagents, reaction[previous:i])
			previous = i + 1
		} else if c == '+' && !is_reagents {
			products = append(products, reaction[previous:i])
			previous = i + 1
		}
		if c == '=' {
			reagents = append(reagents, reaction[previous:i])
			previous = i + 1
			is_reagents = false
		}
	}
	products = append(products, reaction[previous:])
	result, err := parsing.BalanceEquation(reagents, products)
	if err != nil {
		return c.String(http.StatusInternalServerError, "parsing error")
	}
	response.Result = result
	response.Reagents = reagents
	response.Products = products
	return c.Render(http.StatusOK, "balance", response)
}