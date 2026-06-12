package usecaseBooking

import (
	"errors"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestBookingDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/request"
	responseBookingDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

// amadeusContentType is the media type required by the Amadeus booking endpoints
// for request bodies.
const amadeusContentType = "application/vnd.amadeus+json"

type BookingUsecase interface {
	// Create books a new hotel order (Hotel Booking v2.3).
	Create(request requestBookingDTO.HotelBookingRequest) (*responseBookingDTO.HotelOrder, error)
	// GetByReference retrieves a hotel order by its booking reference.
	GetByReference(reference string) (*responseBookingDTO.HotelOrder, error)
	// GetByID retrieves a hotel order by its hotelOrderId (Hotel Booking Retrieve v2.1).
	GetByID(hotelOrderId string) (*responseBookingDTO.HotelOrder, error)
	// Cancel cancels a single hotel booking within an order (Hotel Booking Manage v2.2).
	Cancel(hotelOrderId, hotelBookingId string) (*responseBookingDTO.HotelOrder, error)
	// Modify updates a single hotel booking within an order (Hotel Booking Manage v2.2).
	Modify(hotelOrderId, hotelBookingId string, request requestBookingDTO.UpdateHotelBookingRequest) (*responseBookingDTO.HotelBookingUpdateResponse, error)
	// Delete deletes a single hotel booking within an order (Hotel Booking Manage v2.2).
	Delete(hotelOrderId, hotelBookingId string) (*responseBookingDTO.DeleteBookingResult, error)
}

type bookingUsecase struct {
	client *resty.Client
}

func NewBookingUsecase() BookingUsecase {
	return &bookingUsecase{
		client: amadeusIntegration.NewClient(constants.BOOKING_BASE_URL),
	}
}

func (b *bookingUsecase) Create(request requestBookingDTO.HotelBookingRequest) (*responseBookingDTO.HotelOrder, error) {
	var response sharedResponseDTO.BaseResponse[responseBookingDTO.HotelOrder]

	res, err := b.client.R().
		SetHeader("Content-Type", amadeusContentType).
		SetBody(request).
		SetResult(&response).
		Post("/booking/hotel-orders")
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

	res, err := b.client.R().
		SetResult(&response).
		Get("/booking/hotel-orders/by-reference/" + reference)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}

func (b *bookingUsecase) GetByID(hotelOrderId string) (*responseBookingDTO.HotelOrder, error) {
	var response responseBookingDTO.HotelOrderResponse

	res, err := b.client.R().
		SetResult(&response).
		Get("/booking/hotel-orders/" + hotelOrderId)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Detail)
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}

func (b *bookingUsecase) Cancel(hotelOrderId, hotelBookingId string) (*responseBookingDTO.HotelOrder, error) {
	var response responseBookingDTO.HotelOrderResponse

	res, err := b.client.R().
		SetResult(&response).
		Post("/booking/hotel-orders/" + hotelOrderId + "/hotel-bookings/" + hotelBookingId + "/cancel")
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Detail)
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}

func (b *bookingUsecase) Modify(hotelOrderId, hotelBookingId string, request requestBookingDTO.UpdateHotelBookingRequest) (*responseBookingDTO.HotelBookingUpdateResponse, error) {
	var response responseBookingDTO.HotelBookingUpdateResponse

	res, err := b.client.R().
		SetHeader("Content-Type", amadeusContentType).
		SetBody(request).
		SetResult(&response).
		Patch("/booking/hotel-orders/" + hotelOrderId + "/hotel-bookings/" + hotelBookingId)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Detail)
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response, nil
}

func (b *bookingUsecase) Delete(hotelOrderId, hotelBookingId string) (*responseBookingDTO.DeleteBookingResult, error) {
	var response responseBookingDTO.DeleteBookingResponse

	res, err := b.client.R().
		SetResult(&response).
		Delete("/booking/hotel-orders/" + hotelOrderId + "/hotel-bookings/" + hotelBookingId)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, errors.New(response.Errors[0].Detail)
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Included, nil
}
