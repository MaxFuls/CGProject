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

// func MolarPostHandler(c echo.Context) error {
//
// formula := c.FormValue("formula")
// parsed := parsing.ParseFormula(formula)
// elements := make(map[string]float64)
//
// total := 0.0
// for key, count := range parsed {
// mass, err := parsing.GetMolarMass(key)
// if err != nil {
// return c.String(400, "pizdec")
// }
// elements[key] = mass * float64(count)
// total += elements[key]
// }
//
// var disc ElementDiscription
// disc.Total = fmt.Sprintf("%.3f g/mol", total)
// disc.Formula = formula
// for key, count := range parsed {
// russian, _ := parsing.GetRussianName(key)
// percent := (elements[key] / total) * 100;
// disc.Elements = append(disc.Elements,
// ElementInfo{key, []string{"Название: " + russian,
//   "Масса в соединении: " + fmt.Sprintf("%.3f g/mol", elements[key]),
//   "Количество атомов: " + strconv.Itoa(count),
//   "Процент от общей массы: " + fmt.Sprintf("%.3f %%", percent)}})
// }
// return c.Render(http.StatusOK, "molar", disc)
// }

func MolarPostHandler(c echo.Context) error {
	config := config.LoadConfig()
	db, closeFunc, err := database.OpenDB(config.Driver, "chem.db")
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
