package main

import (
	"math"
	"fmt"
)

const inflationRate = 2.5

func main()  {
	var investmentAmount float64
	var expectedReturnRate float64
	years := 10.0

	// fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter expected return rate: ")
	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter number of years: ")
	outputText("Enter number of years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formatterFV := fmt.Sprintf("Future value is %.2f\n", futureValue)

	// fmt.Println("Future value is ", futureValue)
	fmt.Print(formatterFV)
	fmt.Println("Future real value is ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv = fv / math.Pow(1 + inflationRate / 100, years)
	return fv, frv
	// return // named return
}
