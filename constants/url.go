package constants

// All endpoints target the Amadeus Enterprise ("travel") host, matching the
// provided swagger specifications. These APIs require Enterprise credentials;
// self-service (api.amadeus.com) credentials are rejected by this host.
//
// To target production, change the TRAVEL_*_URL host prefixes from
// "test.travel.api.amadeus.com" to "travel.api.amadeus.com".
const (
	// Host roots (versioned) on the Enterprise travel gateway.
	TRAVEL_BASE_V1_URL = "https://test.travel.api.amadeus.com/v1"
	TRAVEL_BASE_V2_URL = "https://test.travel.api.amadeus.com/v2"
	TRAVEL_BASE_V3_URL = "https://test.travel.api.amadeus.com/v3"

	// AUTH_BASE_URL is the OAuth2 (client credentials) host. It must match the
	// host family of the API endpoints so the issued token is accepted.
	AUTH_BASE_URL = TRAVEL_BASE_V1_URL

	// Per-module base URLs.
	//
	// OFFERS_BASE_URL  -> Hotel Search v3.5  (/shopping/hotel-offers, /{offerId})
	// CONTENT_BASE_URL -> Hotel Content v3.1 (/reference-data/locations/by-hotel)
	// LIST_BASE_URL    -> Hotel List v1.2    (/by-city, /by-geocode, /by-hotels)
	// BOOKING_BASE_URL -> Hotel Booking v2.x (create, retrieve, manage)
	OFFERS_BASE_URL  = TRAVEL_BASE_V3_URL + "/shopping/hotel-offers"
	CONTENT_BASE_URL = TRAVEL_BASE_V3_URL
	LIST_BASE_URL    = TRAVEL_BASE_V1_URL + "/reference-data/locations/hotels"
	BOOKING_BASE_URL = TRAVEL_BASE_V2_URL
)
