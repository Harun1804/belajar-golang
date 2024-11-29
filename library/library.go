package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func SayHello() {
	fmt.Println("Hello")
}

func Introduce(name string) {
	fmt.Println("Hello", name)
}
