package zendesk

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	if _, err := NewClient(nil); err != nil {
		t.Fatal("Failed to create Client")
	}
}

func TestSetHeader(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetHeader("Header1", "hogehoge")

	if client.headers["Header1"] != "hogehoge" {
		t.Fatal("Header1 is wrong")
	}
}

func TestSetSubdomainSuccess(t *testing.T) {
	validSubdomain := "subdomain"

	client, _ := NewClient(&http.Client{})
	if err := client.SetSubdomain(validSubdomain); err != nil {
		t.Fatal("SetSubdomain should success")
	}
}

func TestSetSubdomainFail(t *testing.T) {
	invalidSubdomain := ".subdomain"

	client, _ := NewClient(&http.Client{})
	if err := client.SetSubdomain(invalidSubdomain); err == nil {
		t.Fatal("SetSubdomain should fail")
	}
}

func TestSetEndpointURL(t *testing.T) {
	client, _ := NewClient(nil)
	if err := client.SetEndpointURL("http://127.0.0.1:3000"); err != nil {
		t.Fatal("SetEndpointURL should success")
	}
}

func TestSetCredential(t *testing.T) {
	client, _ := NewClient(nil)
	cred := NewBasicAuthCredential("john.doe@example.com", "password")
	client.SetCredential(cred)

	if email := client.credential.Email(); email != "john.doe@example.com" {
		t.Fatal("client.credential.Email() returns wrong email: " + email)
	}
	if secret := client.credential.Secret(); secret != "password" {
		t.Fatal("client.credential.Secret() returns wrong secret: " + secret)
	}
}
