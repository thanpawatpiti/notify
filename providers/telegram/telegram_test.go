package telegram

import (
	"context"
	"net/http"
	"testing"

	"github.com/thanpawatpiti/notify"
)

func TestSend(t *testing.T) {
	client := &http.Client{
		Transport: &mockTransport{
			roundTrip: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       http.NoBody,
				}, nil
			},
		},
	}

	p := New("test-token", "test-chat", notify.WithHTTPClient(client))

	// Test 1: CommonMessage
	err := p.Send(context.Background(), notify.CommonMessage{Content: "test"})
	if err != nil {
		t.Errorf("CommonMessage: expected no error, got %v", err)
	}

	// Test 2: Payload
	payload := Payload{
		Text:      "Advanced",
		ParseMode: "Markdown",
	}
	err = p.Send(context.Background(), payload)
	if err != nil {
		t.Errorf("Payload: expected no error, got %v", err)
	}
}

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}
