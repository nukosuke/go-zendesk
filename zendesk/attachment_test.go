package zendesk

import (
	"bytes"
	"context"
	"crypto/sha1"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	file := readFixture(filepath.Join(http.MethodPost, "upload.json"))
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
	w := c.UploadAttachment(ctx, "foo", "bar")

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

func TestWriteCancelledContext(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket.json",  201)
	defer mockAPI.Close()

	client := newTestClient(mockAPI)

	canceled, cancelFunc := context.WithCancel(ctx)
	cancelFunc()
	w := client.UploadAttachment(canceled, "foo", "bar")

	file := []byte("body")
	r := bytes.NewBuffer(file)

	_, err := io.Copy(w, r)
	if err != nil {
		t.Fatal("Received an error from write")
	}

	_, err = w.Close()
	if err == nil {
		t.Fatal("Did not receive error when closing writer")
	} else if err != context.Canceled {
		t.Fatalf("did not receive expected error was: %v", err)
	}
}

func TestDeleteUpload(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}))

	c := newTestClient(mockAPI)
	err := c.DeleteUpload(ctx, "foobar")
	if err != nil {
		t.Fatalf("Failed to delete ticket field: %s", err)
	}
}

func TestGetAttachment(t *testing.T) {
	mockAPI := newMockAPI(http.MethodGet, "attachment.json")
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	attachment, err := client.GetAttachment(ctx, 123)
	if err != nil {
		t.Fatalf("Failed to get attachment: %s", err)
	}

	expectedID := int64(498483)
	if attachment.ID != expectedID {
		t.Fatalf("Returned attachment does not have the expected ID %d. Attachment id is %d", expectedID, attachment.ID)
	}
}
