package main

import (
	"fmt"
)

func main()  {
	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	formattedOutput("EBT", ebt)
	formattedOutput("Profit", profit)
	formattedOutput("Profit ratio", ratio)
}

func getUserInput(infoText string) float64 {
	var input float64
	fmt.Print(infoText)
	fmt.Scan(&input)
	return input
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func formattedOutput(label string, value float64)  {
	fmt.Printf("%s = %.2f\n", label, value)
}
