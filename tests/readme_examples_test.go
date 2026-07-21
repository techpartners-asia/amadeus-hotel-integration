package tests

import (
	"testing"

	requestContent "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
	requestList "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	requestGeocode "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/geocode"
	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

// TestReadmeExamplesCompile builds the request literals shown in README.md.
// Documentation that no longer compiles is worse than none, and the README's
// search examples silently rotted once before: the by-geocode snippet set
// CityCode, a field that request has never had.
//
// Keep these in sync with the README. If a field is renamed here, the snippet
// it mirrors needs the same edit.
func TestReadmeExamplesCompile(t *testing.T) {
	// README: Modules > List Module > Search by City Code
	byCity := requestList.HotelListByCityCodeRequest{
		CityCode:    "PAR",
		Radius:      searchcriteria.Ptr(10),
		RadiusUnit:  searchcriteria.Ptr(searchcriteria.RadiusUnitKM),
		Ratings:     []searchcriteria.Rating{searchcriteria.Rating4, searchcriteria.Rating5},
		Amenities:   []searchcriteria.Amenity{searchcriteria.AmenityWifi, searchcriteria.AmenitySpa},
		HotelSource: searchcriteria.Ptr(searchcriteria.HotelSourceAll),
	}
	if q := byCity.ToQueryParams(); q["amenities"] != "WIFI,SPA" {
		t.Errorf("by-city amenities = %q, want %q", q["amenities"], "WIFI,SPA")
	}

	// README: Modules > List Module > Search by Geocode
	byGeocode := requestGeocode.HotelListByGeocodeRequest{
		Latitude:    48.85,
		Longitude:   2.29,
		Radius:      5,
		RadiusUnit:  searchcriteria.RadiusUnitKM,
		Ratings:     []searchcriteria.Rating{searchcriteria.Rating3, searchcriteria.Rating4, searchcriteria.Rating5},
		HotelSource: searchcriteria.HotelSourceAll,
	}
	if q := byGeocode.ToQueryParams(); q["ratings"] != "3,4,5" {
		t.Errorf("by-geocode ratings = %q, want %q", q["ratings"], "3,4,5")
	}

	// README: Modules > Offers Module > List Offers
	offers := requestOffers.HotelOffersListRequest{
		HotelIDs:     []string{"RTPAREIF"},
		Adults:       2,
		CheckInDate:  "2026-06-01",
		CheckOutDate: "2026-06-05",
		RoomQuantity: 1,
		Currency:     "EUR",
		BoardType:    searchcriteria.BoardTypeBreakfast,
		BestRateOnly: requestOffers.Bool(true),
		Lang:         "EN",
	}
	if q := offers.ToQueryParams(); q["boardType"] != "BREAKFAST" {
		t.Errorf("offers boardType = %q, want %q", q["boardType"], "BREAKFAST")
	}

	// README: Search Criteria — the content view is typed too.
	content := requestContent.ContentByIDRequest{
		ID:   "ADNYCCTB",
		View: searchcriteria.ContentViewFull,
	}
	if q := content.ToQueryParams(); q["view"] != "FULL" {
		t.Errorf("content view = %q, want %q", q["view"], "FULL")
	}
}
