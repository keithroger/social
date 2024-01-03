// Handlers for managing following relationships
package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateFollowerReq struct {
	FollowerId string `json:"followerId"`
	FolloweeId string `json:"followeeId"`
}

// createUserHandler creates a new user.
func createFollowerHandler(w http.ResponseWriter, req *http.Request) {
	// Limit request size
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var createFollowerReq CreateFollowerReq

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&createFollowerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert row into db and get generated user id
	query := "INSERT INTO followers(follower_id, followee_id) VALUES($1, $2)"
	_, err = db.Exec(query, createFollowerReq.FollowerId, createFollowerReq.FolloweeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetFollowersResp struct {
	UserIds   []string `json:"userIds"`
	Usernames []string `json:"usernames"`
}

// getUserHandler gets a user
func getFollowersHandler(w http.ResponseWriter, req *http.Request) {
	// Get query queryParams
	queryParams := mux.Vars(req)

	// Get size query param and convert to string
	size, err := strconv.Atoi(queryParams["size"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getFollowersResp := GetFollowersResp{
		UserIds:   make([]string, 0, size),
		Usernames: make([]string, 0, size),
	}

	// If following param is set to true, show who the user is following
	var query string
	if queryParams["following"] == "true" {
		query = `SELECT users.user_id, users.username FROM followers LEFT JOIN users
			ON followers.followee_id = users.user_id WHERE follower_id = $1 AND followee_id > $2 ORDER BY user_id
			LIMIT $3`
	} else {
		query = `SELECT users.user_id, users.username FROM followers LEFT JOIN users
			ON followers.follower_id = users.user_id WHERE followee_id = $1 AND follower_id > $2 ORDER BY user_id
			LIMIT $3`
	}

	// Query Database
	rows, err := db.Query(query, queryParams["id"], queryParams["last-id"], size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read rows
	var userId, username string
	for rows.Next() {
		err = rows.Scan(&userId, &username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Add query row to returned result
		getFollowersResp.UserIds = append(getFollowersResp.UserIds, userId)
		getFollowersResp.Usernames = append(getFollowersResp.Usernames, username)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getFollowersResp)
}

type DeleteFollowerReq struct {
	FollowerId string `json:"followerId"`
	FolloweeId string `json:"followeeId"`
}

// deleteFollowerHandler deletes a follower relatioship
func deleteFollowerHandler(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var deleteFollowReq DeleteFollowerReq

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&deleteFollowReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete row from db
	query := "DELETE FROM followers WHERE follower_id = $1 and followee_id = $2"
	_, err = db.Exec(query, deleteFollowReq.FollowerId, deleteFollowReq.FolloweeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
