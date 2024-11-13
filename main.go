package main

import (
	"log"
	"math"
	"net/http"
	"time"

	customMiddleware "fhi/middleware"
	"fhi/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Create constant for each score scenario <= 0, 1, 2, 3, 4, 5...
func CalculateFinancialHealth(data models.FinancialData) float64 {
	balance, err := data.GetBalance()
	if err != nil {
		log.Println(err) // Ask Henry
		return 0         // add return here to stop logic
	}
	balanceRatio := balance / data.Income
	balanceScore := calculateBalancePoints(balanceRatio)

	savingsRatio := data.Savings / data.SavingsGoal
	savingsScore := calculateSavingPoints(savingsRatio)

	healthScore := balanceScore + savingsScore
	return healthScore
}

// Function to calculate balance points
func calculateBalancePoints(ratio float64) float64 {
	switch {
	case ratio == 0:
		return 0
	case ratio > 0 && ratio <= 0.25:
		return 1
	case ratio > 0.25 && ratio <= 0.5:
		return 2
	case ratio > 0.5 && ratio <= 0.75:
		return 3
	case ratio > 0.75 && ratio <= 1:
		return 4
	case ratio > 1:
		return 5
	default:
		return 0
	}
}

// Function to calculate saving points
func calculateSavingPoints(ratio float64) float64 {
	switch {
	case ratio == 0:
		return 0
	case ratio > 0 && ratio <= 0.25:
		return 1
	case ratio > 0.25 && ratio <= 0.5:
		return 2
	case ratio > 0.5 && ratio <= 0.75:
		return 3
	case ratio > 0.75 && ratio <= 1:
		return 4
	case ratio > 1:
		return 5
	default:
		return 0
	}
}

func MonthlySavingProjection(data models.FinancialData) int {
	pendingSavings := data.SavingsGoal - data.Savings
	balance, err := data.GetBalance()

	if err != nil || balance <= 0 {
		log.Println("Not able to save money:", err)
		return -1 // Return 0 or a more meaningful value if the user can't save
	}

	months := int(math.Ceil(pendingSavings / balance))

	return months
}

// handler for CalculateFinancialHealth
func financialHealthHandler(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	financialHealthScore := CalculateFinancialHealth(*data)
	return c.JSON(http.StatusOK, map[string]float64{"financialHealthScore": financialHealthScore})
}

// handler for SavingsProjection
func SavingsProjectionHandler(c echo.Context) error {
	data := c.Get("financialData").(*models.FinancialData)
	pendingMonths := MonthlySavingProjection(*data)
	return c.JSON(http.StatusOK, map[string]int{"savingsProjectionInMonths": pendingMonths})
}

func main() {
	const port string = ":8080"

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/financial-health", financialHealthHandler, customMiddleware.ValidateFinancialData)
	e.POST("/api/savings-projection", SavingsProjectionHandler, customMiddleware.ValidateFinancialData)

	server := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(e.StartServer(server))

}
