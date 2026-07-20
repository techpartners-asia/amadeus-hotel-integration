package tests

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	responseContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/response"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	responseHotelListDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/response"
	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
	responseHotelOffersDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/response"
)

// These tests compare what Amadeus actually sends against what the response
// DTOs declare. encoding/json silently drops keys with no matching struct
// field, so a DTO can lose data indefinitely without any test failing. The
// checks below catch both failure modes:
//
//	shape mismatch - API sends an array where the DTO expects an object
//	                 (or vice versa); this surfaces as a decode error
//	dropped field  - API sends a key the DTO has no field for; this is
//	                 silent, and is what the "not covered" check finds

// dtoPaths walks a Go type and records every json path it can decode, mapping
// the path to the kind of value it accepts ("obj", "arr" or "scalar").
// Slices and pointers are flattened to their element type, so a []T and a T
// produce the same child paths.
func dtoPaths(t reflect.Type, prefix string, seen map[reflect.Type]bool, out map[string]string) {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}
	if t.Kind() == reflect.Map {
		dtoPaths(t.Elem(), prefix+".*", seen, out)
		return
	}
	if t == reflect.TypeOf(time.Time{}) {
		return
	}
	if t.Kind() != reflect.Struct {
		return
	}
	if seen[t] && prefix != "" {
		return // guard against self-referential types
	}
	seen[t] = true
	defer delete(seen, t)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			name = f.Name
		}
		p := name
		if prefix != "" {
			p = prefix + "." + name
		}
		k := kindOfType(f.Type)
		if k == "scalar" {
			// record the concrete scalar kind so bool/string/number are distinguished
			ft := f.Type
			for ft.Kind() == reflect.Slice || ft.Kind() == reflect.Array || ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			if sk := scalarKindOf(ft); sk != "" {
				k = sk
			}
		}
		out[p] = k
		dtoPaths(f.Type, p, seen, out)
	}
}

func kindOfType(t reflect.Type) string {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// time.Time is a struct but decodes from a JSON string
	if t == reflect.TypeOf(time.Time{}) {
		return "string"
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		return "arr"
	case reflect.Struct, reflect.Map:
		return "obj"
	default:
		return "scalar"
	}
}

// apiPaths walks decoded JSON and records every path with the kinds observed
// for it across all sampled responses.
func apiPaths(v any, prefix string, out map[string]map[string]bool) {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			p := k
			if prefix != "" {
				p = prefix + "." + k
			}
			if out[p] == nil {
				out[p] = map[string]bool{}
			}
			out[p][kindOfValue(val)] = true
			apiPaths(val, p, out)
		}
	case []any:
		for _, item := range x {
			apiPaths(item, prefix, out)
		}
	}
}

func kindOfValue(v any) string {
	switch v.(type) {
	case map[string]any:
		return "obj"
	case []any:
		return "arr"
	case bool:
		return "bool"
	case float64:
		return "number"
	case string:
		return "string"
	default:
		return "scalar"
	}
}

// scalarKindOf reports the JSON scalar kind a Go type can decode, so that a
// string-typed field is not silently paired with a JSON bool. Amadeus returns
// several booleans as quoted strings ("true"/"false"), which decodes into a
// string field but fails against a bool field.
func scalarKindOf(t reflect.Type) string {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return "number"
	default:
		return ""
	}
}

// compatible reports whether a value of the observed JSON kind can decode into
// a DTO field declared with the given kind.
func compatible(observed, declared string) bool {
	if observed == declared {
		return true
	}
	switch declared {
	case "obj", "arr":
		return false
	case "string":
		// json.Unmarshal only accepts JSON strings into a string field
		return observed == "string"
	case "bool":
		return observed == "bool"
	case "number":
		return observed == "number"
	default:
		// unknown/interface targets accept anything
		return observed != "obj" && observed != "arr"
	}
}

// assertDTOCovers fails if any key present in the sampled responses is absent
// from the DTO, or is declared with an incompatible shape.
func assertDTOCovers(t *testing.T, name string, goType reflect.Type, bodies []string) {
	t.Helper()
	if len(bodies) == 0 {
		t.Skipf("%s: no sample responses collected", name)
	}

	declared := map[string]string{}
	dtoPaths(goType, "", map[reflect.Type]bool{}, declared)

	observed := map[string]map[string]bool{}
	for _, b := range bodies {
		var top map[string]any
		if err := json.Unmarshal([]byte(b), &top); err != nil {
			t.Fatalf("%s: sample is not valid JSON: %v", name, err)
		}
		data, ok := top["data"]
		if !ok {
			continue
		}
		apiPaths(data, "", observed)
	}

	var dropped, mismatched []string
	for p, kinds := range observed {
		want, ok := declared[p]
		if !ok {
			dropped = append(dropped, p)
			continue
		}
		for k := range kinds {
			if !compatible(k, want) {
				mismatched = append(mismatched, p+": API sends "+k+", DTO declares "+want)
			}
		}
	}
	sort.Strings(dropped)
	sort.Strings(mismatched)

	for _, m := range mismatched {
		t.Errorf("%s: shape mismatch %s", name, m)
	}
	for _, d := range dropped {
		t.Errorf("%s: API returns %q but no DTO field decodes it (data is silently dropped)", name, d)
	}
	t.Logf("%s: %d responses, %d API paths, all covered by %d DTO paths",
		name, len(bodies), len(observed), len(declared))
}

func TestListDTOCoversAPIResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	newSDK(t) // authenticate
	client := amadeusIntegration.NewClient(constants.LIST_BASE_URL)

	var bodies []string
	for _, city := range []string{"PAR", "LON", "NYC"} {
		res, err := client.R().SetQueryParams(map[string]string{"cityCode": city}).Get("/by-city")
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertDTOCovers(t, "list/by-city", reflect.TypeOf(responseHotelListDTO.GeneralInfoResponse{}), bodies)
}

func TestContentDTOCoversAPIResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)

	hotels, err := s.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	client := amadeusIntegration.NewClient(constants.CONTENT_BASE_URL)

	n := 25
	if len(hotels) < n {
		n = len(hotels)
	}
	var bodies []string
	for _, h := range hotels[:n] {
		res, err := client.R().SetQueryParams(map[string]string{
			"hotelID": h.HotelId,
			"view":    "FULL",
		}).Get("/reference-data/locations/by-hotel")
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertDTOCovers(t, "content/by-hotel", reflect.TypeOf(responseContentDTO.HotelContentResponse{}), bodies)
}

func TestOffersDTOCoversAPIResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	newSDK(t)
	client := amadeusIntegration.NewClient(constants.OFFERS_BASE_URL)
	checkIn, checkOut := stayDates()

	var bodies []string
	for _, id := range sandboxOfferHotels {
		res, err := client.R().SetQueryParams(map[string]string{
			"hotelIds":     id,
			"checkInDate":  checkIn,
			"checkOutDate": checkOut,
			"adults":       "2",
		}).Get("")
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertDTOCovers(t, "offers/list", reflect.TypeOf([]responseHotelOffersDTO.OffersResponse{}), bodies)
}

func TestOfferByIDDTOCoversAPIResponse(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	var offerIDs []string
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
			for _, offer := range o.Offers {
				offerIDs = append(offerIDs, offer.ID)
			}
		}
	}

	client := amadeusIntegration.NewClient(constants.OFFERS_BASE_URL)
	var bodies []string
	for _, id := range offerIDs {
		res, err := client.R().Get(id)
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertDTOCovers(t, "offers/by-id", reflect.TypeOf(responseHotelOffersDTO.OffersResponse{}), bodies)
}
