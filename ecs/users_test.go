package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

// initData resets and initializes database
func initDbData(db *sql.DB, t *testing.T) {
	// Reset all records
	// _, err := db.Exec("")
	_, err := db.Exec("DELETE FROM users; ALTER SEQUENCE users_user_id_seq RESTART")
	if err != nil {
		t.Fatalf("failed to reset database: %v", err)
	}

	// Read sample data from file
	queryBytes, err := os.ReadFile("testdata/sample_data.sql")
	if err != nil {
		t.Fatalf("failed to read sample_data.sql: %v", err)
	}

	// Add sample data to database
	_, err = db.Exec(string(queryBytes))
	if err != nil {
		t.Fatalf("failed to apply sample data to database: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	db, err := initDbConnection()
	if err != nil {
		t.Fatalf("initial db connection failed: %v", err)
	}
	defer db.Close()

	initDbData(db, t)
	if err != nil {
		t.Errorf("failed to connect to postgres: %v", err)
	}

	testCases := []struct {
		name               string
		method             string
		path               string
		createUserReq      CreateUserReq
		expectedStatusCode int
	}{
		{
			name:   "CreateUser",
			method: http.MethodPost,
			path:   "users",
			createUserReq: CreateUserReq{
				Username:    "John123",
				ProfileName: "Johnny",
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Marshal request to json
			jsonReq, err := json.Marshal(testCase.createUserReq)
			if err != nil {
				t.Fatalf("Failed to marshal create user request: %v", err)
			}

			t.Logf("json request: %v", string(jsonReq))

			// Create test request
			body := bytes.NewReader(jsonReq)
			req, err := http.NewRequest(testCase.method, testCase.path, body)
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}

			// Record response
			respRecorder := httptest.NewRecorder()
			handler := createUserHandler(db)
			handler.ServeHTTP(respRecorder, req)

			// Check status code
			if respRecorder.Code != testCase.expectedStatusCode {
				t.Errorf("got: %v, want: %v", respRecorder.Code, testCase.expectedStatusCode)
			}

			// Continue to check for valid body if status is 200
			if testCase.expectedStatusCode != http.StatusOK {
				return
			}

			// Decode json response
			var createUserResp CreateUserResp
			err = json.NewDecoder(respRecorder.Body).Decode(&createUserResp)
			if err != nil {
				t.Fatalf("failed to decode json response: %v", err)
			}

			t.Logf("json response: %v", createUserResp)
		})
	}
}

func TestGetUser(t *testing.T) {
	db, err := initDbConnection()
	if err != nil {
		t.Fatalf("initial db connection failed: %v", err)
	}
	defer db.Close()

	initDbData(db, t)
	if err != nil {
		t.Errorf("failed to connect to postgres: %v", err)
	}

	testCases := []struct {
		name                string
		method              string
		path                string
		pathVars            map[string]string
		expectedGetUserResp GetUserResp
		expectedStatusCode  int
	}{
		{
			name:     "GetUser",
			method:   http.MethodGet,
			path:     "users/3",
			pathVars: map[string]string{"id": "3"},
			expectedGetUserResp: GetUserResp{
				Username:    "smith",
				ProfileName: "Elizabeth Smith",
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create test request
			req, err := http.NewRequest(testCase.method, testCase.path, nil)
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}

			// Set path variables
			req = mux.SetURLVars(req, testCase.pathVars)

			// Record response
			respRecorder := httptest.NewRecorder()
			handler := getUserHandler(db)
			handler.ServeHTTP(respRecorder, req)

			// Check status code
			if respRecorder.Code != testCase.expectedStatusCode {
				t.Errorf("got: %v, want: %v", respRecorder.Code, testCase.expectedStatusCode)
			}

			// Continue to check for valid body if status is 200
			if testCase.expectedStatusCode != http.StatusOK {
				return
			}

			// Decode json response
			var getUserResp GetUserResp
			err = json.NewDecoder(respRecorder.Body).Decode(&getUserResp)
			if err != nil {
				t.Fatalf("failed to decode json response: %v", err)
			}

			// Check response
			if getUserResp.Username != testCase.expectedGetUserResp.Username {
				t.Errorf("incorrect username got: %v, want: %v",
					getUserResp.Username, testCase.expectedGetUserResp.Username)
			}
			if getUserResp.ProfileName != testCase.expectedGetUserResp.ProfileName {
				t.Errorf("incorrect username got: %v, want: %v",
					getUserResp.ProfileName, testCase.expectedGetUserResp.ProfileName)
			}

			t.Logf("json response: %v", getUserResp)
		})
	}
}
