package zendesk

import (
	"net/http"
	"testing"
)

func TestNewClientSuccess(t *testing.T) {
	validSubdomain := "subdomain"

	_, err := NewClient(&http.Client{}, validSubdomain)
	if err != nil {
		t.Fatal("NewClient with valid params must success")
	}
}

func TestNewClientFail(t *testing.T) {
	invalidSubdomain := ".subdomain"

	_, err := NewClient(&http.Client{}, invalidSubdomain)
	if err == nil {
		t.Fatal("NewClient with invalid params must fail")
	}
}
