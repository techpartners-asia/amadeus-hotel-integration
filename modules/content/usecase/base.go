package usecaseContent

import (
	"errors"

	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	requestContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
	responseContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/response"
	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

	"resty.dev/v3"
)

// TODO : Get App from the Amadeus
type ContentUsecase interface {
	GetByID(request requestContentDTO.ContentByIDRequest) (*responseContentDTO.HotelContentResponse, error)
}

type contentUsecase struct {
	client *resty.Client
}

func NewContentUsecase() ContentUsecase {
	return &contentUsecase{
		client: amadeusIntegration.Client,
	}
}

func (c *contentUsecase) GetByID(request requestContentDTO.ContentByIDRequest) (*responseContentDTO.HotelContentResponse, error) {

	var response sharedResponseDTO.BaseResponse[responseContentDTO.HotelContentResponse]

	res, err := c.client.R().SetQueryParams(request.ToQueryParams()).SetResult(&response).Get(request.ID)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &response.Data, nil
}
