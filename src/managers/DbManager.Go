package Managers

import (
	"database/sql"
	Mydb "github.com/go-sql-driver/mysql"
	"log"
)

type DBController struct {
	DB *sql.DB
}

var G_Mydb DBController

func NewDBController() {
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

	G_Mydb = DBController{DB: db}
}

func GetDBController() DBController {
	return G_Mydb
}

/*regarde comment relancer la connexion si ca échoue la fonction ping existe elle renvoie un nil*/
