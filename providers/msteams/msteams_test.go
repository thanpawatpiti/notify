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

	// Test 1: CommonMessage
	err := p.Send(context.Background(), notify.CommonMessage{Content: "test"})
	if err != nil {
		t.Errorf("CommonMessage: expected no error, got %v", err)
	}

	// Test 2: AdaptiveCard
	card := AdaptiveCard{
		Type: "AdaptiveCard",
		Body: []interface{}{
			TextBlock{Text: "Test"},
		},
	}
	err = p.Send(context.Background(), card)
	if err != nil {
		t.Errorf("AdaptiveCard: expected no error, got %v", err)
	}
}
