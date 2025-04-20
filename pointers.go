package main

import "fmt"

func main() {
	age := 32 // normal variable
	agePtr := &age // pointer to the variable age
	fmt.Println("Age:", age) // prints 32
	fmt.Println("Age pointer:", agePtr) // prints the address of age variable
	fmt.Println("Age pointer value:", *agePtr) // prints the value of age variable using pointer
}