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
	e.GET("/molar", handlers.MolarGetHandler)
	e.POST("/molar", handlers.MolarPostHandler)
	e.GET("/balance", handlers.BalanceGetHandler)
	e.POST("/balance", handlers.BalancePostHandler)
	e.Start(config.Address + ":" + config.Port)
}
