package Managers

import (
	_ "github.com/go-sql-driver/mysql" // Driver MySQL
	"log"
)

func GetUserById(db *DBController, uid int) {
	query := "SELECT id, username, email, permissions FROM users Where id= ?"

	var (
		id          int
		username    string
		email       string
		permissions uint8
	)
	err := db.DB.QueryRow(query, uid).Scan(&id, &username, &email, &permissions)
	if err != nil {
		log.Fatal(err)
	}

}
