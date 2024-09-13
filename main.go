package main

import "fmt"

func main()  {
	var names = []string{"John", "Doe"}
	printName("Hello", names)
}

func printName(word string, arr []string)  {
	for _, name := range arr {
		fmt.Println(word, name)
	}
}
