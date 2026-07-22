package booking

import (
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Price is what a booking was agreed at.
//
// The booking context keeps its own price type rather than sharing the offers
// one. They look alike but mean different things: an offers.Price is a live
// quote that expires, while this is a record of what was agreed and does not
// change. Sharing the type would invite treating one as the other.
type Price struct {
	Currency money.Currency
	// Base is the room rate before taxes.
	Base money.Money
	// Total is the agreed price.
	Total money.Money
	// SellingTotal is Total plus markups and fees.
	SellingTotal money.Money
	// Taxes are the individual tax lines.
	Taxes []Tax
	// Markups are the markups applied.
	Markups []money.Money
	// Variations breaks the stay down by night, when the rate varied.
	Variations *Variations
}

// TaxesTotal sums the tax lines not already included in Base.
func (p Price) TaxesTotal() (money.Money, error) {
	var amounts []money.Money
	for _, tax := range p.Taxes {
		if !tax.Included {
			amounts = append(amounts, tax.Amount)
		}
	}
	return money.Sum(amounts...)
}

// Tax is one tax line on a booked price.
type Tax struct {
	Amount      money.Money
	Code        string
	Description string
	Percentage  string
	// Included reports that the tax is already part of Base.
	Included bool
	// PricingFrequency and PricingMode describe how the tax is assessed.
	PricingFrequency string
	PricingMode      string
	// Applicable is the date range the tax covers, when limited.
	Applicable *DateRange
}

// DateRange is a start/end pair.
//
// Note that the booking API's tax block carries no collectionPoint, unlike the
// search API's. There is deliberately no PayableAtProperty here: the data to
// compute it does not arrive, and a method that always returned zero would read
// as "nothing due on arrival" rather than "not known".
type DateRange struct {
	Start datetime.Date
	End   datetime.Date
}

// Variations breaks a booked stay's price down by period.
type Variations struct {
	// Average is the per-night average across the stay.
	Average PricePeriod
	// Changes are the periods whose rate differed.
	Changes []PricePeriod
}

// PricePeriod is a price covering a date range.
type PricePeriod struct {
	Start        datetime.Date
	End          datetime.Date
	Currency     money.Currency
	Base         money.Money
	Total        money.Money
	SellingTotal money.Money
}

// Policies are the terms a booking was made under.
//
// Cancellation is the field that matters after the fact: it determines whether
// cancelling costs the guest anything, and by when.
type Policies struct {
	// PaymentType is how the booking is paid, e.g. "guarantee", "prepay".
	PaymentType string
	// Cancellation holds the cancellation terms.
	Cancellation []CancellationPolicy
	// Refundable is Amadeus's refundability statement, when it sent one.
	Refundable *RefundPolicy
	// Deposit and Prepay are amounts due before arrival.
	Deposit *AmountDuePolicy
	Prepay  *AmountDuePolicy
	// Guarantee is the card-guarantee requirement.
	Guarantee *GuaranteePolicy
	// HoldTime is when an unguaranteed booking is released.
	HoldTime *time.Time
	// CheckInOut is the property's arrival and departure timing.
	CheckInOut *CheckInOutPolicy
	// LengthOfStay constrains the number of nights.
	LengthOfStay *LengthOfStayPolicy
	// Details are additional policy texts.
	Details []media.Text
}

// RefundStatus is how refundable a booking is.
type RefundStatus string

const (
	RefundNonRefundable RefundStatus = "NON_REFUNDABLE"
	RefundRefundable    RefundStatus = "REFUNDABLE"
	RefundUpToDeadline  RefundStatus = "REFUNDABLE_UP_TO_DEADLINE"
	RefundUnknown       RefundStatus = "UNKNOWN"
)

// RefundPolicy is Amadeus's refundability statement.
type RefundPolicy struct {
	Status RefundStatus
}

// CanCancelFreeOfCharge reports whether the booking can still be cancelled at
// no cost as of now, and whether Amadeus said clearly enough to rely on it.
//
// It answers the question a guest actually asks - "can I still cancel?" - which
// depends on the deadline having not yet passed, not merely on one existing.
// When certain is false, do not tell the guest cancellation is free.
func (p Policies) CanCancelFreeOfCharge(now time.Time) (free bool, certain bool) {
	if p.Refundable != nil && p.Refundable.Status == RefundNonRefundable {
		return false, true
	}

	if len(p.Cancellation) == 0 {
		// A refundable statement with no terms says it is refundable but not
		// until when; that is not enough to promise a free cancellation.
		if p.Refundable != nil && p.Refundable.Status == RefundRefundable {
			return true, true
		}
		return false, false
	}

	for _, policy := range p.Cancellation {
		if !policy.IsFree() {
			continue
		}
		// No deadline on a free policy means free throughout.
		if policy.Deadline == nil {
			return true, true
		}
		if now.Before(*policy.Deadline) {
			return true, true
		}
	}
	return false, true
}

// CancellationPolicy is one set of cancellation terms on a booking.
type CancellationPolicy struct {
	// Amount is the fee to cancel.
	Amount money.Money
	// Percentage is the fee as a proportion.
	Percentage string
	// NumberOfNights is the fee expressed as nights charged.
	NumberOfNights int
	// Deadline is when the policy takes effect.
	Deadline *time.Time
	Type     string
	// PolicyType is what triggers it: CANCELLATION, EARLY_CHECKOUT, NO_SHOW.
	PolicyType  string
	Description *media.Text
}

// IsFree reports whether cancelling under this policy costs nothing, checking
// all three forms Amadeus can express a fee in.
func (c CancellationPolicy) IsFree() bool {
	return c.Amount.Amount().IsZero() &&
		c.NumberOfNights == 0 &&
		(c.Percentage == "" || c.Percentage == "0")
}

// AmountDuePolicy is a sum payable before arrival, with its deadline.
type AmountDuePolicy struct {
	Amount      money.Money
	Deadline    *time.Time
	Description *media.Text
}

// GuaranteePolicy is a card-guarantee requirement.
type GuaranteePolicy struct {
	Description      *media.Text
	AcceptedPayments *AcceptedPayments
}

// AcceptedPayments lists how a policy may be settled.
type AcceptedPayments struct {
	CreditCards []string
	Methods     []string
}

// CheckInOutPolicy is the property's arrival and departure timing.
type CheckInOutPolicy struct {
	CheckIn             string
	CheckOut            string
	CheckInDescription  *media.Text
	CheckOutDescription *media.Text
}

// LengthOfStayPolicy constrains the number of nights.
type LengthOfStayPolicy struct {
	Minimum int
	Maximum int
}
