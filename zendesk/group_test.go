package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGroups(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "groups.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	groups, _, err := client.GetGroups(ctx)
	if err != nil {
		t.Fatalf("Failed to get groups: %s", err)
	}

	if len(groups) != 1 {
		t.Fatalf("expected length of groups is 1, but got %d", len(groups))
	}
}

func TestCreateGroup(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "groups.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateGroup(ctx, Group{})
	if err != nil {
		t.Fatalf("Failed to send request to create group: %s", err)
	}
}

func TestGetGroup(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "group.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	group, err := client.GetGroup(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get group: %s", err)
	}

	expectedID := int64(360002440594)
	if group.ID != expectedID {
		t.Fatalf("Returned group does not have the expected ID %d. Group ID is %d", expectedID, group.ID)
	}
}

func TestUpdateGroup(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "groups.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	updatedGroup, err := client.UpdateGroup(ctx, int64(1234), Group{})
	if err != nil {
		t.Fatalf("Failed to send request to create group: %s", err)
	}

	expectedID := int64(360002440594)
	if updatedGroup.ID != expectedID {
		t.Fatalf("Updated group %v did not have expected id %d", updatedGroup, expectedID)
	}
}

func TestDeleteGroup(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteGroup(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete group: %s", err)
	}
}
