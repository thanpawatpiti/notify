package line

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thanpawatpiti/notify"
)

func TestSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("expected Authorization header 'Bearer test-token', got %s", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Override the API endpoint for testing by using a custom HTTP client that redirects requests?
	// Actually, since the URL is hardcoded in the package, we can't easily swap it without exporting it or using a variable.
	// For "Professional" code, we should probably allow overriding the base URL or just mock the transport.
	// Let's use a custom transport to intercept the request to the hardcoded URL.

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
