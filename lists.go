package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string {"Reading", "Hiking", "Cooking"}
	fmt.Println(hobbies)
	// 2)
	fmt.Println("first hobby:", hobbies[0])
	fmt.Println("new hobbies:", hobbies[1:3])
	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println("main hobbies:", mainHobbies)
	// 4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println("main hobbies after slicing:", mainHobbies)
	// 5)
	courseGoals := []string{"Complete Go course", "Build a project"}
	fmt.Println("course goals:", courseGoals)
	// 6)
	courseGoals[1] = "Build a web app"
	courseGoals = append(courseGoals, "Learn concurrency")
	fmt.Println("updated course goals:", courseGoals)
	// 7)
	products := []Product {
		{id: "1", title: "Laptop", price: 999.99},
		{id: "2", title: "Smartphone", price: 499.99},
	}

	println("products:", products)
	newProduct := Product{id: "3", title: "Tablet", price: 299.99}
	products = append(products, newProduct)
	fmt.Println("updated products:", products)
}