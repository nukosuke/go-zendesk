package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSlaPolicies(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policies.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	slaPolicies, _, err := client.GetSlaPolicies(ctx, &SlaPolicyListOptions{})
	if err != nil {
		t.Fatalf("Failed to get sla policies: %s", err)
	}

	if len(slaPolicies) != 3 {
		t.Fatalf("expected length of sla policies is , but got %d", len(slaPolicies))
	}
}

func TestGetSlaPoliciesWithNil(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policies.json")
	client := newTestClient(mockAPI)

	_, _, err := client.GetSlaPolicies(ctx, nil)
	if err == nil {
		t.Fatal("expected an OptionsError, but no error")
	}

	_, ok := err.(*OptionsError)
	if !ok {
		t.Fatalf("unexpected error type: %v", err)
	}
}

func TestCreateSlaPolicy(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "sla_policies.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.CreateSlaPolicy(ctx, SlaPolicy{})
	if err != nil {
		t.Fatalf("Failed to send request to create sla policy: %s", err)
	}
}

func TestGetSlaPolicy(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policy.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	sla, err := client.GetSlaPolicy(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get sla policy: %s", err)
	}

	expectedID := int64(360000068060)
	if sla.ID != expectedID {
		t.Fatalf("Returned sla policy does not have the expected ID %d. Sla policy id is %d", expectedID, sla.ID)
	}
}

func TestGetSlaPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.GetSlaPolicy(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateSlaPolicy(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "sla_policies.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	sla, err := client.UpdateSlaPolicy(ctx, 123, SlaPolicy{})
	if err != nil {
		t.Fatalf("Failed to get sla policy: %s", err)
	}

	expectedID := int64(360000068060)
	if sla.ID != expectedID {
		t.Fatalf("Returned slaPolicy does not have the expected ID %d. Sla policy id is %d", expectedID, sla.ID)
	}
}

func TestUpdateSlaPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.UpdateSlaPolicy(ctx, 1234, SlaPolicy{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteSlaPolicy(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteSlaPolicy(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete sla policy: %s", err)
	}
}

func TestDeleteSlaPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteSlaPolicy(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}
