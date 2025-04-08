package main

import (
	"fmt"
	"os"
)

func main()  {
	revenue, errRevenue := getUserInput("Enter revenue: ")
	expenses, errExpenses := getUserInput("Enter expenses: ")
	taxRate, errTax := getUserInput("Enter tax rate: ")

	if errRevenue != nil {
		showError(errRevenue)		
	}

	if errExpenses != nil {
		showError(errExpenses)		
	}

	if errTax != nil {
		showError(errTax)		
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	writeIntoFile(ebt, profit, ratio)
	formattedOutput("EBT", ebt)
	formattedOutput("Profit", profit)
	formattedOutput("Profit ratio", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var input float64

	fmt.Print(infoText)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, fmt.Errorf("%s input must be greater than zero", infoText)
	}

	return input, nil
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

func showError(err error) {
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------------")
		return
	}
}

func writeIntoFile(ebt, profit, ratio float64) {
	file, err := os.Create("profit.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Ebt: %.2f\nProfit: %.2f\nRatio Rate: %.2f\n", ebt, profit, ratio)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to profit.txt")
}
