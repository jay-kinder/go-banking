package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var name string
	var accountBalance float64

	fmt.Println("Welcome to the Go Banking Application!")
	fmt.Println("------------------------------------")
	
	fmt.Println("Firstly, what is your name?")
	fmt.Scan(&name)

	for {
		menuInput := homePage(name)

		switch menuInput {
			case 1:
				depositAmount := depositMoney()
				accountBalance += depositAmount	
				afterSelection(name)
			case 2:
				withdrawAmount := withdrawMoney(accountBalance)
				accountBalance -= withdrawAmount
				afterSelection(name)
			case 3:
				getBalance(accountBalance)
				afterSelection(name)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("You need to enter a number between 1 and 4...IS THAT NOT CLEAR??!!")
		}
	}
}

// HomePage
func homePage(name string) uint {
	var menuInput uint
	fmt.Println("------------------------------------")
	fmt.Printf("Cheers %v! Please select from the following services:\n\n1. Deposit Moneys\n2. Withdraw Moneys\n3. See Account Balance\n4. Exit\n", name)
	fmt.Println()
	fmt.Println("Choice: ")
	fmt.Scan(&menuInput)
	return menuInput
}

// Deposit
func depositMoney() float64{
	var depositAmount float64
	for {
		fmt.Println()
		fmt.Println("Great! Let's put some money into your account!")
		fmt.Println("How much would you like to put in (JGs)?")
		fmt.Scan(&depositAmount)
		if !validateNumber(depositAmount) {
			fmt.Println("You need to enter a positive number, you silly sausage!")
			continue
		}
		fmt.Println()
		fmt.Println("This will take a couple of seconds to process...")
		time.Sleep(5 * time.Second)
		fmt.Println("......................................")
		time.Sleep(5 * time.Second)
		fmt.Println()
		fmt.Printf("Cool! You've just put %v JayGolds into your Account\n", depositAmount)
		return depositAmount
	}
}
// Withdraw
func withdrawMoney(accountBalance float64) float64 {
	var withdrawAmount float64
	for {
		fmt.Println()
		fmt.Println("Awesome! Let's take some money out of your account.")
		fmt.Println()
		fmt.Println("How much would you like to take out (JGs)?")
		fmt.Scan(&withdrawAmount)
		if !validateNumber(withdrawAmount) || withdrawAmount > accountBalance {
			fmt.Println("You have either tried to enter a negative number, or you are trying to take out more money than you have...tut tut.")
			continue
		}
		fmt.Println()
		fmt.Println("This will take a couple of seconds to process...")
		time.Sleep(3 * time.Second)
		fmt.Println("......................................")
		time.Sleep(3 * time.Second)
		fmt.Printf("Sweet! You've just taken %v JayGolds out of your Account\n", withdrawAmount)
		return withdrawAmount
	}
}
// Balance
func getBalance(accountBalance float64) float64 {
	fmt.Println()
	fmt.Println("Don't be too disappointed...here is your bank balance!")
	fmt.Println("......................................")
	time.Sleep(3 * time.Second)
	fmt.Printf("Balance: %v JayGolds\n", accountBalance)
	return accountBalance
} 

// AfterSelection (go back to select another service, or exit)
func afterSelection(name string) {
	var inputSelection uint
	selectionLoop:
		for {
			fmt.Println()
			fmt.Printf("Would you like to use another Service?:\n\n1.Yes\n2. No\n")
			fmt.Scan(&inputSelection)

			switch inputSelection {
				case 1:
					break selectionLoop
				case 2:
					os.Exit(0)
				default:
					fmt.Println("You need to enter either 1 or 2...DUH!!")
					continue
			}
		}
}

func validateNumber(input float64) bool {
	if input > 0 {
		return true
	} else {
		return false
	}
}

//Handle non numeric input...
// Helper functions for all the functions
// Concurrency on deposit - increase time for processing
