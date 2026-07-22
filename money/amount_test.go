package money

import (
	"errors"
	"math"
	"testing"
)

func TestParseAmountRoundTrip(t *testing.T) {
	// input -> canonical rendering. Trailing zeros are dropped at parse time,
	// so "120.50" renders as "120.5"; the value is what must survive, not the
	// spelling.
	cases := []struct{ in, want string }{
		{"0", "0"},
		{"0.00", "0"},
		{"-0.0", "0"},
		{"1", "1"},
		{"120.50", "120.5"},
		{"120.55", "120.55"},
		{".75", "0.75"},
		{"5.", "5"},
		{"-3", "-3"},
		{"-0.01", "-0.01"},
		{"+42.10", "42.1"},
		{"000123.4500", "123.45"},
		{"999999999.999999999", "999999999.999999999"},
	}

	for _, c := range cases {
		got, err := ParseAmount(c.in)
		if err != nil {
			t.Errorf("ParseAmount(%q): unexpected error: %v", c.in, err)
			continue
		}
		if got.String() != c.want {
			t.Errorf("ParseAmount(%q).String() = %q, want %q", c.in, got.String(), c.want)
		}
		// The canonical form must itself parse back to an equal value.
		again, err := ParseAmount(got.String())
		if err != nil {
			t.Errorf("re-parsing %q: %v", got.String(), err)
			continue
		}
		if !again.Equal(got) {
			t.Errorf("%q did not survive a round trip", c.in)
		}
	}
}

func TestParseAmountRejectsMalformed(t *testing.T) {
	// Exponent notation and separators are rejected deliberately: silently
	// accepting "1e3" or "1,000" would let a wrong price through as a right one.
	for _, in := range []string{"", ".", "-", "abc", "1e3", "1,000", "1 2", "1.2.3", "12.5x", " 1"} {
		if got, err := ParseAmount(in); err == nil {
			t.Errorf("ParseAmount(%q) = %v, want error", in, got)
		} else if !errors.Is(err, ErrMalformedAmount) && !errors.Is(err, ErrAmountOverflow) {
			t.Errorf("ParseAmount(%q): error %v is neither malformed nor overflow", in, err)
		}
	}
}

func TestParseAmountRejectsExcessPrecision(t *testing.T) {
	if _, err := ParseAmount("1.0123456789"); !errors.Is(err, ErrAmountOverflow) {
		t.Errorf("10 fraction digits: got %v, want ErrAmountOverflow", err)
	}
	// A long run of trailing zeros is not excess precision, only excess spelling.
	if _, err := ParseAmount("1.50000000000000"); err != nil {
		t.Errorf("trailing zeros beyond maxScale should be trimmed, got %v", err)
	}
}

func TestAddIsExactAcrossScales(t *testing.T) {
	// 0.1+0.2 is the canonical float64 failure: it yields 0.30000000000000004.
	// Getting this exactly right is the entire reason this type exists.
	sum := MustParseAmount("0.1").Add(MustParseAmount("0.2"))
	if sum.String() != "0.3" {
		t.Errorf("0.1+0.2 = %s, want exactly 0.3", sum)
	}

	cases := []struct{ a, b, want string }{
		{"120.50", "9.99", "130.49"},
		{"100", "0.005", "100.005"},
		{"-5.5", "5.5", "0"},
		{"0", "42.42", "42.42"},
		{"1.999", "0.001", "2"},
	}
	for _, c := range cases {
		got := MustParseAmount(c.a).Add(MustParseAmount(c.b))
		if got.String() != c.want {
			t.Errorf("%s + %s = %s, want %s", c.a, c.b, got, c.want)
		}
	}
}

func TestSub(t *testing.T) {
	got := MustParseAmount("130.49").Sub(MustParseAmount("9.99"))
	if got.String() != "120.5" {
		t.Errorf("130.49 - 9.99 = %s, want 120.5", got)
	}
	if got := MustParseAmount("1").Sub(MustParseAmount("3")); got.String() != "-2" {
		t.Errorf("1 - 3 = %s, want -2", got)
	}
}

func TestCompareAcrossScales(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.10", "1.1", 0},
		{"1.10", "1.2", -1},
		{"2", "1.999999", 1},
		{"-1", "1", -1},
		{"0", "0.000", 0},
		{"100.00", "100", 0},
	}
	for _, c := range cases {
		if got := MustParseAmount(c.a).Compare(MustParseAmount(c.b)); got != c.want {
			t.Errorf("Compare(%s, %s) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestAddSaturatesRatherThanWrapping(t *testing.T) {
	// A wrapped total would present a huge positive charge as a negative one.
	// Saturating keeps it visibly wrong instead.
	max := Amount{units: math.MaxInt64}
	if got := max.Add(Amount{units: 1}); got.units != math.MaxInt64 {
		t.Errorf("overflow wrapped to %d, want saturation at MaxInt64", got.units)
	}
	min := Amount{units: math.MinInt64}
	if got := min.Add(Amount{units: -1}); got.units != math.MinInt64 {
		t.Errorf("underflow wrapped to %d, want saturation at MinInt64", got.units)
	}
}

func TestZeroValueIsZero(t *testing.T) {
	var a Amount
	if !a.IsZero() || a.String() != "0" {
		t.Errorf("zero Amount = %q, IsZero=%v; want \"0\", true", a.String(), a.IsZero())
	}
}

func TestFloat64(t *testing.T) {
	if got := MustParseAmount("120.5").Float64(); got != 120.5 {
		t.Errorf("Float64() = %v, want 120.5", got)
	}
}
