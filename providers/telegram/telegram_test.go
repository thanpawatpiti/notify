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
				expectedURL := "https://api.telegram.org/bottest-token/sendMessage"
				if req.URL.String() != expectedURL {
					t.Errorf("expected URL %s, got %s", expectedURL, req.URL.String())
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       http.NoBody,
				}, nil
			},
		},
	}

	p := New("test-token", "test-chat", notify.WithHTTPClient(client))
	err := p.Send(context.Background(), notify.Message{Content: "test"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}
