package tests

import (
	"testing"
	"time"

	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

// --- request building (no network) ---

// TestOffersByIDQueryParamsOmitsEmpty verifies the by-id request sends no query
// params at all when only the (path-segment) offer id is set.
func TestOffersByIDQueryParamsOmitsEmpty(t *testing.T) {
	req := requestOffers.HotelOffersByIDRequest{OfferID: "ZC7ZKEP98D"}

	q := req.ToQueryParams()

	if len(q) != 0 {
		t.Fatalf("expected no query params, got %v", q)
	}
}

// TestOffersByIDQueryParamsIncludesLang verifies lang is forwarded when set.
// The offer id is a path segment and must never leak into the query string.
func TestOffersByIDQueryParamsIncludesLang(t *testing.T) {
	req := requestOffers.HotelOffersByIDRequest{OfferID: "ZC7ZKEP98D", Lang: "FR"}

	q := req.ToQueryParams()

	if q["lang"] != "FR" {
		t.Errorf("lang = %q, want %q", q["lang"], "FR")
	}
	if _, ok := q["offerId"]; ok {
		t.Error("offerId must be a path segment, not a query param")
	}
}

// TestOffersListQueryParamsBestRateOnlyTrue guards the tri-state *bool: an
// explicit true must be emitted, not treated as unset.
func TestOffersListQueryParamsBestRateOnlyTrue(t *testing.T) {
	req := requestOffers.HotelOffersListRequest{
		HotelIDs:      []string{"RTPAREIF"},
		BestRateOnly:  requestOffers.Bool(true),
		IncludeClosed: requestOffers.Bool(true),
	}

	q := req.ToQueryParams()

	if q["bestRateOnly"] != "true" {
		t.Errorf("bestRateOnly = %q, want \"true\"", q["bestRateOnly"])
	}
	if q["includeClosed"] != "true" {
		t.Errorf("includeClosed = %q, want \"true\"", q["includeClosed"])
	}
}

// TestOffersListQueryParamsPageOffset verifies the bracketed pagination key is
// emitted verbatim, since Amadeus expects the literal "page[offset]".
func TestOffersListQueryParamsPageOffset(t *testing.T) {
	req := requestOffers.HotelOffersListRequest{
		HotelIDs:   []string{"RTPAREIF"},
		PageOffset: "ABC123",
	}

	q := req.ToQueryParams()

	if q["page[offset]"] != "ABC123" {
		t.Errorf("page[offset] = %q, want %q", q["page[offset]"], "ABC123")
	}
}

// TestOffersListQueryParamsRateCodesJoined verifies slice params are joined with
// commas rather than sent as repeated keys.
func TestOffersListQueryParamsRateCodesJoined(t *testing.T) {
	req := requestOffers.HotelOffersListRequest{
		HotelIDs:  []string{"RTPAREIF", "RTPARMAI"},
		RateCodes: []searchcriteria.RateCode{searchcriteria.RateCodeRack, searchcriteria.RateCodeGovernment},
	}

	q := req.ToQueryParams()

	if q["rateCodes"] != "RAC,GOV" {
		t.Errorf("rateCodes = %q, want %q", q["rateCodes"], "RAC,GOV")
	}
	if q["hotelIds"] != "RTPAREIF,RTPARMAI" {
		t.Errorf("hotelIds = %q, want comma-joined", q["hotelIds"])
	}
}

// --- live API (network) ---

// sandboxOfferHotels are the Paris properties that actually carry bookable
// inventory in the Amadeus test environment. Most ids from the Hotel List API
// return "PROPERTY CODE NOT FOUND IN SYSTEM" when queried for offers.
var sandboxOfferHotels = []string{"RTPAREIF", "RTPARMAI", "XKPAR120"}

// stayDates returns a check-in/check-out window far enough ahead to be bookable.
func stayDates() (string, string) {
	in := time.Now().AddDate(0, 0, 21)
	return in.Format("2006-01-02"), in.AddDate(0, 0, 2).Format("2006-01-02")
}

// TestOffersListLive checks that a real hotel-offers response decodes into the
// DTOs and that the core fields a caller depends on are populated.
func TestOffersListLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	var decoded int
	for _, id := range sandboxOfferHotels {
		offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
			HotelIDs:     []string{id},
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			Adults:       2,
		})
		if err != nil {
			// sandbox inventory is not guaranteed for any given date window
			t.Logf("%s: no offers (%v)", id, err)
			continue
		}
		for _, o := range offers {
			decoded++
			if o.Hotel.HotelID == "" {
				t.Errorf("%s: hotel.hotelId empty", id)
			}
			if o.Type != "hotel-offers" {
				t.Errorf("%s: type = %q, want \"hotel-offers\"", id, o.Type)
			}
			for _, offer := range o.Offers {
				if offer.ID == "" {
					t.Errorf("%s: offer id empty", id)
				}
				if offer.Price.Total == "" {
					t.Errorf("%s: offer %s has empty price.total", id, offer.ID)
				}
				if offer.Price.Currency == "" {
					t.Errorf("%s: offer %s has empty price.currency", id, offer.ID)
				}
				if offer.CheckInDate != checkIn {
					t.Errorf("%s: offer %s checkInDate = %q, want %q",
						id, offer.ID, offer.CheckInDate, checkIn)
				}
			}
		}
	}

	if decoded == 0 {
		t.Skip("sandbox returned no offers for any known hotel; nothing to assert")
	}
	t.Logf("decoded %d hotel-offer groups", decoded)
}

// TestOffersGetByIDLive round-trips an offer id from List through GetByID and
// verifies the single-offer response decodes to the same offer.
func TestOffersGetByIDLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	var offerID string
	for _, id := range sandboxOfferHotels {
		offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
			HotelIDs:     []string{id},
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			Adults:       2,
		})
		if err != nil {
			continue
		}
		for _, o := range offers {
			if len(o.Offers) > 0 {
				offerID = o.Offers[0].ID
				break
			}
		}
		if offerID != "" {
			break
		}
	}
	if offerID == "" {
		t.Skip("no bookable offer available in sandbox to retrieve")
	}

	got, err := s.Offers.GetByID(requestOffers.HotelOffersByIDRequest{OfferID: offerID})
	if err != nil {
		t.Fatalf("GetByID(%s): %v", offerID, err)
	}
	if len(got.Offers) == 0 {
		t.Fatalf("GetByID(%s) returned no offers", offerID)
	}
	if got.Offers[0].ID != offerID {
		t.Errorf("offer id = %q, want %q", got.Offers[0].ID, offerID)
	}
	if got.Hotel.HotelID == "" {
		t.Error("hotel.hotelId empty")
	}
}
