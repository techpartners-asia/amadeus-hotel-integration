package amadeusIntegration

import (
	"errors"
	"fmt"
	"log"

	"resty.dev/v3"
)

var (
	Client *resty.Client
)

func Init(id, secret string) {

	authResponse, err := Auth(id, secret)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	fmt.Println(authResponse.AccessToken)

	Client = resty.New().SetHeader("Accept", "application/json").SetHeader("Authorization", "Bearer "+authResponse.AccessToken)
}

func Auth(id, secret string) (*AuthResponse, error) {
	client := resty.New().SetBaseURL("https://test.api.amadeus.com/v1").SetHeader("Accept", "application/json")

	var authResponse AuthResponse

	res, err := client.R().SetFormData(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     id,
		"client_secret": secret,
	}).SetResult(&authResponse).Post("/security/oauth2/token")

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	return &authResponse, nil
}
