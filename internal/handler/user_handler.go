package handler

import (
	"api/internal/model"
	"api/internal/repository"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/users", createUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")
}

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Failed to get users!", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		user, err := repository.GetUserByID(db, id)

		if err != nil {
			http.Error(w, "User not found!", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid input!", http.StatusBadRequest)
			return
		}
		if err := repository.CreateUser(db, u); err != nil {
			http.Error(w, "Failed to create user!", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User

		vars := mux.Vars(r)
		id := vars["id"]

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid input!", http.StatusBadRequest)
			return
		}

		if err := repository.UpdateUser(db, id, u); err != nil {
			http.Error(w, "Failed to update user!", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func deleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := repository.DeleteUser(db, id); err != nil {
			http.Error(w, "Failed to delete user!", http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted!"})
	}

}
