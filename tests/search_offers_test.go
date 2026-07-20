package tests

import (
	"testing"

	sdk "github.com/techpartners-asia/amadeus-hotel-integration"

	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	responseHotelListDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/response"
	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
	responseHotelOffersDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/response"
)

// These tests exercise the flow a real caller uses: search a city for hotels,
// then price the hotels that came back. The two APIs are separate services and
// their inventories do not match - the Hotel List API happily returns
// properties that the Hotel Offers API has never heard of - so the interesting
// behaviour is in how the join degrades, not in the happy path.

// bookableChain is a chain whose Paris properties carry bookable inventory in
// the Amadeus test environment. Searching without a chain filter mostly returns
// properties that the offers API rejects outright.
const bookableChain = "RT"

// searchHotels returns up to limit hotels for the city, filtered to a chain
// with sandbox inventory.
func searchHotels(t *testing.T, s *sdk.SDK, city, chain string, limit int) []responseHotelListDTO.GeneralInfoResponse {
	t.Helper()

	hotels, err := s.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{
		CityCode:   city,
		ChainCodes: []string{chain},
	})
	if err != nil {
		t.Fatalf("search %s/%s: %v", city, chain, err)
	}
	if len(hotels) == 0 {
		t.Skipf("search %s/%s returned no hotels", city, chain)
	}
	if len(hotels) > limit {
		hotels = hotels[:limit]
	}
	return hotels
}

func hotelIDsOf(hotels []responseHotelListDTO.GeneralInfoResponse) []string {
	ids := make([]string, len(hotels))
	for i, h := range hotels {
		ids[i] = h.HotelId
	}
	return ids
}

// TestSearchThenOffersBatched covers the efficient path: one search, then a
// single offers call for every hotel id at once. Amadeus returns offers only
// for the subset that has inventory, so the result is expected to be smaller
// than the search result - but every hotel it does return must be one we asked
// for, with a priced, dated offer attached.
func TestSearchThenOffersBatched(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	hotels := searchHotels(t, s, "PAR", bookableChain, 10)
	ids := hotelIDsOf(hotels)
	t.Logf("search returned %d hotels", len(ids))

	requested := map[string]bool{}
	for _, id := range ids {
		requested[id] = true
	}

	offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
		HotelIDs:     ids,
		CheckInDate:  checkIn,
		CheckOutDate: checkOut,
		Adults:       2,
	})
	if err != nil {
		// A provider-level error on any single id fails the whole batch.
		// TestSearchThenOffersPerHotel covers the fallback for this case.
		t.Skipf("batch offers call failed for all %d hotels: %v", len(ids), err)
	}

	if len(offers) == 0 {
		t.Skip("no hotel in the search result had offers for these dates")
	}
	t.Logf("%d of %d searched hotels had offers", len(offers), len(ids))

	for _, group := range offers {
		if !requested[group.Hotel.HotelID] {
			t.Errorf("offers returned hotel %q that was not in the search result", group.Hotel.HotelID)
		}
		assertOffersPriced(t, group, checkIn, checkOut)
	}
}

// TestSearchThenOffersPerHotel covers the resilient path: price each hotel from
// the search individually, so one unpriceable property cannot take down the
// whole result set. This is the pattern to use when the hotel ids come straight
// from a search and have not been vetted.
func TestSearchThenOffersPerHotel(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	hotels := searchHotels(t, s, "PAR", bookableChain, 8)

	var priced, unpriced int
	for _, h := range hotels {
		offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
			HotelIDs:     []string{h.HotelId},
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			Adults:       2,
		})
		if err != nil || len(offers) == 0 {
			unpriced++
			continue
		}
		priced++
		for _, group := range offers {
			if group.Hotel.HotelID != h.HotelId {
				t.Errorf("asked for %q, offers returned %q", h.HotelId, group.Hotel.HotelID)
			}
			assertOffersPriced(t, group, checkIn, checkOut)
		}
	}

	t.Logf("of %d searched hotels: %d priced, %d without offers", len(hotels), priced, unpriced)
	if priced == 0 {
		t.Skip("no hotel in the search result could be priced for these dates")
	}
}

// TestSearchThenOffersCarriesSearchMetadata checks that the two APIs agree on
// the identity of a hotel: the offers response should echo the same hotelId and
// chainCode the search reported, so callers can join the two safely.
func TestSearchThenOffersCarriesSearchMetadata(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	hotels := searchHotels(t, s, "PAR", bookableChain, 8)
	byID := map[string]responseHotelListDTO.GeneralInfoResponse{}
	for _, h := range hotels {
		byID[h.HotelId] = h
	}

	var checked int
	for _, h := range hotels {
		offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
			HotelIDs:     []string{h.HotelId},
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			Adults:       2,
		})
		if err != nil || len(offers) == 0 {
			continue
		}
		for _, group := range offers {
			from := byID[group.Hotel.HotelID]
			if group.Hotel.ChainCode != from.ChainCode {
				t.Errorf("hotel %s: search chainCode %q, offers chainCode %q",
					group.Hotel.HotelID, from.ChainCode, group.Hotel.ChainCode)
			}
			if group.Hotel.Name == "" {
				t.Errorf("hotel %s: offers response has empty name", group.Hotel.HotelID)
			}
			checked++
		}
	}

	if checked == 0 {
		t.Skip("no hotel could be priced, nothing to cross-check")
	}
	t.Logf("cross-checked %d hotels between search and offers", checked)
}

// assertOffersPriced verifies an offers group carries the fields a booking flow
// needs next: a stable offer id, a price, and the dates that were requested.
func assertOffersPriced(t *testing.T, group responseHotelOffersDTO.OffersResponse, checkIn, checkOut string) {
	t.Helper()

	if len(group.Offers) == 0 {
		t.Errorf("hotel %s: available=%v but no offers attached", group.Hotel.HotelID, group.Available)
		return
	}
	for _, o := range group.Offers {
		if o.ID == "" {
			t.Errorf("hotel %s: offer has no id", group.Hotel.HotelID)
		}
		if o.Price.Total == "" {
			t.Errorf("hotel %s: offer %s has no price.total", group.Hotel.HotelID, o.ID)
		}
		if o.Price.Currency == "" {
			t.Errorf("hotel %s: offer %s has no price.currency", group.Hotel.HotelID, o.ID)
		}
		if o.CheckInDate != checkIn {
			t.Errorf("hotel %s: offer %s checkInDate = %q, want %q",
				group.Hotel.HotelID, o.ID, o.CheckInDate, checkIn)
		}
		if o.CheckOutDate != checkOut {
			t.Errorf("hotel %s: offer %s checkOutDate = %q, want %q",
				group.Hotel.HotelID, o.ID, o.CheckOutDate, checkOut)
		}
	}
}
