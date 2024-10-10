package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Structure to hold the form input
type FinancialData struct {
	Income      float64 `json:"income"`
	Expenses    float64 `json:"expenses"`
	Savings     float64 `json:"savings"`
	SavingsGoal float64 `json:"savings_goal"`
}

func (fd *FinancialData) GetBalance() (float64, error) {
	if fd.Income < fd.Expenses {
		return -1, fmt.Errorf("expenses exceeded income = income: %.0f, expenses: %.0f", fd.Income, fd.Expenses)
	}
	return fd.Income - fd.Expenses, nil
}

// Create constant for each score scenario <= 0, 1, 2, 3, 4, 5...
func CalculateFinancialHealth(data FinancialData) float64 {
	balance, err := data.GetBalance()
	if err != nil {
		log.Println(err) // Ask Henry
		// add return here to stop logic
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

func SavingsProjection(data FinancialData) float64 {
	pendingSavings := data.SavingsGoal - data.Savings
	balance, err := data.GetBalance()
	if err != nil {
		log.Println(err) // Ask Henry
		return 0
	}
	return pendingSavings / balance
}

// handler for CalculateFinancialHealth
func financialHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var data FinancialData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	financialHealthScore := CalculateFinancialHealth(data)
	response := map[string]float64{"financialHealthScore": financialHealthScore}
	json.NewEncoder(w).Encode(response)
}

// handler for SavingsProjection
func SavingsProjectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var data FinancialData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	PendingMonths := SavingsProjection(data)
	response := map[string]float64{"PendingMonths": PendingMonths}
	json.NewEncoder(w).Encode(response)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/financial-health", financialHealthHandler)
	mux.HandleFunc("/api/savings-projection", SavingsProjectionHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	fmt.Println("Server is running on port 8080...")
	log.Fatal(server.ListenAndServe())

}
