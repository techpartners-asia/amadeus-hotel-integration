package tests

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	responseContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/response"
	responseHotelOffersDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"
)

// The coverage checks in dto_fidelity_test.go compare key paths: they prove a
// DTO field exists for every key the API sends. They do not prove the value
// survives the decode. A field can be present, correctly shaped, and still lose
// its contents (a custom UnmarshalJSON that swallows errors, a numeric string
// truncated into an int, an enum type that maps to the empty value).
//
// These checks close that gap by round-tripping: raw JSON -> DTO -> JSON, then
// diffing every leaf value positionally. Anything the DTO cannot hold shows up
// as a missing or altered leaf.

// leaves flattens decoded JSON into path -> value, indexing arrays positionally
// so elements are compared like-for-like rather than as an unordered set.
func leaves(v any, prefix string, out map[string]any) {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			p := k
			if prefix != "" {
				p = prefix + "." + k
			}
			leaves(val, p, out)
		}
	case []any:
		for i, item := range x {
			leaves(item, fmt.Sprintf("%s[%d]", prefix, i), out)
		}
	default:
		out[prefix] = v
	}
}

// sameValue compares two JSON leaves, treating numbers by value rather than by
// formatting so 3 and 3.0 do not read as a difference.
func sameValue(a, b any) bool {
	af, aok := toFloat(a)
	bf, bok := toFloat(b)
	if aok && bok {
		return af == bf
	}
	return a == b
}

func toFloat(v any) (float64, bool) {
	switch x := v.(type) {
	case float64:
		return x, true
	case json.Number:
		f, err := x.Float64()
		return f, err == nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		return f, err == nil
	}
	return 0, false
}

// isZeroish reports whether a leaf holds a JSON zero value. Such a leaf can be
// dropped on re-marshal purely because its DTO field carries `omitempty`, which
// is a marshal-side artifact and not a parsing failure.
func isZeroish(v any) bool {
	switch x := v.(type) {
	case nil:
		return true
	case bool:
		return !x
	case float64:
		return x == 0
	case string:
		return x == ""
	}
	return false
}

// assertRoundTrips decodes each body into T through the SDK's own response
// envelope, re-marshals it, and reports every leaf that did not survive.
func assertRoundTrips[T any](t *testing.T, name string, bodies []string) {
	t.Helper()
	if len(bodies) == 0 {
		t.Skipf("%s: no sample responses collected", name)
	}

	var lost, changed, elided []string
	totalLeaves := 0

	for i, body := range bodies {
		// Decode exactly as the usecases do, so a failure here is a real SDK bug.
		var envelope sharedResponseDTO.BaseResponse[T]
		if err := json.Unmarshal([]byte(body), &envelope); err != nil {
			t.Fatalf("%s: sample %d does not decode into the DTO: %v", name, i, err)
		}

		round, err := json.Marshal(envelope.Data)
		if err != nil {
			t.Fatalf("%s: sample %d does not re-marshal: %v", name, i, err)
		}
		var roundAny any
		if err := json.Unmarshal(round, &roundAny); err != nil {
			t.Fatalf("%s: sample %d re-marshal is not valid JSON: %v", name, i, err)
		}

		var top map[string]any
		if err := json.Unmarshal([]byte(body), &top); err != nil {
			t.Fatalf("%s: sample %d is not valid JSON: %v", name, i, err)
		}
		data, ok := top["data"]
		if !ok {
			continue
		}

		before := map[string]any{}
		leaves(data, "", before)
		after := map[string]any{}
		leaves(roundAny, "", after)
		totalLeaves += len(before)

		for p, want := range before {
			got, present := after[p]
			switch {
			case !present && isZeroish(want):
				elided = append(elided, fmt.Sprintf("%s = %v", p, want))
			case !present:
				lost = append(lost, fmt.Sprintf("sample %d: %s = %v", i, p, want))
			case !sameValue(want, got):
				changed = append(changed, fmt.Sprintf("sample %d: %s: sent %v, decoded %v", i, p, want, got))
			}
		}
	}

	sort.Strings(lost)
	sort.Strings(changed)

	for _, c := range changed {
		t.Errorf("%s: value altered by decode: %s", name, c)
	}
	for _, l := range lost {
		t.Errorf("%s: value lost by decode: %s", name, l)
	}

	t.Logf("%s: %d responses, %d leaf values, %d lost, %d altered, %d omitempty-elided zero values",
		name, len(bodies), totalLeaves, len(lost), len(changed), len(elided))

	// Elided leaves are not failures, but list the distinct ones so a genuine
	// loss cannot hide behind the "harmless zero value" classification.
	if len(elided) > 0 {
		seen := map[string]bool{}
		var distinct []string
		for _, e := range elided {
			if !seen[e] {
				seen[e] = true
				distinct = append(distinct, e)
			}
		}
		sort.Strings(distinct)
		t.Logf("%s: distinct omitempty-elided zero values: %v", name, distinct)
	}
}

func TestOffersListDTORoundTrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	newSDK(t)
	client := amadeusIntegration.NewClient(constants.OFFERS_BASE_URL)
	checkIn, checkOut := stayDates()

	var bodies []string
	for _, id := range sandboxOfferHotels {
		req := allRatesFor(id, checkIn, checkOut)
		res, err := client.R().SetQueryParams(req.ToQueryParams()).Get("")
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertRoundTrips[[]responseHotelOffersDTO.OffersResponse](t, "offers/list", bodies)
}

func TestOfferByIDDTORoundTrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	offerIDs := sampleOfferIDs(s, checkIn, checkOut)

	client := amadeusIntegration.NewClient(constants.OFFERS_BASE_URL)
	var bodies []string
	for _, id := range offerIDs {
		res, err := client.R().Get(id)
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertRoundTrips[responseHotelOffersDTO.OffersResponse](t, "offers/by-id", bodies)
}

func TestContentDTORoundTrip(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)

	hotels := searchHotels(t, s, "PAR", bookableChain, 25)
	client := amadeusIntegration.NewClient(constants.CONTENT_BASE_URL)

	var bodies []string
	for _, h := range hotels {
		res, err := client.R().SetQueryParams(map[string]string{
			"hotelID": h.HotelId,
			"view":    "FULL",
		}).Get("/reference-data/locations/by-hotel")
		if err == nil && res.StatusCode() == 200 {
			bodies = append(bodies, res.String())
		}
	}
	assertRoundTrips[responseContentDTO.HotelContentResponse](t, "content/by-hotel", bodies)
}
