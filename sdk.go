package sdk

import (
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	usecaseBooking "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/usecase"
	usecaseContent "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/usecase"
	usecaseHotelList "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/usecase"
	usecasesHotelOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/usecases"
)

type SDK struct {
	Offers  usecasesHotelOffers.HotelOffersUsecase
	Content usecaseContent.ContentUsecase
	Booking usecaseBooking.BookingUsecase
	List    usecaseHotelList.HotelListUsecase
}

func New(id, secret string) *SDK {

	amadeusIntegration.Init(id, secret)

	return &SDK{
		Offers:  usecasesHotelOffers.NewHotelOffersUsecase(),
		Content: usecaseContent.NewContentUsecase(),
		Booking: usecaseBooking.NewBookingUsecase(),
		List:    usecaseHotelList.NewHotelListUsecase(),
	}
}
