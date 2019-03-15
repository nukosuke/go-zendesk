package zendesk

import (
	"net/http"
	"testing"
)

func TestCreateBrand(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "brands.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateBrand(Brand{})
	if err != nil {
		t.Fatalf("Failed to send request to create brand: %s", err)
	}
}
