package content

import (
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Policies are a property's standing rules.
//
// These are property-level and descriptive: they say what the hotel's policy
// generally is, not what terms a particular booking carries. The terms that
// actually bind a reservation are on the offer, in the offers context, and can
// differ from these. Where the two disagree, the offer wins - so never quote a
// cancellation policy from here to a guest who has booked.
type Policies struct {
	// Payment lists the accepted payment and guarantee arrangements.
	Payment []PaymentPolicy
	// CheckInOut is the property's arrival and departure timing.
	CheckInOut []CheckInOutPolicy
	// Cancellation describes the property's general cancellation terms.
	Cancellation []CancellationPolicy
	// Pets describes the property's terms for animals.
	Pets []PetPolicy
	// Tax describes the taxes and fees the property levies.
	Tax []TaxPolicy
	// Commission describes what the property pays a booker.
	Commission []CommissionPolicy
	// Guest describes rules about who may stay and how children are treated.
	Guest []GuestPolicy
	// Loyalty describes the chain's rewards benefits.
	Loyalty []LoyaltyPolicy
	// StayRequirements are minimum-stay and similar conditions, in prose.
	StayRequirements []media.Text
}

// PaymentPolicy is one accepted payment or guarantee arrangement.
type PaymentPolicy struct {
	// Type classifies the arrangement, e.g. "GUARANTEE", "DEPOSIT".
	Type string
	// Guarantee describes the card guarantee accepted, when the property takes
	// one.
	Guarantee *Guarantee
	// Details are additional terms in prose.
	Details []media.Text
}

// Guarantee describes a card-guarantee arrangement.
type Guarantee struct {
	// AcceptedCards are the card vendor codes accepted, e.g. "VI", "AX".
	AcceptedCards []string
	// AcceptedMethods are the accepted payment methods.
	AcceptedMethods []string
	// Description explains the terms.
	Description *media.Text
}

// CheckInOutPolicy is the property's arrival and departure timing.
type CheckInOutPolicy struct {
	// CheckIn and CheckOut are local times at the property, as Amadeus
	// formatted them. They are strings because Amadeus is inconsistent about
	// whether it sends "15:00", "1500" or prose.
	CheckIn  string
	CheckOut string
	// CheckInDescription and CheckOutDescription hold the prose form, which is
	// frequently the only populated version of the two.
	CheckInDescription  *media.Text
	CheckOutDescription *media.Text
}

// CancellationPolicy is the property's general cancellation terms.
type CancellationPolicy struct {
	// Amount is the fee to cancel.
	Amount money.Money
	// Percentage is the fee as a proportion.
	Percentage string
	// NumberOfNights is the fee expressed as nights charged. A policy may
	// state its fee in any of these three forms.
	NumberOfNights int
	// Deadline is when the policy takes effect, as Amadeus published it.
	Deadline string
	// PolicyType is what triggers it, e.g. "CANCELLATION", "NO_SHOW".
	PolicyType string
	// Description explains the terms.
	Description *media.Text
}

// PetPolicy describes the property's terms for animals.
type PetPolicy struct {
	// Code identifies the policy, e.g. whether pets are permitted at all.
	Code string
	// Description explains the terms in prose. Amadeus publishes no structured
	// "pets allowed" flag, so this text is the answer.
	Description string
	// PricingMethod is how any charge is assessed, e.g. "PER_STAY".
	PricingMethod string
}

// TaxPolicy is one tax or fee the property levies.
type TaxPolicy struct {
	// Code and Description identify it.
	Code        string
	Description string
	// Amount is the charge.
	Amount money.Money
	// Percentage is the rate, when proportional.
	Percentage string
	// Included reports that it is already in quoted rates. When it is not, the
	// guest pays it on top - which is what explains a bill larger than the
	// booking total.
	Included bool
	// Frequency and Mode describe how it is assessed, e.g. "PER_NIGHT" and
	// "PER_ROOM".
	Frequency string
	Mode      string
}

// CommissionPolicy describes what the property pays a booker.
type CommissionPolicy struct {
	// Percentage is the commission rate.
	Percentage string
	// Amount is a flat commission, where one applies.
	Amount money.Money
	// Description explains the terms.
	Description *media.Text
}

// GuestPolicy describes rules about who may stay and how children are treated.
type GuestPolicy struct {
	// MinimumGuestAge is the youngest a guest may be to check in
	// unaccompanied. It is a real constraint in several countries, and worth
	// surfacing before a young traveller books.
	MinimumGuestAge int
	// MaxChildAgeForBedSharing is the oldest a child can be and still share an
	// existing bed rather than needing an extra one.
	MaxChildAgeForBedSharing int
	// ChildStayFree reports that children stay at no extra charge, and
	// ChildStayFreeCutoffAge the age at which that stops.
	ChildStayFree          bool
	ChildStayFreeCutoffAge int
}

// LoyaltyPolicy describes a chain rewards benefit.
type LoyaltyPolicy struct {
	// Eligibility describes who qualifies.
	Eligibility string
	// Membership names the programme and tier.
	Membership *LoyaltyMembership
	// Benefits are what members earn or receive.
	Benefits []LoyaltyBenefit
	// Discount is the member discount, when the programme offers one.
	Discount *LoyaltyDiscount
}

// LoyaltyMembership names a rewards programme.
type LoyaltyMembership struct {
	// ProgramName is the scheme, e.g. "Hilton Honors".
	ProgramName string
	// Tier is the membership level.
	Tier string
	// ProviderCode identifies the programme operator.
	ProviderCode string
}

// LoyaltyBenefit is one thing members earn or receive.
type LoyaltyBenefit struct {
	// Type classifies it, e.g. "POINTS", "UPGRADE".
	Type string
	// Description explains it.
	Description string
	// Quantity is how much is earned, where it is countable.
	Quantity int
	// Unit names what Quantity counts.
	Unit string
}

// LoyaltyDiscount is a member discount.
type LoyaltyDiscount struct {
	// Percentage is the discount rate.
	Percentage string
	// Amount is a flat discount, where one applies.
	Amount money.Money
	// Description explains the terms.
	Description string
}
