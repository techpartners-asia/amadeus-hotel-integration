package offers

import (
	"github.com/techpartners-asia/amadeus-hotel-integration/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Price is what an offer costs.
//
// Amadeus sends every amount as a decimal string with the currency held once on
// the parent object. Here each amount is a money.Money carrying its own
// currency, so a total cannot be added to a figure in another currency by
// accident, and no caller has to parse a price string.
//
// The relationship Amadeus documents is Total = Base + taxes, and
// SellingTotal = Total + markups. Neither is recomputed here: the SDK reports
// what Amadeus sent, because the API is authoritative on what will be charged.
type Price struct {
	// Currency is the currency of this price. Every Money below carries it too.
	Currency money.Currency
	// Base is the room rate before taxes.
	Base money.Money
	// Total is what Amadeus quotes as the price of the stay.
	Total money.Money
	// SellingTotal is Total plus markups and fees, less discounts. It is zero
	// when Amadeus sends no markups, in which case Total is the selling price.
	SellingTotal money.Money
	// RateParityTotal is the total before rate-parity adjustment, when sent.
	RateParityTotal money.Money

	// Taxes are the individual tax lines. Check Tax.Included before adding
	// them to Base: some are already in it, and some are payable at the
	// property rather than at booking.
	Taxes []Tax
	// Markups are the markup amounts applied to reach SellingTotal.
	Markups []money.Money
	// Commissions is the commission breakdown, when the rate pays one.
	Commissions []CommissionValue

	// Variations breaks a multi-night stay into its nightly rates. A stay
	// whose price differs by night has one Change per period.
	Variations Variations
}

// TaxesTotal sums the tax lines that are not already included in Base.
//
// It returns an error only when the lines disagree on currency, which Amadeus
// does occasionally produce by quoting a local tax beside a converted rate.
func (p Price) TaxesTotal() (money.Money, error) {
	var amounts []money.Money
	for _, tax := range p.Taxes {
		if tax.Included {
			continue
		}
		amounts = append(amounts, tax.Amount)
	}
	return money.Sum(amounts...)
}

// PayableAtProperty sums the taxes Amadeus marks as collected at the hotel
// rather than at booking. It is the figure to show a guest as "payable on
// arrival", and it is frequently non-zero for city and tourist taxes.
func (p Price) PayableAtProperty() (money.Money, error) {
	var amounts []money.Money
	for _, tax := range p.Taxes {
		if tax.CollectedAtProperty() {
			amounts = append(amounts, tax.Amount)
		}
	}
	return money.Sum(amounts...)
}

// PerNight returns Total divided across the nights of a stay, for display as a
// nightly rate, together with the remainder that would not divide evenly.
//
// When the stay's price varies by night, prefer Variations.Average: it is
// Amadeus's own figure and reflects the actual nightly rates, where this is a
// flat division.
func (p Price) PerNight(nights int) (perNight money.Money, remainder money.Money, ok bool) {
	if nights <= 0 || p.Total.Amount().IsZero() {
		return money.Money{}, money.Money{}, false
	}
	return p.Total.Split(nights)
}

// Tax is one tax line on a price.
type Tax struct {
	// Amount is the tax charged.
	Amount money.Money
	// Code identifies the tax, e.g. "TOURISM".
	Code string
	// Description explains it in prose.
	Description string
	// Percentage is the rate as Amadeus expressed it, when it is proportional.
	Percentage string
	// Included reports that this tax is already part of Price.Base. Adding an
	// included tax to the base double-charges the guest in your display.
	Included bool
	// PaidInLoyaltyRewards reports a tax settled with loyalty points.
	PaidInLoyaltyRewards bool
	// CollectionPoint is where the tax is collected: "AT_BOOKING_TIME" or
	// "AT_HOTEL_PROPERTY".
	CollectionPoint string
	// PricingFrequency and PricingMode describe how the tax is assessed.
	PricingFrequency string
	PricingMode      string
	// Applicable is the date range the tax applies over, when limited.
	Applicable *DateRange
}

// collectionAtProperty is the CollectionPoint value meaning the guest pays on
// arrival rather than at booking.
const collectionAtProperty = "AT_HOTEL_PROPERTY"

// CollectedAtProperty reports whether the guest pays this tax at the hotel.
func (t Tax) CollectedAtProperty() bool { return t.CollectionPoint == collectionAtProperty }

// DateRange is a start/end pair, used for taxes and policies that apply over a
// limited period.
type DateRange struct {
	Start datetime.Date
	End   datetime.Date
}

// Variations breaks a stay's price down by period.
type Variations struct {
	// Average is Amadeus's own per-night average, and is the figure to display
	// as a nightly rate.
	Average PricePeriod
	// Changes are the periods whose rate differs, in the order Amadeus sent
	// them. A stay priced the same every night has none.
	Changes []PricePeriod
}

// PricePeriod is a price covering a date range, or the average across the stay
// when it is Variations.Average.
type PricePeriod struct {
	// Start and End bound the period. Both are zero on Variations.Average,
	// which covers the whole stay.
	Start datetime.Date
	End   datetime.Date

	Currency     money.Currency
	Base         money.Money
	Total        money.Money
	SellingTotal money.Money
	Markups      []money.Money
}

// IsZero reports whether the period carries no price at all.
func (p PricePeriod) IsZero() bool {
	return p.Base.Amount().IsZero() && p.Total.Amount().IsZero()
}

// CommissionValue is one line of an offer's commission breakdown.
type CommissionValue struct {
	// Amount is the commission paid, when expressed as a sum.
	Amount money.Money
	// Percentage is the commission rate, when expressed as a proportion.
	Percentage float64
	// DecimalPlaces is the precision Amadeus quoted the amount to, preserved
	// because it is the only signal of the currency's minor unit on this block.
	DecimalPlaces int
	// PriceType and IssueCurrencyType are Amadeus's own classifiers for the
	// line.
	PriceType         string
	IssueCurrencyType string
}
