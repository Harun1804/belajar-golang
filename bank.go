package main

import "fmt"

func main()  {
	var acountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		
		fmt.Println("You selected: ", choice)
	
		if choice == 1 {
			fmt.Println("Your account balance is: ", acountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Println("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)
			acountBalance += depositAmount
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}

			fmt.Println("Your new account balance is: ", acountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > acountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				acountBalance -= withdrawAmount
				fmt.Println("Your new account balance is: ", acountBalance)
			}
		} else if choice == 4 {
			fmt.Println("Thank you for using Go Bank!")
			break
		} else {
			fmt.Println("Invalid choice!")
			break
		}
	}

	fmt.Println("Thank for using our bank...")
}