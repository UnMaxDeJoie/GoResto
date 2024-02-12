package http

import (
	"NewGoResto/src/entities"
	"NewGoResto/src/managers"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateUserEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userModel entities.User
		err := json.NewDecoder(r.Body).Decode(&userModel)
		if err != nil {
			fmt.Printf("Error when decoding user model: %v", err)
			http.Error(w, "Invalid user data", http.StatusBadRequest)
			return
		}
		user, err := managers.CreateUser(db, userModel.Username, userModel.PwHash, userModel.Email, userModel.Permission)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func LoginEndpoint(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var credentials struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        err := json.NewDecoder(r.Body).Decode(&credentials)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        userID, err := managers.Login(db, credentials.Email, credentials.Password)
        if err != nil {
            // Ici, ajustez le message d'erreur et le code d'erreur selon les besoins
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Vous pouvez également générer un token JWT ici si nécessaire

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(struct {
            UserID int `json:"user_id"`
        }{
            UserID: userID,
        })
    }
}


func GetUserByIdEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "userID")

		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}
		user, err := managers.GetUserById(db, userID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func DeleteUserEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "userID")

		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}
		managers.DeleteUser(db, userID)
		w.Write([]byte("User deleted"))
	}
}
