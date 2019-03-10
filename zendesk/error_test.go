package zendesk

import (
	"fmt"
	"net/http"
	"testing"
)

func TestError_Response(t *testing.T) {
	resp := &http.Response{}
	err := Error{
		msg:  "foo",
		resp: resp,
	}

	if err.Response() != resp {
		t.Fatal("Response did not  return the provided response")
	}
}

func TestError_Error(t *testing.T) {
	status := http.StatusOK
	resp := &http.Response{
		StatusCode: status,
	}
	body := "foo"
	err := Error{
		msg:  body,
		resp: resp,
	}

	if err.Error() != fmt.Sprintf("%d: %s", status, body) {
		t.Fatal("Error did not have expected value")
	}
}
