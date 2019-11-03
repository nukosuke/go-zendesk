package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTargets(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "targets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	targets, _, err := client.GetTargets(ctx)
	if err != nil {
		t.Fatalf("Failed to get targets: %s", err)
	}

	if len(targets) != 2 {
		t.Fatalf("expected length of targets is , but got %d", len(targets))
	}
}

func TestGetTarget(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "target.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	target, err := client.GetTarget(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get targets: %s", err)
	}

	expectedID := int64(360000217439)
	if target.ID != expectedID {
		t.Fatalf("Returned target does not have the expected ID %d. Ticket id is %d", expectedID, target.ID)
	}
}

func TestCreateTarget(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "target.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateTarget(ctx, Target{})
	if err != nil {
		t.Fatalf("Failed to send request to create target: %s", err)
	}
}

func TestUpdateTarget(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "target.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	updatedField, err := client.UpdateTarget(ctx, int64(1234), Target{})
	if err != nil {
		t.Fatalf("Failed to send request to create target: %s", err)
	}

	expectedID := int64(360000217439)
	if updatedField.ID != expectedID {
		t.Fatalf("Updated field %v did not have expected id %d", updatedField, expectedID)
	}
}

func TestDeleteTarget(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTarget(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete target: %s", err)
	}
}
