package main

import (
	"fmt"
	"runtime"
	"first-app/calculate"
)

// import "first-app/interfaces"
// import "first-app/geometry"

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 2, 3, 5, 6, 7, 4, 3, 2, 4, 5, 3, 2, 4, 6, 7, 8, 9, 0, 8, 7, 6, 5, 4, 3, 2, 1}

	var ch1 = make(chan float64)
	go calculate.GetAverage(numbers, ch1)

	var ch2 = make(chan int)
	go calculate.GetMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
			case avg := <-ch1:
				fmt.Printf("Average \t: %.2f\n", avg)
			case max := <-ch2:
				fmt.Printf("Max \t: %d\n", max)
		}
	}

	// messages := make(chan int, 2)

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

