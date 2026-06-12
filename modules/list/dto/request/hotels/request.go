package requestHotelListHotelsDTO

import "strings"

type (
	// HotelListByHotelsRequest retrieves hotels by their Amadeus property codes.
	// Endpoint: GET /reference-data/locations/hotels/by-hotels
	HotelListByHotelsRequest struct {
		// HotelIds - Amadeus 8-character property codes (chain + city + property).
		// Up to a comma-separated list is supported. Example: ["MCLONGHM", "ACPAR419"]. (required)
		HotelIds []string `json:"hotelIds"`
	}
)

func (r *HotelListByHotelsRequest) ToQueryParams() map[string]string {
	return map[string]string{
		"hotelIds": strings.Join(r.HotelIds, ","),
	}
}
