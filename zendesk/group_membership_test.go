package zendesk

import (
	"net/http"
	"testing"
)

func TestGetGroupMemberships(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "group_memberships.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	groupMemberships, _, err := client.GetGroupMemberships(ctx, &GroupMembershipListOptions{GroupID: 123})
	if err != nil {
		t.Fatalf("Failed to get group memberships: %s", err)
	}

	if len(groupMemberships) != 2 {
		t.Fatalf("expected length of group memberships is 2, but got %d", len(groupMemberships))
	}
}
