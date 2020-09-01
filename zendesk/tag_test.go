package zendesk

import (
	"net/http"
	"testing"
)

func TestGetTicketTags(t *testing.T) {
	mockApi := newMockAPI(http.MethodGet, "tags.json")
	client := newTestClient(mockApi)
	defer mockApi.Close()

	tags, err := client.GetTicketTags(ctx, int64(2))
	if err != nil {
		t.Fatalf("Failed to get ticket tags: %s", err)
	}

	expectedLength := 2
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
}

func TestGetOrganizationTags(t *testing.T) {
	mockApi := newMockAPI(http.MethodGet, "tags.json")
	client := newTestClient(mockApi)
	defer mockApi.Close()

	tags, err := client.GetOrganizationTags(ctx, int64(2))
	if err != nil {
		t.Fatalf("Failed to get organization tags: %s", err)
	}

	expectedLength := 2
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
}

func TestGetUserTags(t *testing.T) {
	mockApi := newMockAPI(http.MethodGet, "tags.json")
	client := newTestClient(mockApi)
	defer mockApi.Close()

	tags, err := client.GetUserTags(ctx, int64(2))
	if err != nil {
		t.Fatalf("Failed to get user tags: %s", err)
	}

	expectedLength := 2
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
}

func TestAddTicketTags(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "tags.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tag := Tag("example")

	tags, err := client.AddTicketTags(ctx, 2, []Tag{tag})
	if err != nil {
		t.Fatalf("Failed to add ticket tags: %s", err)
	}

	expectedLength := 3
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
	if tags[2] != tag {
		t.Fatalf("Returned tags does not have the expexted tag %s. %s given", "important", tags[0])
	}
}

func TestAddOrganizationTags(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "tags.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tag := Tag("example")

	tags, err := client.AddOrganizationTags(ctx, 2, []Tag{tag})
	if err != nil {
		t.Fatalf("Failed to add ticket tags: %s", err)
	}

	expectedLength := 3
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
	if tags[2] != tag {
		t.Fatalf("Returned tags does not have the expexted tag %s. %s given", "important", tags[0])
	}
}

func TestAddUserTags(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "tags.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	tag := Tag("example")

	tags, err := client.AddUserTags(ctx, 2, []Tag{tag})
	if err != nil {
		t.Fatalf("Failed to add ticket tags: %s", err)
	}

	expectedLength := 3
	if len(tags) != expectedLength {
		t.Fatalf("Returned tags does not have the expexted length %d. Tags length is %d", expectedLength, len(tags))
	}
	if tags[2] != tag {
		t.Fatalf("Returned tags does not have the expexted tag %s. %s given", "important", tags[0])
	}
}
