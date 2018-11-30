package zendesk

import (
	"net/http"
	"testing"
)

func TestGetGroups(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "groups.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	groups, _, err := client.GetGroups()
	if err != nil {
		t.Fatalf("Failed to get groups: %s", err)
	}

	if len(groups) != 1 {
		t.Fatalf("expected length of groups is 1, but got %d", len(groups))
	}
}
