package advance

import "fmt"

type TransaformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	multiplier := 4
	doubleMultiplier := multiply(2)
	transformedNumbers := transformNumbers(&numbers, multiply(multiplier))
	doubleTransformedNumbers := transformNumbers(&numbers, doubleMultiplier)
	fmt.Println("Original numbers:", numbers)
	fmt.Println("Double transformed numbers:", doubleTransformedNumbers)
	fmt.Println("Transformed numbers:", transformedNumbers)
}

func transformNumbers(numbers *[]int, transform TransaformFn) []int {
	transformed := make([]int, len(*numbers))
	for i, v := range *numbers {
		transformed[i] = transform(v)
	}
	return transformed
}

func multiply(multiplier int) TransaformFn {
	return func(i int) int {
		return i * multiplier
	}
}