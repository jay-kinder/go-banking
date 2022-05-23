package database

import (
	"banking-app/helper"
	"fmt"
)

type User struct {
    ID     int64
    Email string
	Balance int
}

func CurrentUserCheck(email string, name string) {
	var usr User
	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)
	if err := row.Scan(&usr.Email); err != nil { 
		helper.SignUp(email, name)
	} else {
		fmt.Printf("\nWelcome Back %v!\n", name)
		helper.HomePage(name)
	}
}