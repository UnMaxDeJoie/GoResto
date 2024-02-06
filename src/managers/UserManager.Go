package Managers

import (
	"database/sql"
	"fmt"
	"github.com/UnMaxDeJoie/GoResto/entities"
	"log"
)

func GetUserById(Mydb *DBController, uid int) (entities.User, error) {
	query := "SELECT id, username, email, permissions FROM users WHERE id= ?"

	var user entities.User

	err := Mydb.DB.QueryRow(query, uid).Scan(&user.ID, &user.Username, &user.Email, &user.Permission)
	if err != nil {

		return entities.User{}, err
	}

	return user, nil
}

func DeleteUser(Mydb *DBController, uid int) {
	query := "DELETE FROM users WHERE id= ?"

	result, err := Mydb.DB.Exec(query, uid)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {

		log.Fatal(err)
	}

	if rowsAffected == 0 {

		log.Fatal(fmt.Errorf("aucun utilisateur trouvé avec l'ID %d", uid))
	}
}

func CreateUser(Mydb *sql.DB, Name string, Password string, email string, permission uint8) (entities.User, error) {

	user := entities.User{
		Username:   Name,
		PwHash:     Password,
		Email:      email,
		Permission: permission,
	}

	_, err := Mydb.Exec("INSERT INTO users (username, pw_hash, email, permissions) VALUES (?, ?, ?, ?)", user.Username, user.PwHash, user.Email, user.Permission)
	if err != nil {
		log.Printf("Erreur lors de la création de l'utilisateur : %v", err)
		return entities.User{}, fmt.Errorf("could not create user: %v", err)
	}

	return user, err
}

func rank(db *DBController, uid int) (uint8, error) {
	user, err := GetUserById(db, uid)
	if err != nil {
		return 0, fmt.Errorf("could not find user: %v", err)
	}

	return user.Permission, nil
}
