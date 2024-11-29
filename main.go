package main

import "fmt"

func main() {
	var name string = "John"
	age := 27

	count := 10
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			println("I am counting", i)
		}
	}

	switch age {
		case 25:
			println("I am 25")
		case 26:
			println("I am 26")
		default:
			println("I am not 25 or 26")
	}

	var fruits = []string{"apple", "banana", "orange"}
	for i, fruit := range fruits {
		println(i, fruit)
	}

	fmt.Println("Hello, my name is", name, "and I am", age, "years old.")
	var result int = add(10, 5)
	println("10 + 5 =", result)

	var avg float32 = calculate(10, 20, 30, 40, 50, 60, 70)
	var msg = fmt.Sprintf("The average is %.2f", avg)
	println(msg)
}

func add(a int, b int) int {
	return a + b
}

func calculate(numbers ...int) float32 {
	var total int = 0;

	for _, number := range numbers {
		total += number
	}

	var avg = float32(total) / float32(len(numbers))
	return avg
}
