package Managers

import (
	"database/sql"
	Mydb "github.com/go-sql-driver/mysql"
	"log"
)

type DBController struct {
	DB *sql.DB
}

func NewDBController() *DBController {
	conf := Mydb.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "db",
		DBName:               "data",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &DBController{DB: db}

}
