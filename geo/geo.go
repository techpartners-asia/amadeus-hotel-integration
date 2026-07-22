// Package geo holds the geographic value objects shared by the SDK's bounded
// contexts.
//
// Amadeus returns latitude and longitude as two independent float fields, and a
// radius as a number beside a separate unit string. Both pairs travel together
// and are meaningless apart, so the SDK binds each into one value.
package geo

import (
	"fmt"
	"math"
	"slices"
)

// Coordinates is a point on the earth in decimal degrees (WGS 84).
//
// The zero Coordinates is the valid point 0,0 in the Gulf of Guinea, which is
// indistinguishable from "no coordinates". Amadeus omits the pair entirely for
// properties it cannot locate, so responses model an absent position as a nil
// *Coordinates rather than a zero value.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// ErrOutOfRange is returned for coordinates outside the valid degree ranges.
var ErrOutOfRange = fmt.Errorf("geo: coordinates out of range")

// NewCoordinates returns the point at lat, lon, rejecting values outside
// [-90,90] and [-180,180].
func NewCoordinates(lat, lon float64) (Coordinates, error) {
	c := Coordinates{Latitude: lat, Longitude: lon}
	if err := c.Validate(); err != nil {
		return Coordinates{}, err
	}
	return c, nil
}

// Validate reports whether the coordinates fall within the valid degree ranges.
func (c Coordinates) Validate() error {
	if math.IsNaN(c.Latitude) || math.IsNaN(c.Longitude) {
		return fmt.Errorf("%w: latitude/longitude is not a number", ErrOutOfRange)
	}
	if c.Latitude < -90 || c.Latitude > 90 {
		return fmt.Errorf("%w: latitude %g outside [-90,90]", ErrOutOfRange, c.Latitude)
	}
	if c.Longitude < -180 || c.Longitude > 180 {
		return fmt.Errorf("%w: longitude %g outside [-180,180]", ErrOutOfRange, c.Longitude)
	}
	return nil
}

// String renders the point as "48.8566,2.3522".
func (c Coordinates) String() string {
	return fmt.Sprintf("%g,%g", c.Latitude, c.Longitude)
}

// Unit is the unit a Distance is measured in. Amadeus accepts these four in
// hotel search, and rejects anything else with a 400.
type Unit string

const (
	// Kilometers is Amadeus's default radius unit.
	Kilometers Unit = "KM"
	Miles      Unit = "MILE"
	Meters     Unit = "METER"
	Feet       Unit = "FEET"
)

// AllUnits returns every distance unit Amadeus accepts, in a stable order, for
// callers rendering a unit selector.
func AllUnits() []Unit { return []Unit{Kilometers, Miles, Meters, Feet} }

// IsValid reports whether u is a unit Amadeus accepts.
func (u Unit) IsValid() bool { return slices.Contains(AllUnits(), u) }

// String returns the unit code as Amadeus spells it.
func (u Unit) String() string { return string(u) }

// Distance is a length with its unit attached, used for search radii and for
// the hotel-to-landmark distances Amadeus reports.
//
// The zero Distance is zero kilometers.
type Distance struct {
	Value float64
	Unit  Unit
}

// NewDistance returns a Distance, defaulting a missing unit to kilometers to
// match Amadeus's own default.
func NewDistance(value float64, unit Unit) Distance {
	if unit == "" {
		unit = Kilometers
	}
	return Distance{Value: value, Unit: unit}
}

// String renders the distance as "5 KM".
func (d Distance) String() string {
	return fmt.Sprintf("%g %s", d.Value, d.unitOrDefault())
}

// Meters returns the distance converted to meters, for comparing two distances
// that were quoted in different units.
//
// The conversion is exact for metric units and uses the international
// definitions for the others (1 mile = 1609.344 m, 1 foot = 0.3048 m).
func (d Distance) Meters() float64 {
	switch d.unitOrDefault() {
	case Kilometers:
		return d.Value * 1000
	case Miles:
		return d.Value * 1609.344
	case Feet:
		return d.Value * 0.3048
	default:
		return d.Value
	}
}

func (d Distance) unitOrDefault() Unit {
	if d.Unit == "" {
		return Kilometers
	}
	return d.Unit
}
