package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwskme/goEcommerceBackend/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid.", func(t *testing.T) {
		payload := types.RegisterUserPayLoad{
			FirstName: "user",
			LastName:  "joe",
			Email:     "",
			Password:  "123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handlRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}