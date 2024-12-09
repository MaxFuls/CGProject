package main

import (
	"ChemistryPR/internal/config"
	"ChemistryPR/internal/handlers"
	"ChemistryPR/internal/logger"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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
		templates: template.Must(template.ParseGlob("web/templates/molar.html")),
	}
	e.Renderer = t;
	e.Use(middleware.Static(config.Root))
	e.GET("/", handlers.RootHandlerFunc)
	e.GET("/molar", handlers.MolarGetHandler)
	e.POST("/molar", handlers.MolarPostHandler)
	e.GET("/balance", handlers.BalanceGetHandler)
	e.POST("/balance", handlers.BalancePostHandler)
	e.GET("/sicret", handlers.SicretPage)
	e.Start(config.Address + ":" + config.Port)
}
