package main

import "fmt"

func main() {
	age := 32 // normal variable
	agePtr := &age // pointer to the variable age
	fmt.Println("Age:", age) // prints 32
	fmt.Println("Age pointer:", agePtr) // prints the address of age variable
	fmt.Println("Age pointer value:", *agePtr) // prints the value of age variable using pointer

	editYearAgeToAdultAge(agePtr) // passing the address of age variable to the function
	fmt.Println("Adult year:", *agePtr) // prints 14 (32 - 18)
}

func editYearAgeToAdultAge(age *int) {
	*age = *age - 18
}