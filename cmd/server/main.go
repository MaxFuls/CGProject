package main

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/handlers"
	"ChemistryPR/internal/logger"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	
	config := config.LoadConfig()
	log := logger.SetupLogger(config.Env)
	log.Info("Starting server")
	log.Debug("Debug messages are enabled")

	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("web/templates/*.html"))}
	e.Renderer = t;

	e.Use(middleware.Static(config.Root))
	e.GET("/", handlers.RootHandlerFunc)
	e.GET("/molar", handlers.MolarGetHandler)
	e.POST("/molar", handlers.MolarPostHandler)
	e.GET("/balance", handlers.BalanceGetHandler)
	e.POST("/balance", handlers.BalancePostHandler)
	e.GET("/fortune", func (c echo.Context) error {
		content, err := os.ReadFile("web/fortune.html")
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.HTMLBlob(200, content)
	})
	e.Start(config.Address + ":" + config.Port)
}
