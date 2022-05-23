package database

import "fmt"

type User struct {
    ID     int64
    Email string
	Balance int
}

func CurrentUserCheck(email string) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return nil, fmt.Errorf("CurrentUserCheck %q: %v", email, err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Email); err != nil {
			return nil, fmt.Errorf("CurrentUserCheck %q: %v", email, err)
		}
		users = append(users, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("CurrentUserCheck %q: %v", email, err)
	}
	return users, nil
}