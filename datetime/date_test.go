package datetime

import (
	"encoding/json"
	"testing"
	"time"
)

func TestParseDateRoundTrip(t *testing.T) {
	for _, s := range []string{"2026-07-22", "2026-01-01", "2026-12-31", "2024-02-29"} {
		d, err := ParseDate(s)
		if err != nil {
			t.Errorf("ParseDate(%q): %v", s, err)
			continue
		}
		if d.String() != s {
			t.Errorf("ParseDate(%q).String() = %q", s, d.String())
		}
	}
}

func TestParseEmptyIsZeroNotError(t *testing.T) {
	d, err := ParseDate("")
	if err != nil {
		t.Fatalf("ParseDate(\"\"): %v", err)
	}
	if !d.IsZero() || d.String() != "" {
		t.Errorf("empty date = %q, IsZero=%v", d.String(), d.IsZero())
	}
}

func TestParseDateRejectsMalformed(t *testing.T) {
	for _, s := range []string{"22-07-2026", "2026/07/22", "2026-13-01", "2026-02-30", "today"} {
		if _, err := ParseDate(s); err == nil {
			t.Errorf("ParseDate(%q) succeeded, want error", s)
		}
	}
}

// A calendar date must not move when it is rendered in another timezone. This
// is the failure the type exists to prevent: a check-in date parsed as a
// time.Time in UTC and formatted in Pacific/Auckland lands on the next day.
func TestDateDoesNotShiftAcrossTimezones(t *testing.T) {
	const checkIn = "2026-07-22"
	d := MustParseDate(checkIn)

	for _, zone := range []string{"UTC", "Pacific/Auckland", "America/Los_Angeles", "Asia/Kolkata"} {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			t.Skipf("timezone database unavailable: %v", err)
		}
		if got := d.String(); got != checkIn {
			t.Errorf("date rendered as %q", got)
		}
		if got := d.Time(loc).Format("2006-01-02"); got != checkIn {
			t.Errorf("in %s the date became %q, want %q", zone, got, checkIn)
		}
	}
}

func TestCompareAndOrdering(t *testing.T) {
	early := MustParseDate("2026-07-22")
	late := MustParseDate("2026-07-25")

	if !early.Before(late) || !late.After(early) {
		t.Error("ordering is wrong for two dates in the same month")
	}
	if early.Compare(early) != 0 {
		t.Error("a date should compare equal to itself")
	}
	// Ordering must be driven by the calendar, not by field-by-field magnitude:
	// December 2025 precedes January 2026 despite the larger month number.
	if !MustParseDate("2025-12-31").Before(MustParseDate("2026-01-01")) {
		t.Error("year should outrank month in ordering")
	}
}

func TestAddDaysCrossesMonthAndYearBoundaries(t *testing.T) {
	cases := []struct {
		from string
		days int
		want string
	}{
		{"2026-07-22", 3, "2026-07-25"},
		{"2026-07-31", 1, "2026-08-01"},
		{"2026-12-31", 1, "2027-01-01"},
		{"2026-01-01", -1, "2025-12-31"},
		{"2024-02-28", 1, "2024-02-29"}, // leap year
		{"2025-02-28", 1, "2025-03-01"}, // non-leap
	}
	for _, c := range cases {
		if got := MustParseDate(c.from).AddDays(c.days).String(); got != c.want {
			t.Errorf("%s + %d days = %s, want %s", c.from, c.days, got, c.want)
		}
	}
}

func TestDaysUntilIsUnaffectedByDaylightSaving(t *testing.T) {
	// Europe/Paris shifts on 2026-03-29. Counting these nights in local time
	// would give a non-integer number of days and truncate to 2.
	nights := MustParseDate("2026-03-28").DaysUntil(MustParseDate("2026-03-31"))
	if nights != 3 {
		t.Errorf("nights across a DST boundary = %d, want 3", nights)
	}
	if got := MustParseDate("2026-07-25").DaysUntil(MustParseDate("2026-07-22")); got != -3 {
		t.Errorf("backwards DaysUntil = %d, want -3", got)
	}
}

func TestJSONRoundTrip(t *testing.T) {
	type stay struct {
		CheckIn  Date `json:"checkInDate"`
		CheckOut Date `json:"checkOutDate"`
	}

	encoded, err := json.Marshal(stay{
		CheckIn:  MustParseDate("2026-07-22"),
		CheckOut: MustParseDate("2026-07-25"),
	})
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	const want = `{"checkInDate":"2026-07-22","checkOutDate":"2026-07-25"}`
	if string(encoded) != want {
		t.Errorf("Marshal = %s, want %s", encoded, want)
	}

	var decoded stay
	if err := json.Unmarshal([]byte(want), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded.CheckIn.String() != "2026-07-22" || decoded.CheckOut.String() != "2026-07-25" {
		t.Errorf("decoded = %+v", decoded)
	}
}

func TestUnmarshalAcceptsNullAndEmpty(t *testing.T) {
	for _, raw := range []string{`null`, `""`} {
		var d Date
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Errorf("Unmarshal(%s): %v", raw, err)
			continue
		}
		if !d.IsZero() {
			t.Errorf("Unmarshal(%s) = %v, want zero Date", raw, d)
		}
	}

	var d Date
	if err := json.Unmarshal([]byte(`12345`), &d); err == nil {
		t.Error("Unmarshal of a JSON number succeeded, want error")
	}
}

func TestNewDateNormalises(t *testing.T) {
	if got := NewDate(2026, 13, 1).String(); got != "2027-01-01" {
		t.Errorf("NewDate(2026, 13, 1) = %s, want 2027-01-01", got)
	}
}
