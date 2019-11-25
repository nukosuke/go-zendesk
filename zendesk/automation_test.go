package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAutomations(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "automations.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	automations, _, err := client.GetAutomations(ctx, &AutomationListOptions{})
	if err != nil {
		t.Fatalf("Failed to get automations: %s", err)
	}

	if len(automations) != 3 {
		t.Fatalf("expected length of automations is , but got %d", len(automations))
	}
}

func TestGetAutomationsWithNil(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "automations.json")
	client := newTestClient(mockAPI)

	_, _, err := client.GetAutomations(ctx, nil)
	if err == nil {
		t.Fatal("expected an OptionsError, but no error")
	}

	_, ok := err.(*OptionsError)
	if !ok {
		t.Fatalf("unexpected error type: %v", err)
	}
}

func TestCreateAutomation(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "automations.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateAutomation(ctx, Automation{})
	if err != nil {
		t.Fatalf("Failed to send request to create automation: %s", err)
	}
}

func TestGetAutomation(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "automation.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	trg, err := client.GetAutomation(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get automation: %s", err)
	}

	expectedID := int64(360017421099)
	if trg.ID != expectedID {
		t.Fatalf("Returned automation does not have the expected ID %d. Automation id is %d", expectedID, trg.ID)
	}
}

func TestGetAutomationFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.GetAutomation(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateAutomation(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "automations.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	trg, err := client.UpdateAutomation(ctx, 123, Automation{})
	if err != nil {
		t.Fatalf("Failed to get automation: %s", err)
	}

	expectedID := int64(360017421099)
	if trg.ID != expectedID {
		t.Fatalf("Returned automation does not have the expected ID %d. Automation id is %d", expectedID, trg.ID)
	}
}

func TestUpdateAutomationFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.UpdateAutomation(ctx, 1234, Automation{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteAutomation(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteAutomation(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete automation: %s", err)
	}
}

func TestDeleteAutomationFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteAutomation(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}
