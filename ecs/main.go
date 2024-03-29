package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	// Postgres Variables
	pgReadWriteEndpoint = os.Getenv("POSTGRES_READ_WRITE_ENDPOINT")
	pgReadEndpoint      = os.Getenv("POSTGRES_READ_ENDPOINT")
	pgDB                = os.Getenv("POSTGRES_DB")
	pgUser              = os.Getenv("POSTGRES_USER")
	pgPassword          = os.Getenv("POSTGRES_PASSWORD")

	// DB connection to be shared by handlers
	db *sql.DB
)

func main() {
	fmt.Println("Initializing container")

	// Initialize postgres connection
	err := initDbConnection()
	if err != nil {
		log.Fatalf("initial db connection failed: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to database")

	r := mux.NewRouter()

	// /users api handlers
	r.HandleFunc("/users", createUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", updateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods(http.MethodDelete)

	// /posts api handlers
	r.HandleFunc("/posts", getPostsHandler).Methods(http.MethodGet).Queries("id", "", "last-id", "", "size", "")
	r.HandleFunc("/posts", createPostHandler).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id:[0-9]+}", updatePostHandler).Methods(http.MethodPut)
	r.HandleFunc("/posts/{id:[0-9]+}", deletePostHandler).Methods(http.MethodDelete)

	// /comments api handlers
	r.HandleFunc("/comments", getCommentsHandler).Methods(http.MethodGet).Queries("id", "", "last-id", "", "size", "")
	r.HandleFunc("/comments", createCommentHandler).Methods(http.MethodPost)
	r.HandleFunc("/comments/{id:[0-9]+}", updateCommentHandler).Methods(http.MethodPut)
	r.HandleFunc("/comments/{id:[0-9]+}", deleteCommentHandler).Methods(http.MethodDelete)

	// /follower api handlers
	r.HandleFunc("/followers", getFollowersHandler).Methods(http.MethodGet).Queries("id", "", "last-id", "", "size", "", "following", "")
	r.HandleFunc("/followers", createFollowerHandler).Methods(http.MethodPost)
	r.HandleFunc("/followers", deleteFollowerHandler).Methods(http.MethodDelete)

	// TODO create feed api handler
	// /feed api handler

	// health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %v\n", err)
	}
}

// initDbConnection creates a new connection to the postgres database.
// Should call db.Close() when finished.
func initDbConnection() error {
	// Apply connection configuration
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		pgReadWriteEndpoint, pgUser, pgPassword, pgDB)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to configure connection: %v", err)
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping connection: %v", err)
	}

	return nil
}
