package main

import (
	"fmt"
)

// Structure to hold the form input
type FinancialData struct {
	Income      float64
	Expenses    float64
	Savings     float64
	SavingsGoal float64
}

func (fd *FinancialData) GetBalance() float64 {
	return fd.Income - fd.Expenses
}

func CalculateFinancialHealth(data FinancialData) float64 {
	balance := data.GetBalance()
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

func SavingsProjection(data FinancialData) {
	pendingSavings := data.SavingsGoal - data.Savings
	balance := data.GetBalance()
	if balance <= 0 {
		fmt.Println("Not able to save money")
		return
	}
	PendingMonths := pendingSavings / balance
	fmt.Printf("You're able to reach your saving goals in %.0f months\n", PendingMonths)
}

func main() {
	var data FinancialData
	fmt.Println("Enter your income: ")
	fmt.Scan(&data.Income)

	fmt.Println("Enter your expenses: ")
	fmt.Scan(&data.Expenses)

	fmt.Println("Enter your current savings: ")
	fmt.Scan(&data.Savings)

	fmt.Println("Enter your savings goal: ")
	fmt.Scan(&data.SavingsGoal)

	// Calculate and display financial health score
	financialHealth := CalculateFinancialHealth(data)
	fmt.Printf("Your Financial Health Score: %.0f\n", financialHealth)
	SavingsProjection(data)

}
