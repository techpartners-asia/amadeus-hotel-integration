package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
)

// TestRawOffersProbe prints the unparsed JSON returned by the two Hotel Offers
// endpoints (list and by-id), bypassing the SDK's response decoding. It exists
// to inspect fields by eye; it asserts nothing. Whether the DTOs actually
// capture that payload is checked by TestOffersListDTORoundTrip and the
// coverage tests in dto_fidelity_test.go.
//
// Skipped by default: it is an inspection tool, not a regression test, and it
// costs live API calls. Run it with -v, or the Go tool discards its output:
//
//	AMADEUS_PROBE_RAW=1 go test ./tests/ -run TestRawOffersProbe -count=1 -v
func TestRawOffersProbe(t *testing.T) {
	if os.Getenv("AMADEUS_PROBE_RAW") != "1" {
		t.Skip("set AMADEUS_PROBE_RAW=1 to dump raw offers responses")
	}
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	ids := hotelIDsOf(searchHotels(t, s, "PAR", bookableChain, 5))
	client := amadeusIntegration.NewClient(constants.OFFERS_BASE_URL)

	listReq := requestOffers.HotelOffersListRequest{
		HotelIDs:     ids,
		CheckInDate:  checkIn,
		CheckOutDate: checkOut,
		Adults:       2,
		BestRateOnly: new(bool), // false to see all offers, not just the cheapest
	}
	res, err := client.R().SetQueryParams(listReq.ToQueryParams()).Get("")
	if err != nil {
		t.Fatalf("offers list: %v", err)
	}
	// fmt, not t.Logf: t.Logf is swallowed unless the test runs with -v.
	fmt.Printf("\n=== LIST: GET %s -> %d ===\n%s\n", res.Request.URL, res.StatusCode(), pretty(res.String()))

	// Pull an offer id out of the list response to probe the by-id endpoint.
	var list struct {
		Data []struct {
			Offers []struct {
				ID string `json:"id"`
			} `json:"offers"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(res.String()), &list); err != nil {
		t.Fatalf("decode list: %v", err)
	}
	offerID := ""
	for _, h := range list.Data {
		if len(h.Offers) > 0 {
			offerID = h.Offers[0].ID
			break
		}
	}
	if offerID == "" {
		t.Skip("no offer id in list response; skipping by-id probe")
	}

	byID := requestOffers.HotelOffersByIDRequest{OfferID: offerID}
	res2, err := client.R().SetQueryParams(byID.ToQueryParams()).Get(byID.OfferID)
	if err != nil {
		t.Fatalf("offers by-id: %v", err)
	}
	fmt.Printf("\n=== BY-ID: GET %s -> %d ===\n%s\n", res2.Request.URL, res2.StatusCode(), pretty(res2.String()))
}

// pretty indents JSON for readability, returning the input unchanged when it is
// not valid JSON (e.g. an HTML error page from the gateway).
func pretty(body string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(body), "", "  "); err != nil {
		return body
	}
	return buf.String()
}
