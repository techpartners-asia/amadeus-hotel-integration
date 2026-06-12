package requestHotelOffersDTO

// Bool returns a pointer to b, a convenience for setting optional *bool request
// fields such as HotelOffersListRequest.BestRateOnly and IncludeClosed.
func Bool(b bool) *bool { return &b }

type (
	// HotelOffersByIDRequest fetches a single offer by its id.
	// Endpoint: GET /shopping/hotel-offers/{offerId}
	HotelOffersByIDRequest struct {
		// OfferID - the offer id, used as the path segment (required).
		OfferID string `json:"offerId" required:"true"`
		// Lang - requested language of descriptive texts (e.g. "EN", "FR"). Optional.
		Lang string `json:"lang"`
	}
)

func (r *HotelOffersByIDRequest) ToQueryParams() map[string]string {
	queryParams := map[string]string{}

	if r.Lang != "" {
		queryParams["lang"] = r.Lang
	}

	return queryParams
}
