package tests

import (
	"testing"

	requestHotelListGeocodeDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/geocode"
	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

// TestGeocodeQueryParamsSendsCoordinates guards the bug that made this endpoint
// unusable: the request used to carry cityCode and no coordinates, so Amadeus
// answered "INVALID GEOGRAPHICAL ZONE - Missing coordinates" every time.
func TestGeocodeQueryParamsSendsCoordinates(t *testing.T) {
	req := requestHotelListGeocodeDTO.HotelListByGeocodeRequest{
		Latitude:  48.85,
		Longitude: 2.29,
	}

	q := req.ToQueryParams()

	if q["latitude"] != "48.85" {
		t.Errorf("latitude = %q, want %q", q["latitude"], "48.85")
	}
	if q["longitude"] != "2.29" {
		t.Errorf("longitude = %q, want %q", q["longitude"], "2.29")
	}
	if len(q) != 2 {
		t.Errorf("minimal request should send only coordinates, got %v", q)
	}
}

// TestGeocodeQueryParamsOmitsUnset verifies optional params are omitted rather
// than sent as zero values (radius=0, radiusUnit=), which Amadeus rejects.
func TestGeocodeQueryParamsOmitsUnset(t *testing.T) {
	req := requestHotelListGeocodeDTO.HotelListByGeocodeRequest{
		Latitude:  48.85,
		Longitude: 2.29,
	}

	q := req.ToQueryParams()

	for _, k := range []string{"radius", "radiusUnit", "chainCodes", "amenities", "ratings", "hotelSource"} {
		if v, ok := q[k]; ok {
			t.Errorf("unset optional param %q should be omitted, got %q", k, v)
		}
	}
}

// TestGeocodeQueryParamsIncludesSet verifies set optional params are emitted.
func TestGeocodeQueryParamsIncludesSet(t *testing.T) {
	req := requestHotelListGeocodeDTO.HotelListByGeocodeRequest{
		Latitude:   48.85,
		Longitude:  2.29,
		Radius:     10,
		RadiusUnit: searchcriteria.RadiusUnitKM,
		ChainCodes: []string{"RT", "XK"},
		Ratings:    []searchcriteria.Rating{searchcriteria.Rating4, searchcriteria.Rating5},
	}

	q := req.ToQueryParams()

	want := map[string]string{
		"radius":     "10",
		"radiusUnit": "KM",
		"chainCodes": "RT,XK",
		"ratings":    "4,5",
	}
	for k, v := range want {
		if q[k] != v {
			t.Errorf("param %q = %q, want %q", k, q[k], v)
		}
	}
}

// TestGeocodeLive confirms the endpoint actually returns hotels around a point.
func TestGeocodeLive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)

	hotels, err := s.List.HotelListByGeocode(requestHotelListGeocodeDTO.HotelListByGeocodeRequest{
		Latitude:  48.85,
		Longitude: 2.29,
		Radius:    5,
	})
	if err != nil {
		t.Fatalf("by-geocode: %v", err)
	}
	if len(hotels) == 0 {
		t.Fatal("by-geocode returned no hotels near the Eiffel Tower")
	}
	for _, h := range hotels {
		if h.HotelId == "" {
			t.Error("hotel with empty hotelId")
			break
		}
	}
	t.Logf("by-geocode returned %d hotels", len(hotels))
}
