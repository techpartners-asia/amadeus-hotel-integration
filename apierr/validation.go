package apierr

import (
	"errors"
	"fmt"
	"strings"
)

// ErrValidation is the sentinel every ValidationError wraps, so callers can tell
// "I built the request wrong" apart from "Amadeus rejected it" without knowing
// which field was at fault:
//
//	if errors.Is(err, apierr.ErrValidation) { ... }
var ErrValidation = errors.New("invalid argument")

// ValidationError reports a request the SDK rejected before sending it.
//
// These checks exist where the SDK can be certain: a missing required field, a
// date range that runs backwards, a code outside the set Amadeus documents.
// Catching them locally turns a round trip and an opaque 400 into an immediate
// error naming the field. The SDK does not attempt to replicate Amadeus's full
// validation - anything uncertain is left to the API, which is authoritative.
type ValidationError struct {
	// Field is the domain field at fault, named as the caller wrote it
	// (e.g. "CheckOut", "Guests.Adults").
	Field string
	// Reason explains what is wrong with it.
	Reason string
}

// Invalid returns a ValidationError for field.
func Invalid(field, reason string) *ValidationError {
	return &ValidationError{Field: field, Reason: reason}
}

// Invalidf returns a ValidationError whose reason is formatted.
func Invalidf(field, format string, args ...any) *ValidationError {
	return &ValidationError{Field: field, Reason: fmt.Sprintf(format, args...)}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Reason)
}

// Unwrap makes errors.Is(err, ErrValidation) report true.
func (e *ValidationError) Unwrap() error { return ErrValidation }

// ValidationErrors is a set of validation failures reported together, so one
// call surfaces every problem with a request rather than only the first.
type ValidationErrors []*ValidationError

// Error lists every failure, one per clause.
func (e ValidationErrors) Error() string {
	switch len(e) {
	case 0:
		return "invalid request"
	case 1:
		return e[0].Error()
	}

	parts := make([]string, len(e))
	for i, err := range e {
		parts[i] = err.Error()
	}
	return fmt.Sprintf("%d validation errors: %s", len(e), strings.Join(parts, "; "))
}

// Unwrap returns the individual errors, so errors.Is and errors.As reach them.
// This is the multi-error form Go 1.20 added to errors.Join.
func (e ValidationErrors) Unwrap() []error {
	out := make([]error, len(e))
	for i, err := range e {
		out[i] = err
	}
	return out
}

// Append adds a failure to the set when reason is non-empty, which lets a
// validate method read as a straight list of checks.
func (e ValidationErrors) Append(field, reason string) ValidationErrors {
	if reason == "" {
		return e
	}
	return append(e, Invalid(field, reason))
}

// OrNil returns the set as an error, or a nil error when it is empty. Returning
// e directly would produce a non-nil error interface holding an empty slice,
// which is the classic Go nil-interface trap.
func (e ValidationErrors) OrNil() error {
	if len(e) == 0 {
		return nil
	}
	return e
}
