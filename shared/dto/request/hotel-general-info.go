package sharedRequestDTO

type (
// HotelGeneralInfoRequest struct {
// 	Latitude    float64  `json:"latitude"`    // The latitude of the searched geographical point expressed in decimal degrees. Example: 41.397158
// 	Longitude   float64  `json:"longitude"`   // The longitude of the searched geographical point expressed in decimal degrees. Example: 2.160873
// 	Radius      int      `json:"radius"`      // Maximum distance from the geographical coordinates expressed in defined units. The default unit is metric kilometer. Default value: 5
// 	RadiusUnit  string   `json:"radiusUnit"`  // Unit of measurement used to express the radius. It can be either metric kilometer or imperial mile. Available values: KM, MILE. Default value: KM
// 	ChainCodes  []string `json:"chainCodes"`  // Array of hotel chain codes. Each code is a string consisted of 2 capital alphabetic characters. The code is either a chain or a brand. The response includes all the hotels of the selected chain, or all the hotels of the sub chains of the selected brand.
// 	Amenities   []string `json:"amenities"`   // List of amenities. Available values: SWIMMING_POOL, SPA, FITNESS_CENTER, AIR_CONDITIONING, RESTAURANT, PARKING, PETS_ALLOWED, AIRPORT_SHUTTLE, BUSINESS_CENTER, DISABLED_FACILITIES, WIFI, MEETING_ROOMS, NO_KID_ALLOWED, TENNIS, GOLF, KITCHEN, ANIMAL_WATCHING, BABY-SITTING, BEACH, CASINO, JACUZZI, SAUNA, SOLARIUM, MASSAGE, VALET_PARKING, BAR or LOUNGE, KIDS_WELCOME, NO_PORN_FILMS, MINIBAR, TELEVISION, WI-FI_IN_ROOM, ROOM_SERVICE, GUARDED_PARKG, SERV_SPEC_MENU
// 	Ratings     []string `json:"ratings"`     // Hotel stars. Up to four values can be requested at the same time in a comma separated list. The response includes all the hotels with the requested rating and all hotels with an Amadeus self rating matching the requested rating. Available values: 1, 2, 3, 4, 5
// 	HotelSource string   `json:"hotelSource"` // Hotel source with values BEDBANK for aggregators, DIRECTCHAIN for GDS/Distribution and ALL for both. Available values: BEDBANK, DIRECTCHAIN, ALL. Default value: ALL
// }
)
