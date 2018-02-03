package zendesk

import (
	"github.com/zenform/go-zendesk/common"
	"net/http"
	"testing"
)

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

func TestSetCredential(t *testing.T) {
	cred := &common.Credential{
		AuthType: common.APIToken,
		Email:    "zenform@example.com",
		APIToken: "0123456789abcdefgh",
	}

	client, _ := NewClient(&http.Client{})
	if err := client.SetCredential(cred); err != nil {
		t.Fatal("SetCredential should success")
	}
}
