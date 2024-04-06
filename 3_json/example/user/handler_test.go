package user_test

import (
	"github.com/d-arken/workshop-go/tree/main/3_json/example/router"
	"github.com/stretchr/testify/assert"
	"strings"
)

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a new Gin router
	r := router.Setup()

	// Define a POST request with JSON payload
	jsonStr := `{"name":"Alice","age":25}`
	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expected := `{"name":"Alice","age":25}`

	assert.Equal(t, w.Body.String(), expected)
	assert.Equal(t, w.Code, http.StatusOK)

}

func TestCreateUserShouldFailValidation(t *testing.T) {
	// Create a new Gin router
	r := router.Setup()

	// Define a POST request with JSON payload
	jsonStr := `{"name":"Alice"}`
	req, err := http.NewRequest("POST", "/", strings.NewReader(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}
