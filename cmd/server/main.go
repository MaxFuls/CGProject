package main

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/handlers"
	"ChemistryPR/internal/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	config := config.LoadConfig()
	log := logger.SetupLogger(config.Env)
	log.Info("Starting server")
	log.Debug("Debug messages are enabled")

	e := echo.New()
	e.Use(middleware.Static("frontend"))
	e.GET("/", handlers.RootHandlerFunc)
	e.GET("/molar", handlers.MolarHandlerFunc)
	e.GET("/balance", handlers.BalanceHandlerFunc)
	e.POST("/submit", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(200, name+email)
	})
	e.Start(config.Address + ":" + config.Port)
}
