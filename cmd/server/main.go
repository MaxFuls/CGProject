package main

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/handlers"
	"ChemistryPR/internal/logger"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	config := config.LoadConfig()
	log := logger.SetupLogger(config.Env)
	log.Info("Starting server")
	log.Debug("Debug messages are enabled")
	e := echo.New()
	e.Static("/frontend", "frontend")
	e.GET("/", handlers.RootHandlerFunc)
	e.GET("/molar", handlers.MolarHandlerFunc)
	e.GET("/balance", handlers.BalanceHandlerFunc)
	e.Start(config.Address + ":" + config.Port)
}
