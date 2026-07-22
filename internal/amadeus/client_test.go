package amadeus

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
)

// stubServer stands in for Amadeus. It answers the token endpoint itself and
// delegates everything else to handler, which is what lets these tests exercise
// the transport with no network and no credentials.
type stubServer struct {
	*httptest.Server
	tokenRequests atomic.Int32
	apiRequests   atomic.Int32

	mu       sync.Mutex
	lastPath string
	lastAuth string
	lastType string
	lastBody string
}

func newStubServer(t *testing.T, handler http.HandlerFunc) *stubServer {
	t.Helper()
	stub := &stubServer{}

	stub.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == tokenPath {
			stub.tokenRequests.Add(1)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"access_token":"token-%d","token_type":"Bearer","expires_in":1799}`,
				stub.tokenRequests.Load())
			return
		}

		stub.apiRequests.Add(1)
		body := make([]byte, r.ContentLength)
		if r.ContentLength > 0 {
			_, _ = r.Body.Read(body)
		}

		stub.mu.Lock()
		stub.lastPath = r.URL.RequestURI()
		stub.lastAuth = r.Header.Get("Authorization")
		stub.lastType = r.Header.Get("Content-Type")
		stub.lastBody = string(body)
		stub.mu.Unlock()

		handler(w, r)
	}))
	t.Cleanup(stub.Close)
	return stub
}

func (s *stubServer) client() *Client {
	return NewClient(Options{
		ClientID:     "id",
		ClientSecret: "secret",
		Host:         s.URL,
		HTTPClient:   s.Server.Client(),
		UserAgent:    "test-agent",
	})
}

func (s *stubServer) snapshot() (path, auth, contentType, body string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.lastPath, s.lastAuth, s.lastType, s.lastBody
}

func writeJSON(w http.ResponseWriter, status int, body string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(body))
}

type payload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestDoDecodesEnvelope(t *testing.T) {
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":[{"id":"MC","name":"Marriott"}],"meta":{"count":1}}`)
	})

	got, err := Do[[]payload](context.Background(), stub.client(), Request{Path: "/v3/hotels"})
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if len(got.Data) != 1 || got.Data[0].Name != "Marriott" {
		t.Errorf("data = %+v", got.Data)
	}
	if got.Meta.Count != 1 {
		t.Errorf("meta.count = %d, want 1", got.Meta.Count)
	}
}

func TestRequestCarriesAuthAndHeaders(t *testing.T) {
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{}}`)
	})

	_, err := Do[payload](context.Background(), stub.client(), Request{
		Path:  "/v1/reference-data/locations/hotels/by-city",
		Query: map[string][]string{"cityCode": {"PAR"}, "radius": {"5"}},
	})
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	path, auth, _, _ := stub.snapshot()
	if auth != "Bearer token-1" {
		t.Errorf("Authorization = %q, want %q", auth, "Bearer token-1")
	}
	if !strings.Contains(path, "cityCode=PAR") || !strings.Contains(path, "radius=5") {
		t.Errorf("query parameters missing from %q", path)
	}
}

func TestBookingContentTypeIsAmadeusJSON(t *testing.T) {
	// The booking endpoints reject a request body sent as application/json, so
	// this header is load-bearing rather than cosmetic.
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 201, `{"data":{"id":"ORDER1"}}`)
	})

	_, err := Do[payload](context.Background(), stub.client(), Request{
		Method:      http.MethodPost,
		Path:        "/v2/booking/hotel-orders",
		Body:        map[string]string{"hello": "world"},
		AmadeusJSON: true,
	})
	if err != nil {
		t.Fatalf("Do: %v", err)
	}

	_, _, contentType, body := stub.snapshot()
	if contentType != amadeusContentType {
		t.Errorf("Content-Type = %q, want %q", contentType, amadeusContentType)
	}
	if !strings.Contains(body, `"hello":"world"`) {
		t.Errorf("body = %q, want the encoded payload", body)
	}
}

func TestTokenIsCachedAcrossRequests(t *testing.T) {
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{}}`)
	})
	client := stub.client()

	for range 5 {
		if _, err := Do[payload](context.Background(), client, Request{Path: "/v3/x"}); err != nil {
			t.Fatalf("Do: %v", err)
		}
	}

	if got := stub.tokenRequests.Load(); got != 1 {
		t.Errorf("token fetched %d times across 5 requests, want 1", got)
	}
}

func TestConcurrentRequestsFetchOneToken(t *testing.T) {
	// A burst of goroutines starting cold must produce one token request, not
	// one per goroutine.
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{}}`)
	})
	client := stub.client()

	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := Do[payload](context.Background(), client, Request{Path: "/v3/x"}); err != nil {
				t.Errorf("Do: %v", err)
			}
		}()
	}
	wg.Wait()

	if got := stub.tokenRequests.Load(); got != 1 {
		t.Errorf("token fetched %d times for 20 concurrent requests, want 1", got)
	}
}

func TestExpiredTokenIsRetriedOnce(t *testing.T) {
	// Amadeus can revoke a token before its stated expiry. The first 401 should
	// trigger one silent refresh and retry rather than surfacing to the caller.
	var apiCalls atomic.Int32
	stub := newStubServer(t, func(w http.ResponseWriter, r *http.Request) {
		if apiCalls.Add(1) == 1 {
			writeJSON(w, 401, `{"errors":[{"status":401,"code":38191,"title":"Invalid access token"}]}`)
			return
		}
		if got := r.Header.Get("Authorization"); got != "Bearer token-2" {
			t.Errorf("retry used %q, want the refreshed token", got)
		}
		writeJSON(w, 200, `{"data":{"id":"ok"}}`)
	})

	got, err := Do[payload](context.Background(), stub.client(), Request{Path: "/v3/x"})
	if err != nil {
		t.Fatalf("Do after refresh: %v", err)
	}
	if got.Data.ID != "ok" {
		t.Errorf("data = %+v", got.Data)
	}
	if n := stub.tokenRequests.Load(); n != 2 {
		t.Errorf("token fetched %d times, want 2 (initial + refresh)", n)
	}
}

func TestPersistent401IsReturned(t *testing.T) {
	// A second 401 is a real credential problem, and retrying forever would
	// hide it.
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 401, `{"errors":[{"status":401,"code":38191,"title":"Invalid access token"}]}`)
	})

	_, err := Do[payload](context.Background(), stub.client(), Request{Path: "/v3/x"})
	if !errors.Is(err, apierr.ErrUnauthorized) {
		t.Fatalf("err = %v, want ErrUnauthorized", err)
	}
	if n := stub.apiRequests.Load(); n != 2 {
		t.Errorf("made %d API attempts, want exactly 2", n)
	}
}

func TestErrorStatusesMapToSentinels(t *testing.T) {
	cases := []struct {
		status int
		want   error
	}{
		{400, apierr.ErrInvalidRequest},
		{403, apierr.ErrForbidden},
		{404, apierr.ErrNotFound},
		{429, apierr.ErrRateLimited},
		{500, apierr.ErrServer},
		{503, apierr.ErrServer},
	}

	for _, c := range cases {
		stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
			writeJSON(w, c.status, fmt.Sprintf(`{"errors":[{"status":%d,"code":123,"title":"nope"}]}`, c.status))
		})

		_, err := Do[payload](context.Background(), stub.client(), Request{Path: "/v3/x"})
		if !errors.Is(err, c.want) {
			t.Errorf("status %d produced %v, want %v", c.status, err, c.want)
		}

		var apiErr *apierr.APIError
		if !errors.As(err, &apiErr) {
			t.Errorf("status %d: error is not an *APIError", c.status)
			continue
		}
		if apiErr.StatusCode != c.status || len(apiErr.Details) != 1 {
			t.Errorf("status %d: APIError = %+v", c.status, apiErr)
		}
	}
}

func TestErrorsEnvelopeOnA200IsStillAnError(t *testing.T) {
	// Amadeus returns 200 with an errors array for some failures. Treating that
	// as success hands the caller an empty result and no explanation.
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"errors":[{"status":400,"code":477,"title":"INVALID FORMAT"}]}`)
	})

	_, err := Do[payload](context.Background(), stub.client(), Request{Path: "/v3/x"})
	if err == nil {
		t.Fatal("a 200 carrying an errors array was treated as success")
	}
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) || len(apiErr.Details) != 1 {
		t.Fatalf("err = %v, want an APIError carrying the detail", err)
	}
	if apiErr.Details[0].Title != "INVALID FORMAT" {
		t.Errorf("detail = %+v", apiErr.Details[0])
	}
}

func TestNonJSONErrorBodyIsPreserved(t *testing.T) {
	// A misrouted request can return an HTML page from a proxy. It must not
	// crash the decoder, and the body must survive for diagnosis.
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(502)
		_, _ = w.Write([]byte("<html><body>Bad Gateway</body></html>"))
	})

	_, err := Do[payload](context.Background(), stub.client(), Request{Path: "/v3/x"})
	if !errors.Is(err, apierr.ErrServer) {
		t.Fatalf("err = %v, want ErrServer", err)
	}
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) || !strings.Contains(apiErr.Body, "Bad Gateway") {
		t.Errorf("raw body was lost: %+v", apiErr)
	}
}

func TestContextCancellationIsHonoured(t *testing.T) {
	// The whole point of threading context through: a caller must be able to
	// abandon a slow search.
	stub := newStubServer(t, func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(5 * time.Second):
			writeJSON(w, 200, `{"data":{}}`)
		case <-r.Context().Done():
		}
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := Do[payload](ctx, stub.client(), Request{Path: "/v3/slow"})
	if err == nil {
		t.Fatal("expected the request to be cancelled")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want it to wrap context.DeadlineExceeded", err)
	}
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("cancellation took %v, want it to be prompt", elapsed)
	}
}

func TestBadCredentialsSurfaceAsUnauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 401, `{"error":"invalid_client","error_description":"Client credentials are invalid"}`)
	}))
	defer server.Close()

	client := NewClient(Options{
		ClientID: "wrong", ClientSecret: "wrong",
		Host: server.URL, HTTPClient: server.Client(),
	})

	err := client.Ping(context.Background())
	if !errors.Is(err, apierr.ErrUnauthorized) {
		t.Fatalf("Ping with bad credentials = %v, want ErrUnauthorized", err)
	}
}

func TestClientsAreIndependent(t *testing.T) {
	// The singleton this replaced made two clients impossible: initialising the
	// second silently repointed the first.
	first := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{"id":"first"}}`)
	})
	second := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{"id":"second"}}`)
	})

	clientA, clientB := first.client(), second.client()

	gotA, err := Do[payload](context.Background(), clientA, Request{Path: "/v3/x"})
	if err != nil {
		t.Fatalf("client A: %v", err)
	}
	gotB, err := Do[payload](context.Background(), clientB, Request{Path: "/v3/x"})
	if err != nil {
		t.Fatalf("client B: %v", err)
	}

	if gotA.Data.ID != "first" || gotB.Data.ID != "second" {
		t.Errorf("clients interfered: A=%q B=%q", gotA.Data.ID, gotB.Data.ID)
	}
}

func TestPingReportsHealthyCredentials(t *testing.T) {
	stub := newStubServer(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, 200, `{"data":{}}`)
	})
	if err := stub.client().Ping(context.Background()); err != nil {
		t.Fatalf("Ping: %v", err)
	}
	if stub.tokenRequests.Load() != 1 {
		t.Error("Ping should have fetched exactly one token")
	}
}
