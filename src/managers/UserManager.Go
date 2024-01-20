package managers

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserById(int uid) {
	query := "SELECT id, username, email, permissions FROM users Where id= ?"

	var (
		id          int
		username    string
		email       string
		permissions uint8
	)
	err = db.QueryRow(query, uid).scan(&id, &username, &email, &permissions)
	if err != nil {
		log.Fatal(err)
	}
}
