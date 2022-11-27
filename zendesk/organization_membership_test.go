package zendesk

import (
	"net/http"
	"testing"
)

func TestGetOrganizationMemberships(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "organization_memberships.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	orgMemberships, _, err := client.GetOrganizationMemberships(ctx, &OrganizationMembershipListOptions{})
	if err != nil {
		t.Fatalf("Failed to get organization memberships: %s", err)
	}

	expectedOrgMemberships := 2

	if len(orgMemberships) != expectedOrgMemberships {
		t.Fatalf("expected length of organization memberships is %d, but got %d", expectedOrgMemberships, len(orgMemberships))
	}
}

func TestCreateOrganizationMembership(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "organization_membership.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateOrganizationMembership(ctx, OrganizationMembershipOptions{})
	if err != nil {
		t.Fatalf("Failed to send request to create organization membership: %s", err)
	}
}
