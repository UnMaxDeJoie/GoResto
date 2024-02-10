package managers

import (
	"database/sql"
	Mydb "github.com/go-sql-driver/mysql"
	"log"
)

type DBController struct {
	DB *sql.DB
}

var G_Mydb *DBController

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

	G_Mydb = &DBController{DB: db}
}

func GetDBController() *sql.DB {
	if G_Mydb == nil || G_Mydb.DB == nil {
		NewDBController() // Initialise la connexion si elle n'existe pas.
	} else {
		if err := G_Mydb.DB.Ping(); err != nil {
			log.Println("La connexion à la base de données a été perdue. Tentative de reconnexion.")
			NewDBController() // Réinitialise la connexion en cas d'échec du ping.
		}
	}

	return G_Mydb.DB
}

/*regarde comment relancer la connexion si ca échoue la fonction ping existe elle renvoie un nil*/
