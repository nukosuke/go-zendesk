package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateOrganization(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "organization.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateOrganization(ctx, Organization{})
	if err != nil {
		t.Fatalf("Failed to send request to create organization: %s", err)
	}
}

func TestGetOrganization(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "organization.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	org, err := client.GetOrganization(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get organization: %s", err)
	}

	expectedID := int64(361898904439)
	if org.ID != expectedID {
		t.Fatalf("Returned organization does not have the expected ID %d. Organization ID is %d", expectedID, org.ID)
	}
}

func TestGetOrganizations(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "organizations.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	orgs, _, err := client.GetOrganizations(ctx, &OrganizationListOptions{})
	if err != nil {
		t.Fatalf("Failed to get organizations: %s", err)
	}

	if len(orgs) != 2 {
		t.Fatalf("expected length of organizations is , but got %d", len(orgs))
	}
}

func TestUpdateOrganization(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "organization.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	updatedOrg, err := client.UpdateOrganization(ctx, int64(1234), Organization{})
	if err != nil {
		t.Fatalf("Failed to send request to create organization: %s", err)
	}

	expectedID := int64(361898904439)
	if updatedOrg.ID != expectedID {
		t.Fatalf("Updated organization %v did not have expected id %d", updatedOrg, expectedID)
	}
}

func TestDeleteOrganization(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteOrganization(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete organization: %s", err)
	}
}
