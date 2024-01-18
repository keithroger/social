package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreatePostReq struct {
	UserId string `json:"userId"`
	Body   string `json:"body"`
}

type CreatePostResp struct {
	PostId string `json:"postId"`
}

// createPostHandler creates a new user.
func createPostHandler(w http.ResponseWriter, req *http.Request) {
	// Limit request size
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var createPostReq CreatePostReq
	var createPostResp CreatePostResp

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&createPostReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert row into db and get generated user id
	query := "INSERT INTO posts(user_id, body) VALUES($1, $2) RETURNING post_id"
	err = db.QueryRow(query, createPostReq.UserId, createPostReq.Body).
		Scan(&createPostResp.PostId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createPostResp)
}

type GetPostsResp struct {
	PostId string `json:"postId"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
}

// getPostsHandler gets posts made by a user
func getPostsHandler(w http.ResponseWriter, req *http.Request) {
	// Get query parameters
	id := req.URL.Query().Get("id")
	lastId := req.URL.Query().Get("last-id")
	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getPostsResp := make([]GetPostsResp, 0, size)

	// If following param is set to true, show who the user is following
	query := `SELECT posts.post_id, posts.body, posts.likes
		FROM posts
		WHERE user_id = $1 AND post_id > $2
		ORDER BY post_id DESC
		LIMIT $3`

	// Query Database
	rows, err := db.Query(query, id, lastId, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read rows
	for rows.Next() {
		currPostResp := GetPostsResp{}
		err = rows.Scan(&currPostResp.PostId, &currPostResp.Body, &currPostResp.Likes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Add query row to returned result
		getPostsResp = append(getPostsResp, currPostResp)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getPostsResp)
}

type UpdatePostReq struct {
	Body string `json:"body"`
}

// updatePostHandler updates a post
func updatePostHandler(w http.ResponseWriter, req *http.Request) {
	// Get query queryParams
	queryParams := mux.Vars(req)

	// Limit request size
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var updatePostReq UpdatePostReq

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&updatePostReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert row into db and get generated user id
	query := "UPDATE posts SET body = $1 WHERE post_id=$2"
	_, err = db.Exec(query, updatePostReq.Body, queryParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteUserHandler deletes a post
func deletePostHandler(w http.ResponseWriter, req *http.Request) {
	// Get query queryParams
	queryParams := mux.Vars(req)

	// Delete row from db
	query := "DELETE FROM posts WHERE post_id = $1"
	_, err := db.Exec(query, queryParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
