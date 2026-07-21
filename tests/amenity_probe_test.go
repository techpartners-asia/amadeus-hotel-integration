package tests

import (
	"os"
	"strings"
	"testing"

	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

// TestAmenityCodesAcceptedByAmadeus sends each declared amenity code to the live
// by-city endpoint one at a time and reports any the API rejects as unknown. It
// is the only check that can catch a code that was mis-documented, renamed or
// retired upstream, since every other test compares the SDK against a
// transcription of the same documentation.
//
// This test earned its keep on the first run: Amadeus documents "BABY-SITTING"
// and "BAR or LOUNGE", and the live API rejects both. See reference/amenity.go.
//
// Skipped by default: it costs one request per amenity and depends on Amadeus
// being reachable. Run with AMADEUS_PROBE_AMENITIES=1.
//
// A pass means Amadeus recognises the code, not that the list is exhaustive:
// this cannot discover a code the SDK has never heard of.
func TestAmenityCodesAcceptedByAmadeus(t *testing.T) {
	if os.Getenv("AMADEUS_PROBE_AMENITIES") != "1" {
		t.Skip("set AMADEUS_PROBE_AMENITIES=1 to probe each amenity against the live API")
	}

	client := newSDK(t)

	for _, amenity := range searchcriteria.AllAmenities() {
		t.Run(string(amenity), func(t *testing.T) {
			_, err := client.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{
				CityCode:  "NYC",
				Amenities: []searchcriteria.Amenity{amenity},
			})
			if err == nil {
				return
			}

			// Two 400s mean opposite things here. 7211 is the API saying it does
			// not know the code, which makes the constant wrong. 895 means the
			// code was understood and simply matched no property in the test
			// environment, which says nothing about the constant.
			msg := err.Error()
			switch {
			case strings.Contains(msg, "7211"):
				t.Errorf("Amadeus does not recognise amenity %q (%s): %v", amenity, amenity.Label(), err)
			case strings.Contains(msg, "895"):
				t.Logf("amenity %q accepted, no matching properties in the test environment", amenity)
			default:
				t.Errorf("unexpected error probing amenity %q: %v", amenity, err)
			}
		})
	}
}
