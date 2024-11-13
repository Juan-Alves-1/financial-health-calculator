package middleware

import (
	"fhi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ValidateFinancialData checks the validity of the input data
func ValidateFinancialData(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(models.FinancialData)
		if err := c.Bind(data); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input format"})
		}

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
