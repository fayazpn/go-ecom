package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fayazpn/ecom/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStruct{}
	handler := NewHandler(userStore)

	t.Run("should fail if the payload is valid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test name",
			LastName:  "test last name",
			Email:     "invalid",
			Password:  "password123",
		}
		marshell, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshell))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail if the payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "test name",
			LastName:  "test last name",
			Email:     "email@gmail.com",
			Password:  "password123",
		}
		marshell, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshell))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStruct struct {
}

func (m *mockUserStruct) GetUserByEmail(email string) (*types.User, error) {
	// return &types.User{}, nil
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStruct) GetUserById(id int) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStruct) CreateUser(types.User) error {
	return nil
}
