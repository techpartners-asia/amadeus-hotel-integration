package tests

import (
	"testing"

	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
)

// TestOffersListQueryParamsOmitsEmpty verifies that a minimal request only emits
// the required hotelIds param and never sends empty/zero-valued optional params
// (e.g. adults=0 or currency=) that Amadeus would reject.
func TestOffersListQueryParamsOmitsEmpty(t *testing.T) {
	req := requestOffers.HotelOffersListRequest{
		HotelIDs: []string{"MCLONGHM", "ACPAR419"},
	}

	q := req.ToQueryParams()

	if got := q["hotelIds"]; got != "MCLONGHM,ACPAR419" {
		t.Fatalf("hotelIds = %q, want comma-joined ids", got)
	}
	if len(q) != 1 {
		t.Fatalf("expected only hotelIds, got %d params: %v", len(q), q)
	}
	for _, k := range []string{"adults", "roomQuantity", "currency", "checkInDate", "bestRateOnly", "includeClosed", "lang"} {
		if _, ok := q[k]; ok {
			t.Errorf("unset optional param %q should be omitted, got %q", k, q[k])
		}
	}
}

// TestOffersListQueryParamsIncludesSet verifies set fields are emitted, including
// the tri-state *bool flags.
func TestOffersListQueryParamsIncludesSet(t *testing.T) {
	req := requestOffers.HotelOffersListRequest{
		HotelIDs:     []string{"MCLONGHM"},
		Adults:       2,
		RoomQuantity: 1,
		CheckInDate:  "2026-08-01",
		Currency:     "EUR",
		ChildAges:    []int{6, 9},
		BestRateOnly: requestOffers.Bool(false),
	}

	q := req.ToQueryParams()

	want := map[string]string{
		"hotelIds":     "MCLONGHM",
		"adults":       "2",
		"roomQuantity": "1",
		"checkInDate":  "2026-08-01",
		"currency":     "EUR",
		"childAges":    "6,9",
		"bestRateOnly": "false",
	}
	for k, v := range want {
		if q[k] != v {
			t.Errorf("param %q = %q, want %q", k, q[k], v)
		}
	}
}
