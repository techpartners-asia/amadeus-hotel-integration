package sdk

import (
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	usecaseBooking "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/usecase"
	usecaseContent "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/usecase"
	usecaseHotelList "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/usecase"
	usecasesHotelOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/usecases"
	"github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"
)

type SDK struct {
	Offers  usecasesHotelOffers.HotelOffersUsecase
	Content usecaseContent.ContentUsecase
	Booking usecaseBooking.BookingUsecase
	List    usecaseHotelList.HotelListUsecase
	// SearchCriteria lists the values Amadeus accepts in search filters
	// (amenities, star ratings, board types...). It is static data compiled into
	// the SDK, so its methods never call Amadeus and never fail. The equivalent
	// searchcriteria.All* functions need no SDK value or credentials.
	SearchCriteria searchcriteria.Catalog
}

// New authenticates with Amadeus and returns a ready-to-use SDK. It returns an
// error if authentication fails so callers can handle invalid credentials.
func New(id, secret string) (*SDK, error) {
	if err := amadeusIntegration.Init(id, secret); err != nil {
		return nil, err
	}

	return &SDK{
		Offers:         usecasesHotelOffers.NewHotelOffersUsecase(),
		Content:        usecaseContent.NewContentUsecase(),
		Booking:        usecaseBooking.NewBookingUsecase(),
		List:           usecaseHotelList.NewHotelListUsecase(),
		SearchCriteria: searchcriteria.NewCatalog(),
	}, nil
}
