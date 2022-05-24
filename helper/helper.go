package helper

import (
	"banking-app/database"
	"fmt"
	"os"
	"time"
)

// Validate Number
func ValidateNumber(input float64) bool {
	return input > 0
}

// Main Menu Choice
func HomePage(name string, email string){
	var menuInput uint

	fmt.Println("------------------------------------")
	fmt.Println("\nPlease select from the following services:\n\n1. Deposit Moneys\n2. Withdraw Moneys\n3. See Account Balance\n4. Exit")
	fmt.Println()
	fmt.Println("Choice: ")
	fmt.Scan(&menuInput)

	for {
		switch menuInput {
			case 1:
				depositAmount := DepositMoney()
				go ProcessDeposit(depositAmount, email)	
				AfterSelection(name, email)
			case 2:
				WithdrawMoney(email)				
				AfterSelection(name, email)
			case 3:
				GetBalance(email)
				AfterSelection(name, email)
			case 4:
				os.Exit(0)
			default:
				fmt.Println("You need to enter a number between 1 and 4...IS THAT NOT CLEAR??!!")
		}
	}
}

// Deposit
func DepositMoney() float64 {
	var depositAmount float64
	for {
		fmt.Println()
		fmt.Println("Great! Let's put some money into your account!")
		fmt.Println("How much would you like to put in (JGs)?")
		fmt.Scan(&depositAmount)
		if !ValidateNumber(depositAmount) {
			fmt.Println("You need to enter a positive number, you silly sausage!")
			continue
		}
		fmt.Println()
		fmt.Println("This will take a couple of seconds to process...")
		fmt.Println("......................................")
		return depositAmount
	}
}

// Process Deposit
func ProcessDeposit(depositAmount float64, email string) float64 {	
	time.Sleep(20 * time.Second)
	fmt.Println()
	database.DepositMoney(depositAmount, email)
	fmt.Println("**********************************************************************")
	fmt.Printf("Cool! Your deposit of %v JayGolds into your Account is now available\n", depositAmount)
	fmt.Println("**********************************************************************")
	return depositAmount
}

// Withdraw
func WithdrawMoney(email string) {
	var withdrawAmount float64
	accountBalance, _ := database.GetBalance(email)
	for {
		fmt.Println()
		fmt.Println("Awesome! Let's take some money out of your account.")
		fmt.Println()
		fmt.Println("How much would you like to take out (JGs)?")
		fmt.Scan(&withdrawAmount)
		if !ValidateNumber(withdrawAmount) || withdrawAmount > accountBalance {
			fmt.Println("You have either tried to enter a negative number, or you are trying to take out more money than you have...tut tut.")
			continue
		}
		fmt.Println()
		fmt.Println("This will take a couple of seconds to process...")
		time.Sleep(3 * time.Second)
		fmt.Println("......................................")
		time.Sleep(3 * time.Second)
		database.WithdrawMoney(withdrawAmount, email)
		fmt.Printf("Sweet! You've just taken %v JayGolds out of your Account\n", withdrawAmount)
	}
}

// Balance
func GetBalance(email string) {
	fmt.Println()
	fmt.Println("Don't be too disappointed...here is your bank balance!")
	fmt.Println("......................................")
	time.Sleep(3 * time.Second)
	accountBalance, err := database.GetBalance(email)
	if err != nil {
		fmt.Printf("Balance: %v JayGolds\n", accountBalance)
	} else {
		fmt.Printf("Could not process your request. Error: %v", err)
		os.Exit(1)
	}	
}

// AfterSelection (go back to select another service, or exit)
func AfterSelection(name string, email string) {
	var inputSelection uint
	for {
		fmt.Println()
		fmt.Printf("Would you like to use another Service?:\n\n1.Yes\n2. No\n")
		fmt.Scan(&inputSelection)

		switch inputSelection {
			case 1:
				HomePage(name, email)
			case 2:
				os.Exit(0)
			default:
				fmt.Println("You need to enter either 1 or 2...DUH!!")
				continue
		}
	}
}

// Sign Up
func SignUp(email string) {
	var name string
	fmt.Println("\nYou do not have an account with us!")
	fmt.Println("Please let us know what we should call you!")
	fmt.Scan(&name)
	fmt.Println("\nYou will now be enrolled with the Bank of Jay")
	fmt.Println("..........................................................")
	time.Sleep(5 * time.Second)

	// Insert User into the Database
	database.AddUser(name, email)

	fmt.Printf("\nCheers %v!\n", name)
	HomePage(name, email)
}