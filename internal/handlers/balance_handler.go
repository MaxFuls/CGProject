package handlers

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/database"
	"ChemistryPR/internal/services"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// type SubstanceDiscription struct {
// 	Name string
// 	Details []string
// }

type BalanceResponse struct {
	Reaction string
	Result   string
	Reagents []string //SubstanceDiscription
	Products []string //SubstanceDiscription
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
	config := config.LoadConfig()
	db, closeFunc, err := database.OpenDB(config.Driver, "chem.db")
	if err != nil {
		log.Debug("pizda bd nakrilas")
		panic("pizda bd nakrilas")
	}
	defer closeFunc()
	reaction := c.FormValue("reaction")
	service := services.BalanceService{}
	service.Store = database.NewStore(db)
	response, err := service.GetResponse(reaction)
	for err != nil {
		c.String(http.StatusBadRequest, "pizda huinu napisal")
	}
	return c.Render(http.StatusOK, "balance", response)
}
