package recursive

import "fmt"

func main() {
	facNumber := 6
	fac := factorial(facNumber)
	fmt.Printf("Factorial of %d is: %d \n", facNumber, fac)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}