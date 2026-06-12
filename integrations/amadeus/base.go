package amadeusIntegration

import (
	"errors"
	"sync"
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	"resty.dev/v3"
)

// authPath is the OAuth2 client-credentials token endpoint (relative to AUTH_BASE_URL).
const authPath = "/security/oauth2/token"

// tokenSafetyWindow refreshes the token slightly before it actually expires,
// so an in-flight request never races the expiry.
const tokenSafetyWindow = 30 * time.Second

// manager is the process-wide token manager, initialised by Init.
var manager *tokenManager

// tokenManager caches the OAuth2 access token and transparently refreshes it
// when it is missing or about to expire.
type tokenManager struct {
	mu        sync.Mutex
	id        string
	secret    string
	token     string
	expiresAt time.Time
	client    *resty.Client // bare client used only for the auth call (no auth middleware)
}

// Init authenticates with Amadeus and prepares the token manager. It must be
// called once before constructing use-cases. It returns an error instead of
// terminating the process so callers can handle credential failures.
func Init(id, secret string) error {
	m := &tokenManager{
		id:     id,
		secret: secret,
		client: resty.New().
			SetBaseURL(constants.AUTH_BASE_URL).
			SetHeader("Accept", "application/json"),
	}

	// Validate credentials eagerly so misconfiguration fails fast.
	if _, err := m.validToken(); err != nil {
		return err
	}

	manager = m
	return nil
}

// NewClient returns a resty client bound to baseURL that injects a fresh Bearer
// token before every request. Each module gets its own client, so setting a base
// URL on one never affects another.
func NewClient(baseURL string) *resty.Client {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json")

	client.AddRequestMiddleware(func(_ *resty.Client, r *resty.Request) error {
		token, err := manager.validToken()
		if err != nil {
			return err
		}
		r.SetHeader("Authorization", "Bearer "+token)
		return nil
	})

	return client
}

// validToken returns a non-expired access token, refreshing it if necessary.
func (m *tokenManager) validToken() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.token != "" && time.Now().Before(m.expiresAt.Add(-tokenSafetyWindow)) {
		return m.token, nil
	}

	auth, err := m.authenticate()
	if err != nil {
		return "", err
	}

	m.token = auth.AccessToken
	m.expiresAt = time.Now().Add(time.Duration(auth.ExpiresIn) * time.Second)
	return m.token, nil
}

// authenticate performs the OAuth2 client-credentials exchange.
func (m *tokenManager) authenticate() (*AuthResponse, error) {
	var authResponse AuthResponse

	res, err := m.client.R().
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     m.id,
			"client_secret": m.secret,
		}).
		SetResult(&authResponse).
		Post(authPath)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &authResponse, nil
}
