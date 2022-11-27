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
