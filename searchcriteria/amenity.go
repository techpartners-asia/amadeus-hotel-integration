package searchcriteria

// Amenity is an amenity filter code accepted by the Hotel List API (by-city and
// by-geocode) in the `amenities` query parameter.
//
// The codes are Amadeus' own and are not internally consistent: most use
// underscores, but WI-FI_IN_ROOM uses a hyphen, and GUARDED_PARKG and
// SERV_SPEC_MENU are abbreviated. They are reproduced verbatim because Amadeus
// matches on the exact string.
//
// Two codes deliberately differ from the published Amadeus documentation, which
// lists them as "BABY-SITTING" and "BAR or LOUNGE". Both spellings are rejected
// by the live API with error 7211 INVALID FACILITY CODE; the values below are
// the ones it actually accepts. TestAmenityCodesAcceptedByAmadeus in
// tests/amenity_probe_test.go verifies this against the real endpoint.
type Amenity string

const (
	AmenitySwimmingPool       Amenity = "SWIMMING_POOL"
	AmenitySpa                Amenity = "SPA"
	AmenityFitnessCenter      Amenity = "FITNESS_CENTER"
	AmenityAirConditioning    Amenity = "AIR_CONDITIONING"
	AmenityRestaurant         Amenity = "RESTAURANT"
	AmenityParking            Amenity = "PARKING"
	AmenityPetsAllowed        Amenity = "PETS_ALLOWED"
	AmenityAirportShuttle     Amenity = "AIRPORT_SHUTTLE"
	AmenityBusinessCenter     Amenity = "BUSINESS_CENTER"
	AmenityDisabledFacilities Amenity = "DISABLED_FACILITIES"
	AmenityWifi               Amenity = "WIFI"
	AmenityMeetingRooms       Amenity = "MEETING_ROOMS"
	AmenityNoKidAllowed       Amenity = "NO_KID_ALLOWED"
	AmenityTennis             Amenity = "TENNIS"
	AmenityGolf               Amenity = "GOLF"
	AmenityKitchen            Amenity = "KITCHEN"
	AmenityAnimalWatching     Amenity = "ANIMAL_WATCHING"
	// AmenityBabySitting is documented by Amadeus as "BABY-SITTING", which the
	// API rejects. The underscore form is what works.
	AmenityBabySitting  Amenity = "BABY_SITTING"
	AmenityBeach        Amenity = "BEACH"
	AmenityCasino       Amenity = "CASINO"
	AmenityJacuzzi      Amenity = "JACUZZI"
	AmenitySauna        Amenity = "SAUNA"
	AmenitySolarium     Amenity = "SOLARIUM"
	AmenityMassage      Amenity = "MASSAGE"
	AmenityValetParking Amenity = "VALET_PARKING"
	// AmenityBarOrLounge is documented by Amadeus as "BAR or LOUNGE", which the
	// API rejects, as does BAR_OR_LOUNGE. BAR_LOUNGE is what works.
	AmenityBarOrLounge    Amenity = "BAR_LOUNGE"
	AmenityKidsWelcome    Amenity = "KIDS_WELCOME"
	AmenityNoPornFilms    Amenity = "NO_PORN_FILMS"
	AmenityMinibar        Amenity = "MINIBAR"
	AmenityTelevision     Amenity = "TELEVISION"
	AmenityWifiInRoom     Amenity = "WI-FI_IN_ROOM"
	AmenityRoomService    Amenity = "ROOM_SERVICE"
	AmenityGuardedParking Amenity = "GUARDED_PARKG"
	AmenityServSpecMenu   Amenity = "SERV_SPEC_MENU"
)

// amenityCatalog lists the codes in the order Amadeus documents them.
//
// Labels are spelled out rather than derived from the code because derivation
// mangles the abbreviated ones: GUARDED_PARKG and SERV_SPEC_MENU would render
// as "Guarded Parkg" and "Serv Spec Menu".
var amenityCatalog = []entry[Amenity]{
	{AmenitySwimmingPool, "Swimming Pool"},
	{AmenitySpa, "Spa"},
	{AmenityFitnessCenter, "Fitness Center"},
	{AmenityAirConditioning, "Air Conditioning"},
	{AmenityRestaurant, "Restaurant"},
	{AmenityParking, "Parking"},
	{AmenityPetsAllowed, "Pets Allowed"},
	{AmenityAirportShuttle, "Airport Shuttle"},
	{AmenityBusinessCenter, "Business Center"},
	{AmenityDisabledFacilities, "Disabled Facilities"},
	{AmenityWifi, "Wi-Fi"},
	{AmenityMeetingRooms, "Meeting Rooms"},
	{AmenityNoKidAllowed, "No Children Allowed"},
	{AmenityTennis, "Tennis"},
	{AmenityGolf, "Golf"},
	{AmenityKitchen, "Kitchen"},
	{AmenityAnimalWatching, "Animal Watching"},
	{AmenityBabySitting, "Baby-Sitting"},
	{AmenityBeach, "Beach"},
	{AmenityCasino, "Casino"},
	{AmenityJacuzzi, "Jacuzzi"},
	{AmenitySauna, "Sauna"},
	{AmenitySolarium, "Solarium"},
	{AmenityMassage, "Massage"},
	{AmenityValetParking, "Valet Parking"},
	{AmenityBarOrLounge, "Bar or Lounge"},
	{AmenityKidsWelcome, "Kids Welcome"},
	{AmenityNoPornFilms, "No Adult Films"},
	{AmenityMinibar, "Minibar"},
	{AmenityTelevision, "Television"},
	{AmenityWifiInRoom, "Wi-Fi in Room"},
	{AmenityRoomService, "Room Service"},
	{AmenityGuardedParking, "Guarded Parking"},
	{AmenityServSpecMenu, "Special Menu Service"},
}

// AllAmenities returns every amenity code Amadeus accepts, in a stable order.
func AllAmenities() []Amenity { return codes(amenityCatalog) }

// Label returns a human-readable name for a, or "" when a is not a known code.
func (a Amenity) Label() string { return labelOf(amenityCatalog, a) }

// IsValid reports whether a is a code Amadeus accepts.
func (a Amenity) IsValid() bool { return isValid(amenityCatalog, a) }
