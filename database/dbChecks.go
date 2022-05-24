package database

type User struct {
    ID     int64
    Email string
	Balance int
	Name string
}

func CurrentUserCheck(email string) (User, error) {
	var usr User
	row := db.QueryRow("SELECT * FROM users WHERE email=?", email)
	if err := row.Scan(&usr.ID, &usr.Email, &usr.Balance, &usr.Name); err != nil { 	
		return usr, err
	} else {
		return usr, nil
	}
}

func BalanceCheck(email string) (float64, error) {
	var usr User
	row := db.QueryRow("SELECT balance FROM users WHERE email=?")
	if err := row.Scan(&usr.ID, &usr.Email, &usr.Balance, &usr.Name); err != nil {
		return float64(usr.Balance), err
	} else {
		return float64(usr.Balance), nil
	}
}