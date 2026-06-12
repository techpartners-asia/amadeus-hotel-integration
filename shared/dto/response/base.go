package sharedResponseDTO

import (
	"encoding/json"
	"fmt"
)

type (
	BaseResponse[T any] struct {
		Data   T               `json:"data"`
		Errors []ErrorResponse `json:"errors"`
		Meta   MetaResponse    `json:"meta"`
	}

	LinkResponse struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		Self string `json:"self"`
	}

	MetaResponse struct {
		Count int `json:"count"`
		// Links []LinkResponse `json:"links"`
	}

	// ErrorResponse is a single Amadeus error object. It mirrors the error schema
	// shared by all of the Amadeus hotel APIs.
	ErrorResponse struct {
		// Status - HTTP status code of the response.
		Status int `json:"status"`
		// Code - machine-readable Amadeus error code (e.g. 38196).
		Code int `json:"code"`
		// Title - short, human-readable summary with 1:1 correspondence to Code.
		Title string `json:"title"`
		// Detail - human-readable explanation specific to this occurrence.
		Detail string `json:"detail"`
		// Source - identifies the request element that caused the error.
		Source ErrorSource `json:"source"`
		// Documentation - link to further documentation about the error.
		Documentation string `json:"documentation,omitempty"`
	}

	// ErrorSource identifies the request element that triggered an error.
	ErrorSource struct {
		// Parameter - the URI path or query parameter key that caused the error.
		Parameter string `json:"parameter,omitempty"`
		// Pointer - JSON Pointer (RFC 6901) to the offending entity in the request body.
		Pointer string `json:"pointer,omitempty"`
		// Example - a sample value to guide the user when resolving the issue.
		Example string `json:"example,omitempty"`
	}

	GeoCodeResponse struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)

// APIError is the typed error returned by every SDK call that fails. It wraps the
// HTTP status and the structured Amadeus error objects so callers can inspect them
// (use errors.As to recover it):
//
//	var apiErr *sharedResponseDTO.APIError
//	if errors.As(err, &apiErr) {
//	    if apiErr.StatusCode == 404 { ... }
//	    for _, e := range apiErr.Errors { ... }
//	}
type APIError struct {
	// StatusCode - HTTP status code of the failed response.
	StatusCode int
	// Errors - structured Amadeus error objects parsed from the response body
	// (may be empty if the body was not a standard error envelope).
	Errors []ErrorResponse
	// Raw - the raw response body, kept as a fallback for non-standard errors.
	Raw string
}

func (e *APIError) Error() string {
	if len(e.Errors) > 0 {
		first := e.Errors[0]
		msg := fmt.Sprintf("amadeus: [%d] %s", first.Code, first.Title)
		if first.Detail != "" {
			msg += " - " + first.Detail
		}
		status := first.Status
		if status == 0 {
			status = e.StatusCode
		}
		msg += fmt.Sprintf(" (status %d)", status)
		if len(e.Errors) > 1 {
			msg += fmt.Sprintf(" (+%d more)", len(e.Errors)-1)
		}
		return msg
	}
	return fmt.Sprintf("amadeus: request failed (status %d): %s", e.StatusCode, e.Raw)
}

// ErrorFromResponse returns a typed *APIError when a response represents a failure,
// or nil otherwise. It parses the Amadeus "errors" array from the body regardless of
// the HTTP status, so it also catches 200 responses that carry an errors envelope.
func ErrorFromResponse(statusCode int, isError bool, body string) error {
	var env struct {
		Errors []ErrorResponse `json:"errors"`
	}
	_ = json.Unmarshal([]byte(body), &env)

	if len(env.Errors) > 0 {
		return &APIError{StatusCode: statusCode, Errors: env.Errors, Raw: body}
	}
	if isError {
		return &APIError{StatusCode: statusCode, Raw: body}
	}
	return nil
}
