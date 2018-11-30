package zendesk

import (
	"net/http"
	"testing"
)

func TestGetUsers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "users.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	users, _, err := client.GetUsers()
	if err != nil {
		t.Fatalf("Failed to get users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected length of triggers is 2, but got %d", len(users))
	}
}

func TestCreateUser(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.CreateUser(User{
		Email: "test@example.com",
		Name:  "testuser",
	})
	if err != nil {
		t.Fatalf("Failed to get valid response: %s", err)
	}
	if user.ID == 0 {
		t.Fatal("Failed to create user")
	}
}
