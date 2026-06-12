package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	sdk "github.com/techpartners-asia/amadeus-hotel-integration"

	requestContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	requestHotelListHotelsDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/hotels"
)

// default test credentials; override with AMADEUS_CLIENT_ID / AMADEUS_CLIENT_SECRET.
const (
	defaultClientID     = "2Xt3NOH0ezVWVFp4MIVWjw9sGJSxxhQP"
	defaultClientSecret = "UljgNTvUNW5Vy7ge"
)

// newSDK builds an authenticated SDK, skipping the test when authentication
// cannot be established (e.g. no network or invalid credentials in CI).
func newSDK(t *testing.T) *sdk.SDK {
	t.Helper()

	id := envOr("AMADEUS_CLIENT_ID", defaultClientID)
	secret := envOr("AMADEUS_CLIENT_SECRET", defaultClientSecret)

	s, err := sdk.New(id, secret)
	if err != nil {
		t.Skipf("skipping: cannot authenticate with Amadeus: %v", err)
	}
	return s
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func TestHotelSearch(t *testing.T) {
	s := newSDK(t)

	hotels, err := s.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{
		CityCode: "PAR",
	})
	if err != nil {
		t.Fatalf("Error getting hotels: %v", err)
	}

	for _, hotel := range hotels {
		fmt.Println(hotel.HotelId)
		content, err := s.Content.GetByID(requestContentDTO.ContentByIDRequest{
			ID: hotel.HotelId,
		})
		if err != nil {
			t.Fatalf("Error getting content: %v", err)
		}
		b, _ := json.Marshal(content)
		fmt.Println(string(b))
		fmt.Println("--------------------------------")
	}
}

func TestHotelListByHotelIds(t *testing.T) {
	s := newSDK(t)

	// Resolve a couple of real hotel ids from a city search first.
	hotels, err := s.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{
		CityCode: "PAR",
	})
	if err != nil {
		t.Fatalf("Error getting hotels: %v", err)
	}
	if len(hotels) == 0 {
		t.Skip("no hotels returned for city to resolve ids")
	}

	ids := []string{hotels[0].HotelId}
	if len(hotels) > 1 {
		ids = append(ids, hotels[1].HotelId)
	}

	byIds, err := s.List.HotelListByHotelIds(requestHotelListHotelsDTO.HotelListByHotelsRequest{
		HotelIds: ids,
	})
	if err != nil {
		t.Fatalf("Error getting hotels by ids: %v", err)
	}

	for _, hotel := range byIds {
		fmt.Println(hotel.HotelId, hotel.Name)
	}
}

// TestBookingRetrieve exercises the Retrieve endpoint. It needs an existing
// hotel order id; provide one via AMADEUS_HOTEL_ORDER_ID to run it.
func TestBookingRetrieve(t *testing.T) {
	orderID := os.Getenv("AMADEUS_HOTEL_ORDER_ID")
	if orderID == "" {
		t.Skip("set AMADEUS_HOTEL_ORDER_ID to run the booking retrieve test")
	}

	s := newSDK(t)

	order, err := s.Booking.GetByID(orderID)
	if err != nil {
		t.Fatalf("Error retrieving order: %v", err)
	}

	b, _ := json.Marshal(order)
	fmt.Println(string(b))
}
