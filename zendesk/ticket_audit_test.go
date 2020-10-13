package zendesk

import (
	"net/http"
	"testing"
)

func TestGetAllTicketAudits(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_audits.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketAudits, _, err := client.GetAllTicketAudits(ctx, CursorOption{})
	if err != nil {
		t.Fatalf("Failed to get ticket audits: %s", err)
	}

	if len(ticketAudits) != 1 {
		t.Fatalf("expected length of ticket audit is %d, but got %d", 1, len(ticketAudits))
	}
}

func TestGetTicketAudits(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_audits.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketAudits, _, err := client.GetTicketAudits(ctx, 666, PageOptions{})
	if err != nil {
		t.Fatalf("Failed to get ticket audits: %s", err)
	}

	if len(ticketAudits) != 1 {
		t.Fatalf("expected length of ticket audit is %d, but got %d", 1, len(ticketAudits))
	}
}

func TestGetTicketAudit(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_audit.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketAudit, err := client.GetTicketAudit(ctx, 666, 2127301143)
	if err != nil {
		t.Fatalf("Failed to get ticket audit: %s", err)
	}

	expectedID := int64(2127301143)
	if ticketAudit.ID != expectedID {
		t.Fatalf("Returned ticket audit does not have the expected ID %d. Ticket audit id is %d", expectedID, ticketAudit.ID)
	}
}


