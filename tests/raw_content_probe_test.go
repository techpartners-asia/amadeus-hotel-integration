package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
)

// contentPath mirrors the unexported constant in the content usecase; this
// probe talks to the endpoint directly rather than through the SDK.
const rawContentPath = "/reference-data/locations/by-hotel"

// TestRawContentProbe prints the unparsed JSON returned by the Hotel Content
// endpoint, bypassing the SDK's response decoding. It exists to inspect fields
// by eye; it asserts nothing. Whether the DTOs actually capture that payload is
// checked by TestContentDTORoundTrip and TestContentDTOCoversAPIResponse.
//
// Skipped by default: it is an inspection tool, not a regression test, and one
// FULL-view response runs to ~10k lines. Run it with -v, or the Go tool
// discards its output:
//
//	AMADEUS_PROBE_RAW=1 go test ./tests/ -run TestRawContentProbe -count=1 -v
func TestRawContentProbe(t *testing.T) {
	if os.Getenv("AMADEUS_PROBE_RAW") != "1" {
		t.Skip("set AMADEUS_PROBE_RAW=1 to dump raw content responses")
	}
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)

	ids := hotelIDsOf(searchHotels(t, s, "PAR", bookableChain, 3))
	client := amadeusIntegration.NewClient(constants.CONTENT_BASE_URL)

	for _, id := range ids {
		req := requestContentDTO.ContentByIDRequest{ID: id}

		res, err := client.R().SetQueryParams(req.ToQueryParams()).Get(rawContentPath)
		if err != nil {
			t.Fatalf("content %s: %v", id, err)
		}

		fmt.Printf("\n=== CONTENT %s: GET %s -> %d ===\n%s\n",
			id, res.Request.URL, res.StatusCode(), pretty(res.String()))
	}
}
