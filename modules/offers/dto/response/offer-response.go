package responseHotelOffersDTO

type (
	OffersResponse struct {
		Type      string          `json:"type"`
		Hotel     HotelResponse   `json:"hotel"`
		Available bool            `json:"available"`
		Offers    []OfferResponse `json:"offers"`
		Self      string          `json:"self"`
	}

	HotelResponse struct {
		Type      string  `json:"type"`
		HotelID   string  `json:"hotelId"`
		ChainCode string  `json:"chainCode"`
		DupeID    string  `json:"dupeId"`
		Name      string  `json:"name"`
		CityCode  string  `json:"cityCode"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	OfferResponse struct {
		ID                  string                      `json:"id"`
		CheckInDate         string                      `json:"checkInDate"`
		CheckOutDate        string                      `json:"checkOutDate"`
		RateCode            string                      `json:"rateCode"`
		RateFamilyEstimated RateFamilyEstimatedResponse `json:"rateFamilyEstimated"`
		Room                RoomResponse                `json:"room"`
		Guests              GuestsResponse              `json:"guests"`
		Price               PriceResponse               `json:"price"`
		Policies            PoliciesResponse            `json:"policies"`
		Self                string                      `json:"self"`
	}

	RateFamilyEstimatedResponse struct {
		Code string `json:"code"`
		Type string `json:"type"`
	}

	RoomResponse struct {
		Type          string                `json:"type"`
		TypeEstimated TypeEstimatedResponse `json:"typeEstimated"`
		Description   DescriptionResponse   `json:"description"`
	}

	TypeEstimatedResponse struct {
		Category string `json:"category"`
		Beds     int    `json:"beds"`
		BedType  string `json:"bedType"`
	}

	DescriptionResponse struct {
		Text string `json:"text"`
		Lang string `json:"lang"`
	}

	GuestsResponse struct {
		Adults    int   `json:"adults"`
		ChildAges []int `json:"childAges"`
	}

	PriceResponse struct {
		Currency   string             `json:"currency"`
		Base       string             `json:"base"`
		Total      string             `json:"total"`
		Variations VariationsResponse `json:"variations"`
	}

	VariationsResponse struct {
		Average AverageResponse  `json:"average"`
		Changes []ChangeResponse `json:"changes"`
	}

	AverageResponse struct {
		Base string `json:"base"`
	}

	ChangeResponse struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Total     string `json:"total"`
	}

	PoliciesResponse struct {
		PaymentType  string               `json:"paymentType"`
		Cancellation CancellationResponse `json:"cancellation"`
	}

	CancellationResponse struct {
		Description DescriptionResponse `json:"description"`
		Type        string              `json:"type"`
	}

	DetailResponse struct {
		Message    string            `json:"message"`
		Parameters map[string]string `json:"parameters"`
	}
)
