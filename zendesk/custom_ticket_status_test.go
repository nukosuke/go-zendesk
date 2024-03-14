package zendesk

import (
	"net/http"
	"testing"
)

func TestGetCustomTicketStatuses(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "custom_ticket_statuses.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	items, err := client.GetCustomTicketStatuses(ctx, &CustomTicketStatusListOptions{})
	if err != nil {
		t.Fatalf("Failed to get custom ticket statuses: %s", err)
	}

	if len(items) != 2 {
		t.Fatalf("expected length of custom ticket statuses is 2, but got %d", len(items))
	}
}

func TestGetCustomTicketStatus(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "custom_ticket_status.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	item, err := client.GetCustomTicketStatus(ctx, 35436)
	if err != nil {
		t.Fatalf("Failed to get custom ticket status: %s", err)
	}

	expectedID := int64(35436)
	if item.ID != expectedID {
		t.Fatalf("Returned custom ticket status id does not have the expected ID %d. Custom ticket status is %d", expectedID, item.ID)
	}
}

func TestUpdateCustomTicketStatus(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "custom_ticket_status.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	item, err := client.UpdateCustomTicketStatus(ctx, 35436, CustomTicketStatusUpdateOption{
		Description: "Customer needs a response quickly",
	})
	if err != nil {
		t.Fatalf("Failed to update custom ticket status: %s", err)
	}

	expectedDesc := "Customer needs a response quickly"
	if item.Description != expectedDesc {
		t.Fatalf("Returned custom ticket status description not updated successfully")
	}
}

