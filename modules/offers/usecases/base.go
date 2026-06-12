package usecasesHotelOffers

import (
	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestHotelOffersDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
	responseHotelOffersDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

type HotelOffersUsecase interface {
	List(request requestHotelOffersDTO.HotelOffersListRequest) ([]responseHotelOffersDTO.OffersResponse, error)
	GetByID(request requestHotelOffersDTO.HotelOffersByIDRequest) (*responseHotelOffersDTO.OffersResponse, error)
}

type hotelOffersUsecase struct {
	client *resty.Client
}

func NewHotelOffersUsecase() HotelOffersUsecase {
	return &hotelOffersUsecase{
		client: amadeusIntegration.NewClient(constants.OFFERS_BASE_URL),
	}
}

func (h *hotelOffersUsecase) List(request requestHotelOffersDTO.HotelOffersListRequest) ([]responseHotelOffersDTO.OffersResponse, error) {
	var response sharedResponseDTO.BaseResponse[[]responseHotelOffersDTO.OffersResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&response).Get("")
	if err != nil {
		return nil, err
	}

	if apiErr := sharedResponseDTO.ErrorFromResponse(res.StatusCode(), res.IsError(), res.String()); apiErr != nil {
		return nil, apiErr
	}

	return response.Data, nil
}

func (h *hotelOffersUsecase) GetByID(request requestHotelOffersDTO.HotelOffersByIDRequest) (*responseHotelOffersDTO.OffersResponse, error) {

	var response sharedResponseDTO.BaseResponse[responseHotelOffersDTO.OffersResponse]

	res, err := h.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&response).Get(request.OfferID)
	if err != nil {
		return nil, err
	}

	if apiErr := sharedResponseDTO.ErrorFromResponse(res.StatusCode(), res.IsError(), res.String()); apiErr != nil {
		return nil, apiErr
	}

	return &response.Data, nil

}
