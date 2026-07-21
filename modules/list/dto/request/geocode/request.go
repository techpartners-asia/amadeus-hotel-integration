package requestHotelListGeocodeDTO

import (
	"strconv"
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
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
		// RadiusUnit - unit for Radius. Default KM. See searchcriteria.AllRadiusUnits. Optional.
		RadiusUnit searchcriteria.RadiusUnit `json:"radiusUnit,omitempty"`
		// ChainCodes - hotel chain or brand codes, 2 capital letters each. Optional.
		ChainCodes []string `json:"chainCodes,omitempty"`
		// Amenities - filter by amenity. See searchcriteria.AllAmenities. Optional.
		Amenities []searchcriteria.Amenity `json:"amenities,omitempty"`
		// Ratings - hotel stars, up to searchcriteria.MaxRatings values. See searchcriteria.AllRatings. Optional.
		Ratings []searchcriteria.Rating `json:"ratings,omitempty"`
		// HotelSource - which inventory to search. Default ALL. See searchcriteria.AllHotelSources. Optional.
		HotelSource searchcriteria.HotelSource `json:"hotelSource,omitempty"`
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
		queryParams["radiusUnit"] = string(r.RadiusUnit)
	}
	if len(r.ChainCodes) > 0 {
		queryParams["chainCodes"] = strings.Join(r.ChainCodes, ",")
	}
	if len(r.Amenities) > 0 {
		queryParams["amenities"] = searchcriteria.Join(r.Amenities)
	}
	if len(r.Ratings) > 0 {
		queryParams["ratings"] = searchcriteria.Join(r.Ratings)
	}
	if r.HotelSource != "" {
		queryParams["hotelSource"] = string(r.HotelSource)
	}

	return queryParams
}
