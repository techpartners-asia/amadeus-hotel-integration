package responseHotelListDTO

import sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

type (
	HotelListResponse struct {
		Meta sharedResponseDTO.MetaResponse `json:"meta"`
		Data []GeneralInfoResponse          `json:"data"`
	}

	GeneralInfoResponse struct {
		ChainCode       string         `json:"chainCode"`
		BrandCode       string         `json:"brandCode"`
		MasterChainCode string         `json:"masterChainCode"`
		IataCode        string         `json:"iataCode"`
		DupeId          int            `json:"dupeId"`
		Name            string         `json:"name"`
		HotelId         string         `json:"hotelId"`
		GeoCode         GeoCodeModel   `json:"geoCode"`
		Address         AddressModel   `json:"address"`
		Distance        DistanceModel  `json:"distance"`
		LastUpdate      string         `json:"lastUpdate"`
		Retailing       RetailingModel `json:"retailing"`
	}

	GeoCodeModel = sharedResponseDTO.GeoCodeResponse

	AddressModel struct {
		Lines       []string `json:"lines"`
		PostalCode  string   `json:"postalCode"`
		CityName    string   `json:"cityName"`
		CountryCode string   `json:"countryCode"`
		StateCode   string   `json:"stateCode"`
	}

	// RetailingModel carries retailing metadata such as sponsored placement.
	RetailingModel struct {
		Sponsorship SponsorshipModel `json:"sponsorship"`
	}

	SponsorshipModel struct {
		IsSponsored bool `json:"isSponsored"`
	}

	DistanceModel struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}
)
