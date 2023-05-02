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
		t.Fatalf("expected length of users is 2, but got %d", len(users))
	}
}

func TestGetOrganizationUsers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "users.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	users, _, err := client.GetOrganizationUsers(ctx, 1000006909040, nil)
	if err != nil {
		t.Fatalf("Failed to get users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected length of users is 2, but got %d", len(users))
	}
}

func TestGetManyUsers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "users.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	users, _, err := client.GetManyUsers(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get many users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected length of many users is 2, but got %d", len(users))
	}
}

func TestSearchUsers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "users.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	users, _, err := client.SearchUsers(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to get many users: %s", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected length of many users is 2, but got %d", len(users))
	}
}

func TestGetUser(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "user.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.GetUser(ctx, 369531345753)
	if err != nil {
		t.Fatalf("Failed to get user: %s", err)
	}

	expectedID := int64(369531345753)
	if user.ID != expectedID {
		t.Fatalf("Returned user does not have the expected ID %d. User id is %d", expectedID, user.ID)
	}
}

func TestGetUserFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "user.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.GetUser(ctx, 369531345753)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
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

func TestCreateOrUpdateUserCreated(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.CreateOrUpdateUser(ctx, User{
		Email: "test@example.com",
		Name:  "testuser",
	})
	if err != nil {
		t.Fatalf("Failed to get valid response: %s", err)
	}
	if user.ID == 0 {
		t.Fatal("Failed to create or update user")
	}
}

func TestCreateOrUpdateUserUpdated(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.CreateOrUpdateUser(ctx, User{
		Email: "test@example.com",
		Name:  "testuser",
	})
	if err != nil {
		t.Fatalf("Failed to get valid response: %s", err)
	}
	if user.ID == 0 {
		t.Fatal("Failed to create or update user")
	}
}

func TestCreateOrUpdateUserFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "users.json", http.StatusInternalServerError)

	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateOrUpdateUser(ctx, User{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateUser(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "user.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	user, err := client.UpdateUser(ctx, 369531345753, User{})
	if err != nil {
		t.Fatalf("Failed to update user: %s", err)
	}

	expectedID := int64(369531345753)
	if user.ID != expectedID {
		t.Fatalf("Returned user does not have the expected ID %d. User id is %d", expectedID, user.ID)
	}
}

func TestUpdateUserFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "user.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.UpdateUser(ctx, 369531345753, User{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestGetUserRelated(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "user_related.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	userRelated, err := client.GetUserRelated(ctx, 369531345753)
	if err != nil {
		t.Fatalf("Failed to get user related information: %s", err)
	}

	expectedAssignedTickets := int64(5)
	if userRelated.AssignedTickets != expectedAssignedTickets {
		t.Fatalf("Returned user does not have the expected assigned tickets %d. It is %d", expectedAssignedTickets, userRelated.AssignedTickets)
	}
}
