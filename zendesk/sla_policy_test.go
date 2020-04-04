package zendesk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSLAPolicies(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policies.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	slaPolicies, _, err := client.GetSLAPolicies(ctx, &SLAPolicyListOptions{})
	if err != nil {
		t.Fatalf("Failed to get sla policies: %s", err)
	}

	if len(slaPolicies) != 3 {
		t.Fatalf("expected length of sla policies is , but got %d", len(slaPolicies))
	}
}

func TestGetSLAPoliciesWithNil(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policies.json")
	client := newTestClient(mockAPI)

	_, _, err := client.GetSLAPolicies(ctx, nil)
	if err == nil {
		t.Fatal("expected an OptionsError, but no error")
	}

	_, ok := err.(*OptionsError)
	if !ok {
		t.Fatalf("unexpected error type: %v", err)
	}
}

func TestCreateSLAPolicy(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "sla_policies.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	policy, err := client.CreateSLAPolicy(ctx, SLAPolicy{})
	if err != nil {
		t.Fatalf("Failed to send request to create sla policy: %s", err)
	}

	if len(policy.PolicyMetrics) == 0 {
		t.Fatal("Failed to set the policy metrics from the json response")
	}
}

func TestGetSLAPolicy(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "sla_policy.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	sla, err := client.GetSLAPolicy(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get sla policy: %s", err)
	}

	expectedID := int64(360000068060)
	if sla.ID != expectedID {
		t.Fatalf("Returned sla policy does not have the expected ID %d. Sla policy id is %d", expectedID, sla.ID)
	}
}

func TestGetSLAPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.GetSLAPolicy(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestUpdateSLAPolicy(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "sla_policies.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	sla, err := client.UpdateSLAPolicy(ctx, 123, SLAPolicy{})
	if err != nil {
		t.Fatalf("Failed to get sla policy: %s", err)
	}

	expectedID := int64(360000068060)
	if sla.ID != expectedID {
		t.Fatalf("Returned slaPolicy does not have the expected ID %d. Sla policy id is %d", expectedID, sla.ID)
	}
}

func TestUpdateSLAPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	_, err := c.UpdateSLAPolicy(ctx, 1234, SLAPolicy{})
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}

func TestDeleteSLAPolicy(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteSLAPolicy(ctx, 1234)
	if err != nil {
		t.Fatalf("Failed to delete sla policy: %s", err)
	}
}

func TestDeleteSLAPolicyFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteSLAPolicy(ctx, 1234)
	if err == nil {
		t.Fatal("Client did not return error when api failed")
	}
}
