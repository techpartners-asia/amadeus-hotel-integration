package offers

import (
	"fmt"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/money"
)

// ConversionRates are the exchange rates a search returned, keyed by the
// currency being converted from.
//
// They exist because asking for a currency does not get you prices in it.
// Setting SearchQuery.Currency to "MNT" returns offers still priced in the
// hotel's own currency, plus the rate to convert them:
//
//	{"currency": "EUR", "total": "1410.61"}
//	dictionaries.currencyConversionLookupRates.EUR = {rate: "4099.19...", target: "MNT"}
//
// Without applying the rate there is no way to show the price the guest asked
// for. Convert does it:
//
//	results, _ := client.Offers.Search(ctx, offers.SearchQuery{
//	    HotelIDs: ids, Currency: "MNT",
//	})
//	local, err := results[0].Rates.Convert(offer.Price.Total)
type ConversionRates map[money.Currency]ConversionRate

// ConversionRate converts one currency into another.
type ConversionRate struct {
	// From is the currency being converted out of.
	From money.Currency
	// To is the currency being converted into.
	To money.Currency
	// Rate is the multiplier. Amadeus quotes it to sixteen decimal places, of
	// which only the first few are meaningful; it is held here rounded to the
	// precision money.Amount can represent exactly.
	Rate money.Amount
	// DecimalPlaces is the minor-unit precision of the target currency: 2 for
	// EUR, 0 for MNT and JPY. A converted amount is rounded to it, because a
	// price of 5,781,940.7 tögrög is not a price anyone can pay.
	DecimalPlaces int
	// RawRate is the rate exactly as Amadeus sent it, kept so nothing is lost
	// to the rounding above.
	RawRate string
}

// Convert returns amount expressed in the target currency, rounded to the
// target's minor unit.
//
// It fails when amount is in a currency this rate does not convert from, which
// is a real possibility: a search across several hotels can return prices in
// more than one currency while Amadeus supplies a rate for only some of them.
func (r ConversionRate) Convert(amount money.Money) (money.Money, error) {
	if amount.Currency() != r.From {
		return money.Money{}, fmt.Errorf(
			"offers: rate converts from %s, but the amount is in %s", r.From, amount.Currency())
	}

	product, ok := amount.Amount().MulAmount(r.Rate)
	if !ok {
		return money.Money{}, fmt.Errorf(
			"offers: converting %s at %s overflows", amount, r.Rate)
	}

	return money.New(product.Round(r.DecimalPlaces), r.To), nil
}

// Convert finds the rate for the amount's currency and applies it.
//
// It reports an error when the search returned no rate for that currency,
// rather than returning the unconverted amount, since silently handing back
// euros to a caller who asked for tögrög is how a guest is shown a price three
// orders of magnitude out.
func (rates ConversionRates) Convert(amount money.Money) (money.Money, error) {
	rate, ok := rates[amount.Currency()]
	if !ok {
		return money.Money{}, fmt.Errorf(
			"offers: no conversion rate for %s in this response", amount.Currency())
	}
	return rate.Convert(amount)
}

// ConvertOrOriginal returns the converted amount, or the original unchanged
// when no rate applies.
//
// Use it for display where showing the price in its own currency is an
// acceptable fallback. The bool says which happened, so the currency label can
// be right either way.
func (rates ConversionRates) ConvertOrOriginal(amount money.Money) (money.Money, bool) {
	converted, err := rates.Convert(amount)
	if err != nil {
		return amount, false
	}
	return converted, true
}

// Target returns the currency these rates convert into, and false when there
// are none. Every rate in one response shares a target: the currency asked for.
func (rates ConversionRates) Target() (money.Currency, bool) {
	for _, rate := range rates {
		return rate.To, true
	}
	return "", false
}
