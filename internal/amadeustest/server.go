// Package amadeustest provides the fixture-backed HTTP server the context
// tests run against.
//
// It exists so mapper tests can assert against real Amadeus payloads without a
// network, credentials, or the sandbox being up. Fixtures are captured from the
// live API by the capture tool (see internal/capture), so these tests check the
// mappers against what Amadeus actually sends rather than what the SDK's own
// structs imply it sends.
package amadeustest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/internal/amadeus"
)

// Server is a stand-in for Amadeus that answers the token endpoint itself and
// serves recorded responses for everything else.
type Server struct {
	*httptest.Server

	mu       sync.Mutex
	routes   map[string]response
	fallback *response
	requests []Recorded
}

// response is what the server returns for a route.
type response struct {
	status int
	body   []byte
}

// Recorded is one request the server received, for tests that assert on what
// the SDK sent rather than on what it decoded.
type Recorded struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   string
}

// New returns a running Server registered for cleanup with t.
func New(t *testing.T) *Server {
	t.Helper()

	s := &Server{routes: make(map[string]response)}
	s.Server = httptest.NewServer(http.HandlerFunc(s.handle))
	t.Cleanup(s.Close)
	return s
}

// handle answers the token endpoint directly and looks up everything else.
func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/security/oauth2/token") {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"test-token","token_type":"Bearer","expires_in":1799}`))
		return
	}

	body := readBody(r)

	s.mu.Lock()
	s.requests = append(s.requests, Recorded{
		Method: r.Method,
		Path:   r.URL.Path,
		Query:  r.URL.Query(),
		Header: r.Header.Clone(),
		Body:   body,
	})
	res, ok := s.routes[routeKey(r.Method, r.URL.Path)]
	if !ok && s.fallback != nil {
		res, ok = *s.fallback, true
	}
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"errors":[{"status":404,"code":1797,"title":"NOT FOUND","detail":"no fixture registered for this route"}]}`))
		return
	}

	w.WriteHeader(res.status)
	_, _ = w.Write(res.body)
}

// Fixture registers the contents of testdata/<name>.json as the reply to
// method and path.
func (s *Server) Fixture(t *testing.T, method, path, name string) *Server {
	t.Helper()
	return s.JSON(method, path, http.StatusOK, string(Load(t, name)))
}

// JSON registers a literal body as the reply to method and path. Use it for
// error cases and edge shapes that are awkward to capture from the live API.
func (s *Server) JSON(method, path string, status int, body string) *Server {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.routes[routeKey(method, path)] = response{status: status, body: []byte(body)}
	return s
}

// Always registers a reply for every route, for tests that do not care which
// endpoint was called.
func (s *Server) Always(status int, body string) *Server {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.fallback = &response{status: status, body: []byte(body)}
	return s
}

// Client returns an SDK transport pointed at this server.
func (s *Server) Client() *amadeus.Client {
	return amadeus.NewClient(amadeus.Options{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		Host:         s.URL,
		HTTPClient:   s.Server.Client(),
	})
}

// Requests returns every request the server received, in order.
func (s *Server) Requests() []Recorded {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Recorded(nil), s.requests...)
}

// LastRequest returns the most recent request, failing the test if there was
// none.
func (s *Server) LastRequest(t *testing.T) Recorded {
	t.Helper()
	requests := s.Requests()
	if len(requests) == 0 {
		t.Fatal("no request reached the server")
	}
	return requests[len(requests)-1]
}

func routeKey(method, path string) string { return method + " " + path }

func readBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}
	defer r.Body.Close()

	var sb strings.Builder
	buf := make([]byte, 4096)
	for {
		n, err := r.Body.Read(buf)
		sb.Write(buf[:n])
		if err != nil {
			break
		}
	}
	return sb.String()
}

// Load reads testdata/<name>.json relative to the calling package.
//
// It fails the test rather than skipping when a fixture is missing: a silently
// skipped mapper test is indistinguishable from a passing one, and the point of
// these tests is that they run everywhere.
func Load(t *testing.T, name string) []byte {
	t.Helper()

	path := filepath.Join("testdata", name+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading fixture %s: %v\n\nCapture it with:\n\tgo run ./internal/capture", path, err)
	}
	if !json.Valid(data) {
		t.Fatalf("fixture %s is not valid JSON", path)
	}
	return data
}
