package zendesk

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSearchTickets(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "search_ticket.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	results, _, err := client.Search(ctx, &SearchOptions{})
	if err != nil {
		t.Fatalf("Failed to get search results: %s", err)
	}

	list := results.List()
	if len(list) != 1 {
		t.Fatalf("expected length of sla policies is , but got %d", len(list))
	}

	ticket, ok := list[0].(Ticket)
	if !ok {
		t.Fatalf("Cannot assert %v as a ticket", list[0])
	}

	if ticket.ID != 4 {
		t.Fatalf("Ticket did not have the expected id %v", ticket)
	}
}

func BenchmarkUnmarshalSearchResults(b *testing.B) {
	file := readFixture("ticket_result.json")
	for i := 0; i < b.N; i++ {
		var result SearchResults
		err := json.Unmarshal(file, &result)
		if err != nil {
			b.Fatalf("Recieved error when unmarshalling. %v", err)
		}
	}
}

func TestSearchGroup(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "search_group.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	results, _, err := client.Search(ctx, &SearchOptions{})
	if err != nil {
		t.Fatalf("Failed to get search results: %s", err)
	}

	list := results.List()
	if len(list) != 1 {
		t.Fatalf("expected length of sla policies is , but got %d", len(list))
	}

	result, ok := list[0].(Group)
	if !ok {
		t.Fatalf("Cannot assert %v as a group", list[0])
	}

	if result.ID != 360007194452 {
		t.Fatalf("Group did not have the expected id %v", result)
	}
}

func TestSearchUser(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "search_user.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	results, _, err := client.Search(ctx, &SearchOptions{})
	if err != nil {
		t.Fatalf("Failed to get search results: %s", err)
	}

	list := results.List()
	if len(list) != 1 {
		t.Fatalf("expected length of sla policies is , but got %d", len(list))
	}

	result, ok := list[0].(User)
	if !ok {
		t.Fatalf("Cannot assert %v as a group", list[0])
	}

	if result.ID != 1234 {
		t.Fatalf("Group did not have the expected id %v", result)
	}
}