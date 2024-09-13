package main

import "fmt"

func main()  {
	var names = []string{"John", "Doe"}
	fmt.Println(calculate(1, 2))
	fmt.Println(getAverage(1, 2, 3, 4, 5, 6 ,7, 8, 9, 10))
	printName("Hello", names)
}

func printName(word string, arr []string)  {
	for _, name := range arr {
		fmt.Println(word, name)
	}
}

func calculate(a, b int) int {
	return a + b
}

func getAverage(numbers ...int) float64  {
	var total int = 0;
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
