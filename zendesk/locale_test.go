package zendesk

import (
	"net/http"
	"testing"
)

func TestGetLocales(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "locales.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	locales, err := client.GetLocales(ctx)
	if err != nil {
		t.Fatalf("Failed to get locales: %s", err)
	}

	if len(locales) != 3 {
		t.Fatalf("expected length of groups is 3, but got %d", len(locales))
	}
}
