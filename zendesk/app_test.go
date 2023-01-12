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

	expected := []AppInstallation{
		{
			ID:      42,
			AppID:   913,
			Product: "support",
			Settings: struct {
				Name  string `json:"name"`
				Title string `json:"title"`
			}{
				Name:  "Mystery App",
				Title: "Mystery App",
			},
			SettingsObjects: []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}{
				{
					Name:  "setting-one",
					Value: "value-one",
				},
				{
					Name:  "setting-two",
					Value: "value-two",
				},
			},
			Enabled:   true,
			Paid:      false,
			UpdatedAt: time.Date(2023, 1, 1, 1, 1, 1, 0, time.UTC),
			CreatedAt: time.Date(2023, 1, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			ID:      42,
			AppID:   917,
			Product: "support",
			Settings: struct {
				Name  string `json:"name"`
				Title string `json:"title"`
			}{
				Name:  "Mystery App 2",
				Title: "Mystery App 2",
			},
			SettingsObjects: []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			}{
				{
					Name:  "foo",
					Value: "bar",
				},
			},
			Enabled:   true,
			Paid:      false,
			UpdatedAt: time.Date(2023, 2, 2, 2, 2, 2, 0, time.UTC),
			CreatedAt: time.Date(2023, 2, 2, 2, 2, 2, 0, time.UTC),
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
