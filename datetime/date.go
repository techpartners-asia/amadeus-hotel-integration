// Package datetime holds the calendar value objects the SDK exchanges with
// Amadeus.
//
// A check-in date is a calendar date, not an instant. Parsing "2026-07-22" into
// a time.Time forces a timezone onto it, and the moment that value is formatted
// in a different zone it can shift to the 21st or the 23rd - a booking silently
// moved by a day. Date carries the three fields Amadeus actually sent and
// nothing more.
package datetime

import (
	"fmt"
	"time"
)

// wireLayout is the ISO 8601 calendar-date layout used by every Amadeus hotel
// endpoint.
const wireLayout = "2006-01-02"

// Date is a calendar date with no time and no timezone.
//
// The zero Date is not a real date; use IsZero to test for one. Amadeus omits
// optional dates rather than sending a placeholder, so the zero value is how an
// absent date is represented.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// NewDate returns the given calendar date. It normalises out-of-range
// components the way time.Date does, so NewDate(2026, 13, 1) is January 2027.
func NewDate(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return Date{Year: t.Year(), Month: t.Month(), Day: t.Day()}
}

// ParseDate parses the "2006-01-02" form Amadeus sends. An empty string parses
// to the zero Date without error, since an omitted optional date is not a
// malformed one.
func ParseDate(s string) (Date, error) {
	if s == "" {
		return Date{}, nil
	}
	t, err := time.Parse(wireLayout, s)
	if err != nil {
		return Date{}, fmt.Errorf("datetime: parsing date %q: %w", s, err)
	}
	return Date{Year: t.Year(), Month: t.Month(), Day: t.Day()}, nil
}

// MustParseDate is ParseDate for dates known to be well-formed, such as test
// data. It panics on malformed input.
func MustParseDate(s string) Date {
	d, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return d
}

// Today returns the current date in loc, which callers need when defaulting a
// check-in date. Pass time.Local for the user's calendar; the answer genuinely
// depends on the zone, so there is no sensible default.
func Today(loc *time.Location) Date {
	now := time.Now().In(loc)
	return Date{Year: now.Year(), Month: now.Month(), Day: now.Day()}
}

// IsZero reports whether d is the zero Date, which is how an absent date is
// represented.
func (d Date) IsZero() bool { return d == Date{} }

// String renders the date as "2026-07-22", the form Amadeus expects in query
// parameters and request bodies. The zero Date renders as "".
func (d Date) String() string {
	if d.IsZero() {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, int(d.Month), d.Day)
}

// Time returns midnight on d in loc, for callers that must hand the date to an
// API demanding a time.Time. Choosing the zone is the caller's decision, which
// is exactly the decision this type exists to avoid making implicitly.
func (d Date) Time(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

// Compare returns -1, 0 or +1 as d is before, equal to or after other.
func (d Date) Compare(other Date) int {
	switch {
	case d.Year != other.Year:
		return sign(d.Year - other.Year)
	case d.Month != other.Month:
		return sign(int(d.Month) - int(other.Month))
	default:
		return sign(d.Day - other.Day)
	}
}

// Before reports whether d falls before other.
func (d Date) Before(other Date) bool { return d.Compare(other) < 0 }

// After reports whether d falls after other.
func (d Date) After(other Date) bool { return d.Compare(other) > 0 }

// AddDays returns the date n days after d, and accepts a negative n.
func (d Date) AddDays(n int) Date {
	t := d.Time(time.UTC).AddDate(0, 0, n)
	return Date{Year: t.Year(), Month: t.Month(), Day: t.Day()}
}

// DaysUntil returns the number of days from d to other, negative when other is
// earlier. It is computed in UTC, where every day is 24 hours, so daylight
// saving cannot skew the count.
func (d Date) DaysUntil(other Date) int {
	const day = 24 * time.Hour
	return int(other.Time(time.UTC).Sub(d.Time(time.UTC)) / day)
}

// MarshalJSON renders the date as a JSON string in Amadeus's wire format, and
// the zero Date as null.
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.String() + `"`), nil
}

// UnmarshalJSON parses Amadeus's wire format, accepting null and "" as the zero
// Date.
func (d *Date) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		*d = Date{}
		return nil
	}
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return fmt.Errorf("datetime: date must be a JSON string, got %s", s)
	}

	parsed, err := ParseDate(s[1 : len(s)-1])
	if err != nil {
		return err
	}
	*d = parsed
	return nil
}

func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	default:
		return 0
	}
}
