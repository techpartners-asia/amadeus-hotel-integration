package usecaseHotelList

import (
	"errors"
	"fmt"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	requestHotelListGeocodeDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/geocode"
	responseHotelListDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

type HotelListUsecase interface {
	HotelListByGeocode(request requestHotelListGeocodeDTO.HotelListByGeocodeRequest) ([]responseHotelListDTO.HotelListResponse, error)
	HotelListByCityCode(request requestHotelListCityDTO.HotelListByCityCodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error)
}

type hotelListUsecase struct {
	client *resty.Client
}

func NewHotelListUsecase() HotelListUsecase {

	baseClient := amadeusIntegration.Client

	baseClient.SetBaseURL(constants.BASE_V1_URL + "/reference-data/locations/hotels")

	return &hotelListUsecase{
		client: baseClient,
	}
}

// * : Hotel List By Geocode
func (h *hotelListUsecase) HotelListByGeocode(request requestHotelListGeocodeDTO.HotelListByGeocodeRequest) ([]responseHotelListDTO.HotelListResponse, error) {

	var responses []responseHotelListDTO.HotelListResponse

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&responses).Get("/by-geocode")
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return responses, nil
}

// * : Hotel List By City Code
func (h *hotelListUsecase) HotelListByCityCode(request requestHotelListCityDTO.HotelListByCityCodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error) {

	var responses sharedResponseDTO.BaseResponse[[]responseHotelListDTO.GeneralInfoResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&responses).Get("/by-city")
	if err != nil {
		return nil, err
	}

	fmt.Println(res.RawResponse.Body)

	if len(responses.Errors) > 0 {
		return nil, errors.New(responses.Errors[0].Detail)
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return responses.Data, nil
}
