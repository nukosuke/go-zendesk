package zendesk

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestUserRoleText(t *testing.T) {
	for key := UserRoleEndUser; key <= UserRoleAdmin; key++ {
		if text := UserRoleText(key); text == "" {
			t.Fatalf("key=%d is undefined", key)
		}
	}
}

func TestGetUsers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "users.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	users, _, err := client.GetUsers(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected length of triggers is 2, but got %d", len(users))
	}
}

func TestGetUsersRolesEncodeCorrectly(t *testing.T) {
	expected := "role%5B%5D=admin&role%5B%5D=end-user"
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryString := r.URL.Query().Encode()
		if queryString != expected {
			t.Fatalf(`Did not get the expect query string: "%s". Was: "%s"`, expected, queryString)
		}
		w.Write(readFixture(filepath.Join(http.MethodGet, "users.json")))
	}))

	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	opts := UserListOptions{
		Roles: []string{
			"admin",
			"end-user",
		},
	}

	_, _, err := client.GetUsers(ctx, &opts)
	if err != nil {
		t.Fatalf("Failed to get users: %s", err)
	}
}

func TestCreateUser(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.CreateUser(ctx, User{
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

func TestCreateUOrUpdateUser(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.CreateUser(ctx, User{
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
