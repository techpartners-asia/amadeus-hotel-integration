package responseBookingDTO

// ==================== Hotel Booking Retrieve / Manage envelopes ====================

type (
	// HotelOrderResponse is the envelope returned by Retrieve (v2.1) and Cancel
	// (v2.2). The data is the full hotel order; for Cancel the affected booking's
	// bookingStatus is updated to CANCELLED.
	HotelOrderResponse struct {
		Data     HotelOrder `json:"data"`
		Warnings []Warning  `json:"warnings,omitempty"`
		Errors   []Error    `json:"errors,omitempty"`
	}

	// HotelBookingUpdateResponse is the envelope returned by Modify (PATCH, v2.2).
	// Data is a light reference to the order; Included carries the full updated order.
	HotelBookingUpdateResponse struct {
		Data     UpdatedOrderReference `json:"data"`
		Included HotelOrder            `json:"included"`
		Warnings []Warning             `json:"warnings,omitempty"`
		Errors   []Error               `json:"errors,omitempty"`
	}

	// UpdatedOrderReference is the light {type,id} reference returned by Modify.
	UpdatedOrderReference struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	}

	// DeleteBookingResponse is the envelope returned by Delete (v2.2).
	DeleteBookingResponse struct {
		Included DeleteBookingResult `json:"included"`
		Warnings []Warning           `json:"warnings,omitempty"`
		Errors   []Error             `json:"errors,omitempty"`
	}

	// DeleteBookingResult holds the cancellation reference for a deleted booking.
	DeleteBookingResult struct {
		// CancellationNumber - provider cancellation reference for the deleted booking.
		CancellationNumber string `json:"cancellationNumber"`
	}
)
