package main

import (
	"fmt"
	"example.com/bank/media"
	"github.com/Pallinder/go-randomdata"
)

const fileBalance = "balance.txt"

func main()  {
	var acountBalance, err = media.ReadFloatFromFile(fileBalance)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic(err) // same as dd() in laravel
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("We are online 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()
	
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		
		fmt.Println("You selected: ", choice)
	
		switch choice {
			case 1:
				fmt.Println("Your account balance is: ", acountBalance)
			case 2:
				var depositAmount float64
				fmt.Println("Enter the amount to deposit: ")
				fmt.Scan(&depositAmount)
				acountBalance += depositAmount
				if depositAmount <= 0 {
					fmt.Println("Invalid deposit amount!")
					continue
				}

				fmt.Println("Your new account balance is: ", acountBalance)
				media.WriteFloatToFile(acountBalance, fileBalance)
			case 3:
				var withdrawAmount float64
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&withdrawAmount)
				if withdrawAmount > acountBalance {
					fmt.Println("Insufficient funds!")
				} else {
					acountBalance -= withdrawAmount
					fmt.Println("Your new account balance is: ", acountBalance)
					media.WriteFloatToFile(acountBalance, fileBalance)
				}
			default:
				fmt.Println("Thank you for using Go Bank!")
				return
		}
	}
}
