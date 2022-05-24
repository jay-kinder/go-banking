package database

import (
	"fmt"
	"os"
)

func AddUser(name string, email string) {
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

func GetBalance(email string) (User, error) {
	var usr User
	row := db.QueryRow("SELECT balance FROM users WHERE email=?", email)
	if err := row.Scan(&usr.ID, &usr.Email, &usr.Balance, &usr.Name); err != nil {
		return usr, err
	} else {
		return usr, nil
	}
}