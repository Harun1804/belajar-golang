package main

import (
	"math"
	"fmt"
)

func main()  {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	fmt.Println("Future value is", futureValue)
}
