package main

import (
	"fmt"
)

type str string

func (s str) log() {
	fmt.Println("Logging: " + s)
}

func main() {
	var text str = "Hello, World!"
	text.log()
}
