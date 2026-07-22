package codes

// HotelSource selects which inventory the Hotel List API searches, via the
// `hotelSource` query parameter. Defaults to HotelSourceAll when omitted.
type HotelSource string

const (
	// HotelSourceBedbank restricts results to aggregator inventory.
	HotelSourceBedbank HotelSource = "BEDBANK"
	// HotelSourceDirectChain restricts results to GDS/Distribution inventory.
	HotelSourceDirectChain HotelSource = "DIRECTCHAIN"
	// HotelSourceAll searches both. This is the Amadeus default.
	HotelSourceAll HotelSource = "ALL"
)

var hotelSourceCatalog = []entry[HotelSource]{
	{HotelSourceBedbank, "Aggregators"},
	{HotelSourceDirectChain, "GDS / Distribution"},
	{HotelSourceAll, "All sources"},
}

// AllHotelSources returns every hotel source Amadeus accepts.
func AllHotelSources() []HotelSource { return allOf(hotelSourceCatalog) }

// Label returns a human-readable name for h, or "" when h is not a known code.
func (h HotelSource) Label() string { return labelOf(hotelSourceCatalog, h) }

// IsValid reports whether h is a code Amadeus accepts.
func (h HotelSource) IsValid() bool { return isValid(hotelSourceCatalog, h) }
