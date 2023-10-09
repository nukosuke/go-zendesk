package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetTicketsCBP(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, _, err := client.GetTicketsCBP(ctx, &CBPOptions{
		CursorPagination: CursorPagination{
			PageSize:  10,
			PageAfter: "",
		},
	})
	if err != nil {
		t.Fatalf("Failed to get tickets: %s", err)
	}

	expectedLength := 2
	if len(tickets) != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, len(tickets))
	}
}

func TestGetTicketsIteratorCBPDefault(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ops := NewPaginationOptions()
	it := client.GetTicketsIterator(ctx, ops)

	expectedLength := 2
	ticketCount := 0
	for it.HasMore() {
		tickets, err := it.GetNext()
		if err == nil {
			for _, ticket := range tickets {
				println(ticket.Subject)
				ticketCount++
			}
		}
	}
	if ticketCount != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, ticketCount)
	}
}

func TestGetTicketsIteratorOBPOptional(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ops := NewPaginationOptions()
	ops.IsCBP = false
	it := client.GetTicketsIterator(ctx, ops)

	expectedLength := 2
	ticketCount := 0
	for it.HasMore() {
		tickets, err := it.GetNext()
		if err == nil {
			for _, ticket := range tickets {
				println(ticket.Subject)
				ticketCount++
			}
		}
	}
	if ticketCount != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, ticketCount)
	}
}

func TestGetOrganizationTickets(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, _, err := client.GetOrganizationTickets(ctx, 360363695492, &TicketListOptions{
		PageOptions: PageOptions{
			Page:    1,
			PerPage: 10,
		},
		SortBy:    "created_at",
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

func TestGetOrganizationTicketsOBP(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, _, err := client.GetOrganizationTicketsOBP(ctx, &OBPOptions{
		PageOptions: PageOptions{
			Page:    1,
			PerPage: 10,
		},
		CommonOptions: CommonOptions{
			SortBy:    "created_at",
			SortOrder: "asc",
			Id:        360363695492,
		},
	})
	if err != nil {
		t.Fatalf("Failed to get tickets: %s", err)
	}

	expectedLength := 2
	if len(tickets) != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, len(tickets))
	}
}

func TestGetOrganizationTicketsCBP(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tickets, _, err := client.GetOrganizationTicketsCBP(ctx, &CBPOptions{
		CursorPagination: CursorPagination{
			PageSize:  10,
			PageAfter: "",
		},
		CommonOptions: CommonOptions{
			SortBy:    "created_at",
			SortOrder: "asc",
			Id:        360363695492,
		},
	})
	if err != nil {
		t.Fatalf("Failed to get tickets: %s", err)
	}

	expectedLength := 2
	if len(tickets) != expectedLength {
		t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, len(tickets))
	}
}

func TestGetOrganizationTicketsIterator(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "tickets.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ops := NewPaginationOptions()
	ops.Sort = "updated_at"
	ops.PageSize = 10

	it := client.GetOrganizationTicketsIterator(ctx, ops)

	expectedLength := 2
	ticketCount := 0
	for it.HasMore() {
		tickets, err := it.GetNext()
		if len(tickets) != expectedLength {
			t.Fatalf("Returned tickets does not have the expected length %d. Tickets length is %d", expectedLength, len(tickets))
		}
		if err == nil {
			for _, ticket := range tickets {
				println(ticket.Subject)
				ticketCount++
			}
		}
		if err != nil {
			t.Fatalf("Failed to get tickets: %s", err)
		}
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

	expectedVia := &Via{
		Channel: "email",
		Source: struct {
			From map[string]interface{} `json:"from"`
			To   map[string]interface{} `json:"to"`
			Rel  string                 `json:"rel"`
		}{
			From: map[string]interface{}{
				"address": "nukosuke@lavabit.com",
				"name":    "Yosuke Tamura",
			},
			To: map[string]interface{}{
				"name":    "Terraform Zendesk provider",
				"address": "support@d3v-terraform-provider.zendesk.com",
			},
			Rel: "",
		},
	}

	if !reflect.DeepEqual(ticket.Via, expectedVia) {
		t.Fatal(fmt.Sprintf("Expected ticket via object to be %v but got %v", expectedVia, ticket.Via))
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
		Comment: &TicketComment{
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

func TestUpdateTicket(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "ticket.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticket, err := client.UpdateTicket(ctx, 2, Ticket{})
	if err != nil {
		t.Fatalf("Failed to update ticket: %s", err)
	}

	expectedID := int64(2)
	if ticket.ID != expectedID {
		t.Fatalf("Returned ticket does not have the expected ID %d. Ticket id is %d", expectedID, ticket.ID)
	}
}

func TestUpdateTicketFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "ticket.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.UpdateTicket(ctx, 2, Ticket{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteTicket(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteTicket(ctx, 437)
	if err != nil {
		t.Fatalf("Failed to delete ticket field: %s", err)
	}
}

func TestTicketMarshalling(t *testing.T) {
	var src, dst Ticket

	marshalled, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("Failed to json-marshal ticket: %v", err)
	}

	err = json.Unmarshal(marshalled, &dst)
	if err != nil {
		t.Fatalf("Failed to json-unmarshal ticket: %v", err)
	}

	if !reflect.DeepEqual(src, dst) {
		t.Fatalf("remarshalling is inconsistent")
	}

}
