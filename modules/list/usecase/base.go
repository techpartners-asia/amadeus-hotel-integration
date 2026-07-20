package usecaseHotelList

import (
	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
	requestHotelListGeocodeDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/geocode"
	requestHotelListHotelsDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/hotels"
	responseHotelListDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

type HotelListUsecase interface {
	HotelListByGeocode(request requestHotelListGeocodeDTO.HotelListByGeocodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error)
	HotelListByCityCode(request requestHotelListCityDTO.HotelListByCityCodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error)
	HotelListByHotelIds(request requestHotelListHotelsDTO.HotelListByHotelsRequest) ([]responseHotelListDTO.GeneralInfoResponse, error)
}

type hotelListUsecase struct {
	client *resty.Client
}

func NewHotelListUsecase() HotelListUsecase {
	return &hotelListUsecase{
		client: amadeusIntegration.NewClient(constants.LIST_BASE_URL),
	}
}

// * : Hotel List By Geocode
func (h *hotelListUsecase) HotelListByGeocode(request requestHotelListGeocodeDTO.HotelListByGeocodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error) {

	var responses sharedResponseDTO.BaseResponse[[]responseHotelListDTO.GeneralInfoResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&responses).Get("/by-geocode")
	if err != nil {
		return nil, err
	}

	if apiErr := sharedResponseDTO.ErrorFromResponse(res.StatusCode(), res.IsError(), res.String()); apiErr != nil {
		return nil, apiErr
	}

	return responses.Data, nil
}

// * : Hotel List By City Code
func (h *hotelListUsecase) HotelListByCityCode(request requestHotelListCityDTO.HotelListByCityCodeRequest) ([]responseHotelListDTO.GeneralInfoResponse, error) {

	var responses sharedResponseDTO.BaseResponse[[]responseHotelListDTO.GeneralInfoResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&responses).Get("/by-city")
	if err != nil {
		return nil, err
	}

	if apiErr := sharedResponseDTO.ErrorFromResponse(res.StatusCode(), res.IsError(), res.String()); apiErr != nil {
		return nil, apiErr
	}

	return responses.Data, nil
}

// * : Hotel List By Hotel Ids
func (h *hotelListUsecase) HotelListByHotelIds(request requestHotelListHotelsDTO.HotelListByHotelsRequest) ([]responseHotelListDTO.GeneralInfoResponse, error) {

	var responses sharedResponseDTO.BaseResponse[[]responseHotelListDTO.GeneralInfoResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&responses).Get("/by-hotels")
	if err != nil {
		return nil, err
	}

	if apiErr := sharedResponseDTO.ErrorFromResponse(res.StatusCode(), res.IsError(), res.String()); apiErr != nil {
		return nil, apiErr
	}

	return responses.Data, nil
}
