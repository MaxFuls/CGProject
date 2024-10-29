package main

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/logger"
	midLog "ChemistryPR/internal/middleware/logger"
	"bufio"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("button clicked - 1")
		return next(c)
	}
}

func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("pisa popa")
	return func(c echo.Context) error {
		fmt.Println("button clicked - 2")
		return next(c)
	}
}

func main() {
	config := config.LoadConfig()
	log := logger.SetupLogger(config.Env)
	log.Info("Starting server")
	log.Debug("Debug messages are enabled")
	buffer := make([]byte, 1024)
	file, _ := os.Open("../../index.html")
	reader := bufio.NewReader(file)
	reader.Read(buffer)
	e := echo.New()
	e.Use(midLog.LogMiddleware(log))
	e.GET("/", func(c echo.Context) error { return c.HTMLBlob(200, buffer) })
	e.Start(config.Address + ":" + config.Port)
}
