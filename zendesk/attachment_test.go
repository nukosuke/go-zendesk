package zendesk

import (
	"bytes"
	ctx "context"
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
	w := c.UploadAttachment(ctx.Background(), "foo")
	w.SetToken("bar")

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

func TestDeleteUpload(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteUpload(ctx.Background(), "foobar")
	if err != nil {
		t.Fatalf("Failed to delete ticket field: %s", err)
	}
}

func TestGetAttachment(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "attachment.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	attachment, err := client.GetAttachment(ctx.Background(), 123)
	if err != nil {
		t.Fatalf("Failed to get attachment: %s", err)
	}

	expectedID := int64(498483)
	if attachment.ID != expectedID {
		t.Fatalf("Returned attachment does not have the expected ID %d. Attachment id is %d", expectedID, attachment.ID)
	}
}
