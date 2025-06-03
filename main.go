package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	number:= sumup(1, 2, 3, 4, 5)
	anotherNumber := sumup(numbers...)
	println("Sum of numbers is:", number)
	println("Sum of another numbers is:", anotherNumber)
}

func sumup(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}