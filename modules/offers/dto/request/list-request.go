package requestHotelOffersDTO

import (
	"strconv"
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

type (
	HotelOffersListRequest struct {
		HotelIDs           []string                     `json:"hotelIds" required:"true"` // Amadeus property codes on 8 chars. Mandatory parameter for a search by predefined list of hotels.
		Adults             int                          `json:"adults"`                   // Number of adult guests (1-9) per room. Default value: 1
		CheckInDate        string                       `json:"checkInDate"`              // Check-in date of the stay (hotel local date). Format YYYY-MM-DD. The lowest accepted value is the present date (no dates in the past). If not present, the default value will be today's date in the GMT time zone.
		CheckOutDate       string                       `json:"checkOutDate"`             // Check-out date of the stay (hotel local date). Format YYYY-MM-DD. The lowest accepted value is checkInDate+1. If not present, it will default to checkInDate +1.
		ChildAges          []int                        `json:"childAges"`                // Comma separated list of ages of each child at the time of check-out from the hotel. If several children have the same age, their ages should be repeated in the list.
		CountryOfResidence string                       `json:"countryOfResidence"`       // Code of the country of residence of the traveler expressed using ISO 3166-1 format.
		RoomQuantity       int                          `json:"roomQuantity"`             // Number of rooms requested (1-9). Default value: 1
		RateCodes          []searchcriteria.RateCode    `json:"rateCodes"`                // Special rates (Amadeus 3 chars codes): Public Rate (PRO...), Qualified Rate (GOV...) or Corporate Rate (IBM...). When a corporate rate is entered, Amadeus performs a check to verify which chains are authorized to view the rate, only authorized chains will be queried for availability. Warning: The availability response can also contain public rates. See searchcriteria.AllRateCodes for the documented codes; corporate codes are account-specific and are not enumerated there.
		PriceRange         string                       `json:"priceRange"`               // Filter hotel offers by price per night interval (ex: 200-300 or -300 or 100). It is mandatory to include a currency when this field is set.
		Currency           string                       `json:"currency"`                 // Use this parameter to request a specific currency. ISO currency code (http://www.iso.org/iso/home/standards/currency_codes.htm). If a hotel does not support the requested currency, the prices for the hotel will be returned in the local currency of the hotel.
		PaymentPolicy      searchcriteria.PaymentPolicy `json:"paymentPolicy"`            // Filter the response based on a specific payment type. NONE means all types (default). See searchcriteria.AllPaymentPolicies.
		BoardType          searchcriteria.BoardType     `json:"boardType"`                // Filter response based on available meals. HALF_BOARD, FULL_BOARD and ALL_INCLUSIVE apply to Aggregators only. See searchcriteria.AllBoardTypes.
		IncludeClosed      *bool                        `json:"includeClosed"`            // Show all properties (include sold out) or available only. For sold out properties, please check availability on other dates. Optional.
		BestRateOnly       *bool                        `json:"bestRateOnly"`             // Used to return only the cheapest offer per hotel or all available offers. Default value: true. Optional.
		PageOffset         string                       `json:"page[offset]"`             // Represents the next value from which the scrolling will re-start. Up to 2000 rates can be returned in one shot.
		Lang               string                       `json:"lang"`                     // Requested language of descriptive texts. Examples: FR, fr, fr-FR. If a language is not available the text will be returned in english. ISO language code (https://www.iso.org/iso-639-language-codes.html).
	}
)

func (r *HotelOffersListRequest) ToQueryParams() map[string]string {
	// hotelIds is the only required parameter; everything else is sent only when
	// set, so a minimal request does not emit empty/zero-valued query params
	// (e.g. adults=0 or currency=) that Amadeus would reject.
	queryParams := map[string]string{
		"hotelIds": strings.Join(r.HotelIDs, ","),
	}

	if r.Adults > 0 {
		queryParams["adults"] = strconv.Itoa(r.Adults)
	}
	if r.CheckInDate != "" {
		queryParams["checkInDate"] = r.CheckInDate
	}
	if r.CheckOutDate != "" {
		queryParams["checkOutDate"] = r.CheckOutDate
	}
	if len(r.ChildAges) > 0 {
		childAges := make([]string, len(r.ChildAges))
		for i, age := range r.ChildAges {
			childAges[i] = strconv.Itoa(age)
		}
		queryParams["childAges"] = strings.Join(childAges, ",")
	}
	if r.CountryOfResidence != "" {
		queryParams["countryOfResidence"] = r.CountryOfResidence
	}
	if r.RoomQuantity > 0 {
		queryParams["roomQuantity"] = strconv.Itoa(r.RoomQuantity)
	}
	if len(r.RateCodes) > 0 {
		queryParams["rateCodes"] = searchcriteria.Join(r.RateCodes)
	}
	if r.PriceRange != "" {
		queryParams["priceRange"] = r.PriceRange
	}
	if r.Currency != "" {
		queryParams["currency"] = r.Currency
	}
	if r.PaymentPolicy != "" {
		queryParams["paymentPolicy"] = string(r.PaymentPolicy)
	}
	if r.BoardType != "" {
		queryParams["boardType"] = string(r.BoardType)
	}
	if r.IncludeClosed != nil {
		queryParams["includeClosed"] = strconv.FormatBool(*r.IncludeClosed)
	}
	if r.BestRateOnly != nil {
		queryParams["bestRateOnly"] = strconv.FormatBool(*r.BestRateOnly)
	}
	if r.PageOffset != "" {
		queryParams["page[offset]"] = r.PageOffset
	}
	if r.Lang != "" {
		queryParams["lang"] = r.Lang
	}

	return queryParams
}
