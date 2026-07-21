// Package searchcriteria holds the sets of values Amadeus accepts in hotel
// search filters: amenities, star ratings, board types and so on.
//
// The request DTOs used to declare these as plain strings, which left callers
// copying codes out of doc comments. Each set here is a named string type with
// its constants, an All* function for enumerating it (rendering a filter UI,
// say) and Label/IsValid methods, so a wrong code fails to compile instead of
// coming back as an Amadeus 400.
//
// The codes are what the live API accepts, which is not always what Amadeus
// documents; see the note on Amenity.
package searchcriteria

import "strings"

// entry pairs a code with its human-readable label. One ordered catalog per
// type drives All*, Label and IsValid, so there is no second list to drift.
type entry[T ~string] struct {
	Code  T
	Label string
}

// codes returns the catalog's codes in declaration order.
func codes[T ~string](catalog []entry[T]) []T {
	out := make([]T, len(catalog))
	for i, e := range catalog {
		out[i] = e.Code
	}
	return out
}

// labelOf returns the label for code, or "" when the catalog has no such code.
func labelOf[T ~string](catalog []entry[T], code T) string {
	for _, e := range catalog {
		if e.Code == code {
			return e.Label
		}
	}
	return ""
}

// isValid reports whether code appears in the catalog.
func isValid[T ~string](catalog []entry[T], code T) bool {
	for _, e := range catalog {
		if e.Code == code {
			return true
		}
	}
	return false
}

// Ptr returns a pointer to v. HotelListByCityCodeRequest takes its optional
// scalars as pointers to tell "unset" apart from the zero value, and a constant
// is not addressable, so searchcriteria.Ptr(searchcriteria.RadiusUnitKM) saves
// callers a temporary variable.
func Ptr[T any](v T) *T { return &v }

// Join renders typed codes as the comma-separated list Amadeus expects in a
// query parameter. It returns "" for an empty slice, which callers use to skip
// the parameter entirely.
func Join[T ~string](values []T) string {
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = string(v)
	}
	return strings.Join(parts, ",")
}
