package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTriggers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "triggers.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	triggers, _, err := client.GetTriggers(&TriggerListOptions{})
	if err != nil {
		t.Fatalf("Failed to get triggers: %s", err)
	}

	if len(triggers) != 8 {
		t.Fatalf("expected length of triggers is , but got %d", len(triggers))
	}
}

func TestCreateTrigger(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "triggers.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateTrigger(Trigger{})
	if err != nil {
		t.Fatalf("Failed to send request to create trigger: %s", err)
	}
}

func TestGetTrigger(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "trigger.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	trg, err := client.GetTrigger(123)
	if err != nil {
		t.Fatalf("Failed to get trigger: %s", err)
	}

	expectedID := int64(360056295714)
	if trg.ID != expectedID {
		t.Fatalf("Returned trigger does not have the expected ID %d. Trigger id is %d", expectedID, trg.ID)
	}
}

func TestGetTriggerFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.GetTrigger(1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateTrigger(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "triggers.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	trg, err := client.UpdateTrigger(123, Trigger{})
	if err != nil {
		t.Fatalf("Failed to get trigger: %s", err)
	}

	expectedID := int64(360056295714)
	if trg.ID != expectedID {
		t.Fatalf("Returned trigger does not have the expected ID %d. Trigger id is %d", expectedID, trg.ID)
	}
}

func TestUpdateTriggerFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.UpdateTrigger(1234, Trigger{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteTrigger(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTrigger(1234)
	if err != nil {
		t.Fatalf("Failed to delete trigger: %s", err)
	}
}

func TestDeleteTriggerFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTrigger(1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}
