package middleware

import (
	"fhi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ValidateFinancialData checks the validity of the input data
func ValidateFinancialData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create a FinancialData struct
		data := new(models.FinancialData)

		// Parse form values
		income, err := strconv.ParseFloat(c.FormValue("income"), 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid income format"})
		}

		expenses, err := strconv.ParseFloat(c.FormValue("expenses"), 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expenses format"})
		}

		savings, err := strconv.ParseFloat(c.FormValue("savings"), 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid savings format"})
		}

		savingsGoal, err := strconv.ParseFloat(c.FormValue("savingsGoal"), 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid savings goal format"})
		}

		data.Income = income
		data.Expenses = expenses
		data.Savings = savings
		data.SavingsGoal = savingsGoal

		if data.Income < 0 || data.Expenses < 0 || data.Savings < 0 || data.SavingsGoal < 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Values cannot be negative"})
		}

		if data.Income == 0 || data.Expenses == 0 || data.SavingsGoal == 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Income, Expenses or Savings goal cannot be zero"})
		}

		if data.Savings >= data.SavingsGoal {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Savings cannot be or exceed the savings goal"})
		}

		c.Set("financialData", data)
		return next(c)
	}
}
