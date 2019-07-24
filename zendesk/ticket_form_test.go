package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTicketForms(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_forms.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketForms, _, err := client.GetTicketForms(ctx, nil)
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

	_, err := client.CreateTicketForm(ctx, TicketForm{})
	if err != nil {
		t.Fatalf("Failed to send request to create ticket form: %s", err)
	}
}

func TestDeleteTicketForm(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTicketForm(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete ticket field: %s", err)
	}
}

func TestDeleteTicketFormFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTicketForm(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestGetTicketForm(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_form.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	f, err := client.GetTicketForm(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get ticket fields: %s", err)
	}

	expectedID := int64(47)
	if f.ID != expectedID {
		t.Fatalf("Returned ticket form does not have the expected ID %d. Ticket id is %d", expectedID, f.ID)
	}
}

func TestGetTicketFormFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.GetTicketForm(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateTicketForm(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "ticket_form.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	f, err := client.UpdateTicketForm(ctx, 123, TicketForm{})
	if err != nil {
		t.Fatalf("Failed to get ticket fields: %s", err)
	}

	expectedID := int64(47)
	if f.ID != expectedID {
		t.Fatalf("Returned ticket form does not have the expected ID %d. Ticket id is %d", expectedID, f.ID)
	}
}

func TestUpdateTicketFormFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.UpdateTicketForm(ctx, 1234, TicketForm{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}
