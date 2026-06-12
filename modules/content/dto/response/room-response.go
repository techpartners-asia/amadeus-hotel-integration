package responseContentDTO

import sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"

type HotelRoomClassification string

const (
	PhysicallyChallengedRooms       HotelRoomClassification = "PHYSICALLY_CHALLENGED_ROOMS"
	NonsmokingRooms                 HotelRoomClassification = "NONSMOKING_ROOMS"
	Suites                          HotelRoomClassification = "SUITES"
	BungalowsAndVillas              HotelRoomClassification = "BUNGALOWS_AND_VILLAS"
	Floors                          HotelRoomClassification = "FLOORS"
	ExecutiveFloor                  HotelRoomClassification = "EXECUTIVE_FLOOR"
	RoomsThatWork                   HotelRoomClassification = "ROOMS_THAT_WORK"
	AvailableRooms                  HotelRoomClassification = "AVAILABLE_ROOMS"
	AvailableSuites                 HotelRoomClassification = "AVAILABLE_SUITES"
	DoubleBedrooms                  HotelRoomClassification = "DOUBLE_BEDROOMS"
	KingBedrooms                    HotelRoomClassification = "KING_BEDROOMS"
	TotalRooms                      HotelRoomClassification = "TOTAL_ROOMS"
	Apartments                      HotelRoomClassification = "APARTMENTS"
	QueenBedrooms                   HotelRoomClassification = "QUEEN_BEDROOMS"
	Penthouses                      HotelRoomClassification = "PENTHOUSES"
	Studios                         HotelRoomClassification = "STUDIOS"
	FirstFloorRooms                 HotelRoomClassification = "FIRST_FLOOR_ROOMS"
	SmokingRooms                    HotelRoomClassification = "SMOKING_ROOMS"
	TwinBedrooms                    HotelRoomClassification = "TWIN_BEDROOMS"
	DriveUpRooms                    HotelRoomClassification = "DRIVE_UP_ROOMS"
	RoomsWithInternetAccess         HotelRoomClassification = "ROOMS_WITH_INTERNET_ACCESS"
	FreestandingUnits               HotelRoomClassification = "FREESTANDING_UNITS"
	AirConditionedGuestRooms        HotelRoomClassification = "AIR_CONDITIONED_GUEST_ROOMS"
	ConciergeLevels                 HotelRoomClassification = "CONCIERGE_LEVELS"
	Condos                          HotelRoomClassification = "CONDOS"
	ClubLevels                      HotelRoomClassification = "CLUB_LEVELS"
	TotalAvailableRoomsNadSuites    HotelRoomClassification = "TOTAL_AVAILABLE_ROOMS_NAD_SUITES"
	TotalRoomsAndSuites             HotelRoomClassification = "TOTAL_ROOMS_AND_SUITES"
	EmployeesOnProperty             HotelRoomClassification = "EMPLOYEES_ON_PROPERTY"
	EmployeesWorkingForProperty     HotelRoomClassification = "EMPLOYEES_WORKING_FOR_PROPERTY"
	SeparateFloorsForWomen          HotelRoomClassification = "SEPARATE_FLOORS_FOR_WOMEN"
	Buildings                       HotelRoomClassification = "BUILDINGS"
	AccommodationsWithBalcony       HotelRoomClassification = "ACCOMMODATIONS_WITH_BALCONY"
	AdjoiningRoomsOrSuites          HotelRoomClassification = "ADJOINING_ROOMS_OR_SUITES"
	ConnectingRoomsOrSuites         HotelRoomClassification = "CONNECTING_ROOMS_OR_SUITES"
	FamilyOrOversizedAccommodations HotelRoomClassification = "FAMILY_OR_OVERSIZED_ACCOMMODATIONS"
	SingleBeddedAccommodations      HotelRoomClassification = "SINGLE_BEDDED_ACCOMMODATIONS"
	Cabin                           HotelRoomClassification = "CABIN"
	Cottage                         HotelRoomClassification = "COTTAGE"
	Loft                            HotelRoomClassification = "LOFT"
	Parlour                         HotelRoomClassification = "PARLOUR"
	Room                            HotelRoomClassification = "ROOM"
	Lanai                           HotelRoomClassification = "LANAI"
	Bungalow                        HotelRoomClassification = "BUNGALOW"
	Villa                           HotelRoomClassification = "VILLA"
	Efficiency                      HotelRoomClassification = "EFFICIENCY"
	AllRoomsNonSmoking              HotelRoomClassification = "ALL_ROOMS_NON_SMOKING"
	DoubleDoubleRooms               HotelRoomClassification = "DOUBLE_DOUBLE_ROOMS"
	KingKingBedrooms                HotelRoomClassification = "KING_KING_BEDROOMS"
	QueenQueenBedrooms              HotelRoomClassification = "QUEEN_QUEEN_BEDROOMS"
	TwinTwinBedrooms                HotelRoomClassification = "TWIN_TWIN_BEDROOMS"
	ApartmentFor1                   HotelRoomClassification = "APARTMENT_FOR_1"
	ApartmentFor2                   HotelRoomClassification = "APARTMENT_FOR_2"
	ApartmentFor3                   HotelRoomClassification = "APARTMENT_FOR_3"
	ApartmentFor4                   HotelRoomClassification = "APARTMENT_FOR_4"
	ApartmentFor6                   HotelRoomClassification = "APARTMENT_FOR_6"
	Cabin1Room                      HotelRoomClassification = "1_ROOM_CABIN"
	Cabin1Bedroom                   HotelRoomClassification = "1_BEDROOM_CABIN"
	Cabin2Bedroom                   HotelRoomClassification = "2_BEDROOM_CABIN"
	JuniorSuite                     HotelRoomClassification = "JUNIOR_SUITE"
	JacuzziSuite                    HotelRoomClassification = "JACUZZI_SUITE"
	RunOfTheHouse                   HotelRoomClassification = "RUN_OF_THE_HOUSE"
	LargeSuite                      HotelRoomClassification = "LARGE_SUITE"
	Bedroom1                        HotelRoomClassification = "1_BEDROOM"
	Bedroom2                        HotelRoomClassification = "2_BEDROOMS"
	Bedroom3                        HotelRoomClassification = "3_BEDROOMS"
	VillaFor1                       HotelRoomClassification = "VILLA_FOR_1"
	VillaFor2                       HotelRoomClassification = "VILLA_FOR_2"
	VillaFor3                       HotelRoomClassification = "VILLA_FOR_3"
	VillaFor6                       HotelRoomClassification = "VILLA_FOR_6"
	VillaFor8                       HotelRoomClassification = "VILLA_FOR_8"
	SingleWithPullout               HotelRoomClassification = "SINGLE_WITH_PULLOUT"
	BusinessPlan                    HotelRoomClassification = "BUSINESS_PLAN"
	BusinessClass                   HotelRoomClassification = "BUSINESS_CLASS"
	Classic                         HotelRoomClassification = "CLASSIC"
	Comfort                         HotelRoomClassification = "COMFORT"
	Deluxe                          HotelRoomClassification = "DELUXE"
	DeluxeSuite                     HotelRoomClassification = "DELUXE_SUITE"
	Economy                         HotelRoomClassification = "ECONOMY"
	Luxury                          HotelRoomClassification = "LUXURY"
	Premier                         HotelRoomClassification = "PREMIER"
	Standard                        HotelRoomClassification = "STANDARD"
	Superior                        HotelRoomClassification = "SUPERIOR"
	Dormitory                       HotelRoomClassification = "DORMITORY"
	Elevator                        HotelRoomClassification = "ELEVATOR"
)

type HotelRoomCategory string

const (
	SegAllSuite                   HotelRoomCategory = "ALL_SUITE"
	SegBudget                     HotelRoomCategory = "BUDGET"
	SegCorporateBusinessTransient HotelRoomCategory = "CORPORATE_BUSINESS_TRANSIENT"
	SegDeluxe                     HotelRoomCategory = "DELUXE"
	SegEconomy                    HotelRoomCategory = "ECONOMY"
	SegExtendedStay               HotelRoomCategory = "EXTENDED_STAY"
	SegFirstClass                 HotelRoomCategory = "FIRST_CLASS"
	SegLuxury                     HotelRoomCategory = "LUXURY"
	SegMeetingOrConvention        HotelRoomCategory = "MEETING_OR_CONVENTION"
	SegModerate                   HotelRoomCategory = "MODERATE"
	SegResidentialApartment       HotelRoomCategory = "RESIDENTIAL_APARTMENT"
	SegResort                     HotelRoomCategory = "RESORT"
	SegTourist                    HotelRoomCategory = "TOURIST"
	SegUpscale                    HotelRoomCategory = "UPSCALE"
	SegEfficiency                 HotelRoomCategory = "EFFICIENCY"
	SegStandard                   HotelRoomCategory = "STANDARD"
	SegMidscale                   HotelRoomCategory = "MIDSCALE"
	SegQuality                    HotelRoomCategory = "QUALITY"
	SegUnknown                    HotelRoomCategory = "UNKNOWN"
	SegMidscaleWithoutFAndB       HotelRoomCategory = "MIDSCALE_WITHOUT_F_AND_B"
	SegUpperUpscale               HotelRoomCategory = "UPPER_UPSCALE"
)

type BedType string

const (
	BedDouble                BedType = "DOUBLE"
	BedFuton                 BedType = "FUTON"
	BedKing                  BedType = "KING"
	BedMurphyBed             BedType = "MURPHY_BED"
	BedQueen                 BedType = "QUEEN"
	BedSofaBed               BedType = "SOFA_BED"
	BedTatamiMats            BedType = "TATAMI_MATS"
	BedTwin                  BedType = "TWIN"
	BedSingle                BedType = "SINGLE"
	BedPullOut               BedType = "PULL_OUT"
	BedWaterBed              BedType = "WATER_BED"
	BedUnknownOrOtherBedType BedType = "UNKNOWN_OR_OTHER_BED_TYPE"
	BedSuperKing             BedType = "SUPER_KING"
	BedDormBed               BedType = "DORM_BED"
	BedFull                  BedType = "FULL"
	BedRunOfTheHouse         BedType = "RUN_OF_THE_HOUSE"
)

type RoomViewType string

const (
	ViewAirport      RoomViewType = "AIRPORT_VIEW"
	ViewBay          RoomViewType = "BAY_VIEW"
	ViewCity         RoomViewType = "CITY_VIEW"
	ViewCourtyard    RoomViewType = "COURTYARD_VIEW"
	ViewGolf         RoomViewType = "GOLF_VIEW"
	ViewHarbor       RoomViewType = "HARBOR_VIEW"
	ViewIntercostals RoomViewType = "INTERCOSTALS_VIEW"
	ViewLake         RoomViewType = "LAKE_VIEW"
	ViewMarina       RoomViewType = "MARINA_VIEW"
	ViewMountain     RoomViewType = "MOUNTAIN_VIEW"
	ViewOcean        RoomViewType = "OCEAN_VIEW"
	ViewPool         RoomViewType = "POOL_VIEW"
	ViewRiver        RoomViewType = "RIVER_VIEW"
	ViewWater        RoomViewType = "WATER_VIEW"
	ViewBeach        RoomViewType = "BEACH_VIEW"
	ViewGarden       RoomViewType = "GARDEN_VIEW"
	ViewPark         RoomViewType = "PARK_VIEW"
	ViewForest       RoomViewType = "FOREST_VIEW"
	ViewRainForest   RoomViewType = "RAIN_FOREST_VIEW"
	ViewVarious      RoomViewType = "VARIOUS_VIEW"
	ViewLimited      RoomViewType = "LIMITED_VIEW"
	ViewSlope        RoomViewType = "SLOPE_VIEW"
	ViewStrip        RoomViewType = "STRIP_VIEW"
	ViewCountryside  RoomViewType = "COUNTRYSIDE_VIEW"
	ViewSea          RoomViewType = "SEA_VIEW"
	ViewValley       RoomViewType = "VALLEY_VIEW"
	ViewDesert       RoomViewType = "DESERT_VIEW"
	ViewCanal        RoomViewType = "CANAL_VIEW"
	ViewLagoon       RoomViewType = "LAGOON_VIEW"
	ViewResort       RoomViewType = "RESORT_VIEW"
	ViewVineyard     RoomViewType = "VINEYARD_VIEW"
)

type ArchitectureCode string

const (
	ArchitectureCodeArtDeco       ArchitectureCode = "ART_DECO"
	ArchitectureCodeBrazilian     ArchitectureCode = "BRAZILIAN"
	ArchitectureCodeContemporary  ArchitectureCode = "CONTEMPORARY"
	ArchitectureCodeHighRise      ArchitectureCode = "HIGH_RISE"
	ArchitectureCodeHistoric      ArchitectureCode = "HISTORIC"
	ArchitectureCodeMediterranean ArchitectureCode = "MEDITERRANEAN"
	ArchitectureCodeModern        ArchitectureCode = "MODERN"
	ArchitectureCodeOriental      ArchitectureCode = "ORIENTAL"
	ArchitectureCodeSouthwest     ArchitectureCode = "SOUTHWEST"
	ArchitectureCodeTraditional   ArchitectureCode = "TRADITIONAL"
	ArchitectureCodeVictorian     ArchitectureCode = "VICTORIAN"
	ArchitectureCodeWestern       ArchitectureCode = "WESTERN"
	ArchitectureCodeAncient       ArchitectureCode = "ANCIENT"
	ArchitectureCodeThemed        ArchitectureCode = "THEMED"
)

type PricingTimeWindow string

const (
	PricingTimeWindowHourly     PricingTimeWindow = "HOURLY"
	PricingTimeWindowDaily      PricingTimeWindow = "DAILY"
	PricingTimeWindowMonthly    PricingTimeWindow = "MONTHLY"
	PricingTimeWindowWeekend    PricingTimeWindow = "WEEKEND"
	PricingTimeWindowWeekly     PricingTimeWindow = "WEEKLY"
	PricingTimeWindowFullPeriod PricingTimeWindow = "FULL_PERIOD"
)

type Unit string

const (
	// Length/Distance
	Miles       Unit = "MILES"
	Kilometers  Unit = "KILOMETERS"
	Meters      Unit = "METERS"
	Millimeters Unit = "MILLIMETERS"
	Centimeters Unit = "CENTIMETERS"
	Yards       Unit = "YARDS"
	Feet        Unit = "FEET"
	Inches      Unit = "INCHES"
	Pixels      Unit = "PIXELS"
	Block       Unit = "BLOCK"

	// Area
	SquareFeet       Unit = "SQUARE_FEET"
	SquareMeters     Unit = "SQUARE_METERS"
	SquareInch       Unit = "SQUARE_INCH"
	SquareYard       Unit = "SQUARE_YARD"
	SquareMillimeter Unit = "SQUARE_MILLIMETER"
	SquareCentimeter Unit = "SQUARE_CENTIMETER"
	Acre             Unit = "ACRE"
	Hectare          Unit = "HECTARE"

	// Weight/Mass
	Pounds    Unit = "POUNDS"
	Kilograms Unit = "KILOGRAMS"
	Ounce     Unit = "OUNCE"
	Gram      Unit = "GRAM"

	// Volume
	Gallons     Unit = "GALLONS"
	Liters      Unit = "LITERS"
	CubicMeters Unit = "CUBIC_METERS"

	// Digital/Energy
	Megabytes Unit = "MEGABYTES"
	Gigabytes Unit = "GIGABYTES"
	Kilowatts Unit = "KILOWATTS"
)

type (
	RoomResponse struct {
		Name                     QualifiedFreeTextResponse        `json:"name"`
		HotelRoomClassification  HotelRoomClassification          `json:"hotelRoomClassification"`  // Indicates the type of the hotel room. Enum values are taken from the OTA Code list GRI
		ProviderContentReference ProviderContentReferenceResponse `json:"providerContentReference"` // Identifies the provider content reference for the room.
		HotelRoomCategory        HotelRoomCategory                `json:"hotelRoomCategory"`        //
		Beds                     int                              `json:"beds"`                     // example: 2. Number of beds in the room
		BedType                  BedType                          `json:"bedType"`                  // Type of the bed.Enum values are taken from the OTA code list BED Here are the list of values
		Description              QualifiedFreeTextResponse        `json:"description"`              //
		Quantity                 int                              `json:"quantity"`                 // Indicates the number of rooms under the given type
		BedRoomsPerRoom          int                              `json:"bedRoomsPerRoom"`          // Indicates the number of bed rooms under the room
		BathroomsPerRoom         int                              `json:"bathroomsPerRoom"`         // Indicates the number of bathrooms under the room
		PolicyDescriptions       []string                         `json:"policyDescriptions"`       // Lists out the set of policies for the room
		ViewCode                 RoomViewType                     `json:"viewCode"`                 // Indicates the view of the room. Enum values are taken from the OTA Code list VIE
		ArchitectureCode         ArchitectureCode                 `json:"architectureCode"`         // Indicates the architecture of the room. Enum values are taken from the OTA Code list ARC
		HotelRoomLocation        string                           `json:"hotelRoomLocation"`        // Indicates the location of the room within the hotel property
		SortOrder                int                              `json:"sortOrder"`                // Absolute order in which to display the room
		Media                    []MediaResponse                  `json:"media"`                    // List of media associated to the room
		IsNonSmoking             bool                             `json:"isNonSmoking"`             // Indicates whether the room is a smoking room or not
		StandardPersonCapacity   int                              `json:"standardPersonCapacity"`   // Indicates the typical capacity for the room
		MaxPersonCapacity        MaxPersonCapacityResponse        `json:"maxPersonCapacity"`        // Capacity of hotel room ( total of number of adult, children)
		MaxSleepFurnishings      MaxSleepFurnishingResponse       `json:"maxSleepFurnishings"`      // Indicates the maximum number of extra beds, cribs that can be accomodated in the room
		Dimensions               DimensionsResponse               `json:"dimensions"`               // A dimension is a measurement such as length, width, or height. Dimensions of a place refers to its size and proportions. The value of height, length and width can be collectively used to calculate multiple values such as total surface area and weight, volume and density.
	}

	RoomDimensionsResponse struct {
		Length        int  `json:"length"`        // Indicates the length of the room
		Width         int  `json:"width"`         // Indicates the width of the room
		Height        int  `json:"height"`        // Indicates the height of the room
		DecimalPlaces int  `json:"decimalPlaces"` // Indicates the number of decimal places for the length, width and height
		Unit          Unit `json:"unit"`          // Indicates the unit of the length, width and height
		Area          int  `json:"area"`          // Indicates the area of the room
		AreaUnit      Unit `json:"areaUnit"`      // Indicates the unit of the area
	}

	MaxSleepFurnishingResponse = sharedResponseDTO.MaxSleepFurnishingsResponse

	MaxPersonCapacityResponse = sharedResponseDTO.MaxPersonCapacityResponse

	// * Indicate relationships from one entity to many other entities of any kind (e.g. from one passenger to their flight segments).
	ProviderContentReferenceResponse struct {
		ID string `json:"id"` // example: 1234567890. ID
	}

	// * Base data model related to Amenity,
	// * inherited amenities can be created using this baseline for extension ( using type for polymorphism )
	// * amenityType : Enum related to generic Amenity model to identify which amenityType
	// * amenityAttribute : String related to specify an Attribute for a given amenityType
	// * amenityProvider : Gives the information to the source of the amenity content
	AmenityResponse struct {
		Code             string                  `json:"code"`             // Indicates the unique code to represent the amenity. There are different types of amenity types. Based on each amenity type , there are various amenities.
		Description      string                  `json:"description"`      //
		IsChargeable     bool                    `json:"isChargeable"`     // default: false
		Price            PriceResponse           `json:"price"`            // Indicates the price of the amenity
		AmenityType      string                  `json:"amenityType"`      // Describes the type / Category of the Amenity Following are different types of amenity types BusinessFacilities DiningFacilities DisabilityFacilities HotelAmenities RoomFacilities SafetyFacilities
		AmenityAttribute string                  `json:"amenityAttribute"` // Describes the Attribute related to the Amenity Type (e.g. Amenity Power can have "Power Outlet" or "USB Outlet)
		AmenityProvider  AmenityProviderResponse `json:"amenityProvider"`  // Provides more information on the Source of the Amenity Content
		Media            []MediaResponse         `json:"media"`            // List of media associated to the amenity
		Quantity         int                     `json:"quantity"`         // Indicates the quantity of the amenity, i.e; how many counts are available for a particular amenity type
		PricingMethod    PricingMethod           `json:"pricingMethod"`    // indicates the pricing method used to asses the amenity's usage cost
	}

	AmenityProviderResponse = sharedResponseDTO.AmenityProvider

	PriceResponse struct {
		Currency          CurrencyResponse  `json:"currency"`          // Indicates the currency of the price.
		SellingTotal      string            `json:"sellingTotal"`      // sellingTotal = Total + margins + markup + totalFees - discounts
		Total             string            `json:"total"`             // total = base + totalTaxes
		PricingTimeWindow PricingTimeWindow `json:"pricingTimeWindow"` // Unit timing Price for which the rate plan pricing applies. It can contain values such as
	}

	CurrencyResponse = sharedResponseDTO.CurrencyResponse
)
