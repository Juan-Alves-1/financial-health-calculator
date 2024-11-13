package models

import "fmt"

type FinancialData struct {
	Income      float64 `json:"income"`
	Expenses    float64 `json:"expenses"`
	Savings     float64 `json:"savings"`
	SavingsGoal float64 `json:"savingsGoal"`
}

func (fd *FinancialData) GetBalance() (float64, error) {
	if fd.Income < fd.Expenses {
		return -1, fmt.Errorf("expenses exceeded income = income: %.0f, expenses: %.0f", fd.Income, fd.Expenses)
	}
	return fd.Income - fd.Expenses, nil
}
