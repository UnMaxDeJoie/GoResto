package managers

import (
	"GoResto/entities"
	"database/sql"
	"errors"
	"fmt"
)

var secretKey = []byte("secret-key")

func login(Mydb *DBController, pwHash string, email string) (string, error) {
	var user entities.User
	err := Mydb.DB.QueryRow("SELECT id, pw_hash, email FROM users WHERE email = ?", user.Email).Scan(&user.Email, &user.PwHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}
	if pwHash == user.PwHash {
		tokenString, err := createToken(email)
		if err != nil {
			fmt.Errorf("No username found")
		}
		return tokenString, err
	} else {
		return "", fmt.Errorf("password not match")
	}
}
