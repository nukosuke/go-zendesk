package zendesk

import (
	"net/http"
	"testing"
)

func TestNewPublicTicketComment(t *testing.T) {
	publicComment := NewPublicTicketComment("public comment", 12345)

	// Both true and nil are public comments
	if *publicComment.Public == false {
		t.Fatalf("Returned comment is not marked as public. Comment public is %v", *publicComment.Public)
	}
}

func TestNewPrivateTicketComment(t *testing.T) {
	privateComment := NewPrivateTicketComment("private comment", 12345)

	// Both true and nil are public comments
	if *privateComment.Public != false {
		t.Fatalf("Returned comment is not marked as private. Comment public is %v", *privateComment.Public)
	}
}

func TestCreateTicketComment(t *testing.T) {
	mockAPI := newMockAPI(http.MethodPut, "ticket.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	publicComment := NewPublicTicketComment("public comment", 12345)

	err := client.CreateTicketComment(ctx, 2, publicComment)
	if err != nil {
		t.Fatalf("Failed to create ticket comment: %s", err)
	}
}

func TestListTicketComments(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_comments.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	ticketComments, err := client.ListTicketComments(ctx, 2)
	if err != nil {
		t.Fatalf("Failed to list ticket comments: %s", err)
	}

	expectedLength := 2
	if len(ticketComments) != expectedLength {
		t.Fatalf("Returned ticket comments does not have the expected length %d. Ticket comments length is %d", expectedLength, len(ticketComments))
	}
}
