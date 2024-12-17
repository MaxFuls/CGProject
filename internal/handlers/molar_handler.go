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

type ElementInfo struct {
	Symbol  string
	Details []string
}

type ElementDiscription struct {
	Formula  string
	Total    string
	Elements []ElementInfo
}

func MolarGetHandler(c echo.Context) error {
	config := config.LoadConfig()
	content, err := os.ReadFile(config.Root + "/molar.html")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.HTMLBlob(200, content)
}

func MolarPostHandler(c echo.Context) error {
	config := config.LoadConfig()
	db, closeFunc, err := database.OpenDB(config.Driver, config.Dns)
	if err != nil {
		log.Debug("pizda bd nakrilas")
		panic("pizda bd nakrilas")
	}
	defer closeFunc()
	formula := c.FormValue("formula")
	s := services.MolarMassService{}
	s.Store = database.NewStore(db)
	response, err := s.GetResponse(formula)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.Render(http.StatusOK, "molar", response)
}
