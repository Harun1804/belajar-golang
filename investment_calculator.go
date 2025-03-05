package main

import (
	"math"
	"fmt"
)

func main()  {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)
}
