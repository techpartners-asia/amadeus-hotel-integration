package geo

import (
	"errors"
	"math"
	"testing"
)

func TestNewCoordinatesAcceptsValidPoints(t *testing.T) {
	cases := []struct{ lat, lon float64 }{
		{48.8566, 2.3522}, // Paris
		{0, 0},
		{-90, -180},
		{90, 180},
	}
	for _, c := range cases {
		if _, err := NewCoordinates(c.lat, c.lon); err != nil {
			t.Errorf("NewCoordinates(%g, %g): %v", c.lat, c.lon, err)
		}
	}
}

func TestNewCoordinatesRejectsOutOfRange(t *testing.T) {
	cases := []struct{ lat, lon float64 }{
		{91, 0},
		{-90.1, 0},
		{0, 181},
		{0, -180.5},
		{math.NaN(), 0},
		{0, math.NaN()},
	}
	for _, c := range cases {
		if _, err := NewCoordinates(c.lat, c.lon); !errors.Is(err, ErrOutOfRange) {
			t.Errorf("NewCoordinates(%g, %g) = %v, want ErrOutOfRange", c.lat, c.lon, err)
		}
	}
}

func TestCoordinatesString(t *testing.T) {
	c := Coordinates{Latitude: 48.8566, Longitude: 2.3522}
	if got := c.String(); got != "48.8566,2.3522" {
		t.Errorf("String() = %q", got)
	}
}

func TestUnitValidity(t *testing.T) {
	for _, u := range AllUnits() {
		if !u.IsValid() {
			t.Errorf("%s should be valid", u)
		}
	}
	for _, u := range []Unit{"", "YARDS", "km"} {
		if u.IsValid() {
			t.Errorf("%q should not be valid", u)
		}
	}
}

func TestDistanceDefaultsToKilometers(t *testing.T) {
	// Amadeus's own default is KM, so an unset unit must mean KM and not an
	// unitless number.
	d := NewDistance(5, "")
	if d.Unit != Kilometers {
		t.Errorf("unit = %q, want KM", d.Unit)
	}
	if got := (Distance{Value: 5}).Meters(); got != 5000 {
		t.Errorf("zero-unit Distance{5}.Meters() = %g, want 5000", got)
	}
}

func TestDistanceMetersConversion(t *testing.T) {
	cases := []struct {
		d    Distance
		want float64
	}{
		{Distance{5, Kilometers}, 5000},
		{Distance{1, Miles}, 1609.344},
		{Distance{250, Meters}, 250},
		{Distance{100, Feet}, 30.48},
	}
	for _, c := range cases {
		if got := c.d.Meters(); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("%s = %g m, want %g", c.d, got, c.want)
		}
	}
}

func TestDistanceMetersEnablesCrossUnitComparison(t *testing.T) {
	// The point of Meters: Amadeus quotes one hotel's distance in miles and
	// another's in kilometers, and the caller still has to sort them.
	oneMile := Distance{1, Miles}
	oneKilometer := Distance{1, Kilometers}
	if !(oneMile.Meters() > oneKilometer.Meters()) {
		t.Error("a mile should be longer than a kilometer")
	}
}

func TestDistanceString(t *testing.T) {
	if got := (Distance{5, Kilometers}).String(); got != "5 KM" {
		t.Errorf("String() = %q, want %q", got, "5 KM")
	}
}
