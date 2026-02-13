package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	sdk "github.com/techpartners-asia/amadeus-hotel-integration"

	requestContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
)

func TestHotelSearch(t *testing.T) {
	sdk := sdk.New("2Xt3NOH0ezVWVFp4MIVWjw9sGJSxxhQP", "UljgNTvUNW5Vy7ge")

	offers, err := sdk.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{
		CityCode: "PAR",
	})
	if err != nil {
		t.Fatalf("Error getting offers: %v", err)
	}

	for _, offer := range offers {
		fmt.Println(offer.HotelId)
		content, err := sdk.Content.GetByID(requestContentDTO.ContentByIDRequest{
			ID: offer.HotelId,
		})
		if err != nil {
			t.Fatalf("Error getting content: %v", err)
		}
		b, _ := json.Marshal(content)
		fmt.Println(string(b))
		fmt.Println("--------------------------------")
	}

}
