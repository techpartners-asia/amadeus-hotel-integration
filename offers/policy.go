package offers

import (
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Policies are the terms attached to an offer: what it costs to cancel, when
// payment is due, and how long a guest may stay.
//
// Read Cancellation before presenting an offer as refundable. Amadeus expresses
// refundability three different ways - a Refundable block, a cancellation
// deadline, and a cancellation fee - and they do not always agree. Refundable
// below reconciles them conservatively.
type Policies struct {
	// PaymentType is how the offer is paid for, e.g. "guarantee", "deposit".
	PaymentType string

	// Cancellation is the set of cancellation terms. Amadeus sends both a
	// single "cancellation" object and a "cancellations" array, often with the
	// same content; both are merged here, deduplicated, so callers read one
	// list rather than guessing which field was populated.
	Cancellation []CancellationPolicy

	// Refundable is Amadeus's own refundability statement, when it sends one.
	Refundable *RefundPolicy

	// Deposit and Prepay are amounts due before arrival.
	Deposit *PaymentPolicy
	Prepay  *PaymentPolicy
	// Guarantee is the card-guarantee requirement, which takes no payment but
	// holds a card against no-show.
	Guarantee *GuaranteePolicy

	// HoldTime is when an unguaranteed booking is released.
	HoldTime *time.Time

	// CheckInOut is the property's check-in and check-out timing.
	CheckInOut *CheckInOutPolicy

	// LengthOfStay constrains how many nights may be booked.
	LengthOfStay *LengthOfStayPolicy

	// Details are additional policy texts Amadeus supplies without structure.
	Details []media.Text
}

// RefundStatus is how refundable an offer is.
type RefundStatus string

const (
	// RefundNonRefundable means cancelling forfeits the payment entirely.
	RefundNonRefundable RefundStatus = "NON_REFUNDABLE"
	// RefundRefundable means the offer can be cancelled for a full refund.
	RefundRefundable RefundStatus = "REFUNDABLE"
	// RefundUpToDeadline means it is refundable until a stated deadline and
	// not after. Check CancellationPolicy.Deadline for when.
	RefundUpToDeadline RefundStatus = "REFUNDABLE_UP_TO_DEADLINE"
	// RefundUnknown means Amadeus did not say. Treat it as non-refundable when
	// the answer matters, since the alternative misleads a guest about their
	// money.
	RefundUnknown RefundStatus = "UNKNOWN"
)

// RefundPolicy is Amadeus's refundability statement for an offer.
type RefundPolicy struct {
	Status RefundStatus
}

// IsRefundable reports the offer's refundability, and whether Amadeus stated it
// clearly enough to rely on.
//
// It returns certain=false when Amadeus supplied no refundability block and no
// cancellation terms, in which case the caller must not tell a guest the
// booking is refundable. The conservative reading is deliberate: presenting a
// non-refundable rate as refundable costs a real person real money.
func (p Policies) IsRefundable() (refundable bool, certain bool) {
	if p.Refundable != nil {
		switch p.Refundable.Status {
		case RefundNonRefundable:
			return false, true
		case RefundRefundable, RefundUpToDeadline:
			return true, true
		case RefundUnknown:
			// Fall through to the cancellation terms, which may be clearer.
		}
	}

	if len(p.Cancellation) == 0 {
		return false, false
	}

	// A cancellation policy with a deadline in the future, or with no fee at
	// all, means the offer can still be cancelled without losing everything.
	for _, policy := range p.Cancellation {
		if policy.Deadline != nil || policy.Amount.Amount().IsZero() {
			return true, true
		}
	}
	return false, true
}

// FreeCancellationUntil returns the latest deadline before which the offer can
// be cancelled without charge, and false when there is none.
//
// This is the date to show beside "free cancellation". Policies carrying a fee
// are ignored, because a deadline attached to a charge is not free cancellation.
func (p Policies) FreeCancellationUntil() (time.Time, bool) {
	var latest time.Time
	found := false

	for _, policy := range p.Cancellation {
		if policy.Deadline == nil || !policy.IsFree() {
			continue
		}
		if !found || policy.Deadline.After(latest) {
			latest, found = *policy.Deadline, true
		}
	}
	return latest, found
}

// CancellationPolicy is one set of cancellation terms.
type CancellationPolicy struct {
	// Amount is the fee charged to cancel, zero when there is none.
	Amount money.Money
	// Percentage is the fee as a proportion, when Amadeus expressed it that way.
	Percentage string
	// NumberOfNights is the fee expressed as nights charged, when Amadeus
	// expressed it that way. A policy may state its fee in any of these three
	// forms, so check all of them before concluding cancellation is free.
	NumberOfNights int
	// Deadline is the instant after which this policy takes effect. It is nil
	// when the policy applies from booking.
	Deadline *time.Time
	// Type is the scope, typically "FULL_STAY".
	Type string
	// PolicyType is what triggers it: "CANCELLATION", "EARLY_CHECKOUT" or
	// "NO_SHOW".
	PolicyType string
	// Description is the terms in prose.
	Description *media.Text
}

// IsFree reports whether cancelling under this policy costs nothing, in any of
// the three forms Amadeus can express a fee.
func (c CancellationPolicy) IsFree() bool {
	return c.Amount.Amount().IsZero() &&
		c.NumberOfNights == 0 &&
		(c.Percentage == "" || c.Percentage == "0")
}

// PaymentPolicy is an amount due before arrival, with its deadline.
type PaymentPolicy struct {
	Amount           money.Money
	Deadline         *time.Time
	Description      *media.Text
	AcceptedPayments *AcceptedPayments
}

// GuaranteePolicy is a card-guarantee requirement, which reserves against
// no-show without taking payment.
type GuaranteePolicy struct {
	Description      *media.Text
	AcceptedPayments *AcceptedPayments
}

// AcceptedPayments lists how a policy may be settled.
type AcceptedPayments struct {
	// CreditCards are the accepted vendor codes, e.g. "VI", "CA", "AX".
	CreditCards []string
	// Methods are the accepted payment methods more broadly.
	Methods []string
	// CardPolicies are per-vendor requirements, including which cardholder
	// fields must be supplied.
	CardPolicies []CreditCardPolicy
}

// Accepts reports whether the given credit card vendor code is accepted.
func (a AcceptedPayments) Accepts(vendorCode string) bool {
	for _, code := range a.CreditCards {
		if code == vendorCode {
			return true
		}
	}
	for _, policy := range a.CardPolicies {
		if policy.VendorCode == vendorCode {
			return true
		}
	}
	return false
}

// CreditCardPolicy is one vendor's requirements.
type CreditCardPolicy struct {
	VendorCode string
	// Inputs are the fields the guest must supply for this vendor.
	Inputs []InputParameter
}

// InputParameter is a field a guest must supply to pay by card.
type InputParameter struct {
	// Label names the field, e.g. "cardHolderName".
	Label string
	// Optional reports whether it may be omitted. Amadeus sends this as the
	// string "true"/"false"; the mapper normalises it.
	Optional bool
}

// CheckInOutPolicy is the property's arrival and departure timing.
type CheckInOutPolicy struct {
	// CheckIn and CheckOut are local times at the property, as Amadeus
	// formatted them. They are kept as strings because Amadeus is inconsistent
	// about whether it sends "15:00", "15:00:00" or prose.
	CheckIn  string
	CheckOut string
	// CheckInDescription and CheckOutDescription hold the prose form, which is
	// often the only usable version.
	CheckInDescription  *media.Text
	CheckOutDescription *media.Text
}

// LengthOfStayPolicy constrains how many nights may be booked.
type LengthOfStayPolicy struct {
	Minimum int
	Maximum int
	// MinimumDescription and MaximumDescription explain the constraint in
	// prose, when Amadeus supplies it.
	MinimumDescription *media.Text
	MaximumDescription *media.Text
}

// Permits reports whether a stay of the given number of nights satisfies the
// policy. A zero Minimum or Maximum means unconstrained in that direction.
func (l LengthOfStayPolicy) Permits(nights int) bool {
	if l.Minimum > 0 && nights < l.Minimum {
		return false
	}
	if l.Maximum > 0 && nights > l.Maximum {
		return false
	}
	return true
}
