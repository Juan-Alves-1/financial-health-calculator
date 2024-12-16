package handlers

import (
	services "fhi/internal/service"
	"fhi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DisplayHomepage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func FinancialHealth(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	financialHealthScore := services.CalculateFinancialHealth(*data)

	switch {
	case financialHealthScore <= 1:
		return c.Redirect(http.StatusFound, "/financial-health/survivor")
	case financialHealthScore <= 2:
		return c.Redirect(http.StatusFound, "/financial-health/dreamer")
	case financialHealthScore <= 3:
		return c.Redirect(http.StatusFound, "/financial-health/fighter")
	case financialHealthScore <= 4:
		return c.Redirect(http.StatusFound, "/financial-health/ninja")
	default:
		return c.Redirect(http.StatusFound, "/financial-health/champion")
	}
}

func SavingProjection(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	pendingMonths := services.MonthlySavingProjection(*data)
	return c.JSON(http.StatusOK, map[string]int{"savingProjectionInMonths": pendingMonths})
}

func RenderSurvivor(c echo.Context) error {
	return c.Render(http.StatusOK, "survivor.html", nil)
}

func RenderDreamer(c echo.Context) error {
	return c.Render(http.StatusOK, "dreamer.html", nil)
}

func RenderFighter(c echo.Context) error {
	return c.Render(http.StatusOK, "fighter.html", nil)
}

func RenderNinja(c echo.Context) error {
	return c.Render(http.StatusOK, "ninja.html", nil)
}

func RenderChampion(c echo.Context) error {
	return c.Render(http.StatusOK, "champion.html", nil)
}
