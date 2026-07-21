package requestContentDTO

import (
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

type (
	// ContentByIDRequest fetches hotel content details for a single property.
	// Endpoint: GET /reference-data/locations/by-hotel
	ContentByIDRequest struct {
		// ID - Amadeus property code (hotelID). Example: "ADNYCCTB". (required)
		ID string `json:"id" required:"true"`
		// Fields - restrict the response to the listed content blocks
		// (e.g. "hotel", "rooms", "facilities"). Optional.
		Fields []string `json:"fields,omitempty"`
		// Lang - language for textual content (ISO 639-1, e.g. "EN", "FR"). Optional.
		Lang string `json:"lang,omitempty"`
		// View - response detail level. See searchcriteria.AllContentViews. Optional.
		View searchcriteria.ContentView `json:"view,omitempty"`
	}
)

func (r *ContentByIDRequest) ToQueryParams() map[string]string {
	queryParams := map[string]string{
		"hotelID": r.ID,
	}

	if len(r.Fields) > 0 {
		queryParams["fields"] = strings.Join(r.Fields, ",")
	}

	if r.Lang != "" {
		queryParams["lang"] = r.Lang
	}

	// Amadeus defaults to a light view that only returns the "basic" block.
	// Ask for FULL unless the caller wants something narrower, so rooms,
	// facilities, policies, awards and pointOfInterest come back populated.
	if r.View != "" {
		queryParams["view"] = string(r.View)
	} else {
		queryParams["view"] = string(searchcriteria.ContentViewFull)
	}

	return queryParams
}
