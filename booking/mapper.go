package booking

import (
	"strconv"

	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto/bookingreq"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto/bookingres"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/mapping"
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// This file is the anti-corruption layer for the booking context, in both
// directions: toRequest builds Amadeus's wire body from a domain Reservation,
// and mapOrder translates the response back.
//
// The outbound direction is the more delicate one. A field dropped on the way
// out is a booking made on terms the caller did not choose, so every domain
// field is written through explicitly rather than by any reflective shortcut.

// toRequest builds the Amadeus create-booking body from a Reservation.
func toRequest(r Reservation) bookingreq.HotelBookingRequest {
	data := bookingreq.BookingData{
		Type:        "hotel-order",
		Guests:      make([]bookingreq.Guest, 0, len(r.Guests)),
		Payment:     toPayment(r.Payment),
		TravelAgent: toAgent(r.Agent),
	}

	for _, guest := range r.Guests {
		wire := bookingreq.Guest{
			Tid:       guest.ID,
			Title:     guest.Title,
			FirstName: guest.FirstName,
			LastName:  guest.LastName,
			Phone:     guest.Phone,
			Email:     guest.Email,
			ChildAge:  guest.ChildAge,
		}
		for _, traveler := range guest.FrequentTraveler {
			wire.FrequentTraveler = append(wire.FrequentTraveler, bookingreq.FrequentTraveler{
				AirlineCode:        traveler.AirlineCode,
				FrequentTravelerId: traveler.MembershipID,
			})
		}
		data.Guests = append(data.Guests, wire)
	}

	for _, room := range r.Rooms {
		association := bookingreq.RoomAssociation{
			HotelOfferId:   room.OfferID,
			SpecialRequest: room.SpecialRequest,
		}
		for _, id := range room.GuestIDs {
			association.GuestReferences = append(association.GuestReferences, bookingreq.GuestReference{
				GuestReference: strconv.Itoa(id),
				HotelLoyaltyId: room.LoyaltyIDs[id],
			})
		}
		if room.ManualMarkup != nil {
			association.TravelAgentManualMarkup = &bookingreq.TravelAgentManualMarkup{
				Amount:   room.ManualMarkup.Amount().String(),
				Currency: string(room.ManualMarkup.Currency()),
			}
		}
		data.RoomAssociations = append(data.RoomAssociations, association)
	}

	if r.Arrival != nil {
		flight := &bookingreq.ArrivalFlightDetails{
			CarrierCode: r.Arrival.CarrierCode,
			Number:      r.Arrival.FlightNumber,
			Departure:   &bookingreq.Departure{IataCode: r.Arrival.DepartureAirport},
		}
		if r.Arrival.ArrivalAirport != "" || r.Arrival.Terminal != "" || r.Arrival.ArrivingAt != nil {
			arrival := &bookingreq.Arrival{
				IataCode: r.Arrival.ArrivalAirport,
				Terminal: r.Arrival.Terminal,
			}
			if r.Arrival.ArrivingAt != nil {
				arrival.At = *r.Arrival.ArrivingAt
			}
			flight.Arrival = arrival
		}
		data.ArrivalInformation = &bookingreq.ArrivalInformation{ArrivalFlightDetails: flight}
	}

	if r.AddToPNR != nil {
		origin := r.AddToPNR.OriginSystemCode
		if origin == "" {
			origin = "GDS"
		}
		data.AssociatedRecord = &bookingreq.AssociatedRecord{
			Reference:        r.AddToPNR.Reference,
			OriginSystemCode: origin,
		}
	}

	return bookingreq.HotelBookingRequest{Data: data}
}

func toPayment(p Payment) bookingreq.Payment {
	payment := bookingreq.Payment{
		Method:              string(p.Method),
		PaymentInstructions: p.Instructions,
		PayerCode:           p.PayerCode,
	}

	if p.IATANumber != "" {
		payment.IataTravelAgency = &bookingreq.IataTravelAgency{IataNumber: p.IATANumber}
	}
	if p.Supplier != nil {
		payment.HotelSupplierInformation = &bookingreq.HotelSupplierInformation{
			Phone: p.Supplier.Phone,
			Fax:   p.Supplier.Fax,
			Email: p.Supplier.Email,
		}
	}
	if p.BillBack != nil {
		payment.BillBack = &bookingreq.BillBack{
			BillbackProviderCode:          p.BillBack.ProviderCode,
			BillbackProviderAccountNumber: p.BillBack.AccountNumber,
			TravelAgencyId:                p.BillBack.TravelAgencyID,
			BookerId:                      p.BillBack.BookerID,
		}
	}
	if p.Card != nil {
		payment.PaymentCard = toCard(*p.Card)
	}

	return payment
}

func toCard(c Card) *bookingreq.PaymentCard {
	card := &bookingreq.PaymentCard{
		PaymentCardInfo: bookingreq.PaymentCardInfo{
			VendorCode:   c.VendorCode,
			CardNumber:   c.Number,
			ExpiryDate:   c.Expiry,
			SecurityCode: c.SecurityCode,
			HolderName:   c.HolderName,
		},
	}

	if c.ThreeDS != nil {
		card.ThreeDomainSecure = &bookingreq.ThreeDomainSecure{
			Version:         c.ThreeDS.Version,
			Eci:             c.ThreeDS.ECI,
			CryptogramValue: c.ThreeDS.CryptogramValue,
			DsTransactionId: c.ThreeDS.DSTransactionID,
			Xid:             c.ThreeDS.XID,
			TransStatus:     c.ThreeDS.TransStatus,
			ParesStatus:     c.ThreeDS.ParesStatus,
			VeresStatus:     c.ThreeDS.VeresStatus,
		}
	}
	if c.BillingAddress != nil {
		card.Address = &bookingreq.Address{
			Lines:       c.BillingAddress.Lines,
			PostalCode:  c.BillingAddress.PostalCode,
			CityName:    c.BillingAddress.CityName,
			PostalBox:   c.BillingAddress.PostalBox,
			StateCode:   c.BillingAddress.StateCode,
			CountryCode: c.BillingAddress.CountryCode,
		}
	}

	return card
}

func toAgent(a Agent) bookingreq.TravelAgent {
	return bookingreq.TravelAgent{
		TravelAgentId: a.ID,
		Contact: bookingreq.Contact{
			Email: a.Email,
			Phone: a.Phone,
			Fax:   a.Fax,
		},
	}
}

// toUpdateRequest builds the PATCH body for a modification.
func toUpdateRequest(m Modification) bookingreq.UpdateHotelBookingRequest {
	data := bookingreq.UpdateHotelBookingData{}

	if m.SpecialRequest != nil || len(m.GuestIDs) > 0 {
		association := &bookingreq.UpdateRoomAssociation{}
		if m.SpecialRequest != nil {
			association.SpecialRequest = *m.SpecialRequest
		}
		for _, id := range m.GuestIDs {
			association.GuestReferences = append(association.GuestReferences, bookingreq.GuestReference{
				GuestReference: strconv.Itoa(id),
				HotelLoyaltyId: m.LoyaltyIDs[id],
			})
		}
		data.RoomAssociation = association
	}

	if m.OfferID != "" || m.Stay != nil || m.Guests != nil || m.RateCode != "" {
		offer := &bookingreq.UpdateHotelOffer{Id: m.OfferID}
		product := &bookingreq.UpdateHotelOfferProduct{RateCode: string(m.RateCode)}

		if m.Stay != nil {
			product.CheckInDate = m.Stay.CheckIn.String()
			product.CheckOutDate = m.Stay.CheckOut.String()
		}
		if m.Guests != nil {
			product.Guests = &bookingreq.UpdateOfferGuests{
				Adults:    m.Guests.Adults,
				ChildAges: m.Guests.ChildAges,
			}
		}
		if *product != (bookingreq.UpdateHotelOfferProduct{}) || product.Guests != nil {
			offer.Product = product
		}
		data.HotelOffer = offer
	}

	if m.Card != nil {
		data.Payment = &bookingreq.UpdatePayment{PaymentCard: toCard(*m.Card)}
	}

	return bookingreq.UpdateHotelBookingRequest{Data: bookingreq.UpdateHotelBooking{HotelBooking: data}}
}

// mapOrder translates an Amadeus hotel order into the domain aggregate.
func mapOrder(o bookingres.HotelOrder) Order {
	order := Order{
		ID:   OrderID(o.Id),
		Self: o.Self,
	}

	for _, booking := range o.HotelBookings {
		order.Bookings = append(order.Bookings, mapBooking(booking))
	}
	for _, guest := range o.Guests {
		order.Guests = append(order.Guests, mapGuest(guest))
	}
	for _, record := range o.AssociatedRecords {
		order.Records = append(order.Records, Record{
			Reference:        record.Reference,
			OriginSystemCode: record.OriginSystemCode,
		})
	}

	return order
}

func mapBooking(b bookingres.HotelBooking) Booking {
	booking := Booking{
		ID:            BookingID(b.Id),
		Status:        Status(b.BookingStatus),
		Offer:         mapBookedOffer(b.HotelOffer),
		TravelAgentID: b.TravelAgentId,
		Hotel: Hotel{
			ID:                 b.Hotel.HotelId,
			Name:               b.Hotel.Name,
			ChainCode:          b.Hotel.ChainCode,
			TermsAndConditions: b.Hotel.TermsAndConditions,
			Self:               b.Hotel.Self,
		},
	}

	for _, provider := range b.HotelProviderInformation {
		booking.Providers = append(booking.Providers, ProviderReference{
			ProviderCode:       provider.HotelProviderCode,
			ConfirmationNumber: provider.ConfirmationNumber,
			CancellationNumber: provider.CancellationNumber,
			OnRequestNumber:    provider.OnRequestNumber,
		})
	}

	for _, association := range b.RoomAssociations {
		room := RoomAssignment{
			OfferID:        association.HotelOfferId,
			SpecialRequest: association.SpecialRequest,
		}
		for _, reference := range association.GuestReferences {
			room.Guests = append(room.Guests, GuestReference{
				GuestID:        reference.GuestReference,
				HotelLoyaltyID: reference.HotelLoyaltyId,
			})
		}
		if association.TravelAgentManualMarkup != nil {
			markup := mapping.Money(
				association.TravelAgentManualMarkup.Amount,
				association.TravelAgentManualMarkup.Currency,
			)
			room.ManualMarkup = &markup
		}
		booking.Rooms = append(booking.Rooms, room)
	}

	if b.Payment != nil {
		booking.Payment = mapPaymentSummary(*b.Payment)
	}
	if b.ArrivalInformation != nil && b.ArrivalInformation.ArrivalFlightDetails != nil {
		booking.Arrival = mapArrival(*b.ArrivalInformation.ArrivalFlightDetails)
	}

	return booking
}

// offerCurrency returns the currency the booked offer was priced in, which the
// nested price blocks inherit when they carry none of their own.
func offerCurrency(o bookingres.HotelOffer) string {
	if o.Price != nil {
		return o.Price.Currency
	}
	return ""
}

func mapGuest(g bookingres.ResponseGuest) BookedGuest {
	guest := BookedGuest{
		ID:        g.Id,
		TempID:    g.Tid,
		Title:     g.Title,
		FirstName: g.FirstName,
		LastName:  g.LastName,
		Phone:     g.Phone,
		Email:     g.Email,
		ChildAge:  g.ChildAge,
	}
	for _, traveler := range g.FrequentTraveler {
		guest.FrequentTraveler = append(guest.FrequentTraveler, FrequentTraveler{
			AirlineCode:  traveler.AirlineCode,
			MembershipID: traveler.FrequentTravelerId,
		})
	}
	return guest
}

func mapBookedOffer(o bookingres.HotelOffer) BookedOffer {
	offer := BookedOffer{
		ID: o.Id,
		Stay: Stay{
			CheckIn:  mapping.Date(o.CheckInDate),
			CheckOut: mapping.Date(o.CheckOutDate),
		},
		RoomQuantity: o.RoomQuantity,
		BoardType:    codes.BoardType(o.BoardType),
		RateCode:     codes.RateCode(o.RateCode),
		Category:     o.Category,
	}

	currency := offerCurrency(o)

	if o.Guests != nil {
		offer.Guests = Guests{Adults: o.Guests.Adults, ChildAges: o.Guests.ChildAges}
	}
	if o.RateFamilyEstimated != nil {
		offer.RateFamily = &RateFamily{Code: o.RateFamilyEstimated.Code, Type: o.RateFamilyEstimated.Type}
	}
	if o.Price != nil {
		price := mapPrice(*o.Price)
		offer.Price = &price
	}
	if o.Policies != nil {
		policies := mapPolicies(*o.Policies, currency)
		offer.Policies = &policies
	}
	if o.Room != nil {
		offer.Room = &Room{
			Type:        o.Room.Type,
			Description: mapQualifiedText(o.Room.Description),
		}
		if o.Room.TypeEstimated != nil {
			offer.Room.Category = o.Room.TypeEstimated.Category
			offer.Room.Beds = o.Room.TypeEstimated.Beds
			offer.Room.BedType = o.Room.TypeEstimated.BedType
		}
	}
	if o.RoomInformation != nil {
		details := mapRoomDetails(*o.RoomInformation)
		if offer.Room == nil {
			offer.Room = &Room{}
		}
		offer.Room.Details = details
	}
	if o.Commission != nil {
		offer.Commission = &Commission{
			Amount:      mapping.Money(o.Commission.Amount, currency),
			Percentage:  o.Commission.Percentage,
			Description: mapQualifiedText(o.Commission.Description),
		}
	}
	for _, service := range o.Services {
		extra := Extra{
			Code:          service.Code,
			Description:   service.Description,
			IsChargeable:  service.IsChargeable,
			PricingMethod: service.PricingMethod,
			Quantity:      service.Quantity,
			Attribute:     service.ServiceAttribute,
		}
		if service.Price != nil {
			price := mapPrice(*service.Price)
			extra.Price = &price
		}
		offer.Extras = append(offer.Extras, extra)
	}

	return offer
}

func mapPrice(p bookingres.HotelPrice) Price {
	price := Price{
		Currency:     money.Currency(p.Currency),
		Base:         mapping.Money(p.Base, p.Currency),
		Total:        mapping.Money(p.Total, p.Currency),
		SellingTotal: mapping.Money(p.SellingTotal, p.Currency),
	}

	for _, markup := range p.Markups {
		price.Markups = append(price.Markups, mapping.Money(markup.Amount, p.Currency))
	}
	for _, tax := range p.Taxes {
		currency := mapping.Or(tax.Currency, p.Currency)
		price.Taxes = append(price.Taxes, Tax{
			Amount:           mapping.Money(tax.Amount, currency),
			Code:             tax.Code,
			Description:      tax.Description,
			Percentage:       tax.Percentage,
			Included:         tax.Included,
			PricingFrequency: tax.PricingFrequency,
			PricingMode:      tax.PricingMode,
			Applicable:       mapApplicable(tax.ApplicableDate),
		})
	}
	if p.Variations != nil {
		price.Variations = mapVariations(*p.Variations, p.Currency)
	}

	return price
}

func mapVariations(v bookingres.PriceVariations, fallbackCurrency string) *Variations {
	// The booking API sends only the per-period changes; unlike Hotel Search it
	// has no average block, so Variations here carries Changes alone.
	variations := &Variations{}

	for _, change := range v.Changes {
		currency := mapping.Or(change.Currency, fallbackCurrency)
		variations.Changes = append(variations.Changes, PricePeriod{
			Start:        mapping.Date(change.StartDate),
			End:          mapping.Date(change.EndDate),
			Currency:     money.Currency(currency),
			Base:         mapping.Money(change.Base, currency),
			Total:        mapping.Money(change.Total, currency),
			SellingTotal: mapping.Money(change.SellingTotal, currency),
		})
	}

	return variations
}

func mapPolicies(p bookingres.PolicyDetails, currency string) Policies {
	policies := Policies{PaymentType: p.PaymentType}

	for _, cancellation := range p.Cancellations {
		policies.Cancellation = append(policies.Cancellation, CancellationPolicy{
			Amount:         mapping.Money(cancellation.Amount, currency),
			Percentage:     cancellation.Percentage,
			NumberOfNights: cancellation.NumberOfNights,
			Deadline:       mapping.Timestamp(cancellation.Deadline),
			Type:           cancellation.Type,
			PolicyType:     cancellation.PolicyType,
			Description:    mapQualifiedText(cancellation.Description),
		})
	}

	if p.Refundable != nil {
		policies.Refundable = &RefundPolicy{Status: RefundStatus(p.Refundable.CancellationRefund)}
	}
	if p.Deposit != nil {
		policies.Deposit = mapAmountDue(*p.Deposit, currency)
	}
	if p.Prepay != nil {
		policies.Prepay = mapAmountDue(*p.Prepay, currency)
	}
	if p.Guarantee != nil {
		policies.Guarantee = &GuaranteePolicy{
			Description: mapQualifiedText(p.Guarantee.Description),
		}
		if p.Guarantee.AcceptedPayments != nil {
			policies.Guarantee.AcceptedPayments = &AcceptedPayments{
				CreditCards: p.Guarantee.AcceptedPayments.CreditCards,
				Methods:     p.Guarantee.AcceptedPayments.Methods,
			}
		}
	}
	if p.HoldTime != nil {
		policies.HoldTime = mapping.Timestamp(p.HoldTime.Deadline)
	}
	if p.CheckInOut != nil {
		policies.CheckInOut = &CheckInOutPolicy{
			CheckIn:             p.CheckInOut.CheckIn,
			CheckOut:            p.CheckInOut.CheckOut,
			CheckInDescription:  mapQualifiedText(p.CheckInOut.CheckInDescription),
			CheckOutDescription: mapQualifiedText(p.CheckInOut.CheckOutDescription),
		}
	}

	// The structured block wins over the two deprecated top-level fields, which
	// Amadeus still sends alongside it. Note the wire typo "minimuLengthOfStay",
	// reproduced in the DTO because matching it is what makes the field decode.
	minimum, maximum := p.MinimumLengthOfStay, p.MaximumLengthOfStay
	if p.LengthOfStay != nil {
		if p.LengthOfStay.MinimumLengthOfStay > 0 {
			minimum = p.LengthOfStay.MinimumLengthOfStay
		}
		if p.LengthOfStay.MaximumLengthOfStay > 0 {
			maximum = p.LengthOfStay.MaximumLengthOfStay
		}
	}
	if minimum > 0 || maximum > 0 {
		policies.LengthOfStay = &LengthOfStayPolicy{Minimum: minimum, Maximum: maximum}
	}

	for _, detail := range p.AdditionalDetails {
		for i := range detail.Description {
			if text := mapQualifiedText(&detail.Description[i]); text != nil {
				policies.Details = append(policies.Details, *text)
			}
		}
	}

	return policies
}

func mapAmountDue(d bookingres.DepositPolicy, currency string) *AmountDuePolicy {
	return &AmountDuePolicy{
		Amount:      mapping.Money(d.Amount, currency),
		Deadline:    mapping.Timestamp(d.Deadline),
		Description: mapQualifiedText(d.Description),
	}
}

func mapRoomDetails(r bookingres.RoomInformation) *RoomDetails {
	details := &RoomDetails{
		ID:             r.Id,
		Type:           r.Type,
		Description:    textValue(r.Description),
		Category:       r.HotelRoomCategory,
		Classification: r.HotelRoomClassification,
		BedType:        r.BedType,
		Beds:           r.Beds,
		Dimensions:     mapping.Dimensions(r.Dimensions),
		Media:          mapping.MediaAssets(r.Media),
	}

	details.Name = mapQualifiedText(r.Name)
	if r.MaxPersonCapacity != nil {
		details.MaxOccupancy = &Occupancy{
			Adults:   r.MaxPersonCapacity.Adults,
			Children: r.MaxPersonCapacity.Children,
			Total:    r.MaxPersonCapacity.Total,
		}
	}

	return details
}

func mapPaymentSummary(p bookingres.PaymentOutput) *PaymentSummary {
	summary := &PaymentSummary{
		Method:       PaymentMethod(p.Method),
		Instructions: p.PaymentInstructions,
	}

	if p.IataTravelAgency != nil {
		summary.IATANumber = p.IataTravelAgency.IataNumber
	}
	if p.HotelSupplierInformation != nil {
		summary.Supplier = &SupplierContact{
			Phone: p.HotelSupplierInformation.Phone,
			Fax:   p.HotelSupplierInformation.Fax,
			Email: p.HotelSupplierInformation.Email,
		}
	}
	if p.PaymentCard != nil {
		summary.Card = &MaskedCard{
			VendorCode:   p.PaymentCard.PaymentCardInfo.VendorCode,
			MaskedNumber: p.PaymentCard.PaymentCardInfo.CardNumber,
			ExpiryDate:   p.PaymentCard.PaymentCardInfo.ExpiryDate,
			HolderName:   p.PaymentCard.PaymentCardInfo.HolderName,
		}
	}
	if p.B2bWallet != nil {
		summary.VirtualCard = &VirtualCard{
			Reference: p.B2bWallet.VirtualCreditCardId,
			Provider:  p.B2bWallet.PaymentProvider,
		}
	}

	return summary
}

func mapArrival(f bookingres.ArrivalFlightDetails) *ArrivalDetails {
	arrival := &ArrivalDetails{
		CarrierCode:  f.CarrierCode,
		FlightNumber: f.Number,
	}

	if f.Departure != nil {
		arrival.DepartureAirport = f.Departure.IataCode
	}
	if f.Arrival != nil {
		arrival.ArrivalAirport = f.Arrival.IataCode
		arrival.Terminal = f.Arrival.Terminal
		if !f.Arrival.At.IsZero() {
			at := f.Arrival.At
			arrival.ArrivingAt = &at
		}
	}

	return arrival
}

// textValue returns a text block's content, or "" when absent.
func textValue(q *bookingres.QualifiedFreeText) string {
	if q == nil {
		return ""
	}
	return q.Text
}

// mapQualifiedText translates the booking API's text block, which carries only
// a value and a language.
func mapQualifiedText(q *bookingres.QualifiedFreeText) *media.Text {
	if q == nil || q.Text == "" {
		return nil
	}
	return &media.Text{Value: q.Text, Lang: q.Lang}
}

// mapApplicable translates a tax's date range.
func mapApplicable(d *bookingres.TaxApplicableDate) *DateRange {
	if d == nil {
		return nil
	}
	return &DateRange{Start: mapping.Date(d.Start), End: mapping.Date(d.End)}
}
