package zendesk

import (
	"net/http"
	"testing"
)

func TestGetOrganizationFields(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "organization_fields.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketFields, _, err := client.GetOrganizationFields(ctx)
	if err != nil {
		t.Fatalf("Failed to get organization fields: %s", err)
	}

	if len(ticketFields) != 2 {
		t.Fatalf("expected length of organization fields is , but got %d", len(ticketFields))
	}
}

func TestOrganizationField(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "organization_fields.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateOrganizationField(ctx, OrganizationField{})
	if err != nil {
		t.Fatalf("Failed to send request to create organization field: %s", err)
	}
}
