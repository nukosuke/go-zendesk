package zendesk

import "testing"

func TestNewBasicAuthCredential(t *testing.T) {
	cred := NewBasicAuthCredential("john.doe@example.com", "password")

	if cred.Email() != "john.doe@example.com" {
		t.Fatalf("BasicAuthCredential: email not match")
	}
	if cred.Secret() != "password" {
		t.Fatalf("BasicAuthCredential: secret not match")
	}
}

func TestNewAPITokenCredential(t *testing.T) {
	cred := NewAPITokenCredential("john.doe@example.com", "apitoken")

	if cred.Email() != "john.doe@example.com"+"/token" {
		t.Fatalf("APITokenCredential: email not match")
	}
	if cred.Secret() != "apitoken" {
		t.Fatalf("APITokenCredential: secret not match")
	}
}
