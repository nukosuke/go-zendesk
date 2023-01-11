package zendesk

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListAppInstallations(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "apps.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	actual, err := client.ListInstallations(ctx)
	if err != nil {
		t.Fatalf("Failed to send request to list app installations: %s", err)
	}

	expected := []App{
		{
			ID:                 42,
			OwnerID:            913,
			Name:               "Mystery App",
			AuthorName:         "John Doe",
			AuthorEmail:        "john@doe.me",
			ShortDescription:   "Does mysterious things",
			Enabled:            true,
			CreatedAt:          time.Date(2023, 1, 1, 1, 1, 1, 0, time.UTC),
			UpdatedAt:          time.Date(2023, 1, 1, 1, 1, 1, 0, time.UTC),
			Version:            "v1.0.1",
			TermsConditionsURL: "https://example.com/terms",
		},
		{
			ID:                 47,
			OwnerID:            913,
			Name:               "Mystery App 2",
			AuthorName:         "Jane Doe",
			AuthorEmail:        "jane@doe.me",
			ShortDescription:   "Does *more* mysterious things",
			Enabled:            true,
			CreatedAt:          time.Date(2023, 2, 2, 2, 2, 2, 0, time.UTC),
			UpdatedAt:          time.Date(2023, 2, 2, 2, 2, 2, 0, time.UTC),
			Version:            "v1.0.2",
			TermsConditionsURL: "https://example.com/terms",
		},
	}

	if len(actual) != 2 {
		t.Fatalf("expected 2 apps, got %d", len(actual))
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("apps not equal")
	}
}

func TestListAppInstallationsCanceledContext(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "apps.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	canceled, cancel := context.WithCancel(ctx)
	cancel()

	_, err := client.ListInstallations(canceled)
	if err == nil {
		t.Fatalf("did not get expected error")
	}
}
