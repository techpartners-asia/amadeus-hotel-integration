package requestHotelListCityDTO

import (
	"strconv"
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

type (
	HotelListByCityCodeRequest struct {
		CityCode    string                      `json:"cityCode"`    // Destination city code or airport code. In case of city code, the search will be done around the city center. Available codes can be found in IATA table codes (3 chars IATA Code). Example: PAR
		Radius      *int                        `json:"radius"`      // Maximum distance from the geographical coordinates expressed in defined units. The default unit is metric kilometer. Default value: 5
		RadiusUnit  *searchcriteria.RadiusUnit  `json:"radiusUnit"`  // Unit of measurement used to express the radius. See searchcriteria.AllRadiusUnits. Default value: KM
		ChainCodes  []string                    `json:"chainCodes"`  // Array of hotel chain codes. Each code is a string consisted of 2 capital alphabetic characters. The code is either a chain or a brand. The response includes all the hotels of the selected chain, or all the hotels of the sub chains of the selected brand.
		Amenities   []searchcriteria.Amenity    `json:"amenities"`   // List of amenities to filter on. See searchcriteria.AllAmenities for the full set.
		Ratings     []searchcriteria.Rating     `json:"ratings"`     // Hotel stars. Up to searchcriteria.MaxRatings values can be requested at the same time. The response includes all the hotels with the requested rating and all hotels with an Amadeus self rating matching the requested rating. See searchcriteria.AllRatings.
		HotelSource *searchcriteria.HotelSource `json:"hotelSource"` // Which inventory to search. See searchcriteria.AllHotelSources. Default value: ALL
	}
)

func (r *HotelListByCityCodeRequest) ToQueryParams() map[string]string {
	queryParams := map[string]string{
		"cityCode": r.CityCode,
	}

	if r.Radius != nil {
		queryParams["radius"] = strconv.Itoa(*r.Radius)
	}

	if r.RadiusUnit != nil {
		queryParams["radiusUnit"] = string(*r.RadiusUnit)
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

	if r.HotelSource != nil {
		queryParams["hotelSource"] = string(*r.HotelSource)
	}

	return queryParams

}
