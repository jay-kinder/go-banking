package database

import (
	"fmt"
	"os"
)

func AddUser(email string, name string) {
	result, err := db.Exec("INSERT INTO users (email, balance, name) VALUES (?, 0, ?)", email, name)
	if err == nil {
		result.LastInsertId()
		fmt.Printf("Success! Your account is set up with the email: %v \n", email)
	} else {
		fmt.Printf("Error adding user to DB. Erorr: %v", err)
		fmt.Println()
		os.Exit(1)
	}
}