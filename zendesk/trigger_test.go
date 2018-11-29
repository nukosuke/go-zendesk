package zendesk

import (
	"net/http"
	"testing"
)

func TestGetTriggers(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "triggers.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	triggers, _, err := client.GetTriggers()
	if err != nil {
		t.Fatalf("Failed to get triggers: %s", err)
	}

	if len(triggers) != 8 {
		t.Fatalf("expected length of triggers is , but got %d", len(triggers))
	}
}
