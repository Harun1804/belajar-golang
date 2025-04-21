package main

import "fmt"

func main() {
	age := 32 // normal variable
	agePtr := &age // pointer to the variable age
	fmt.Println("Age:", age) // prints 32
	fmt.Println("Age pointer:", agePtr) // prints the address of age variable
	fmt.Println("Age pointer value:", *agePtr) // prints the value of age variable using pointer

	adultYear := getAdultYear(agePtr) // passing the address of age variable to the function
	fmt.Println("Adult year:", adultYear) // prints 14 (32 - 18)
}

func getAdultYear(age *int) int {
	return *age - 18
}