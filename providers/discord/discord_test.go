package discord

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

	// Test 1: CommonMessage
	err := p.Send(context.Background(), notify.CommonMessage{Content: "test"})
	if err != nil {
		t.Errorf("CommonMessage: expected no error, got %v", err)
	}

	// Test 2: Embed
	embed := Embed{
		Title:       "Test Embed",
		Description: "Desc",
	}
	err = p.Send(context.Background(), embed)
	if err != nil {
		t.Errorf("Embed: expected no error, got %v", err)
	}
}
