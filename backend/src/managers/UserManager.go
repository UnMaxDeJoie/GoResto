package managers

import (
	"NewGoResto/src/entities"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetUserById(Mydb *sql.DB, uid int) (entities.User, error) {
	query := "SELECT id, username, email, permissions FROM users WHERE id= ?"

	var user entities.User

	err := Mydb.QueryRow(query, uid).Scan(&user.ID, &user.Username, &user.Email, &user.Permission)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func DeleteUser(Mydb *sql.DB, uid int) {
	query := "DELETE FROM users WHERE id= ?"

	result, err := Mydb.Exec(query, uid)
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

func CreateUser(Mydb *sql.DB, Name, Password, Email string, Permission uint8) (entities.User, error) {
	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erreur lors du hashage du mot de passe : %v", err)
		return entities.User{}, fmt.Errorf("could not hash password: %v", err)
	}

	user := entities.User{
		Username:   Name,
		PwHash:     string(hashedPassword),
		Email:      Email,
		Permission: Permission,
	}

	_, err = Mydb.Exec("INSERT INTO users (username, pw_hash, email, permissions) VALUES (?, ?, ?, ?)", user.Username, user.PwHash, user.Email, user.Permission)
	if err != nil {
		log.Printf("Erreur lors de la création de l'utilisateur : %v", err)
		return entities.User{}, fmt.Errorf("could not create user: %v", err)
	}
	return user, nil
}

func rank(db *sql.DB, uid int) (uint8, error) {
	user, err := GetUserById(db, uid)
	if err != nil {
		return 0, fmt.Errorf("could not find user: %v", err)
	}

	return user.Permission, nil
}
