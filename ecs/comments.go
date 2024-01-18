package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateCommentReq struct {
	UserId string `json:"userId"`
	PostId string `json:"postId"`
	Body   string `json:"body"`
}

type CreateCommentResp struct {
	CommentId string `json:"commentId"`
}

// createCommentHandler creates a new comment.
func createCommentHandler(w http.ResponseWriter, req *http.Request) {
	// Limit request size
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var createCommentReq CreateCommentReq
	var createCommentResp CreateCommentResp

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&createCommentReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert row into db and get generated user id
	query := "INSERT INTO comments(post_id, user_id, body) VALUES($1, $2, $3) RETURNING comment_id"
	err = db.QueryRow(query, createCommentReq.PostId, createCommentReq.UserId, createCommentReq.Body).
		Scan(&createCommentResp.CommentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createCommentResp)
}

type GetCommentResp struct {
	PostId   string `json:"postId"`
	Username string `json:"username"`
	Body     string `json:"body"`
}

// getCommentsHandler gets a comments associated with a post id
func getCommentsHandler(w http.ResponseWriter, req *http.Request) {
	// Get query parameters
	id := req.URL.Query().Get("id")
	lastId := req.URL.Query().Get("last-id")
	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	getCommentsResp := make([]GetCommentResp, 0, size)

	// If following param is set to true, show who the user is following
	query := `SELECT comments.post_id, comments.body, users.username
		FROM comments
		LEFT JOIN users
		ON comments.user_id = users.user_id
		WHERE comments.post_id = $1 AND comments.comment_id > $2
		ORDER BY comments.comment_id DESC
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
		currCommentResp := GetCommentResp{}
		err = rows.Scan(&currCommentResp.PostId, &currCommentResp.Body, &currCommentResp.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Add query row to returned result
		getCommentsResp = append(getCommentsResp, currCommentResp)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCommentsResp)
}

type UpdateCommentReq struct {
	Body string `json:"body"`
}

// updateCommentHandler updates a comment
func updateCommentHandler(w http.ResponseWriter, req *http.Request) {
	// Get query queryParams
	queryParams := mux.Vars(req)

	// Limit request size
	req.Body = http.MaxBytesReader(w, req.Body, 1<<10) // 1KB

	var updateCommentReq UpdateCommentReq

	// Parse json request
	err := json.NewDecoder(req.Body).Decode(&updateCommentReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert row into db and get generated user id
	query := "UPDATE comments SET body = $1 WHERE comment_id=$2"
	_, err = db.Exec(query, updateCommentReq.Body, queryParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// deleteCommentHandler deletes a comment
func deleteCommentHandler(w http.ResponseWriter, req *http.Request) {
	// Get query queryParams
	queryParams := mux.Vars(req)

	// Delete row from db
	query := "DELETE FROM comments WHERE comment_id = $1"
	_, err := db.Exec(query, queryParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
