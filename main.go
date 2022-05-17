package main

import (
	"banking-app/helper"
	"fmt"
	"os"
)

func main() {
	var name string
	

	fmt.Println("Welcome to the Go Banking Application!")
	fmt.Println("------------------------------------")
	
	fmt.Println("Firstly, what is your name?")
	fmt.Scan(&name)

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
