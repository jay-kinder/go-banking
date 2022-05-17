package main

import (
	"banking-app/database"
	"banking-app/helper"
	"fmt"
	"log"
	"os"
)

func main() {
	var name string
	var email string
	

	database.DbConnection()

	fmt.Println()
	fmt.Println("Welcome to the Go Banking Application!")
	fmt.Println("------------------------------------")
	
	fmt.Println("Firstly, what is your name?")
	fmt.Scan(&name)

	// Checking DB for User
	fmt.Println("Next, what is your email")
	fmt.Scan(&email)
	users, err := database.CurrentUserCheck(email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", users)
	// --------------------------------------

	for {
		menuInput := helper.HomePage(name)

		switch menuInput {
			case 1:
				depositAmount := helper.DepositMoney()
				go helper.ProcessDeposit(depositAmount)	
				helper.AfterSelection(name)
			case 2:
				helper.WithdrawMoney()				
				helper.AfterSelection(name)
			case 3:
				helper.GetBalance()
				helper.AfterSelection(name)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("You need to enter a number between 1 and 4...IS THAT NOT CLEAR??!!")
		}
	}
}
