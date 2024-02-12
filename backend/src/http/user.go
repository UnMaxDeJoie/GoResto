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
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

func GetUserByIdEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIDString := chi.URLParam(r, "userID")

		userID, err := strconv.Atoi(userIDString)
		if err != nil {
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
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}
		managers.DeleteUser(db, userID)
		w.Write([]byte("User deleted"))
	}
}
