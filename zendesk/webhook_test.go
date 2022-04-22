package zendesk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	mockAPI := newMockAPI(http.MethodPost, "webhooks.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	hook, err := client.CreateWebhook(context.Background(), &Webhook{
		Authentication: &WebhookAuthentication{
			AddPosition: "header",
			Data: map[string]string{
				"password": "hello_123",
				"username": "john_smith",
			},
			Type: "basic_auth",
		},
		Endpoint:      "https://example.com/status/200",
		HTTPMethod:    http.MethodGet,
		Name:          "Example Webhook",
		RequestFormat: "json",
		Status:        "active",
		Subscriptions: []string{"conditional_ticket_events"},
	})
	if err != nil {
		t.Fatalf("Failed to create webhook: %v", err)
	}

	if len(hook.Subscriptions) != 1 || hook.Authentication.AddPosition != "header" {
		t.Fatalf("Invalid response of webhook: %v", hook)
	}
}

func TestGetWebhook(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "webhook.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	hook, err := client.GetWebhook(ctx, "01EJFTSCC78X5V07NPY2MHR00M")
	if err != nil {
		t.Fatalf("Failed to get webhook: %s", err)
	}

	expectedID := "01EJFTSCC78X5V07NPY2MHR00M"
	if hook.ID != expectedID {
		t.Fatalf("Returned webhook does not have the expected ID %s. Webhook ID is %s", expectedID, hook.ID)
	}
}

func TestUpdateWebhook(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	err := client.UpdateWebhook(ctx, "01EJFTSCC78X5V07NPY2MHR00M", &Webhook{})
	if err != nil {
		t.Fatalf("Failed to send request to create webhook: %s", err)
	}
}

func TestDeleteWebhook(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	err := client.DeleteWebhook(ctx, "01EJFTSCC78X5V07NPY2MHR00M")
	if err != nil {
		t.Fatalf("Failed to delete webhook: %s", err)
	}
}
