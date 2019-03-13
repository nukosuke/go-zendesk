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

	if err.Error() != fmt.Sprintf("%d: %s", status, body) {
		t.Fatal("Error did not have expected value")
	}
}
