package zendesk

import (
	"net/http"
	"testing"
)

func TestGetUserFields(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "user_fields.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	fields, page, err := client.GetUserFields(ctx, nil)
	if err != nil {
		t.Fatalf("Received error calling API: %v", err)
	}

	if page.Count != 1 {
		t.Fatalf("Did not receive the correct count in the page field. Was %d expected 1", page.Count)
	}

	n := len(fields)
	if n != 1 {
		t.Fatalf("Expected 1 entry in fields list. Got %d", n)
	}

	id := fields[0].ID
	if id != 7 {
		t.Fatalf("Field did not have the expected id. Was %d", id)
	}
}
