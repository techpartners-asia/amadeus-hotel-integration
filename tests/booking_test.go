package tests

import (
	"os"
	"testing"

	requestBookingDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/request"
	requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
)

// Booking tests create real hotel orders in the Amadeus test environment, so
// they only run when explicitly opted in:
//
//	AMADEUS_ALLOW_BOOKING=1 go test ./tests/ -run TestBooking
//
// Never point these at production credentials.

// testCard is the Amadeus-published sandbox test card.
const (
	testCardVendor = "VI"
	testCardNumber = "4151289722471370"
	testCardExpiry = "2030-08"
	testCardHolder = "BOB SMITH"
)

func requireBookingOptIn(t *testing.T) {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	if os.Getenv("AMADEUS_ALLOW_BOOKING") != "1" {
		t.Skip("set AMADEUS_ALLOW_BOOKING=1 to run tests that create real sandbox bookings")
	}
}

// bookableOfferID finds an offer that can actually be booked in the sandbox.
func bookableOfferID(t *testing.T) string {
	t.Helper()
	s := newSDK(t)
	checkIn, checkOut := stayDates()

	for _, id := range sandboxOfferHotels {
		offers, err := s.Offers.List(requestOffers.HotelOffersListRequest{
			HotelIDs:     []string{id},
			CheckInDate:  checkIn,
			CheckOutDate: checkOut,
			Adults:       1,
		})
		if err != nil {
			continue
		}
		for _, o := range offers {
			if len(o.Offers) > 0 {
				return o.Offers[0].ID
			}
		}
	}
	t.Skip("no bookable offer available in sandbox")
	return ""
}

// minimalBooking builds the smallest valid hotel-order request. It deliberately
// leaves every optional field unset, which is the case that used to fail: the
// request DTOs serialised empty optional fields and Amadeus rejected them one at
// a time ("Wrong parameter: hotelLoyaltyId").
func minimalBooking(offerID string) requestBookingDTO.HotelBookingRequest {
	return requestBookingDTO.HotelBookingRequest{
		Data: requestBookingDTO.BookingData{
			Type: "hotel-order",
			Guests: []requestBookingDTO.Guest{{
				Tid:       1,
				Title:     "MR",
				FirstName: "BOB",
				LastName:  "SMITH",
				Phone:     "+33679278416",
				Email:     "bob.smith@email.com",
			}},
			RoomAssociations: []requestBookingDTO.RoomAssociation{{
				HotelOfferId:    offerID,
				GuestReferences: []requestBookingDTO.GuestReference{{GuestReference: "1"}},
			}},
			Payment: requestBookingDTO.Payment{
				Method: "CREDIT_CARD",
				PaymentCard: &requestBookingDTO.PaymentCard{
					PaymentCardInfo: requestBookingDTO.PaymentCardInfo{
						VendorCode: testCardVendor,
						CardNumber: testCardNumber,
						ExpiryDate: testCardExpiry,
						HolderName: testCardHolder,
					},
				},
			},
			TravelAgent: requestBookingDTO.TravelAgent{
				Contact: requestBookingDTO.Contact{Email: "bob.smith@email.com"},
			},
		},
	}
}

// TestBookingCreateMinimal verifies a booking with no optional fields set is
// accepted, and that the created order decodes with the identifiers a caller
// needs to retrieve or cancel it later.
func TestBookingCreateMinimal(t *testing.T) {
	requireBookingOptIn(t)
	s := newSDK(t)

	order, err := s.Booking.Create(minimalBooking(bookableOfferID(t)))
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	if order.Id == "" {
		t.Error("order has no id; it cannot be retrieved or cancelled later")
	}
	if order.Type != "hotel-order" {
		t.Errorf("type = %q, want %q", order.Type, "hotel-order")
	}
	if len(order.HotelBookings) == 0 {
		t.Error("order has no hotelBookings")
	}
	for _, b := range order.HotelBookings {
		if b.Id == "" {
			t.Error("hotel booking has no id; it cannot be cancelled")
		}
	}
	if len(order.AssociatedRecords) == 0 {
		t.Error("order has no associatedRecords (PNR reference)")
	}
	t.Logf("created order %s with %d booking(s)", order.Id, len(order.HotelBookings))
}

// TestBookingCreateThenRetrieve round-trips a created order through GetByID.
func TestBookingCreateThenRetrieve(t *testing.T) {
	requireBookingOptIn(t)
	s := newSDK(t)

	created, err := s.Booking.Create(minimalBooking(bookableOfferID(t)))
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	got, err := s.Booking.GetByID(created.Id)
	if err != nil {
		t.Fatalf("GetByID(%s): %v", created.Id, err)
	}
	if got.Id != created.Id {
		t.Errorf("retrieved order id = %q, want %q", got.Id, created.Id)
	}
	if len(got.HotelBookings) != len(created.HotelBookings) {
		t.Errorf("retrieved %d bookings, created %d",
			len(got.HotelBookings), len(created.HotelBookings))
	}
}
