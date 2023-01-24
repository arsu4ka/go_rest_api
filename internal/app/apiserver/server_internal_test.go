package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arsu4ka/go_rest_api/internal/app/model"
	"github.com/arsu4ka/go_rest_api/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	server := NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user3@gmail.com",
				"password": "user3pass",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "data",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid password and email",
			payload: map[string]string{
				"email":    "cool",
				"password": "hi user",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			data := &bytes.Buffer{}
			json.NewEncoder(data).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", data)
			server.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}

}

func TestServer_HandleSesssionsCreate(t *testing.T) {
	userTest := model.TestUser(t)
	server := NewServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	server.store.User().Create(userTest)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    userTest.Email,
				"password": userTest.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": userTest.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    userTest.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			data := &bytes.Buffer{}
			json.NewEncoder(data).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", data)
			server.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}

}
