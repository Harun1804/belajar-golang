package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileBalance = "balance.txt"

func writeFileBalance(balance float64) {
	balanceString := fmt.Sprint(balance)
	os.WriteFile(fileBalance, []byte(balanceString), 0644)
}

func readFileBalance() (float64, error) {
	data, err := os.ReadFile(fileBalance)
	if err != nil {
		return 1000, errors.New("balance file not found")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("error converting balance to float")
	}

	return balance, nil
}

func main()  {
	var acountBalance, err = readFileBalance()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic(err) // same as dd() in laravel
	}

	fmt.Println("Welcome to Go Bank!")
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
				writeFileBalance(acountBalance)
			case 3:
				var withdrawAmount float64
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&withdrawAmount)
				if withdrawAmount > acountBalance {
					fmt.Println("Insufficient funds!")
				} else {
					acountBalance -= withdrawAmount
					fmt.Println("Your new account balance is: ", acountBalance)
					writeFileBalance(acountBalance)
				}
			default:
				fmt.Println("Thank you for using Go Bank!")
				return
		}
	}
}
