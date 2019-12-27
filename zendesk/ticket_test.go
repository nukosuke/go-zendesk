package zendesk

import (
	"context"
	"encoding/json"
	"github.com/nukosuke/go-zendesk/zendesk/sideload"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sort"
	"testing"
)

func TestGetTickets(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, _, err := client.GetTickets(ctx, &TicketListOptions{
		PageOptions: PageOptions{
			Page:    1,
			PerPage: 10,
		},
		SortBy:    "id",
		SortOrder: "asc",
	})
	if err != nil {
		t.Fatalf("Failed to get tickets: %s", err)
	}

	expectedLength := 2
	if len(tickets) != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, len(tickets))
	}
}

func TestGetTicket(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.GetTicket(ctx, 2)
	if err != nil {
		t.Fatalf("Failed to get ticket: %s", err)
	}

	expectedID := int64(2)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
}

func TestGetTicketCanceledContext(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()
	canceled, cancelFunc := context.WithCancel(ctx)
	cancelFunc()
	_, err := client.GetTicket(canceled, 2)
	if err == nil {
		t.Fatal("Did not get error when calling with cancelled context")
	}
}

func TestGetTicketSideloaded(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		expectedQuery := `users,groups,dates`
		actual := q.Get("include")
		if actual != expectedQuery {
			t.Fatalf(`Actual query did not match expected. Was "%s" expected "%s"`, actual, expectedQuery)
		}
		w.Write(readFixture(filepath.Join(http.MethodGet, "ticket_sideload.json")))
	}))
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	var users []User
	var groups []Group
	var ticketDates sideload.TicketDates
	ticket, err := client.GetTicket(
		ctx,
		4,
		sideload.IncludeObject("users", &users),
		sideload.IncludeObject("groups", &groups),
		sideload.IncludeTicketDates(&ticketDates),
	)

	if err != nil {
		t.Fatalf("Failed to get ticket: %s", err)
	}

	expectedID := int64(4)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}

	if len(users) != 1 {
		t.Fatalf("Did not have the expected number of users")
	}

	userID := int64(377922500012)
	if users[0].ID != userID {
		t.Fatalf("User did not have the expected userID %d. User was: %v", userID, users[0])
	}

	if len(groups) != 1 {
		t.Fatalf("Did not have the expected number of groups")
	}

	groupID := int64(360004077472)
	if groups[0].ID != groupID {
		t.Fatalf("Group did not have expected group id %d", groupID)
	}

	if ticketDates.RequesterUpdatedAt.Year() != 2019 && ticketDates.RequesterUpdatedAt.Month() != 6 {
		t.Fatalf("Ticket dates did not have expected value. Was %v", ticketDates)
	}
}

func TestTicketSideloadReturnsErrorIfNotPassedPointer(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_sideload.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	var users []User
	_, err := client.GetTicket(ctx, 4, sideload.IncludeObject("users", users))

	if err == nil {
		t.Fatalf("Did not recieve error from error when passing list by copy")
	}
}

// Test the CustomField unmarshalling fails on an invalid value.
// In this case a float64 as CustomField.Value should cause an error.
func TestGetTicketWithInvalidCustomField(t *testing.T) {
	// Test with a number value.
	invalidCustomFieldJson := `{ "id": 360005657120, "value": 123.456 }`
	var customField CustomField
	err := json.Unmarshal([]byte(invalidCustomFieldJson), &customField)
	if err == nil {
		t.Fatalf("Expected an error when parsing a custom field of type number.")
	}

	// Test with an array of numbers.
	invalidCustomFieldJson = `{ "id": 360005657120, "value": [123, 456] }`
	err = json.Unmarshal([]byte(invalidCustomFieldJson), &customField)
	if err == nil {
		t.Fatalf("Expected an error when parsing a custom field of type [number, ...].")
	}
}

func TestGetTicketWithCustomFields(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_custom_field.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.GetTicket(ctx, 4)
	if err != nil {
		t.Fatalf("Failed to get ticket: %s", err)
	}

	expectedID := int64(4)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
	if ticket.CustomFields == nil || len(ticket.CustomFields) == 0 {
		t.Fatalf("Returned ticket does not have the expected custom fields.")
	}
	for _, cf := range ticket.CustomFields {
		switch cf.Value.(type) {
		case string:
			expectedCustomFieldValue := "Custom field value for testing"
			if cf.Value != expectedCustomFieldValue {
				t.Fatalf("Returned custom field value is not the expected value %s", cf.Value)
			}
		case []string:
			expectedCustomFieldValue := []string{"list", "of", "values"}
			sort.Strings(expectedCustomFieldValue)
			// FIXME: This comparison of array contents was necessary because reflect.DeepEqual(cf.Value.([]string), expectedCustomFieldValue) would not work.
			if len(cf.Value.([]string)) != len(expectedCustomFieldValue) {
				t.Fatalf("Expected length comparison failed")
			}
			for _, v := range cf.Value.([]string) {
				i := sort.SearchStrings(expectedCustomFieldValue, v)
				if i >= len(expectedCustomFieldValue) || expectedCustomFieldValue[i] != v {
					t.Fatalf("Expected to find %s in custom fields", v)
				}
			}
		case nil:
			/* Do nothing */
		case bool:
			if !cf.Value.(bool) {
				t.Fatal("Expected to find true in custom fields")
			}
		default:
			t.Fatalf("Invalid value type in custom field:  %v.", cf)
		}
	}
}

func TestGetMultipleTicket(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_show_many.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, err := client.GetMultipleTickets(ctx, []int64{2, 3})
	if err != nil {
		t.Fatalf("Failed to get ticket: %s", err)
	}

	expectedLen := 2
	if len(tickets) != expectedLen {
		t.Fatalf("Returned tickets does not have the length %d. Length is %d", expectedLen, len(tickets))
	}
}

func TestCreateTicket(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.CreateTicket(ctx, Ticket{
		Subject: "nyanyanyanya",
		Comment: TicketComment{
			Body: "(●ↀ ω ↀ )",
		},
	})
	if err != nil {
		t.Fatalf("Failed to create ticket: %s", err)
	}

	expectedID := int64(4)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
}
