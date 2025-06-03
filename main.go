package main

import "fmt"

type TransaformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	multiplier := 4
	transformedNumbers := transformNumbers(&numbers, multiply(multiplier))
	fmt.Println("Original numbers:", numbers)
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