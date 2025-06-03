package main

func main() {
	number:= sumup(1, 2, 3, 4, 5)
	println("Sum of numbers is:", number)
}

func sumup(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}