package requestHotelOffersDTO

type (
	HotelOffersByIDRequest struct {
		OfferID string `json:"offerId" required:"true"`
		Lang    string `json:"lang"`
	}
)

func (r *HotelOffersByIDRequest) ToQueryParams() map[string]string {
	return map[string]string{
		"lang": r.Lang,
	}
}
