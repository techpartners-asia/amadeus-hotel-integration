package sharedResponseDTO

type (
	BaseResponse[T any] struct {
		Data   T               `json:"data"`
		Errors []ErrorResponse `json:"errors"`
		Meta   MetaResponse    `json:"meta"`
	}

	LinkResponse struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		Self string `json:"self"`
	}

	MetaResponse struct {
		Count int `json:"count"`
		// Links []LinkResponse `json:"links"`
	}

	ErrorResponse struct {
		Status int    `json:"status"`
		Code   int    `json:"code"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Source struct {
			Parameter string `json:"parameter"`
		} `json:"source"`
	}

	GeoCodeResponse struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)
