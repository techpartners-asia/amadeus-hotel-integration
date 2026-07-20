package responseContentDTO

type (
	HotelContentResponse struct {
		Promotions      []PromotionResponse       `json:"promotions"`
		Awards          []AwardsResponse          `json:"awards"`
		Policies        PolicyResponse            `json:"policies"`
		Rooms           []RoomResponse            `json:"rooms"`
		Facilities      FacilityResponse          `json:"facilities"`
		PointOfInterest []PointOfInterestResponse `json:"pointOfInterest"`
		Hotel           HotelResponse             `json:"hotel"`
		Basic           BasicResponse             `json:"basic"`
	}
)
