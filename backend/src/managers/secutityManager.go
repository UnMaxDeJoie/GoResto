package managers

import (
	"NewGoResto/src/entities"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Login(Mydb *sql.DB, password, email string) (int, error) {
	var user entities.User

	// Récupération du hash du mot de passe depuis la base de données
	err := Mydb.QueryRow("SELECT id, pw_hash, email FROM users WHERE email = ?", email).Scan(&user.ID, &user.PwHash, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("invalid email or password")
		}
		log.Printf("Erreur lors de la récupération de l'utilisateur : %v", err)
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PwHash), []byte(password))
	if err != nil {
		// bcrypt.CompareHashAndPassword renvoie une erreur si les mots de passe ne correspondent pas
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return 0, errors.New("invalid email or password")
		}
		// Autre erreur bcrypt
		log.Printf("Erreur lors de la comparaison des mots de passe : %v", err)
		return 0, err
	}

	// tokenString, err := CreateToken(user.Email, user.ID)
	// if err != nil {
	// 	log.Printf("Erreur lors de la création du token : %v", err)
	// 	return 0, "", err
	// }

	return user.ID, nil
}
