package responseHotelListDTO

import sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

type (
	HotelListResponse struct {
		Meta sharedResponseDTO.MetaResponse `json:"meta"`
		Data []GeneralInfoResponse          `json:"data"`
	}

	GeneralInfoResponse struct {
		ChainCode string        `json:"chainCode"`
		IataCode  string        `json:"iataCode"`
		DupeId    int           `json:"dupeId"`
		Name      string        `json:"name"`
		HotelId   string        `json:"hotelId"`
		GeoCode   GeoCodeModel  `json:"geoCode"`
		Address   AddressModel  `json:"address"`
		Distance  DistanceModel `json:"distance"`
	}

	GeoCodeModel struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	AddressModel struct {
		CountryCode string `json:"countryCode"`
	}

	DistanceModel struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}
)
