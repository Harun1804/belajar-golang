package main

import (
	"fmt"
	"runtime"
)

// import "first-app/interfaces"
// import "first-app/geometry"

func main() {
	runtime.GOMAXPROCS(2)

	// messages := make(chan int, 2)

	var messages = make(chan string)

	var sayHelloTo = func (who string)  {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// go func ()  {
	// 	for {
	// 		i := <- messages
	// 		fmt.Println("receive data", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("send data", i)
	// 	messages <- i
	// }

	// var bangunDatar interfaces.Hitung
	// bangunDatar = geometry.Persegi{Sisi: 10.0}
	// fmt.Println("===== persegi")
	// fmt.Println("luas      :", bangunDatar.Luas())
	// fmt.Println("keliling  :", bangunDatar.Keliling())

	// bangunDatar = geometry.Lingkaran{Diameter: 14.0}
	// fmt.Println("===== lingkaran")
	// fmt.Println("luas      :", bangunDatar.Luas())
	// fmt.Println("keliling  :", bangunDatar.Keliling())
	// fmt.Println("jari-jari :", bangunDatar.(geometry.Lingkaran).JariJari())
	
	// var s1 = library.Student{Name: "ethan", Grade: 20}
	// fmt.Println(s1.Name)
	// fmt.Println(s1.Grade)
	// library.SayHello()
	// library.Introduce("ethan")
}

