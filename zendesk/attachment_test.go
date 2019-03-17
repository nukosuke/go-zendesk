package zendesk

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	file := readFixture(filepath.Join("POST", "upload.json"))
	h := sha1.New()
	h.Write(file)
	expectedSum := h.Sum(nil)
	r := bytes.NewReader(file)
	var attachmentSum []byte
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := sha1.New()
		io.Copy(h, r.Body)
		attachmentSum = h.Sum(nil)
		w.WriteHeader(http.StatusCreated)
		w.Write(file)
	}))

	c := newTestClient(mockAPI)
	w := c.UploadAttachment("foo", "bar")
	_, err := io.Copy(w, r)
	if err != nil {
		t.Fatal("Received an error from write")
	}

	out, err := w.Close()
	if err != nil {
		t.Fatalf("Received an error from close %v", err)
	}

	expected := "6bk3gql82em5nmf"
	if out.Token != expected {
		t.Fatalf("Received an unexpected token %s expected %s", out.Token, expected)
	}

	if !reflect.DeepEqual(expectedSum, attachmentSum) {
		t.Fatalf("Check sum of the written file does not match the expected checksum")
	}
}
