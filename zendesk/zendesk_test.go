package zendesk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

////////// Helper //////////

func fixture(filename string) string {
	dir, err := filepath.Abs("../fixture")
	if err != nil {
		fmt.Printf("Failed to resolve fixture directory. Check the path: %s", err)
		os.Exit(1)
	}
	return filepath.Join(dir, filename)
}

func readFixture(filename string) []byte {
	bytes, err := ioutil.ReadFile(fixture(filename))
	if err != nil {
		fmt.Printf("Failed to read fixture. Check the path: %s", err)
		os.Exit(1)
	}
	return bytes
}

func newMockAPI(method string, filename string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(readFixture(filepath.Join(method, filename)))
	}))
}

func newTestClient(mockAPI *httptest.Server) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		credential: NewAPITokenCredential("", ""),
	}
	c.SetEndpointURL(mockAPI.URL)
	return c
}

////////// Test //////////

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

func TestNewGetRequest(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetEndpointURL("http://localhost")
	client.SetCredential(NewAPITokenCredential("email", "token"))

	req, err := client.NewGetRequest("/")
	if err != nil {
		t.Fatalf("Failed to create http.Request: %s", err)
	}

	if req.Method != "GET" {
		t.Fatal("HTTP method not match")
	}
}

func TestNewPostRequest(t *testing.T) {
	client, _ := NewClient(nil)
	client.SetEndpointURL("http://localhost")
	client.SetCredential(NewAPITokenCredential("email", "token"))

	req, err := client.NewPostRequest("/", map[string]string{})
	if err != nil {
		t.Fatalf("Failed to create http.Request: %s", err)
	}

	if req.Method != "POST" {
		t.Fatal("HTTP method not match")
	}
}

func TestIncludeHeaders(t *testing.T) {
	client, _ := NewClient(nil)
	client.headers = map[string]string{
		"Header1":      "1",
		"Header2":      "2",
		"Content-Type": "application/json",
	}

	req, _ := http.NewRequest("POST", "localhost", strings.NewReader(""))
	client.includeHeaders(req)

	if len(req.Header) != 3 {
		t.Fatal("req.Header length does not match")
	}

	for k, v := range req.Header {
		switch k {
		case "Header1":
			if v[0] != "1" {
				t.Fatalf(`%s header expect "1", but got "%s"`, k, v[0])
			}
		case "Header2":
			if v[0] != "2" {
				t.Fatalf(`%s header expect "2", but got "%s"`, k, v[0])
			}
		case "Content-Type":
			if v[0] != "application/json" {
				t.Fatalf(`%s header expect "2", but got "%s"`, k, v[0])
			}
		}
	}
}
