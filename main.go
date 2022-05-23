package main

import (
	"banking-app/database"
	"banking-app/helper"
	"fmt"
	"log"
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
	} else if len(users) == 0 {
		//fmt.Printf("Users found: %v\n", users)
		helper.SignUp(email, name)
	} else {
		helper.HomePage(name)
	}	
}
