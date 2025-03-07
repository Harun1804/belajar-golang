package main

import (
	"math"
	"fmt"
)

func main()  {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := math.Round(profit / revenue * 100)

	fmt.Printf("EBT = %.2f\n", ebt)
	fmt.Printf("Profit = %.2f\n", profit)
	fmt.Printf("Profit ratio = %.2f%%\n", ratio)
}
