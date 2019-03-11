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

func TestCreateTicketForm(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket_form.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateTicketForm(TicketForm{})
	if err != nil {
		t.Fatalf("Failed to send request to create ticket form: %s", err)
	}
}
