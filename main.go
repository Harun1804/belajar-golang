package main

import "fmt"

type TransaformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	transformedNumbers := transformNumbers(&numbers, multiplyByThree)
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

func multiplyByTwo(n int) int {
	return n * 2
}

func multiplyByThree(n int) int {
	return n * 3
}