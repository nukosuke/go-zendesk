package zendesk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBrand(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "brands.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateBrand(ctx, Brand{})
	if err != nil {
		t.Fatalf("Failed to send request to create brand: %s", err)
	}
}

func TestCreateBrandCanceledContext(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "brands.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	canceled, cancelFunc := context.WithCancel(ctx)
	cancelFunc()

	_, err := client.CreateBrand(canceled, Brand{})
	if err == nil {
		t.Fatalf("did not get expected error")
	}
}

func TestGetBrand(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "brand.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	brand, err := client.GetBrand(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get brand: %s", err)
	}

	expectedID := int64(360002143133)
	if brand.ID != expectedID {
		t.Fatalf("Returned brand does not have the expected ID %d. Brand ID is %d", expectedID, brand.ID)
	}
}

func TestUpdateBrand(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "brands.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	updatedBrand, err := client.UpdateBrand(ctx, int64(1234), Brand{})
	if err != nil {
		t.Fatalf("Failed to send request to create brand: %s", err)
	}

	expectedID := int64(360002143133)
	if updatedBrand.ID != expectedID {
		t.Fatalf("Updated brand %v did not have expected id %d", updatedBrand, expectedID)
	}
}

func TestUpdateBrandCanceledContext(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "brands.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	canceled, cancelFunc := context.WithCancel(ctx)
	cancelFunc()
	_, err := client.UpdateBrand(canceled, int64(1234), Brand{})
	if err == nil {
		t.Fatalf("did not get expected error")
	}
}

func TestDeleteBrand(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteBrand(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete brand: %s", err)
	}
}
