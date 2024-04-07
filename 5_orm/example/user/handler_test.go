package user_test

import (
	"github.com/d-arken/workshop-go/5_orm/router"
	"github.com/d-arken/workshop-go/5_orm/user"
	mock_user "github.com/d-arken/workshop-go/5_orm/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a new Gin router
	ctrl := gomock.NewController(t)
	mockSvc := mock_user.NewMockServiceInterface(ctrl)

	mockSvc.EXPECT().Create(gomock.Any()).Times(1)
	handler := user.NewHandler(mockSvc)

	r := router.Setup(handler)

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
	ctrl := gomock.NewController(t)
	mockSvc := mock_user.NewMockServiceInterface(ctrl)
	handler := user.NewHandler(mockSvc)
	r := router.Setup(handler)

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
