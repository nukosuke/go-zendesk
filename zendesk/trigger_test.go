package zendesk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: make API mocking helper
func TestGetTriggers(t *testing.T) {
	triggersJSON, err := ioutil.ReadFile("../test/fixtures/triggers.json")
	if err != nil {
		fmt.Printf("Failed to read fixture: %s", err)
	}
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(triggersJSON)
	}))
	defer mockAPI.Close()

	client, _ := NewClient(nil)
	client.SetCredential(NewAPITokenCredential("", ""))
	client.SetEndpointURL(mockAPI.URL)

	triggers, _, err := client.GetTriggers()
	if err != nil {
		t.Fatalf("Failed to get triggers: %s", err)
	}

	if len(triggers) != 8 {
		t.Fatalf("expected length of triggers is , but got %d", len(triggers))
	}
}
