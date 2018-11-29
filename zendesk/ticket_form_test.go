package zendesk

import (
	"net/http"
	"testing"
)

func TestGetTicketForms(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_forms.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketForms, _, err := client.GetTicketForms()
	if err != nil {
		t.Fatalf("Failed to get ticket forms: %s", err)
	}

	if len(ticketForms) != 1 {
		t.Fatalf("expected length of ticket forms is , but got %d", len(ticketForms))
	}
}
