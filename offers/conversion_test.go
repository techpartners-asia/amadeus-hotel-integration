package offers_test

import (
	"context"
	"math"
	"net/http"
	"strings"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/internal/amadeustest"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/money"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/offers"
)

// search-converted.json was captured from the live sandbox with currency=MNT.
// It is the case that matters: Amadeus returns prices in the hotel's own
// currency and supplies a rate, so dropping the dictionaries block makes the
// requested currency impossible to display.

func convertedSearch(t *testing.T) offers.HotelOffers {
	t.Helper()
	server := amadeustest.New(t)
	server.Fixture(t, http.MethodGet, searchPath, "search-converted")

	results, err := offers.NewService(server.Client()).Search(context.Background(),
		offers.SearchQuery{HotelIDs: []string{"RTPAREIF"}, Currency: "MNT"})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("no hotels in the fixture")
	}
	return results[0]
}

func TestConversionRatesAreNotDropped(t *testing.T) {
	result := convertedSearch(t)

	if len(result.Rates) == 0 {
		t.Fatal("the dictionaries block was dropped; a requested currency cannot be displayed without it")
	}

	rate, ok := result.Rates["EUR"]
	if !ok {
		t.Fatalf("no EUR rate; got %v", result.Rates)
	}
	if rate.From != "EUR" || rate.To != "MNT" {
		t.Errorf("rate converts %s->%s, want EUR->MNT", rate.From, rate.To)
	}
	// MNT has no minor unit, so a converted price must land on whole tögrög.
	if rate.DecimalPlaces != 0 {
		t.Errorf("DecimalPlaces = %d, want 0 for MNT", rate.DecimalPlaces)
	}
	if rate.RawRate == "" {
		t.Error("the rate as Amadeus sent it was not preserved")
	}

	if target, ok := result.Rates.Target(); !ok || target != "MNT" {
		t.Errorf("Target() = %q, %v; want MNT", target, ok)
	}
}

func TestSixteenDigitRateIsAccepted(t *testing.T) {
	// Amadeus quotes "4099.1909999999998035" - sixteen decimal places, of which
	// all but three are float noise. money.Amount holds nine, and rejecting the
	// conversion over digits nobody means would be useless, so ParseRate rounds.
	result := convertedSearch(t)
	rate := result.Rates["EUR"]

	if got := rate.Rate.String(); got != "4099.191" {
		t.Errorf("rate = %s, want it rounded to 4099.191", got)
	}
	if rate.RawRate != "4099.1909999999998035" {
		t.Errorf("RawRate = %q, want the original preserved", rate.RawRate)
	}
}

func TestPriceIsConvertedAndRoundedToTheTargetCurrency(t *testing.T) {
	result := convertedSearch(t)
	offer := result.Offers[0]

	// Amadeus did NOT convert: the price is still in the hotel's currency.
	if offer.Price.Total.Currency() != "EUR" {
		t.Fatalf("price came back in %s; the fixture should still be in EUR",
			offer.Price.Total.Currency())
	}

	converted, err := result.Rates.Convert(offer.Price.Total)
	if err != nil {
		t.Fatalf("Convert: %v", err)
	}
	if converted.Currency() != "MNT" {
		t.Errorf("converted to %s, want MNT", converted.Currency())
	}

	// No fractional tögrög may survive: it is not a price anyone can pay.
	if converted.Amount().Scale() != 0 {
		t.Errorf("converted amount has %d decimal places, want 0",
			converted.Amount().Scale())
	}

	// Cross-check the arithmetic in floating point. It is the wrong tool for
	// money, which is the point: it is an independent calculation, so it
	// catches an order-of-magnitude slip or a misplaced rounding step without
	// re-implementing the exact algorithm under test. The fixture is
	// re-captured periodically, so the expected figure is derived rather than
	// written down.
	rate := result.Rates[offer.Price.Total.Currency()]
	want := offer.Price.Total.Amount().Float64() * rate.Rate.Float64()
	got := converted.Amount().Float64()

	if diff := math.Abs(got - want); diff > 1 {
		t.Errorf("converted = %s, but %s x %s should be about %.2f (out by %.2f)",
			converted, offer.Price.Total, rate.Rate, want, diff)
	}
	// A conversion into tögrög must move the magnitude by roughly the rate;
	// returning the euro figure unchanged is the failure that matters.
	if got < offer.Price.Total.Amount().Float64()*100 {
		t.Errorf("converted = %s, which is implausibly close to the original %s",
			converted, offer.Price.Total)
	}
}

func TestConvertRefusesTheWrongCurrency(t *testing.T) {
	// A search across several hotels can return more than one currency while
	// Amadeus supplies a rate for only some. Silently returning the
	// unconverted amount would show a guest a price three orders of magnitude
	// out.
	result := convertedSearch(t)

	if _, err := result.Rates.Convert(money.MustParse("100.00", "USD")); err == nil {
		t.Error("converting USD with only an EUR rate should fail")
	}

	rate := result.Rates["EUR"]
	if _, err := rate.Convert(money.MustParse("100.00", "GBP")); err == nil {
		t.Error("an EUR rate should refuse a GBP amount")
	}
}

func TestConvertOrOriginalFallsBackHonestly(t *testing.T) {
	result := convertedSearch(t)

	converted, ok := result.Rates.ConvertOrOriginal(result.Offers[0].Price.Total)
	if !ok || converted.Currency() != "MNT" {
		t.Errorf("EUR should convert: %s, ok=%v", converted, ok)
	}

	// An unconvertible currency comes back untouched, and ok says so - which
	// is what lets the caller label the figure correctly.
	original := money.MustParse("100.00", "USD")
	same, ok := result.Rates.ConvertOrOriginal(original)
	if ok {
		t.Error("USD should not report as converted")
	}
	if same.Currency() != "USD" || same.String() != original.String() {
		t.Errorf("fallback = %s, want the original %s", same, original)
	}
}

func TestNoDictionariesMeansNoRates(t *testing.T) {
	// The ordinary case: a search in the hotel's own currency has no
	// dictionaries block at all, and Rates must be nil rather than an empty
	// map pretending a conversion is available.
	result := search(t)[0]

	if result.Rates != nil {
		t.Errorf("Rates = %v, want nil when no dictionaries were returned", result.Rates)
	}
	if _, ok := result.Rates.Target(); ok {
		t.Error("Target() reported a currency with no rates")
	}
	if _, err := result.Rates.Convert(result.Offers[0].Price.Total); err == nil {
		t.Error("Convert should fail when no rates were returned")
	}
}

func TestZeroRateIsDroppedRatherThanApplied(t *testing.T) {
	// A zero rate would convert every price to nothing. Dropping it makes the
	// failure visible instead.
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, searchPath, http.StatusOK, `{
	  "data":[{"hotel":{"hotelId":"RTPAREIF","name":"T","cityCode":"PAR"},"available":true,
	    "offers":[{"id":"X","checkInDate":"2026-08-10","checkOutDate":"2026-08-13",
	      "room":{"type":"STD"},"price":{"currency":"EUR","total":"100.00"},"guests":{"adults":2}}]}],
	  "dictionaries":{"currencyConversionLookupRates":{
	    "EUR":{"rate":"0","target":"MNT","targetDecimalPlaces":0},
	    "GBP":{"rate":"not-a-number","target":"MNT","targetDecimalPlaces":0}}}}`)

	results, err := offers.NewService(server.Client()).Search(context.Background(),
		offers.SearchQuery{HotelIDs: []string{"RTPAREIF"}})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}

	if len(results[0].Rates) != 0 {
		t.Errorf("Rates = %v, want both the zero and the malformed rate dropped", results[0].Rates)
	}
}

func TestDisplayConvertsWithoutTouchingTheOriginal(t *testing.T) {
	result := convertedSearch(t)
	offer := result.Offers[0]

	shown := result.DisplayTotal(offer)

	if !shown.Converted {
		t.Error("the price should have converted")
	}
	if shown.Currency() != "MNT" {
		t.Errorf("display currency = %s, want MNT", shown.Currency())
	}
	if strings.Contains(shown.String(), ".") {
		t.Errorf("display = %q, want whole tögrög", shown.String())
	}

	// The crucial guarantee: Display did not mutate the offer. The price the
	// booking is charged is still the euro figure Amadeus quoted.
	if offer.Price.Total.Currency() != "EUR" {
		t.Errorf("offer.Price was changed to %s; it must stay the source of truth",
			offer.Price.Total.Currency())
	}
	if shown.Original.Currency() != "EUR" || shown.Original.String() != offer.Price.Total.String() {
		t.Errorf("Original = %s, want the untouched %s", shown.Original, offer.Price.Total)
	}
}

func TestDisplayFallsBackToOriginalCurrency(t *testing.T) {
	// With no rate for the amount's currency, Display shows the original and
	// says so, so the caller never renders a bare number with the wrong label.
	result := convertedSearch(t)

	shown := result.Display(money.MustParse("100.00", "USD"))
	if shown.Converted {
		t.Error("USD has no rate here and must not report as converted")
	}
	if shown.Currency() != "USD" || shown.String() != "100 USD" {
		t.Errorf("display = %s, want the original 100 USD", shown)
	}
}

func TestDisplayWorksForAnyMoneyOnTheOffer(t *testing.T) {
	// Not just the total: base, taxes and per-night rates convert the same way.
	result := convertedSearch(t)
	offer := result.Offers[0]

	perNight, _, ok := offer.Price.PerNight(offer.Stay.Nights())
	if !ok {
		t.Skip("offer has no divisible total")
	}

	shown := result.Display(perNight)
	if !shown.Converted || shown.Currency() != "MNT" {
		t.Errorf("per-night display = %s (converted=%v), want MNT", shown, shown.Converted)
	}
}
