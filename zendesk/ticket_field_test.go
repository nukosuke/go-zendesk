package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTicketFields(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_fields.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketFields, _, err := client.GetTicketFields()
	if err != nil {
		t.Fatalf("Failed to get ticket fields: %s", err)
	}

	if len(ticketFields) != 15 {
		t.Fatalf("expected length of ticket fields is , but got %d", len(ticketFields))
	}
}

func TestGetTicketField(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_field.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketField, err := client.GetTicketField(123)
	if err != nil {
		t.Fatalf("Failed to get ticket fields: %s", err)
	}

	expectedID := int64(360011737434)
	if ticketField.ID != expectedID {
		t.Fatalf("Returned ticket field does not have the expected ID %d. Ticket id is %d", expectedID, ticketField.ID)
	}
}

func TestCreateTicketField(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket_fields.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateTicketField(TicketField{})
	if err != nil {
		t.Fatalf("Failed to send request to create ticket field: %s", err)
	}
}

func TestUpdateTicketField(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "ticket_field.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	updatedField, err := client.UpdateTicketField(int64(1234), TicketField{})
	if err != nil {
		t.Fatalf("Failed to send request to create ticket field: %s", err)
	}

	expectedID := int64(360011737434)
	if updatedField.ID != expectedID {
		t.Fatalf("Updated field %v did not have expected id %d", updatedField, expectedID)
	}
}

func TestDeleteTicketField(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTicketField(1234)
	if err != nil {
		t.Fatalf("Failed to delete ticket field: %s", err)
	}
}
