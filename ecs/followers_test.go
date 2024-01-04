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
		}, {
			name:   "Get Posts",
			method: http.MethodGet,
			path:   "posts/",
			queryParams: map[string]string{
				"id":      "1",
				"last-id": "0",
				"size":    "5",
			},
			jsonReq:        nil,
			jsonResp:       []GetPostsResp{},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        getPostsHandler,
		}, {
			name:        "Create Post",
			method:      http.MethodPost,
			path:        "posts/",
			queryParams: nil,
			jsonReq: CreatePostReq{
				UserId: "1",
				Body:   "This is a test post",
			},
			jsonResp:       CreatePostResp{},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        createPostHandler,
		}, {
			name:        "Update Posts",
			method:      http.MethodPut,
			path:        "posts/4",
			queryParams: map[string]string{"id": "4"},
			jsonReq: UpdatePostReq{
				Body: "This is an updated post",
			},
			jsonResp:       nil,
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        updatePostHandler,
		}, {
			name:           "Delete Post",
			method:         http.MethodDelete,
			path:           "posts/4",
			queryParams:    map[string]string{"id": "1"},
			jsonReq:        nil,
			jsonResp:       nil,
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        deletePostHandler,
		}, {
			name:           "Get Comments",
			method:         http.MethodGet,
			path:           "comments/",
			queryParams:    map[string]string{"id": "2", "last-id": "1", "size": "5"},
			jsonReq:        nil,
			jsonResp:       []GetCommentResp{},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        getCommentsHandler,
		}, {
			name:        "Create Comment",
			method:      http.MethodPost,
			path:        "comments/",
			queryParams: nil,
			jsonReq: CreateCommentReq{
				UserId: "5",
				PostId: "3",
				Body:   "This is a body comment",
			},
			jsonResp:       CreateCommentResp{},
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        createCommentHandler,
		}, {
			name:        "Update Comment",
			method:      http.MethodPut,
			path:        "comment/2",
			queryParams: map[string]string{"id": "1"},
			jsonReq: UpdateCommentReq{
				Body: "Updated comment",
			},
			jsonResp:       nil,
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        updateCommentHandler,
		}, {
			name:           "Delete Comment",
			method:         http.MethodDelete,
			path:           "comment/4",
			queryParams:    map[string]string{"id": "4"},
			jsonReq:        nil,
			jsonResp:       nil,
			expectedBody:   "",
			expectedStatus: http.StatusOK,
			handler:        deleteCommentHandler,
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
