// Package booking turns an offer into a reservation, and manages it afterwards.
//
// It is the bounded context over Amadeus's Hotel Booking API. Its aggregate is
// the Order - one PNR in the Amadeus GDS - which holds one or more Bookings,
// each covering one or more rooms at one property. The Order ID is the only
// handle to a reservation afterwards: store it, because retrieval, modification
// and cancellation all require it.
//
// This is the only context that spends money. Create sends real payment
// details and, against the production environment, produces a real charge.
// Everything here validates before sending, and the validation is deliberately
// stricter than elsewhere in the SDK: a round trip saved matters less than a
// booking made wrong.
package booking

import (
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/money"
)

// OrderID identifies a hotel order, and is the handle to everything afterwards.
// Persist it at the moment a booking succeeds: without it there is no way to
// retrieve or cancel the reservation.
type OrderID string

// String returns the identifier.
func (id OrderID) String() string { return string(id) }

// BookingID identifies one hotel booking within an order. Modification and
// cancellation act on a booking, not on the whole order.
type BookingID string

// String returns the identifier.
func (id BookingID) String() string { return string(id) }

// Status is the state of a booking.
type Status string

const (
	// StatusConfirmed (HK) means the hotel has confirmed the reservation.
	StatusConfirmed Status = "CONFIRMED"
	// StatusPending (HN) means the request is with the hotel and not yet
	// accepted. It is not a reservation until it becomes CONFIRMED.
	StatusPending Status = "PENDING"
	// StatusOnHold (HO) means held pending deferred payment.
	StatusOnHold Status = "ON_HOLD"
	// StatusCancelled means the booking was cancelled.
	StatusCancelled Status = "CANCELLED"
	// StatusPast means confirmed, with check-out already in the past.
	StatusPast Status = "PAST"
	// StatusUnconfirmed (UC) means the hotel could not confirm it.
	StatusUnconfirmed Status = "UNCONFIRMED"
	// StatusDenied (NO) means the hotel refused it.
	StatusDenied Status = "DENIED"
	// StatusGhost (GK) is a passive record held for information only; it is not
	// a live reservation.
	StatusGhost Status = "GHOST"
	// StatusDeleted means the record was removed.
	StatusDeleted Status = "DELETED"
)

// IsActive reports whether the booking is a live reservation the guest can turn
// up on.
//
// PENDING is deliberately excluded: an on-request booking the hotel has not
// accepted is not a room, and telling a guest otherwise sends them to a
// property with no reservation.
func (s Status) IsActive() bool {
	return s == StatusConfirmed || s == StatusOnHold || s == StatusPast
}

// IsCancelled reports whether the booking is no longer in effect.
func (s Status) IsCancelled() bool {
	return s == StatusCancelled || s == StatusDenied || s == StatusDeleted
}

// Order is a set of hotel bookings made together for a set of guests. It
// corresponds to one PNR in the Amadeus GDS.
type Order struct {
	// ID is the order identifier, required for every later operation.
	ID OrderID
	// Bookings are the hotel bookings in this order, at least one.
	Bookings []Booking
	// Guests are everyone travelling on the order, with the IDs Amadeus
	// assigned them.
	Guests []BookedGuest
	// Records link the order to its GDS record locators.
	Records []Record
	// Self is the Amadeus URL for retrieving this order.
	Self string
}

// Reference returns the first GDS record locator on the order, which is what a
// guest quotes to the hotel or an agent quotes to Amadeus support.
func (o Order) Reference() (string, bool) {
	for _, record := range o.Records {
		if record.Reference != "" {
			return record.Reference, true
		}
	}
	return "", false
}

// IsConfirmed reports whether every booking in the order is a live reservation.
// An order that is only partly confirmed returns false; check the individual
// bookings to see which.
func (o Order) IsConfirmed() bool {
	if len(o.Bookings) == 0 {
		return false
	}
	for _, booking := range o.Bookings {
		if !booking.Status.IsActive() {
			return false
		}
	}
	return true
}

// Booking is one or more rooms reserved at one property.
type Booking struct {
	// ID identifies this booking within the order.
	ID BookingID
	// Status is the reservation's state. Check Status.IsActive before telling
	// a guest they have a room.
	Status Status
	// Providers hold the hotel provider's own confirmation and cancellation
	// numbers, which the property asks for on arrival or over the phone.
	Providers []ProviderReference
	// Rooms are the room-to-guest assignments.
	Rooms []RoomAssignment
	// Offer is the priced product that was booked, as it stood at booking time.
	Offer BookedOffer
	// Hotel identifies the property.
	Hotel Hotel
	// Payment summarises how the booking was paid for. Card numbers are
	// returned masked by Amadeus and are never complete here.
	Payment *PaymentSummary
	// TravelAgentID is the booking source receiving commission.
	TravelAgentID string
	// Arrival holds the guest's inbound flight, when supplied at booking.
	Arrival *ArrivalDetails
}

// ConfirmationNumber returns the provider's confirmation reference, which is
// what the guest quotes at the property.
//
// Amadeus sends "......" for an on-request booking that has no number yet, so a
// pending booking legitimately has none.
func (b Booking) ConfirmationNumber() (string, bool) {
	for _, provider := range b.Providers {
		if provider.ConfirmationNumber != "" && provider.ConfirmationNumber != placeholderNumber {
			return provider.ConfirmationNumber, true
		}
	}
	return "", false
}

// CancellationNumber returns the provider's cancellation reference, which
// exists only after a booking has been cancelled.
//
// Amadeus sends "NONE" when the provider returned no reference, which is a
// successful cancellation without a number rather than a failure.
func (b Booking) CancellationNumber() (string, bool) {
	for _, provider := range b.Providers {
		if provider.CancellationNumber != "" && provider.CancellationNumber != noNumber {
			return provider.CancellationNumber, true
		}
	}
	return "", false
}

// Placeholders Amadeus substitutes for a missing provider reference.
const (
	placeholderNumber = "......"
	noNumber          = "NONE"
)

// ProviderReference holds one hotel provider's references for a booking.
type ProviderReference struct {
	// ProviderCode is the 2-letter provider, e.g. "RT" for Accor.
	ProviderCode string
	// ConfirmationNumber is the provider's booking reference.
	ConfirmationNumber string
	// CancellationNumber is the provider's cancellation reference, set once
	// cancelled.
	CancellationNumber string
	// OnRequestNumber identifies a booking still awaiting hotel acceptance.
	OnRequestNumber string
}

// Record links an order to a record locator in an external system.
type Record struct {
	// Reference is the record locator, e.g. "JKL789".
	Reference string
	// OriginSystemCode names the system holding it, "GDS" for Amadeus.
	OriginSystemCode string
}

// Hotel identifies the booked property. It is deliberately thin: the booking
// context needs to name the property, not describe it. Use the content context
// for a full description.
type Hotel struct {
	ID                 string
	Name               string
	ChainCode          string
	TermsAndConditions string
	Self               string
}

// RoomAssignment correlates one room with the guests occupying it.
type RoomAssignment struct {
	// OfferID is the offer this room was booked from.
	OfferID string
	// Guests reference the guests in the room. The first is the main guest,
	// who holds the reservation and the form of payment.
	Guests []GuestReference
	// SpecialRequest is free text passed to the reception. It is a request,
	// not a guarantee, and the property may ignore it.
	SpecialRequest string
	// ManualMarkup is the agency markup applied, overriding Margin Manager.
	ManualMarkup *money.Money
}

// GuestReference points at a guest on the order, optionally with their hotel
// loyalty membership.
type GuestReference struct {
	// GuestID references the guest.
	GuestID string
	// HotelLoyaltyID is the chain rewards membership, used for points and
	// online check-in. An invalid number is rejected by the chain.
	HotelLoyaltyID string
}

// BookedGuest is a guest on an order, after Amadeus has assigned them an ID.
type BookedGuest struct {
	// ID is Amadeus's identifier for the guest on this order.
	ID int
	// TempID is the id the caller assigned at booking time, which is how a
	// returned guest is matched back to the one that was sent.
	TempID int

	Title     string
	FirstName string
	LastName  string
	Phone     string
	Email     string

	// ChildAge is the guest's age when they are a child, and zero for an
	// adult. Amadeus does not distinguish an infant from an adult here, so a
	// zero means "adult" on the way back even though it means "infant" on the
	// way in.
	ChildAge int
	// FrequentTraveler lists the airline loyalty memberships held.
	FrequentTraveler []FrequentTraveler
}

// FullName returns the guest's name as one string.
func (g BookedGuest) FullName() string {
	switch {
	case g.FirstName == "":
		return g.LastName
	case g.LastName == "":
		return g.FirstName
	default:
		return g.FirstName + " " + g.LastName
	}
}

// FrequentTraveler is an airline loyalty membership.
type FrequentTraveler struct {
	AirlineCode  string
	MembershipID string
}

// BookedOffer is the product as it was booked, preserved with the reservation.
//
// It is a snapshot rather than a live price: it records what was agreed, and
// does not change if the property later reprices the room.
type BookedOffer struct {
	ID string
	// Stay is the booked date range.
	Stay Stay
	// Guests is the occupancy booked.
	Guests Guests
	// RoomQuantity is how many rooms the offer covers.
	RoomQuantity int

	BoardType codes.BoardType
	RateCode  codes.RateCode
	Category  string
	// RateFamily is Amadeus's classification of the rate.
	RateFamily *RateFamily

	// Price is what was agreed.
	Price *Price
	// Policies are the terms the booking was made under, chiefly cancellation.
	Policies *Policies
	// Room describes what was booked.
	Room *Room
	// Extras are the additional services attached.
	Extras []Extra
	// Commission is what the booker earns.
	Commission *Commission
}

// Stay is the booked date range.
type Stay struct {
	CheckIn  datetime.Date
	CheckOut datetime.Date
}

// Nights returns the number of nights booked.
func (s Stay) Nights() int {
	if s.CheckIn.IsZero() || s.CheckOut.IsZero() {
		return 0
	}
	return s.CheckIn.DaysUntil(s.CheckOut)
}

// Guests is the occupancy of a booked offer.
type Guests struct {
	Adults    int
	ChildAges []int
}

// RateFamily classifies a rate.
type RateFamily struct {
	Code string
	Type string
}

// Room is the booked room.
type Room struct {
	Type        string
	Category    string
	Beds        int
	BedType     string
	Description *media.Text
	// Details is Amadeus's extended room block, when supplied.
	Details *RoomDetails
}

// RoomDetails is the extended room description on a booking.
type RoomDetails struct {
	ID             string
	Name           *media.Text
	Type           string
	Description    string
	Category       string
	Classification string
	BedType        string
	Beds           int
	Dimensions     *media.Dimensions
	MaxOccupancy   *Occupancy
	Media          []media.Asset
}

// Occupancy is how many people a room takes.
type Occupancy struct {
	Adults   int
	Children int
	Total    int
}

// Extra is an additional service attached to a booking.
type Extra struct {
	Code          string
	Description   string
	IsChargeable  bool
	PricingMethod string
	Quantity      int
	Attribute     string
	Price         *Price
}

// Commission is what the booker earns on a booking.
type Commission struct {
	Amount      money.Money
	Percentage  string
	Description *media.Text
}

// PaymentSummary describes how a booking was paid for.
//
// It never contains a complete card number: Amadeus masks the number in
// responses and omits the security code entirely.
type PaymentSummary struct {
	// Method is how the booking was paid, e.g. PaymentCreditCard.
	Method PaymentMethod
	// Instructions is free text passed to the hotelier.
	Instructions string
	// Card is the masked card detail, when a card was used.
	Card *MaskedCard
	// IATANumber is the agency number guaranteeing the booking.
	IATANumber string
	// Supplier holds the hotel supplier's contact details.
	Supplier *SupplierContact
	// VirtualCard describes the virtual card generated for a B2B wallet
	// payment.
	VirtualCard *VirtualCard
}

// MaskedCard is a card as Amadeus returns it: enough to recognise, never enough
// to charge.
type MaskedCard struct {
	// VendorCode is the two-letter card type, e.g. "VI" for Visa.
	VendorCode string
	// MaskedNumber is the card number with all but the last digits obscured.
	MaskedNumber string
	// ExpiryDate is the expiry as Amadeus formatted it.
	ExpiryDate string
	// HolderName is the cardholder.
	HolderName string
}

// VirtualCard is a virtual card generated for a booking.
type VirtualCard struct {
	// Reference identifies the generated card. It is a reference, not a card
	// number: Amadeus never returns the virtual card's digits here.
	Reference string
	// Provider is the payment provider that issued it.
	Provider string
}

// SupplierContact holds a hotel supplier's contact details.
type SupplierContact struct {
	Phone string
	Fax   string
	Email string
}

// ArrivalDetails is the guest's inbound flight, passed to the property so it
// knows when to expect them.
type ArrivalDetails struct {
	CarrierCode  string
	FlightNumber string
	// DepartureAirport and ArrivalAirport are IATA codes.
	DepartureAirport string
	ArrivalAirport   string
	Terminal         string
	// ArrivingAt is the local arrival time at the destination.
	ArrivingAt *time.Time
}

// CancellationResult is what Delete returns: the provider's reference for the
// cancellation, when it supplied one.
type CancellationResult struct {
	// CancellationNumber is the provider's reference. It is empty when the
	// provider returned none, which is still a successful cancellation.
	CancellationNumber string
}
