package requestHotelListGeocodeDTO

import (
	"strconv"
	"strings"
)

type (
	// HotelListByGeocodeRequest searches hotels around a geographic point.
	// Endpoint: GET /reference-data/locations/hotels/by-geocode
	//
	// Latitude and Longitude are required; the endpoint rejects a request
	// without coordinates ("INVALID GEOGRAPHICAL ZONE - Missing coordinates").
	HotelListByGeocodeRequest struct {
		// Latitude - decimal coordinate, e.g. 48.85. (required)
		Latitude float64 `json:"latitude" required:"true"`
		// Longitude - decimal coordinate, e.g. 2.29. (required)
		Longitude float64 `json:"longitude" required:"true"`
		// Radius - maximum distance from the coordinates, in RadiusUnit. Default 5. Optional.
		Radius int `json:"radius,omitempty"`
		// RadiusUnit - unit for Radius. Available values: KM, MILE. Default KM. Optional.
		RadiusUnit string `json:"radiusUnit,omitempty"`
		// ChainCodes - hotel chain or brand codes, 2 capital letters each. Optional.
		ChainCodes []string `json:"chainCodes,omitempty"`
		// Amenities - filter by amenity codes such as SWIMMING_POOL, WIFI, PARKING. Optional.
		Amenities []string `json:"amenities,omitempty"`
		// Ratings - hotel stars, up to four values from 1..5. Optional.
		Ratings []string `json:"ratings,omitempty"`
		// HotelSource - BEDBANK, DIRECTCHAIN or ALL. Default ALL. Optional.
		HotelSource string `json:"hotelSource,omitempty"`
	}
)

// ToQueryParams builds the query string. Only latitude and longitude are always
// sent; every optional parameter is emitted solely when set, so a minimal
// request never sends empty or zero values (radius=0, radiusUnit=) that
// Amadeus rejects.
func (r *HotelListByGeocodeRequest) ToQueryParams() map[string]string {
	queryParams := map[string]string{
		"latitude":  strconv.FormatFloat(r.Latitude, 'f', -1, 64),
		"longitude": strconv.FormatFloat(r.Longitude, 'f', -1, 64),
	}

	if r.Radius > 0 {
		queryParams["radius"] = strconv.Itoa(r.Radius)
	}
	if r.RadiusUnit != "" {
		queryParams["radiusUnit"] = r.RadiusUnit
	}
	if len(r.ChainCodes) > 0 {
		queryParams["chainCodes"] = strings.Join(r.ChainCodes, ",")
	}
	if len(r.Amenities) > 0 {
		queryParams["amenities"] = strings.Join(r.Amenities, ",")
	}
	if len(r.Ratings) > 0 {
		queryParams["ratings"] = strings.Join(r.Ratings, ",")
	}
	if r.HotelSource != "" {
		queryParams["hotelSource"] = r.HotelSource
	}

	return queryParams
}
