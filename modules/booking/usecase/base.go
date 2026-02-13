package usecaseBooking

import (
	"errors"

	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestBookingDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/request"
	responseBookingDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

type BookingUsecase interface {
	Create(request requestBookingDTO.HotelBookingRequest) (*responseBookingDTO.HotelOrder, error)
	GetByReference(reference string) (*responseBookingDTO.HotelOrder, error)
}

type bookingUsecase struct {
	client *resty.Client
}

func NewBookingUsecase() BookingUsecase {
	return &bookingUsecase{
		client: amadeusIntegration.Client,
	}
}

func (b *bookingUsecase) Create(request requestBookingDTO.HotelBookingRequest) (*responseBookingDTO.HotelOrder, error) {
	var response sharedResponseDTO.BaseResponse[responseBookingDTO.HotelOrder]

	res, err := b.client.R().SetBody(request).SetResult(&response).Post("/booking/hotel-orders")
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}

func (b *bookingUsecase) GetByReference(reference string) (*responseBookingDTO.HotelOrder, error) {
	var response sharedResponseDTO.BaseResponse[responseBookingDTO.HotelOrder]

	res, err := b.client.R().SetResult(&response).Get("/booking/hotel-orders/by-reference/" + reference)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}
