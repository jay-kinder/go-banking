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
	fmt.Println("Welcome to the Go Banking Application!")
	fmt.Println("------------------------------------")

	fmt.Println("Firstly, what is your name?")
	fmt.Scan(&name)
	fmt.Println("Next, what is your email")
	fmt.Scan(&email)

	// Checking DB for User
	usr, err := database.CurrentUserCheck(email) 
	if err != nil {
		helper.SignUp(email, name)
	} else {
		fmt.Printf("\nWelcome Back %v!\n", usr.Name)
		helper.HomePage(name, email)
	}
}
