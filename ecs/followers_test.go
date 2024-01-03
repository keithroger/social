package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateFollowerHandler(t *testing.T) {
	err := initDbConnection()
	if err != nil {
		t.Fatalf("initial db connection failed: %v", err)
	}
	defer db.Close()

	initDbData(t)
	if err != nil {
		t.Errorf("failed to connect to postgres: %v", err)
	}

	testCases := []struct {
		name               string
		method             string
		path               string
		createFollowerReq  CreateFollowerReq
		expectedStatusCode int
	}{
		{
			name:   "Create Follower",
			method: http.MethodPost,
			path:   "followers/",
			createFollowerReq: CreateFollowerReq{
				FollowerId: "1",
				FolloweeId: "2",
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Marshal request to json
			jsonReq, err := json.Marshal(testCase.createFollowerReq)
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
			createFollowerHandler(respRecorder, req)

			// Check status code
			if respRecorder.Code != testCase.expectedStatusCode {
				t.Errorf("got: %v, want: %v", respRecorder.Code, testCase.expectedStatusCode)
			}
		})
	}
}

func TestHandlers(t *testing.T) {
	err := initDbConnection()
	if err != nil {
		t.Fatalf("initial db connection failed: %v", err)
	}
	defer db.Close()

	initDbData(t)
	if err != nil {
		t.Errorf("failed to connect to postgres: %v", err)
	}

	testCases := []testCase{
		{
			name:        "Create Follower",
			method:      http.MethodPost,
			path:        "followers/",
			queryParams: nil,
			jsonReq: CreateFollowerReq{
				FollowerId: "1",
				FolloweeId: "2",
			},
			jsonResp:       nil,
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        createFollowerHandler,
		}, {
			name:   "Get Following",
			method: http.MethodGet,
			path:   "followers/4",
			queryParams: map[string]string{
				"id":        "4",
				"last-id":   "2",
				"size":      "5",
				"following": "true",
			},
			jsonReq: nil,
			jsonResp: GetFollowersResp{
				UserIds:   []string{},
				Usernames: []string{},
			},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        getFollowersHandler,
		}, {
			name:   "Get Followers",
			method: http.MethodGet,
			path:   "followers/4",
			queryParams: map[string]string{
				"id":      "1",
				"last-id": "0",
				"size":    "5",
			},
			jsonReq: nil,
			jsonResp: GetFollowersResp{
				UserIds:   []string{},
				Usernames: []string{},
			},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        getFollowersHandler,
		}, {
			name:   "Delete Follower",
			method: http.MethodDelete,
			path:   "followers/",
			jsonReq: DeleteFollowerReq{
				FollowerId: "1",
				FolloweeId: "4",
			},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        deleteFollowerHandler,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handlerTester(t, db, tc)
		})
	}
}

type testCase struct {
	name           string
	method         string
	path           string
	queryParams    map[string]string
	jsonReq        interface{}
	jsonResp       interface{} // Expected json result
	expectedBody   string      // Expected body if no expected json
	expectedStatus int
	handler        http.HandlerFunc
}

func handlerTester(t *testing.T, db *sql.DB, tc testCase) {
	// Marshal request to json
	jsonReq, err := json.Marshal(tc.jsonReq)
	if err != nil {
		t.Fatalf("Failed to marshal create user request: %v", err)
	}

	t.Logf("request json: %v", string(jsonReq))
	t.Logf("request query params: %v", tc.queryParams)

	// Create test request
	body := bytes.NewReader(jsonReq)
	req, err := http.NewRequest(tc.method, tc.path, body)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
	}

	// Set query parameters
	req = mux.SetURLVars(req, tc.queryParams)

	// Record response
	respRecorder := httptest.NewRecorder()
	tc.handler.ServeHTTP(respRecorder, req)

	t.Logf("response body: %v", respRecorder.Body.String())

	// Check status code
	if respRecorder.Code != tc.expectedStatus {
		t.Errorf("got: %v, want: %v", respRecorder.Code, tc.expectedStatus)
	}

	// Continue to check for valid body if status is 200
	if tc.expectedStatus != http.StatusOK {
		return
	}

	if tc.jsonResp == nil {
		return
	}

	// Decode json response
	respType := reflect.TypeOf(tc.jsonResp)
	respInterface := reflect.New(respType).Interface()
	err = json.NewDecoder(respRecorder.Body).Decode(&respInterface)
	if err != nil {
		t.Fatalf("failed to decode json response: %v", err)
	}

	t.Logf("response json: %v", respInterface)

	// TODO check response
}
