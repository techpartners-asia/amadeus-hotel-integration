package content

import (
	"context"
	"net/url"
	"strings"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto/contentdto"
)

// contentPath is the Hotel Content (v3.1) endpoint.
const contentPath = "/v3/reference-data/locations/by-hotel"

// Service fetches hotel descriptions. Obtain one from the SDK client:
//
//	client, _ := sdk.New(cfg)
//	hotel, err := client.Content.Get(ctx, content.Query{HotelID: "HLPAR266"})
//
// Content does not change per stay, so it is worth caching. An offer is not.
type Service interface {
	// Get fetches everything Amadeus publishes about one property.
	Get(ctx context.Context, query Query) (*Hotel, error)
}

// Query selects a property and how much of its content to return.
type Query struct {
	// HotelID is the Amadeus 8-character property code. Required.
	HotelID string
	// View sets how much detail to return. The SDK sends FULL when this is
	// unset, because Amadeus's own default returns only the basic block - and
	// a caller asking for content almost never wants just the name.
	View codes.ContentView
	// Fields restricts the response to named content blocks, e.g. "hotel",
	// "rooms", "facilities". Leave it empty for everything the View allows.
	Fields []string
	// Lang requests text in a language, e.g. "FR". Amadeus falls back to
	// English where it holds no translation.
	Lang string
}

func (q Query) validate() error {
	var errs apierr.ValidationErrors

	if strings.TrimSpace(q.HotelID) == "" {
		errs = errs.Append("HotelID", "is required")
	} else if len(q.HotelID) != 8 {
		errs = append(errs, apierr.Invalidf("HotelID",
			"%q is not a property code; they are exactly 8 characters", q.HotelID))
	}
	if q.View != "" && !q.View.IsValid() {
		errs = append(errs, apierr.Invalidf("View", "%q is not a known content view", q.View))
	}

	return errs.OrNil()
}

func (q Query) params() url.Values {
	values := url.Values{"hotelID": {q.HotelID}}

	// Amadeus defaults to a light view returning only the basic block. Ask for
	// FULL unless the caller wanted something narrower, so rooms, facilities,
	// policies, awards and points of interest come back populated.
	view := q.View
	if view == "" {
		view = codes.ContentViewFull
	}
	values.Set("view", string(view))

	if len(q.Fields) > 0 {
		values.Set("fields", strings.Join(q.Fields, ","))
	}
	if q.Lang != "" {
		values.Set("lang", q.Lang)
	}

	return values
}

type service struct {
	client *amadeus.Client
}

// NewService returns the content service backed by client.
func NewService(client *amadeus.Client) Service {
	return &service{client: client}
}

func (s *service) Get(ctx context.Context, query Query) (*Hotel, error) {
	if err := query.validate(); err != nil {
		return nil, err
	}

	envelope, err := amadeus.Do[contentdto.HotelContentResponse](ctx, s.client, amadeus.Request{
		Path:  contentPath,
		Query: query.params(),
	})
	if err != nil {
		return nil, err
	}

	hotel := mapHotel(envelope.Data)
	return &hotel, nil
}
