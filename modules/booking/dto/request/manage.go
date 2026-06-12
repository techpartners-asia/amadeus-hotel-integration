package requestBookingDTO

// ==================== Hotel Booking Manage (v2.2) ====================
//
// UpdateHotelBookingRequest is the PATCH body for modifying a single hotel
// booking within an existing hotel order.
// Endpoint: PATCH /booking/hotel-orders/{hotelOrderId}/hotel-bookings/{hotelBookingId}
//
// All fields are optional: send only the parts of the booking you want to change.

type (
	UpdateHotelBookingRequest struct {
		// Data - the update payload.
		Data UpdateHotelBooking `json:"data"`
	}

	UpdateHotelBooking struct {
		// HotelBooking - the hotel booking changes (room association, offer, payment).
		HotelBooking UpdateHotelBookingData `json:"hotelBooking"`
	}

	UpdateHotelBookingData struct {
		// RoomAssociation - updated special requests and guest references.
		RoomAssociation *UpdateRoomAssociation `json:"roomAssociation,omitempty"`
		// HotelOffer - updated offer / product details (dates, rate code, guests).
		HotelOffer *UpdateHotelOffer `json:"hotelOffer,omitempty"`
		// Payment - updated payment card information.
		Payment *UpdatePayment `json:"payment,omitempty"`
	}

	UpdateRoomAssociation struct {
		// SpecialRequest - free-text special request for the room.
		SpecialRequest string `json:"specialRequest,omitempty"`
		// GuestReferences - references to guests and their loyalty ids. Reuses the
		// creation-time GuestReference type.
		GuestReferences []GuestReference `json:"guestReferences,omitempty"`
	}

	UpdateHotelOffer struct {
		// Id - the offer id to switch to, if changing the rate.
		Id string `json:"id,omitempty"`
		// Product - updated product attributes.
		Product *UpdateHotelOfferProduct `json:"product,omitempty"`
	}

	UpdateHotelOfferProduct struct {
		// CheckInDate - new check-in date (YYYY-MM-DD).
		CheckInDate string `json:"checkInDate,omitempty"`
		// CheckOutDate - new check-out date (YYYY-MM-DD).
		CheckOutDate string `json:"checkOutDate,omitempty"`
		// RateCode - new rate code.
		RateCode string `json:"rateCode,omitempty"`
		// Category - new rate category.
		Category string `json:"category,omitempty"`
		// Guests - updated guest counts.
		Guests *UpdateOfferGuests `json:"guests,omitempty"`
	}

	UpdateOfferGuests struct {
		// Adults - number of adult guests.
		Adults int `json:"adults,omitempty"`
		// ChildAges - ages of children, one entry per child.
		ChildAges []int `json:"childAges,omitempty"`
	}

	UpdatePayment struct {
		// PaymentCard - updated credit card details (reuses the creation-time
		// PaymentCard type: card info, 3DS, billing address).
		PaymentCard *PaymentCard `json:"paymentCard,omitempty"`
	}
)
