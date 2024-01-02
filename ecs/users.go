// Handlers for managing users
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateUserReq struct {
	Username    string `json:"username"`
	ProfileName string `json:"profileName"`
}

type CreateUserResp struct {
	UserId string `json:"userId"`
}

// createUserHandler creates a new user.
func createUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Limit request size
		req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

		var createUserReq CreateUserReq
		var createUserResp CreateUserResp

		// Parse json request
		err := json.NewDecoder(req.Body).Decode(&createUserReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert row into db and get generated user id
		query := "INSERT INTO users(username, profile_name) VALUES($1, $2) RETURNING user_id"
		err = db.QueryRow(query, createUserReq.Username, createUserReq.ProfileName).
			Scan(&createUserResp.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createUserResp)
	}
}

type GetUserResp struct {
	Username    string `json:"username"`
	ProfileName string `json:"profileName"`
}

// getUserHandler gets a user
func getUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		var getUserResp GetUserResp

		query := "SELECT username, profile_name FROM users WHERE user_id = $1"

		err := db.QueryRow(query, vars["id"]).
			Scan(&getUserResp.Username, &getUserResp.ProfileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Return response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getUserResp)
	}
}

// updateUserHandler updates a user
func updateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	}
}

// deleteUserHandler deletes a user
func deleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	}
}
