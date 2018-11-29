package zendesk

import (
	"net/http"
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
