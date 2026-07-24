package amadeus

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"time"
)

// This file is the request/response logging, and its redaction.
//
// Logging raw traffic is a debugging convenience with a sharp edge: the
// create-booking request carries a full card number, its CVV and a 3DS
// cryptogram, and the auth request carries the client secret. Writing those to
// a log is exactly the leak the booking validation is careful never to cause in
// an error message. So request bodies are redacted before they are logged, and
// the auth exchange logs neither its body nor its token.
//
// Response bodies are logged as received: Amadeus returns card numbers already
// masked, and the raw response is the thing worth seeing.
//
// Everything logs at Debug. Traffic logging is verbose and off unless the
// caller's handler is set to emit Debug, and the redaction cost is skipped
// entirely when it is not.

// sensitiveKeys are the JSON fields redacted out of a logged request body,
// lower-cased for a case-insensitive match. These are the values that must
// never reach a log: the card number, its verification code, the 3DS
// cryptogram, and the OAuth credentials.
var sensitiveKeys = map[string]bool{
	"cardnumber":      true,
	"securitycode":    true,
	"cryptogramvalue": true,
	"client_secret":   true,
	"access_token":    true,
	"authorization":   true,
}

// redactPlaceholder replaces a sensitive value in a logged body.
const redactPlaceholder = "[REDACTED]"

// logRequest records an outgoing request. The body is the domain value about to
// be encoded, redacted before it is rendered.
func (c *Client) logRequest(ctx context.Context, method, url string, body any) {
	if !c.logger.Enabled(ctx, slog.LevelDebug) {
		return
	}

	attrs := []slog.Attr{
		slog.String("method", method),
		slog.String("url", url),
	}
	if body != nil {
		attrs = append(attrs, slog.String("body", redactBody(body)))
	}
	c.logger.LogAttrs(ctx, slog.LevelDebug, "amadeus request", attrs...)
}

// logResponse records a completed response. The body is logged as received;
// Amadeus masks card numbers in responses, and the raw payload is the point.
func (c *Client) logResponse(ctx context.Context, method, url string, status int, started time.Time, body []byte) {
	if !c.logger.Enabled(ctx, slog.LevelDebug) {
		return
	}

	c.logger.LogAttrs(ctx, slog.LevelDebug, "amadeus response",
		slog.String("method", method),
		slog.String("url", url),
		slog.Int("status", status),
		slog.Duration("elapsed", time.Since(started)),
		slog.Int("bytes", len(body)),
		slog.String("body", string(body)),
	)
}

// redactBody marshals a request body and blanks its sensitive fields.
//
// It fails safe: a body it cannot parse is reported as omitted rather than
// logged raw, because "cannot parse" must never become "logged a card number
// because the shape surprised me".
func redactBody(body any) string {
	raw, err := json.Marshal(body)
	if err != nil {
		return "[body omitted: not encodable]"
	}

	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return "[body omitted: not JSON]"
	}

	redactValue(decoded)

	out, err := json.Marshal(decoded)
	if err != nil {
		return "[body omitted: not encodable after redaction]"
	}
	return string(out)
}

// redactValue walks a decoded JSON value in place, replacing the value of any
// sensitive key with the placeholder. It recurses through objects and arrays,
// so a card nested under payment.paymentCard.paymentCardInfo is still caught.
func redactValue(v any) {
	switch node := v.(type) {
	case map[string]any:
		for key, child := range node {
			if sensitiveKeys[strings.ToLower(key)] {
				node[key] = redactPlaceholder
				continue
			}
			redactValue(child)
		}
	case []any:
		for _, child := range node {
			redactValue(child)
		}
	}
}
