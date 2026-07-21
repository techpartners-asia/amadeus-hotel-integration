package tests

import (
	"strings"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

// amadeusAmenityCodes is the `amenities` enum for the Hotel List API (by-city
// and by-geocode), transcribed here so the package's constants are checked
// against an independent list rather than against themselves. Editing
// reference/amenity.go without editing this list fails
// TestAmenityCatalogMatchesAmadeus.
//
// This is Amadeus' published list with two corrections. The docs spell them
// "BABY-SITTING" and "BAR or LOUNGE"; the live API rejects both with 7211
// INVALID FACILITY CODE and accepts BABY_SITTING and BAR_LOUNGE instead. The
// values below are the ones that actually work, confirmed by
// TestAmenityCodesAcceptedByAmadeus. Do not "correct" them back to the docs.
//
// The remaining odd spellings are genuine: WI-FI_IN_ROOM carries a hyphen where
// the rest use underscores, and GUARDED_PARKG and SERV_SPEC_MENU are
// abbreviated. All three are accepted as written.
var amadeusAmenityCodes = []string{
	"SWIMMING_POOL", "SPA", "FITNESS_CENTER", "AIR_CONDITIONING", "RESTAURANT",
	"PARKING", "PETS_ALLOWED", "AIRPORT_SHUTTLE", "BUSINESS_CENTER",
	"DISABLED_FACILITIES", "WIFI", "MEETING_ROOMS", "NO_KID_ALLOWED", "TENNIS",
	"GOLF", "KITCHEN", "ANIMAL_WATCHING", "BABY_SITTING", "BEACH", "CASINO",
	"JACUZZI", "SAUNA", "SOLARIUM", "MASSAGE", "VALET_PARKING", "BAR_LOUNGE",
	"KIDS_WELCOME", "NO_PORN_FILMS", "MINIBAR", "TELEVISION", "WI-FI_IN_ROOM",
	"ROOM_SERVICE", "GUARDED_PARKG", "SERV_SPEC_MENU",
}

// TestAmenityCatalogMatchesAmadeus is the reason this file exists: a reference
// list that silently drops a code is worse than no list at all, because callers
// trust AllAmenities to be exhaustive when building a filter UI. A missing code
// makes an amenity unfilterable with no error anywhere.
func TestAmenityCatalogMatchesAmadeus(t *testing.T) {
	got := searchcriteria.AllAmenities()

	if len(got) != len(amadeusAmenityCodes) {
		t.Errorf("AllAmenities returned %d codes, Amadeus documents %d", len(got), len(amadeusAmenityCodes))
	}

	declared := map[string]bool{}
	for _, a := range got {
		declared[string(a)] = true
	}

	for _, want := range amadeusAmenityCodes {
		if !declared[want] {
			t.Errorf("amenity %q is documented by Amadeus but has no constant", want)
		}
		delete(declared, want)
	}
	for extra := range declared {
		t.Errorf("amenity %q has a constant but is not documented by Amadeus", extra)
	}
}

// TestAmenityOrderIsStable pins the declaration order, since AllAmenities feeds
// filter UIs where reordering on every build would be visible to users.
func TestAmenityOrderIsStable(t *testing.T) {
	got := searchcriteria.AllAmenities()
	if len(got) != len(amadeusAmenityCodes) {
		t.Fatalf("length mismatch: %d vs %d", len(got), len(amadeusAmenityCodes))
	}
	for i, want := range amadeusAmenityCodes {
		if string(got[i]) != want {
			t.Errorf("position %d: got %q, want %q", i, got[i], want)
		}
	}
}

// TestAllListsAreNonEmptyAndLabelled checks the shared contract of every
// reference type: every enumerated code validates, and every one has a label,
// because a blank label renders as an empty row in a filter UI.
func TestAllListsAreNonEmptyAndLabelled(t *testing.T) {
	check := func(name string, codes []string, valid []bool, labels []string) {
		t.Helper()
		if len(codes) == 0 {
			t.Errorf("%s: list is empty", name)
		}
		seen := map[string]bool{}
		for i, c := range codes {
			if c == "" {
				t.Errorf("%s[%d]: empty code", name, i)
			}
			if seen[c] {
				t.Errorf("%s: duplicate code %q", name, c)
			}
			seen[c] = true
			if !valid[i] {
				t.Errorf("%s: %q is in the list but IsValid says otherwise", name, c)
			}
			if strings.TrimSpace(labels[i]) == "" {
				t.Errorf("%s: %q has no label", name, c)
			}
		}
	}

	amenities := searchcriteria.AllAmenities()
	ac, av, al := unzip(amenities, func(a searchcriteria.Amenity) (string, bool, string) {
		return string(a), a.IsValid(), a.Label()
	})
	check("Amenity", ac, av, al)

	ratings := searchcriteria.AllRatings()
	rc, rv, rl := unzip(ratings, func(r searchcriteria.Rating) (string, bool, string) {
		return string(r), r.IsValid(), r.Label()
	})
	check("Rating", rc, rv, rl)

	sources := searchcriteria.AllHotelSources()
	sc, sv, sl := unzip(sources, func(h searchcriteria.HotelSource) (string, bool, string) {
		return string(h), h.IsValid(), h.Label()
	})
	check("HotelSource", sc, sv, sl)

	units := searchcriteria.AllRadiusUnits()
	uc, uv, ul := unzip(units, func(u searchcriteria.RadiusUnit) (string, bool, string) {
		return string(u), u.IsValid(), u.Label()
	})
	check("RadiusUnit", uc, uv, ul)

	boards := searchcriteria.AllBoardTypes()
	bc, bv, bl := unzip(boards, func(b searchcriteria.BoardType) (string, bool, string) {
		return string(b), b.IsValid(), b.Label()
	})
	check("BoardType", bc, bv, bl)

	policies := searchcriteria.AllPaymentPolicies()
	pc, pv, pl := unzip(policies, func(p searchcriteria.PaymentPolicy) (string, bool, string) {
		return string(p), p.IsValid(), p.Label()
	})
	check("PaymentPolicy", pc, pv, pl)

	views := searchcriteria.AllContentViews()
	vc, vv, vl := unzip(views, func(v searchcriteria.ContentView) (string, bool, string) {
		return string(v), v.IsValid(), v.Label()
	})
	check("ContentView", vc, vv, vl)

	rates := searchcriteria.AllRateCodes()
	tc, tv, tl := unzip(rates, func(c searchcriteria.RateCode) (string, bool, string) {
		return string(c), c.IsValid(), c.Label()
	})
	check("RateCode", tc, tv, tl)
}

func unzip[T any](items []T, f func(T) (string, bool, string)) ([]string, []bool, []string) {
	codes := make([]string, len(items))
	valid := make([]bool, len(items))
	labels := make([]string, len(items))
	for i, it := range items {
		codes[i], valid[i], labels[i] = f(it)
	}
	return codes, valid, labels
}

// TestIsValidRejectsUnknown guards against IsValid being wired to something
// that always returns true, which would make it useless as an input check.
func TestIsValidRejectsUnknown(t *testing.T) {
	if searchcriteria.Amenity("NOT_AN_AMENITY").IsValid() {
		t.Error("Amenity.IsValid accepted an unknown code")
	}
	if searchcriteria.Amenity("").IsValid() {
		t.Error("Amenity.IsValid accepted an empty code")
	}
	if searchcriteria.Rating("6").IsValid() {
		t.Error("Rating.IsValid accepted 6")
	}
	if searchcriteria.BoardType("BRUNCH").IsValid() {
		t.Error("BoardType.IsValid accepted an unknown code")
	}
	if searchcriteria.Amenity("NOT_AN_AMENITY").Label() != "" {
		t.Error("Label should be empty for an unknown code")
	}
	// Amadeus codes are case-sensitive; a lowercase code is a 400, not a match.
	if searchcriteria.Amenity("wifi").IsValid() {
		t.Error("Amenity.IsValid should be case-sensitive")
	}
}

// TestRateCodeValidatesShapeNotMembership documents the deliberate difference:
// corporate rate codes are negotiated per account, so IsValid can only check
// the 3-character shape. Rejecting unlisted codes would break corporate rates.
func TestRateCodeValidatesShapeNotMembership(t *testing.T) {
	if !searchcriteria.RateCode("IBM").IsValid() {
		t.Error("a well-formed corporate code must be valid even though it is not listed")
	}
	if searchcriteria.RateCode("TOOLONG").IsValid() {
		t.Error("a 7-character code must be rejected")
	}
	if searchcriteria.RateCode("ab1").IsValid() {
		t.Error("a lowercase code must be rejected")
	}
	if searchcriteria.RateCode("").IsValid() {
		t.Error("an empty code must be rejected")
	}
}

// TestJoinRendersCommaSeparated covers the helper every ToQueryParams relies on,
// including BAR_LOUNGE, whose code Amadeus documents incorrectly.
func TestJoinRendersCommaSeparated(t *testing.T) {
	got := searchcriteria.Join([]searchcriteria.Amenity{
		searchcriteria.AmenityWifi,
		searchcriteria.AmenityBarOrLounge,
		searchcriteria.AmenitySwimmingPool,
	})
	want := "WIFI,BAR_LOUNGE,SWIMMING_POOL"
	if got != want {
		t.Errorf("Join = %q, want %q", got, want)
	}

	if got := searchcriteria.Join([]searchcriteria.Rating{}); got != "" {
		t.Errorf("Join of empty slice = %q, want empty so the param is omitted", got)
	}
}

// TestAllReturnsACopy guards the catalogs against mutation through the returned
// slice, which would corrupt the lists for every later caller in the process.
func TestAllReturnsACopy(t *testing.T) {
	first := searchcriteria.AllAmenities()
	original := first[0]
	first[0] = searchcriteria.Amenity("MUTATED")

	if second := searchcriteria.AllAmenities(); second[0] != original {
		t.Errorf("mutating the returned slice changed the catalog: got %q, want %q", second[0], original)
	}
}

// TestCatalogMatchesPackageFunctions verifies the SDK-callable accessor returns
// the same data as the package-level functions, so the two entry points cannot
// disagree.
func TestCatalogMatchesPackageFunctions(t *testing.T) {
	c := searchcriteria.NewCatalog()

	if len(c.Amenities()) != len(searchcriteria.AllAmenities()) {
		t.Error("Catalog.Amenities disagrees with AllAmenities")
	}
	if len(c.Ratings()) != len(searchcriteria.AllRatings()) {
		t.Error("Catalog.Ratings disagrees with AllRatings")
	}
	if len(c.HotelSources()) != len(searchcriteria.AllHotelSources()) {
		t.Error("Catalog.HotelSources disagrees with AllHotelSources")
	}
	if len(c.RadiusUnits()) != len(searchcriteria.AllRadiusUnits()) {
		t.Error("Catalog.RadiusUnits disagrees with AllRadiusUnits")
	}
	if len(c.BoardTypes()) != len(searchcriteria.AllBoardTypes()) {
		t.Error("Catalog.BoardTypes disagrees with AllBoardTypes")
	}
	if len(c.PaymentPolicies()) != len(searchcriteria.AllPaymentPolicies()) {
		t.Error("Catalog.PaymentPolicies disagrees with AllPaymentPolicies")
	}
	if len(c.ContentViews()) != len(searchcriteria.AllContentViews()) {
		t.Error("Catalog.ContentViews disagrees with AllContentViews")
	}
	if len(c.RateCodes()) != len(searchcriteria.AllRateCodes()) {
		t.Error("Catalog.RateCodes disagrees with AllRateCodes")
	}
}
