package zendesk

import (
	"errors"
	"net/http"
	"net/http/httptest"
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

	_, err := client.CreateTicketComment(ctx, 2, publicComment)
	if err != nil {
		t.Fatalf("Failed to create ticket comment: %s", err)
	}
}

func TestListTicketComments(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "ticket_comments.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	result, err := client.ListTicketComments(ctx, 2, nil)
	if err != nil {
		t.Fatalf("Failed to list ticket comments: %s", err)
	}

	expectedLength := 2
	if len(result.TicketComments) != expectedLength {
		t.Fatalf("Returned ticket comments does not have the expected length %d. Ticket comments length is %d", expectedLength, len(result.TicketComments))
	}

	expectedPaginationMeta := CursorPaginationMeta{
		HasMore:      true,
		AfterCursor:  "xxx",
		BeforeCursor: "yyy",
	}

	if result.Meta != expectedPaginationMeta {
		t.Fatalf(`Failed to return correct cursor options.
Expected: %+v
Received: %+v`, expectedPaginationMeta, result.Meta)
	}
}

func TestMakeCommentPrivate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		apiReturn      error
		expectedErrStr string
	}{
		{
			name:           "successfully made private",
			apiReturn:      nil,
			expectedErrStr: "",
		},
		{
			name:           "error making private",
			apiReturn:      errors.New(`{"error":"Couldn't authenticate you"}`),
			expectedErrStr: `401: {"error":"Couldn't authenticate you"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if test.apiReturn != nil {
					w.WriteHeader(http.StatusUnauthorized)
					_, _ = w.Write([]byte(test.apiReturn.Error()))
				} else {
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write(nil)
				}
			}))
			defer mockAPI.Close()

			client := newTestClient(mockAPI)
			err := client.MakeCommentPrivate(ctx, 2, 12841284)
			if err == nil {
				if test.expectedErrStr != "" {
					t.Fatalf("Expected error %s, did not get one", test.expectedErrStr)
				}
			} else {
				if test.expectedErrStr != err.Error() {
					t.Fatalf("Got %s, wanted %s", err.Error(), test.expectedErrStr)
				}
			}
		})
	}
}
