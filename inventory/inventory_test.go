package inventory_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/geo"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeustest"
	"github.com/techpartners-asia/amadeus-hotel-integration/inventory"
)

const (
	pathByCity    = "/v1/reference-data/locations/hotels/by-city"
	pathByGeocode = "/v1/reference-data/locations/hotels/by-geocode"
	pathByHotels  = "/v1/reference-data/locations/hotels/by-hotels"
)

// newService returns a service backed by the by-city fixture.
func newService(t *testing.T) (inventory.Service, *amadeustest.Server) {
	t.Helper()
	server := amadeustest.New(t)
	server.Fixture(t, http.MethodGet, pathByCity, "hotels-by-city")
	server.Fixture(t, http.MethodGet, pathByGeocode, "hotels-by-city")
	server.Fixture(t, http.MethodGet, pathByHotels, "hotels-by-city")
	return inventory.NewService(server.Client()), server
}

func TestByCityMapsEveryField(t *testing.T) {
	service, _ := newService(t)

	hotels, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("ByCity: %v", err)
	}
	if len(hotels) != 4 {
		t.Fatalf("got %d hotels, want 4", len(hotels))
	}

	got := hotels[0]
	if got.ID != "MCPARC12" {
		t.Errorf("ID = %q", got.ID)
	}
	if got.Name != "AC HOTEL BY MARRIOTT PARIS PORTE MAILLOT" {
		t.Errorf("Name = %q", got.Name)
	}
	if got.ChainCode != "MC" || got.BrandCode != "AK" || got.MasterChainCode != "EM" {
		t.Errorf("chain codes = %q/%q/%q", got.ChainCode, got.BrandCode, got.MasterChainCode)
	}
	if got.IATACode != "PAR" {
		t.Errorf("IATACode = %q", got.IATACode)
	}
	// dupeId arrives as a JSON number here and as a string from Hotel Search;
	// the domain normalises both to a string.
	if got.DupeID != "700027723" {
		t.Errorf("DupeID = %q, want the number rendered as a string", got.DupeID)
	}
	if got.LastUpdate != "2023-08-08T00:00:00" {
		t.Errorf("LastUpdate = %q", got.LastUpdate)
	}
	if got.Position == nil || got.Position.Latitude != 48.87825 || got.Position.Longitude != 2.28454 {
		t.Errorf("Position = %+v", got.Position)
	}
	if got.Address == nil || got.Address.PostalCode != "75017" || got.Address.CityName != "PARIS" {
		t.Errorf("Address = %+v", got.Address)
	}
	if got.Address != nil && len(got.Address.Lines) != 1 {
		t.Errorf("Address.Lines = %v", got.Address.Lines)
	}
	if got.DistanceFromSearch == nil || got.DistanceFromSearch.Value != 3.79 ||
		got.DistanceFromSearch.Unit != geo.Kilometers {
		t.Errorf("DistanceFromSearch = %+v", got.DistanceFromSearch)
	}
	if got.Sponsored {
		t.Error("Sponsored = true, want false")
	}
}

func TestSponsoredPlacementIsSurfaced(t *testing.T) {
	// Paid placement changes what result order means, so it must not be dropped.
	service, _ := newService(t)

	hotels, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("ByCity: %v", err)
	}
	if !hotels[1].Sponsored {
		t.Errorf("%s should be marked sponsored", hotels[1].ID)
	}
}

func TestAbsentCoordinatesStayNilRatherThanBecomingNullIsland(t *testing.T) {
	// 0,0 is a real point in the Gulf of Guinea. Mapping an unlocatable
	// property to it would put it in the Atlantic and it would pass a
	// non-zero check.
	service, _ := newService(t)

	hotels, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("ByCity: %v", err)
	}

	unlocated := hotels[3]
	if unlocated.Position != nil {
		t.Errorf("Position = %+v, want nil for a property with no geoCode", unlocated.Position)
	}
	if unlocated.Address != nil {
		t.Errorf("Address = %+v, want nil when Amadeus sent none", unlocated.Address)
	}
	if unlocated.DistanceFromSearch != nil {
		t.Errorf("DistanceFromSearch = %+v, want nil", unlocated.DistanceFromSearch)
	}
	if unlocated.DupeID != "" {
		t.Errorf("DupeID = %q, want empty when absent", unlocated.DupeID)
	}
}

func TestPartialAddressIsPreserved(t *testing.T) {
	service, _ := newService(t)
	hotels, _ := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})

	sparse := hotels[2].Address
	if sparse == nil {
		t.Fatal("address should be present even when only partly filled")
	}
	if sparse.CityName != "PARIS" || sparse.CountryCode != "FR" {
		t.Errorf("address = %+v", sparse)
	}
	if len(sparse.Lines) != 0 || sparse.PostalCode != "" {
		t.Errorf("absent address parts should stay empty, got %+v", sparse)
	}
	if sparse.IsEmpty() {
		t.Error("an address with a city should not report IsEmpty")
	}
}

func TestQueryParametersAreSentCorrectly(t *testing.T) {
	service, server := newService(t)

	_, err := service.ByCity(context.Background(), inventory.CityQuery{
		CityCode: "PAR",
		Filters: inventory.Filters{
			Radius:     10,
			RadiusUnit: geo.Miles,
			ChainCodes: []string{"MC", "HL"},
			Amenities:  []codes.Amenity{codes.AmenitySwimmingPool, codes.AmenitySpa},
			Ratings:    []codes.Rating{codes.Rating4, codes.Rating5},
			Source:     codes.HotelSourceAll,
		},
	})
	if err != nil {
		t.Fatalf("ByCity: %v", err)
	}

	query := server.LastRequest(t).Query
	want := map[string]string{
		"cityCode":    "PAR",
		"radius":      "10",
		"radiusUnit":  "MILE",
		"chainCodes":  "MC,HL",
		"amenities":   "SWIMMING_POOL,SPA",
		"ratings":     "4,5",
		"hotelSource": "ALL",
	}
	for key, expected := range want {
		if got := query.Get(key); got != expected {
			t.Errorf("query[%s] = %q, want %q", key, got, expected)
		}
	}
}

func TestUnsetFiltersAreOmittedEntirely(t *testing.T) {
	// Amadeus rejects radius=0 and radiusUnit=, so an unset filter must not be
	// sent as an empty value.
	service, server := newService(t)

	if _, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"}); err != nil {
		t.Fatalf("ByCity: %v", err)
	}

	query := server.LastRequest(t).Query
	for _, key := range []string{"radius", "radiusUnit", "chainCodes", "amenities", "ratings", "hotelSource"} {
		if _, present := query[key]; present {
			t.Errorf("unset filter %q was sent as %q", key, query.Get(key))
		}
	}
}

func TestByGeocodeSendsCoordinates(t *testing.T) {
	service, server := newService(t)

	_, err := service.ByGeocode(context.Background(), inventory.GeocodeQuery{
		Position: geo.Coordinates{Latitude: 48.85, Longitude: 2.29},
	})
	if err != nil {
		t.Fatalf("ByGeocode: %v", err)
	}

	query := server.LastRequest(t).Query
	if query.Get("latitude") != "48.85" || query.Get("longitude") != "2.29" {
		t.Errorf("coordinates sent as %q, %q", query.Get("latitude"), query.Get("longitude"))
	}
}

func TestByIDsJoinsCodes(t *testing.T) {
	service, server := newService(t)

	_, err := service.ByIDs(context.Background(), inventory.IDsQuery{
		HotelIDs: []string{"MCLONGHM", "ACPAR419"},
	})
	if err != nil {
		t.Fatalf("ByIDs: %v", err)
	}
	if got := server.LastRequest(t).Query.Get("hotelIds"); got != "MCLONGHM,ACPAR419" {
		t.Errorf("hotelIds = %q", got)
	}
}

// Validation happens before any network call, so a caller learns which field is
// wrong instead of decoding an Amadeus 400.
func TestValidationRejectsBadQueriesWithoutCallingAmadeus(t *testing.T) {
	service, server := newService(t)

	cases := []struct {
		name string
		call func() error
	}{
		{"missing city code", func() error {
			_, err := service.ByCity(context.Background(), inventory.CityQuery{})
			return err
		}},
		{"city code of the wrong length", func() error {
			_, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PARIS"})
			return err
		}},
		{"unknown amenity", func() error {
			_, err := service.ByCity(context.Background(), inventory.CityQuery{
				CityCode: "PAR",
				Filters:  inventory.Filters{Amenities: []codes.Amenity{"BABY-SITTING"}},
			})
			return err
		}},
		{"too many ratings", func() error {
			_, err := service.ByCity(context.Background(), inventory.CityQuery{
				CityCode: "PAR",
				Filters: inventory.Filters{Ratings: []codes.Rating{
					codes.Rating1, codes.Rating2, codes.Rating3, codes.Rating4, codes.Rating5,
				}},
			})
			return err
		}},
		{"radius unit Amadeus rejects here", func() error {
			_, err := service.ByCity(context.Background(), inventory.CityQuery{
				CityCode: "PAR",
				Filters:  inventory.Filters{RadiusUnit: geo.Feet},
			})
			return err
		}},
		{"coordinates out of range", func() error {
			_, err := service.ByGeocode(context.Background(), inventory.GeocodeQuery{
				Position: geo.Coordinates{Latitude: 100, Longitude: 0},
			})
			return err
		}},
		{"no hotel IDs", func() error {
			_, err := service.ByIDs(context.Background(), inventory.IDsQuery{})
			return err
		}},
		{"malformed hotel ID", func() error {
			_, err := service.ByIDs(context.Background(), inventory.IDsQuery{HotelIDs: []string{"nope"}})
			return err
		}},
	}

	before := len(server.Requests())
	for _, c := range cases {
		err := c.call()
		if !errors.Is(err, apierr.ErrValidation) {
			t.Errorf("%s: err = %v, want ErrValidation", c.name, err)
		}
	}
	if after := len(server.Requests()); after != before {
		t.Errorf("%d invalid queries reached the network", after-before)
	}
}

func TestValidationReportsEveryProblemAtOnce(t *testing.T) {
	service, _ := newService(t)

	_, err := service.ByCity(context.Background(), inventory.CityQuery{
		CityCode: "",
		Filters: inventory.Filters{
			Radius:     -1,
			Amenities:  []codes.Amenity{"NOT_A_THING"},
			ChainCodes: []string{"TOOLONG"},
		},
	})

	var errs apierr.ValidationErrors
	if !errors.As(err, &errs) {
		t.Fatalf("err = %v (%T), want ValidationErrors", err, err)
	}
	if len(errs) != 4 {
		t.Errorf("reported %d problems, want all 4: %v", len(errs), err)
	}
}

func TestAmadeusErrorsSurfaceTyped(t *testing.T) {
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, pathByCity, http.StatusBadRequest,
		`{"errors":[{"status":400,"code":572,"title":"INVALID OPTION","detail":"unknown city","source":{"parameter":"cityCode"}}]}`)
	service := inventory.NewService(server.Client())

	_, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "ZZZ"})
	if !errors.Is(err, apierr.ErrInvalidRequest) {
		t.Fatalf("err = %v, want ErrInvalidRequest", err)
	}

	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		t.Fatal("error should also be an *APIError")
	}
	if len(apiErr.Details) != 1 || apiErr.Details[0].Code != 572 {
		t.Errorf("details = %+v", apiErr.Details)
	}
	if apiErr.Details[0].Source.Parameter != "cityCode" {
		t.Errorf("source = %+v, want the offending parameter", apiErr.Details[0].Source)
	}
}

func TestEmptyResultIsNotAnError(t *testing.T) {
	// A search matching nothing is a valid answer, not a failure.
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, pathByCity, http.StatusOK, `{"meta":{"count":0},"data":[]}`)
	service := inventory.NewService(server.Client())

	hotels, err := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("empty result returned an error: %v", err)
	}
	if len(hotels) != 0 {
		t.Errorf("got %d hotels, want 0", len(hotels))
	}
}

func TestIDsHelperFeedsTheOtherContexts(t *testing.T) {
	service, _ := newService(t)
	hotels, _ := service.ByCity(context.Background(), inventory.CityQuery{CityCode: "PAR"})

	ids := inventory.IDs(hotels)
	if len(ids) != 4 || ids[0] != "MCPARC12" {
		t.Errorf("IDs = %v", ids)
	}
}

func TestContextCancellationPropagates(t *testing.T) {
	service, _ := newService(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := service.ByCity(ctx, inventory.CityQuery{CityCode: "PAR"}); !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}
