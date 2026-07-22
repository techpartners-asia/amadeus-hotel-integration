package tests

import (
	"testing"

	responseOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/response"
)

// offer is a tiny constructor for readable table-driven fixtures.
func offer(id, roomType, total, currency string) responseOffers.OfferResponse {
	return responseOffers.OfferResponse{
		ID:    id,
		Room:  responseOffers.RoomResponse{Type: roomType},
		Price: responseOffers.PriceResponse{Total: total, Currency: currency},
	}
}

// TestGroupByRoomShape is the core unit test: no network, fully deterministic.
// It pins every behaviour the helper promises - grouping key, group ordering,
// within-group ordering, the cheapest/price-from summary, and the empty-room
// and bad-price edge cases.
func TestGroupByRoomShape(t *testing.T) {
	in := responseOffers.OffersResponse{
		Offers: []responseOffers.OfferResponse{
			offer("C-mid", "C3S", "150.00", "EUR"),
			offer("B-cheap", "B1D", "90.00", "EUR"),
			offer("C-cheap", "C3S", "70.00", "EUR"),
			offer("B-exp", "B1D", "300.00", "EUR"),
			offer("noprice", "C3S", "", "EUR"), // unparseable price
			offer("noroom", "", "50.00", "EUR"), // empty room.type
		},
	}

	groups := in.GroupByRoom()

	// Three room types: "" (50), B1D (from 90), C3S (from 70). Sorted by
	// cheapest price ascending: "" @50, C3S @70, B1D @90.
	if len(groups) != 3 {
		t.Fatalf("got %d groups, want 3: %+v", len(groups), groups)
	}
	wantOrder := []struct {
		room  string
		from  string
		count int
	}{
		{"", "50.00", 1},
		{"C3S", "70.00", 3},
		{"B1D", "90.00", 2},
	}
	for i, w := range wantOrder {
		g := groups[i]
		if g.RoomType != w.room {
			t.Errorf("group %d: RoomType = %q, want %q", i, g.RoomType, w.room)
		}
		if g.PriceFrom != w.from {
			t.Errorf("group %d (%s): PriceFrom = %q, want %q", i, w.room, g.PriceFrom, w.from)
		}
		if len(g.Offers) != w.count {
			t.Errorf("group %d (%s): %d offers, want %d", i, w.room, len(g.Offers), w.count)
		}
	}

	// Within C3S: cheapest first (70, 150), unparseable price (noprice) last.
	c3s := groups[1]
	gotIDs := []string{c3s.Offers[0].ID, c3s.Offers[1].ID, c3s.Offers[2].ID}
	wantIDs := []string{"C-cheap", "C-mid", "noprice"}
	for i := range wantIDs {
		if gotIDs[i] != wantIDs[i] {
			t.Errorf("C3S offer %d: id = %q, want %q (order %v)", i, gotIDs[i], wantIDs[i], gotIDs)
		}
	}

	// Cheapest points at Offers[0] and drives the representative room + price.
	if c3s.Cheapest == nil || c3s.Cheapest.ID != "C-cheap" {
		t.Errorf("C3S Cheapest = %v, want offer C-cheap", c3s.Cheapest)
	}
	if c3s.Room.Type != "C3S" {
		t.Errorf("C3S representative Room.Type = %q, want C3S", c3s.Room.Type)
	}
	if c3s.PriceFromCurrency != "EUR" {
		t.Errorf("C3S PriceFromCurrency = %q, want EUR", c3s.PriceFromCurrency)
	}
}

// TestGroupByRoomEmpty verifies the helper is safe on a hotel with no offers.
func TestGroupByRoomEmpty(t *testing.T) {
	var empty responseOffers.OffersResponse
	if groups := empty.GroupByRoom(); len(groups) != 0 {
		t.Errorf("empty hotel: got %d groups, want 0", len(groups))
	}
}

// TestGroupByRoomAllPricesMissing verifies grouping still works (and preserves
// first-seen order) when no offer carries a parseable price.
func TestGroupByRoomAllPricesMissing(t *testing.T) {
	in := responseOffers.OffersResponse{
		Offers: []responseOffers.OfferResponse{
			offer("a", "ZZZ", "", "EUR"),
			offer("b", "AAA", "", "EUR"),
		},
	}
	groups := in.GroupByRoom()
	if len(groups) != 2 {
		t.Fatalf("got %d groups, want 2", len(groups))
	}
	// Equal (+Inf) prices fall back to RoomType as tie-breaker: AAA before ZZZ.
	if groups[0].RoomType != "AAA" || groups[1].RoomType != "ZZZ" {
		t.Errorf("tie-break order = [%q %q], want [AAA ZZZ]", groups[0].RoomType, groups[1].RoomType)
	}
}

// TestGroupByRoomLive groups a real bestRateOnly=false response and checks the
// invariants that must hold against live data: every offer lands in exactly one
// group, no group is empty, and each group's cheapest really is its minimum.
func TestGroupByRoomLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	hotels, err := s.Offers.List(allRatesFor("RTPARVAL", checkIn, checkOut))
	if err != nil {
		t.Fatalf("offers list: %v", err)
	}
	if len(hotels) == 0 {
		t.Skip("RTPARVAL returned no offers for these dates")
	}

	grouped := responseOffers.GroupByRoom(hotels)
	for _, gh := range grouped {
		var regrouped int
		for _, room := range gh.Rooms {
			if len(room.Offers) == 0 {
				t.Errorf("hotel %s: room %q has an empty group", gh.Hotel.HotelID, room.RoomType)
			}
			regrouped += len(room.Offers)
			for _, o := range room.Offers {
				if o.Room.Type != room.RoomType {
					t.Errorf("hotel %s: offer %s (room %q) grouped under %q",
						gh.Hotel.HotelID, o.ID, o.Room.Type, room.RoomType)
				}
			}
			if room.Cheapest != nil && room.PriceFrom != room.Cheapest.Price.Total {
				t.Errorf("hotel %s room %q: PriceFrom %q != cheapest %q",
					gh.Hotel.HotelID, room.RoomType, room.PriceFrom, room.Cheapest.Price.Total)
			}
		}
		// No offer is dropped or duplicated by grouping.
		var original int
		for _, h := range hotels {
			if h.Hotel.HotelID == gh.Hotel.HotelID {
				original = len(h.Offers)
			}
		}
		if regrouped != original {
			t.Errorf("hotel %s: grouped %d offers, response had %d", gh.Hotel.HotelID, regrouped, original)
		}
		t.Logf("hotel %s: %d offers -> %d room groups", gh.Hotel.HotelID, original, len(gh.Rooms))
	}
}
