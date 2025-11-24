package msteams

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thanpawatpiti/notify"
)

func TestSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	p := New(server.URL)
	err := p.Send(context.Background(), notify.Message{Content: "test"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
