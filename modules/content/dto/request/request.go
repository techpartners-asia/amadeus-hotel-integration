package requestContentDTO

type (
	ContentByIDRequest struct {
		ID string `json:"id" required:"true"`
	}
)

func (r *ContentByIDRequest) ToQueryParams() map[string]string {
	return map[string]string{
		"hotelId": r.ID,
	}
}
