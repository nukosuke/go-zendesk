package zendesk

import (
	"net/http"
	"testing"
)

func TestGetBrands(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "brands.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	brands, _, err := client.GetBrands()
	if err != nil {
		t.Fatalf("Failed to get brands: %s", err)
	}

	if len(brands) != 2 {
		t.Fatalf("expected length of brands is 2, but got %d", len(brands))
	}
}

func TestCreateBrand(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "brands.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateBrand(Brand{})
	if err != nil {
		t.Fatalf("Failed to send request to create brand: %s", err)
	}
}
