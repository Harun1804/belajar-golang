package main

import (
	"math"
	"fmt"
)

func main()  {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	years := 10.0

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formatterFV := fmt.Sprintf("Future value is %.2f\n", futureValue)

	// fmt.Println("Future value is ", futureValue)
	fmt.Print(formatterFV)
	fmt.Println("Future real value is ", futureRealValue)
}
