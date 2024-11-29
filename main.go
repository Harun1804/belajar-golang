package main

import "first-app/library"
import "fmt"

func main() {
	var s1 = library.Student{Name: "ethan", Grade: 20}
	fmt.Println(s1.Name)
	fmt.Println(s1.Grade)
	library.SayHello()
	library.Introduce("ethan")
}

