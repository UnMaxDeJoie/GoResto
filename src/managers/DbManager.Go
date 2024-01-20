package controller

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DBController struct {
	DB *sql.DB
}

func NewDBController() *DBController {
	conf := mysql.Config{
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
