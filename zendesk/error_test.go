package zendesk

import (
	"fmt"
	"net/http"
	"testing"
)

func TestError_Error(t *testing.T) {
	status := http.StatusOK
	resp := &http.Response{
		StatusCode: status,
	}
	body := []byte("foo")
	err := Error{
		body: body,
		resp: resp,
	}

	expected := fmt.Sprintf("%d: %s", status, body)
	if v := err.Error(); v != expected {
		t.Fatalf("Error %s did not have expected value %s", v, expected)
	}
}

func TestError_ErrorEmptyBody(t *testing.T) {
	status := http.StatusOK
	resp := &http.Response{
		StatusCode: status,
	}
	err := Error{
		resp: resp,
	}

	expected := fmt.Sprintf("%d: %s", status, http.StatusText(status))
	if v := err.Error(); v != expected {
		t.Fatalf("Error %s did not have expected value %s", v, expected)
	}
}

func TestError_Headers(t *testing.T) {
	retryAfter := "Retry-After"
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			retryAfter: []string{"92"},
		},
	}

	err := Error{
		resp: resp,
	}

	if _, ok := err.Headers()[retryAfter]; !ok {
		t.Fatal("Could not get header values from zendesk error")
	}
}

func TestError_Status(t *testing.T) {
	retryAfter := "Retry-After"
	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			retryAfter: []string{"92"},
		},
	}

	err := Error{
		resp: resp,
	}

	if status := err.Status(); status != http.StatusTooManyRequests {
		t.Fatal("Status returned from error was not the correct status code")
	}
}
