package handlers

import (
	services "fhi/internal/service"
	"fhi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FinancialHealth(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	financialHealthScore := services.CalculateFinancialHealth(*data)

	var templateName string
	switch {
	case financialHealthScore <= 1:
		templateName = "survivor.html"
	case financialHealthScore <= 2:
		templateName = "dreamer.html"
	case financialHealthScore <= 3:
		templateName = "fighter.html"
	case financialHealthScore <= 4:
		templateName = "ninja.html"
	default:
		templateName = "champion.html"
	}

	// Render the selected template
	return c.Render(http.StatusOK, templateName, nil)
}

func SavingProjection(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	pendingMonths := services.MonthlySavingProjection(*data)
	return c.JSON(http.StatusOK, map[string]int{"savingProjectionInMonths": pendingMonths})
}
