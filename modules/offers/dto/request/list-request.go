package requestHotelOffersDTO

import (
	"strconv"
	"strings"
)

type (
	HotelOffersListRequest struct {
		HotelIDs           []string `json:"hotelIds" required:"true"` // Amadeus property codes on 8 chars. Mandatory parameter for a search by predefined list of hotels.
		Adults             int      `json:"adults"`                   // Number of adult guests (1-9) per room. Default value: 1
		CheckInDate        string   `json:"checkInDate"`              // Check-in date of the stay (hotel local date). Format YYYY-MM-DD. The lowest accepted value is the present date (no dates in the past). If not present, the default value will be today's date in the GMT time zone.
		CheckOutDate       string   `json:"checkOutDate"`             // Check-out date of the stay (hotel local date). Format YYYY-MM-DD. The lowest accepted value is checkInDate+1. If not present, it will default to checkInDate +1.
		ChildAges          []int    `json:"childAges"`                // Comma separated list of ages of each child at the time of check-out from the hotel. If several children have the same age, their ages should be repeated in the list.
		CountryOfResidence string   `json:"countryOfResidence"`       // Code of the country of residence of the traveler expressed using ISO 3166-1 format.
		RoomQuantity       int      `json:"roomQuantity"`             // Number of rooms requested (1-9). Default value: 1
		RateCodes          []string `json:"rateCodes"`                // Special rates (comma separated list of Amadeus 3 chars codes): Public Rate (PRO...), Qualified Rate (GOV...) or Corporate Rate (IBM...). When a corporate rate is entered, Amadeus performs a check to verify which chains are authorized to view the rate, only authorized chains will be queried for availability. Warning: The availability response can also contain public rates. Example Rates: GOV - Government rate, AAA - AAA rate, MIL - Military/veteran rate, SNR - Senior rate, PRO - Promotional rate, COR - Corporate Code
		PriceRange         string   `json:"priceRange"`               // Filter hotel offers by price per night interval (ex: 200-300 or -300 or 100). It is mandatory to include a currency when this field is set.
		Currency           string   `json:"currency"`                 // Use this parameter to request a specific currency. ISO currency code (http://www.iso.org/iso/home/standards/currency_codes.htm). If a hotel does not support the requested currency, the prices for the hotel will be returned in the local currency of the hotel.
		PaymentPolicy      string   `json:"paymentPolicy"`            // Filter the response based on a specific payment type. NONE means all types (default). Available values: GUARANTEE, DEPOSIT, NONE. Default value: NONE
		BoardType          string   `json:"boardType"`                // Filter response based on available meals: ROOM_ONLY = Room Only, BREAKFAST = Breakfast, HALF_BOARD = Diner & Breakfast (only for Aggregators), FULL_BOARD = Full Board (only for Aggregators), ALL_INCLUSIVE = All Inclusive (only for Aggregators). Available values: ROOM_ONLY, BREAKFAST, HALF_BOARD, FULL_BOARD, ALL_INCLUSIVE
		IncludeClosed      bool     `json:"includeClosed"`            // Show all properties (include sold out) or available only. For sold out properties, please check availability on other dates.
		BestRateOnly       bool     `json:"bestRateOnly"`             // Used to return only the cheapest offer per hotel or all available offers. Default value: true
		PageOffset         string   `json:"page[offset]"`             // Represents the next value from which the scrolling will re-start. Up to 2000 rates can be returned in one shot.
		Lang               string   `json:"lang"`                     // Requested language of descriptive texts. Examples: FR, fr, fr-FR. If a language is not available the text will be returned in english. ISO language code (https://www.iso.org/iso-639-language-codes.html).
	}
)

func (r *HotelOffersListRequest) ToQueryParams() map[string]string {

	childAges := make([]string, len(r.ChildAges))
	for i, age := range r.ChildAges {
		childAges[i] = strconv.Itoa(age)
	}

	return map[string]string{
		"hotelIds":           strings.Join(r.HotelIDs, ","),
		"adults":             strconv.Itoa(r.Adults),
		"checkInDate":        r.CheckInDate,
		"checkOutDate":       r.CheckOutDate,
		"childAges":          strings.Join(childAges, ","),
		"countryOfResidence": r.CountryOfResidence,
		"roomQuantity":       strconv.Itoa(r.RoomQuantity),
		"rateCodes":          strings.Join(r.RateCodes, ","),
		"priceRange":         r.PriceRange,
		"currency":           r.Currency,
		"paymentPolicy":      r.PaymentPolicy,
		"boardType":          r.BoardType,
		"includeClosed":      strconv.FormatBool(r.IncludeClosed),
		"bestRateOnly":       strconv.FormatBool(r.BestRateOnly),
		"page[offset]":       r.PageOffset,
		"lang":               r.Lang,
	}
}
