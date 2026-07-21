package searchcriteria

// RadiusUnit is the unit the Hotel List API applies to the `radius` parameter.
// Defaults to RadiusUnitKM when omitted.
type RadiusUnit string

const (
	RadiusUnitKM   RadiusUnit = "KM"
	RadiusUnitMile RadiusUnit = "MILE"
)

var radiusUnitCatalog = []entry[RadiusUnit]{
	{RadiusUnitKM, "Kilometers"},
	{RadiusUnitMile, "Miles"},
}

// AllRadiusUnits returns every radius unit Amadeus accepts.
func AllRadiusUnits() []RadiusUnit { return codes(radiusUnitCatalog) }

// Label returns a human-readable name for u, or "" when u is not a known code.
func (u RadiusUnit) Label() string { return labelOf(radiusUnitCatalog, u) }

// IsValid reports whether u is a code Amadeus accepts.
func (u RadiusUnit) IsValid() bool { return isValid(radiusUnitCatalog, u) }
