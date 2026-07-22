package inventory

import (
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/geo"
)

// radiusUnits are the distance units the Hotel List endpoints accept. The other
// geo.Unit values are valid distances but Amadeus rejects them here, so the
// restriction lives in this context rather than in geo.
var radiusUnits = []geo.Unit{geo.Kilometers, geo.Miles}

// filters are the search criteria the by-city and by-geocode endpoints share.
// by-hotels takes none of them: it is a lookup by key, not a search.
type filters struct {
	// Radius bounds the search around its centre. Zero means Amadeus's own
	// default of 5.
	Radius float64
	// RadiusUnit is the unit for Radius. Only geo.Kilometers and geo.Miles are
	// accepted here. Empty means kilometers.
	RadiusUnit geo.Unit
	// ChainCodes restricts results to these chains or brands, two capital
	// letters each.
	ChainCodes []string
	// Amenities restricts results to properties offering all of these.
	Amenities []codes.Amenity
	// Ratings restricts results to these star ratings. Amadeus accepts at most
	// codes.MaxRatings of them at once.
	Ratings []codes.Rating
	// Source selects which inventory to search. Empty means all of it.
	Source codes.HotelSource
}

// apply writes the filters that are set into q, leaving unset ones out
// entirely. Sending radius=0 or radiusUnit= makes Amadeus reject the request,
// so "omit when unset" is required behaviour rather than tidiness.
func (f filters) apply(q url.Values) {
	if f.Radius > 0 {
		q.Set("radius", strconv.FormatFloat(f.Radius, 'f', -1, 64))
	}
	if f.RadiusUnit != "" {
		q.Set("radiusUnit", string(f.RadiusUnit))
	}
	if len(f.ChainCodes) > 0 {
		q.Set("chainCodes", strings.Join(f.ChainCodes, ","))
	}
	if len(f.Amenities) > 0 {
		q.Set("amenities", codes.Join(f.Amenities))
	}
	if len(f.Ratings) > 0 {
		q.Set("ratings", codes.Join(f.Ratings))
	}
	if f.Source != "" {
		q.Set("hotelSource", string(f.Source))
	}
}

// validate collects every problem with the shared filters.
func (f filters) validate(errs apierr.ValidationErrors) apierr.ValidationErrors {
	if f.Radius < 0 {
		errs = append(errs, apierr.Invalidf("Radius", "must not be negative, got %g", f.Radius))
	}
	if f.RadiusUnit != "" && !slices.Contains(radiusUnits, f.RadiusUnit) {
		errs = append(errs, apierr.Invalidf("RadiusUnit",
			"%q is not accepted here; hotel search takes %s or %s", f.RadiusUnit, geo.Kilometers, geo.Miles))
	}
	for _, code := range f.Amenities {
		if !code.IsValid() {
			errs = append(errs, apierr.Invalidf("Amenities", "%q is not an amenity Amadeus accepts", code))
		}
	}
	if len(f.Ratings) > codes.MaxRatings {
		errs = append(errs, apierr.Invalidf("Ratings",
			"at most %d ratings may be requested at once, got %d", codes.MaxRatings, len(f.Ratings)))
	}
	for _, code := range f.Ratings {
		if !code.IsValid() {
			errs = append(errs, apierr.Invalidf("Ratings", "%q is not a rating; use 1 to 5", code))
		}
	}
	if f.Source != "" && !f.Source.IsValid() {
		errs = append(errs, apierr.Invalidf("Source", "%q is not a known hotel source", f.Source))
	}
	for _, chain := range f.ChainCodes {
		if len(chain) != 2 {
			errs = append(errs, apierr.Invalidf("ChainCodes",
				"%q is not a chain code; they are exactly 2 characters", chain))
		}
	}
	return errs
}

// CityQuery searches for hotels around a city or airport.
//
//	hotels, err := client.Inventory.ByCity(ctx, inventory.CityQuery{
//	    CityCode: "PAR",
//	    Filters:  inventory.Filters{Radius: 10, Ratings: []codes.Rating{codes.Rating5}},
//	})
type CityQuery struct {
	// CityCode is the 3-letter IATA city or airport code, e.g. "PAR".
	// Required. A city code searches around the city centre.
	CityCode string
	// Filters narrows the results. Optional.
	Filters Filters
}

// Filters is the public name for the criteria the searches share.
type Filters = filters

// GeocodeQuery searches for hotels around a geographic point.
type GeocodeQuery struct {
	// Position is the centre of the search. Required: the endpoint rejects a
	// request without coordinates.
	Position geo.Coordinates
	// Filters narrows the results. Optional.
	Filters Filters
}

// IDsQuery looks up specific properties by their Amadeus codes. It is a lookup
// rather than a search, so it takes no filters.
type IDsQuery struct {
	// HotelIDs are the 8-character property codes to fetch. Required.
	HotelIDs []string
}

// maxHotelIDsPerRequest is what Amadeus accepts in one by-hotels call. Beyond
// it the request fails, so the SDK reports the problem rather than letting the
// caller discover it as a 400.
const maxHotelIDsPerRequest = 100

func (q CityQuery) validate() error {
	var errs apierr.ValidationErrors
	if q.CityCode == "" {
		errs = errs.Append("CityCode", "is required")
	} else if len(q.CityCode) != 3 {
		errs = append(errs, apierr.Invalidf("CityCode",
			"%q is not an IATA code; they are exactly 3 characters", q.CityCode))
	}
	return q.Filters.validate(errs).OrNil()
}

func (q CityQuery) params() url.Values {
	values := url.Values{"cityCode": {q.CityCode}}
	q.Filters.apply(values)
	return values
}

func (q GeocodeQuery) validate() error {
	var errs apierr.ValidationErrors
	if err := q.Position.Validate(); err != nil {
		errs = errs.Append("Position", err.Error())
	}
	return q.Filters.validate(errs).OrNil()
}

func (q GeocodeQuery) params() url.Values {
	values := url.Values{
		"latitude":  {strconv.FormatFloat(q.Position.Latitude, 'f', -1, 64)},
		"longitude": {strconv.FormatFloat(q.Position.Longitude, 'f', -1, 64)},
	}
	q.Filters.apply(values)
	return values
}

func (q IDsQuery) validate() error {
	var errs apierr.ValidationErrors
	switch {
	case len(q.HotelIDs) == 0:
		errs = errs.Append("HotelIDs", "at least one hotel ID is required")
	case len(q.HotelIDs) > maxHotelIDsPerRequest:
		errs = append(errs, apierr.Invalidf("HotelIDs",
			"at most %d IDs per request, got %d", maxHotelIDsPerRequest, len(q.HotelIDs)))
	}
	for _, id := range q.HotelIDs {
		if !HotelID(id).IsValid() {
			errs = append(errs, apierr.Invalidf("HotelIDs",
				"%q is not a property code; they are 8 uppercase alphanumeric characters", id))
		}
	}
	return errs.OrNil()
}

func (q IDsQuery) params() url.Values {
	return url.Values{"hotelIds": {strings.Join(q.HotelIDs, ",")}}
}
