package line

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
				if req.URL.String() != lineMessagingAPI {
					t.Errorf("expected URL %s, got %s", lineMessagingAPI, req.URL.String())
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       http.NoBody,
				}, nil
			},
		},
	}

	p := New("test-token", "test-user", notify.WithHTTPClient(client))

	// Test 1: CommonMessage
	err := p.Send(context.Background(), notify.CommonMessage{Content: "test"})
	if err != nil {
		t.Errorf("CommonMessage: expected no error, got %v", err)
	}

	// Test 2: FlexMessage
	flexMsg := FlexMessage{
		AltText: "Flex",
		Contents: BubbleContainer{
			Type: "bubble",
			Body: &BoxComponent{
				Type:   "box",
				Layout: "vertical",
				Contents: []FlexComponent{
					TextComponent{Type: "text", Text: "Hello"},
				},
			},
		},
	}
	err = p.Send(context.Background(), flexMsg)
	if err != nil {
		t.Errorf("FlexMessage: expected no error, got %v", err)
	}
}

type mockTransport struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTrip(req)
}
