package main

import (
	"banking-app/database"
	"banking-app/helper"
	"fmt"
)

func main() {
	var name string
	var email string	

	database.DbConnection()

	fmt.Println()
	fmt.Println("------------------------------------")

	fmt.Println("What is your Email Address?")
	fmt.Scan(&email)

	// Checking DB for User
	usr, err := database.CurrentUserCheck(email) 
	if err != nil {
		helper.SignUp(email)
	} else {
		fmt.Printf("\nWelcome Back %v!\n", usr.Name)
		helper.HomePage(name, email)
	}
}
