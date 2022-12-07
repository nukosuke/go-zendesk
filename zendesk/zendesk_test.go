package zendesk

import (
	"fmt"
	"io"
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
	bytes, err := os.ReadFile(fixture(filename))
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

func newMockAPIWithStatus(method string, filename string, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
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

func TestBearerAuthCredential(t *testing.T) {
	client, _ := NewClient(nil)
	cred := NewBearerTokenCredential("hello")
	client.SetCredential(cred)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer hello" {
			t.Fatalf("unexpected auth header: " + auth)
		}
	}))
	client.SetEndpointURL(server.URL)
	defer server.Close()

	// trigger request, assert in the server code
	_, _ = client.get(ctx, "/groups.json")
}

func TestGet(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "groups.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	body, err := client.get(ctx, "/groups.json")
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}
}

func TestGetFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "groups.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.get(ctx, "/groups.json")
	if err == nil {
		t.Fatal("Did not receive error from client")
	}

	if _, ok := err.(Error); !ok {
		t.Fatalf("Did not return a zendesk error %s", err)
	}
}

func TestGetFailureCanReadErrorBody(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodGet, "groups.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.get(ctx, "/groups.json")
	if err == nil {
		t.Fatal("Did not receive error from client")
	}

	clientErr, ok := err.(Error)
	if !ok {
		t.Fatalf("Did not return a zendesk error %s", err)
	}

	body := clientErr.Body()
	_, err = io.ReadAll(body)
	if err != nil {
		t.Fatal("Client received error while reading client body")
	}

	err = body.Close()
	if err != nil {
		t.Fatal("Client received error while closing body")
	}
}

func TestPost(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "groups.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	body, err := client.post(ctx, "/groups.json", Group{})
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}
}

func TestPostFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "groups.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.post(ctx, "/groups.json", Group{})
	if err == nil {
		t.Fatal("Did not receive error from client")
	}

	if _, ok := err.(Error); !ok {
		t.Fatalf("Did not return a zendesk error %s", err)
	}
}

func TestPut(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "groups.json", http.StatusOK)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	body, err := client.put(ctx, "/groups.json", Group{})
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}
}

func TestPutFailure(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPut, "groups.json", http.StatusInternalServerError)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	_, err := client.put(ctx, "/groups.json", Group{})
	if err == nil {
		t.Fatal("Did not receive error from client")
	}

	if _, ok := err.(Error); !ok {
		t.Fatalf("Did not return a zendesk error %s", err)
	}
}

func TestDelete(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.delete(ctx, "/foo/id")
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}
}

func TestDeleteFailure(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.delete(ctx, "/foo/id")
	if err == nil {
		t.Fatalf("Failed to recieve error from Delete")
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

func TestAddOptions(t *testing.T) {
	ep := "/triggers.json"
	ops := &TriggerListOptions{
		PageOptions: PageOptions{
			PerPage: 10,
			Page:    2,
		},
		Active: true,
	}
	expected := "/triggers.json?active=true&page=2&per_page=10"

	u, err := addOptions(ep, ops)
	if err != nil {
		t.Fatal(err)
	}

	if u != expected {
		t.Fatalf("\nExpect:\t%s\nGot:\t%s", expected, u)
	}
}
