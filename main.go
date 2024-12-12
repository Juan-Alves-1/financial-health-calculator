package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"fhi/internal/handlers"
	customMiddleware "fhi/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	const port string = ":8080"

	e := echo.New()

	// Initialize templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static") // deny any accesses?

	e.GET("/", handlers.DisplayHomepage)
	e.POST("/financial-health", handlers.FinancialHealth, customMiddleware.ValidateFinancialData)
	e.POST("/savings-projection", handlers.SavingProjection, customMiddleware.ValidateFinancialData)

	server := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(e.StartServer(server))

}
