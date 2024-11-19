package services

import (
	"fhi/models"
	"log"
	"math"
)

// Create constant for each score scenario <= 0, 1, 2, 3, 4, 5...
func CalculateFinancialHealth(data models.FinancialData) float64 {
	balance, err := data.GetBalance()
	if err != nil {
		log.Println(err) // Ask Henry
		return 0         // add return here to stop logic
	}
	balanceRatio := balance / data.Income
	balancePoints := calculateBalancePoints(balanceRatio)

	savingRatio := data.Savings / data.SavingsGoal
	savingPoints := calculateSavingPoints(savingRatio)

	weightedScore := (balancePoints * 0.8) + (savingPoints * 0.2)
	return weightedScore
}

// Function to calculate balance points
func calculateBalancePoints(ratio float64) float64 {
	switch {
	case ratio > 0.05 && ratio <= 0.15:
		return 1
	case ratio > 0.15 && ratio <= 0.25:
		return 2
	case ratio > 0.25 && ratio <= 0.4:
		return 3
	case ratio > 0.4 && ratio <= 0.7:
		return 4
	case ratio > 0.7 && ratio <= 1:
		return 5
	default:
		return 0
	}
}

// Function to calculate saving points
func calculateSavingPoints(ratio float64) float64 {
	switch {
	case ratio > 0.05 && ratio <= 0.25:
		return 1
	case ratio > 0.25 && ratio <= 0.5:
		return 2
	case ratio > 0.5 && ratio <= 0.75:
		return 3
	case ratio > 0.75 && ratio < 1:
		return 4
	case ratio == 1:
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
