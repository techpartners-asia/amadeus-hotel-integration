// Package contentdto holds the wire structures for the Hotel Content API
// (v3.1): GET /reference-data/locations/by-hotel.
//
// This is the richest of the four schemas - rooms, facilities, policies,
// awards, points of interest and their media - and the one whose blocks are
// most often absent, since what a property publishes varies by source.
package contentdto

import (
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto"
)

type RatingSystem string

const (
	STAR    RatingSystem = "STAR"
	DIAMOND RatingSystem = "DIAMOND"
)

type (
	// * Indicates a public reward, prize, or title that expresses appreciation for any kind of achievement.
	AwardsResponse struct {
		Name         string       `json:"name"`
		ProviderName string       `json:"providerName"` // example: Michelin Name of the provider who has bestowed the honor or award
		Rating       string       `json:"rating"`       // Describes the rating value and its recognitions for the received honor or award
		Description  string       `json:"description"`  // Describes the rating value and its recognitions for the received honor or award
		DateGranted  string       `json:"dateGranted"`  // date on which the hounor was bestowed upon.Format YYYY-MM-DD (ISO 8601)
		RatingSystem RatingSystem `json:"ratingSystem"` // It is a way to evaluate a restaurant's quality using symbols or other notations. For Instance Star Rating from Michelin Stars or Local Star Rating, Diamond from AAA
	}
)

type (
	HotelContentResponse struct {
		Promotions      []PromotionResponse       `json:"promotions"`
		Awards          []AwardsResponse          `json:"awards"`
		Policies        PolicyResponse            `json:"policies"`
		Rooms           []RoomResponse            `json:"rooms"`
		Facilities      FacilityResponse          `json:"facilities"`
		PointOfInterest []PointOfInterestResponse `json:"pointOfInterest"`
		Hotel           HotelResponse             `json:"hotel"`
		Basic           BasicResponse             `json:"basic"`
	}
)

type MeetingRoomType string

const (
	MeetingRoomTypeBanquet    MeetingRoomType = "BANQUET"     // Round tables for meals
	MeetingRoomTypeClassroom  MeetingRoomType = "CLASSROOM"   // Desks/tables facing front
	MeetingRoomTypeConference MeetingRoomType = "CONFERENCE"  // Large central table
	MeetingRoomTypeTheatre    MeetingRoomType = "THEATRE"     // Rows of chairs, no tables
	MeetingRoomTypeUShape     MeetingRoomType = "U_SHAPE"     // Tables in a 'U' with open front
	MeetingRoomTypeReception  MeetingRoomType = "RECEPTION"   // Standing space for cocktails/social
	MeetingRoomTypeBoardroom  MeetingRoomType = "BOARDROOM"   // Formal executive table
	MeetingRoomTypeOpenSquare MeetingRoomType = "OPEN_SQUARE" // Tables in a square with hollow center
	MeetingRoomTypeOther      MeetingRoomType = "OTHER"       // Custom or undefined layout
)

type Layout string

const (
	// Standard Styles
	LayoutBanquet      Layout = "BANQUET"
	LayoutClassroom    Layout = "CLASSROOM"
	LayoutConference   Layout = "CONFERENCE"
	LayoutTheatre      Layout = "THEATRE"
	LayoutReception    Layout = "RECEPTION"
	LayoutBoardroom    Layout = "BOARDROOM"
	LayoutHollowSquare Layout = "HOLLOW_SQUARE"
	LayoutOpenSquare   Layout = "OPEN_SQUARE"
	LayoutBallroom     Layout = "BALLROOM"
	LayoutUShape       Layout = "U_SHAPE"
	LayoutUShaped      Layout = "U_SHAPED" // Note: Both variants exist in your list
	LayoutTShaped      Layout = "T_SHAPED"
	LayoutEShaped      Layout = "E_SHAPED"

	// Rounds & Social
	LayoutRoundTables        Layout = "ROUND_TABLES"
	LayoutRoundsFor8         Layout = "ROUNDS_FOR_8"
	LayoutRoundsFor10        Layout = "ROUNDS_FOR_10"
	LayoutBanquetRoundsFor12 Layout = "BANQUET_ROUNDS_FOR_12"
	LayoutCocktailRounds     Layout = "COCKTAIL_ROUNDS"
	LayoutCrescentRounds     Layout = "CRESCENT_ROUNDS"
	LayoutCrescentRoundsOf5  Layout = "CRESCENT_ROUNDS_OF_5"
	LayoutCrescentRoundsOf6  Layout = "CRESCENT_ROUNDS_OF_6"

	// Specialized Classroom Variations
	LayoutClassroom2Per6ft        Layout = "CLASSROOM_2_PER_6_FT_TABLES"
	LayoutClassroom3Per6ft        Layout = "CLASSROOM_3_PER_6_FT_TABLES"
	LayoutClassroom3Per8ft        Layout = "CLASSROOM_3_PER_8_FT_TABLES"
	LayoutClassroom4Per8ft        Layout = "CLASSROOM_4_PER_8_FT_TABLES"
	LayoutClassroomChevron2Per6ft Layout = "CLASSROOM_CHEVRON_2_PER_6_FT_TABLES"
	LayoutClassroomChevron3Per6ft Layout = "CLASSROOM_CHEVRON_3_PER_6_FT_TABLES"
	LayoutClassroomChevron3Per8ft Layout = "CLASSROOM_CHEVRON_3_PER_8_FT_TABLES"
	LayoutClassroomChevron4Per8ft Layout = "CLASSROOM_CHEVRON_4_PER_8_FT_TABLES"

	// Exhibits
	LayoutExhibit          Layout = "EXHIBIT"
	LayoutTabletopExhibits Layout = "TABLETOP_EXHIBITS"
	LayoutExhibit8x10      Layout = "8_INCH_X_10_INCH_EXHIBITS"
	LayoutExhibit10x10     Layout = "10_INCH_X_10_INCH_EXHIBITS"
	LayoutIslandExhibit    Layout = "ISLAND_EXHIBIT"
	LayoutPeninsulaExhibit Layout = "PENINSULA_EXHIBIT"
	LayoutPerimeterExhibit Layout = "PERIMETER_EXHIBIT"

	// Miscellaneous
	LayoutRoyalConference   Layout = "ROYAL_CONFERENCE"
	LayoutTalkShow          Layout = "TALK_SHOW"
	LayoutTheaterSemiCircle Layout = "THEATER_SEMI_CIRCLE"
	LayoutTheaterChevron    Layout = "THEATER_CHEVRON"
	LayoutFlowNoFurniture   Layout = "FLOW_NO_TABLES_OR_CHAIRS"
	LayoutFoyer             Layout = "FOYER"
	LayoutCustom            Layout = "CUSTOM"
	LayoutOther             Layout = "OTHER"
)

// RestaurantServiceCategory represents the service style or category of a food outlet.
type RestaurantServiceCategory string

const (
	RestaurantServiceCategoryAllPurpose     RestaurantServiceCategory = "ALL_PURPOSE"
	RestaurantServiceCategoryBeverage       RestaurantServiceCategory = "BEVERAGE"
	RestaurantServiceCategoryBuffet         RestaurantServiceCategory = "BUFFET"
	RestaurantServiceCategoryCafe           RestaurantServiceCategory = "CAFE"
	RestaurantServiceCategoryCafeteria      RestaurantServiceCategory = "CAFETERIA"
	RestaurantServiceCategoryCasual         RestaurantServiceCategory = "CASUAL"
	RestaurantServiceCategoryFamily         RestaurantServiceCategory = "FAMILY"
	RestaurantServiceCategoryFastFood       RestaurantServiceCategory = "FAST_FOOD"
	RestaurantServiceCategoryFineDining     RestaurantServiceCategory = "FINE_DINING"
	RestaurantServiceCategoryKiosk          RestaurantServiceCategory = "KIOSK"
	RestaurantServiceCategoryTakeOut        RestaurantServiceCategory = "TAKE_OUT"
	RestaurantServiceCategoryUpscale        RestaurantServiceCategory = "UPSCALE"
	RestaurantServiceCategoryBarOrLounge    RestaurantServiceCategory = "BAR_OR_LOUNGE"
	RestaurantServiceCategoryBrasserie      RestaurantServiceCategory = "BRASSERIE"
	RestaurantServiceCategoryCoffeeBar      RestaurantServiceCategory = "COFFEE_BAR"
	RestaurantServiceCategoryDessertSnack   RestaurantServiceCategory = "DESSERT_OR_ICE_CREAM_OR_SNACK_BAR"
	RestaurantServiceCategoryFullService    RestaurantServiceCategory = "FULL_SERVICE"
	RestaurantServiceCategoryPub            RestaurantServiceCategory = "PUB"
	RestaurantServiceCategoryDeli           RestaurantServiceCategory = "DELI"
	RestaurantServiceCategoryPrivateDining  RestaurantServiceCategory = "PRIVATE_DINING"
	RestaurantServiceCategorySportsBar      RestaurantServiceCategory = "SPORTS_BAR"
	RestaurantServiceCategoryPianoBar       RestaurantServiceCategory = "PIANO_BAR"
	RestaurantServiceCategoryOutdoorBarCafe RestaurantServiceCategory = "OUTDOOR_BAR_OR_CAFE"
	RestaurantServiceCategoryBeerGarden     RestaurantServiceCategory = "BEER_GARDEN"
	RestaurantServiceCategoryBeachBar       RestaurantServiceCategory = "BEACH_BAR"
	RestaurantServiceCategoryTapasBar       RestaurantServiceCategory = "TAPAS_BAR"
	RestaurantServiceCategoryDessert        RestaurantServiceCategory = "DESSERT"
	RestaurantServiceCategoryFoodTruck      RestaurantServiceCategory = "FOOD_TRUCK"
)

type PhoneCategory string

const (
	PhoneCategoryEmergencyContact PhoneCategory = "EMERGENCY_CONTACT"
	PhoneCategoryTravelArranger   PhoneCategory = "TRAVEL_ARRANGER"
	PhoneCategoryDaytimeContact   PhoneCategory = "DAYTIME_CONTACT"
	PhoneCategoryEveningContact   PhoneCategory = "EVENING_CONTACT"
	PhoneCategoryTollFreeNumber   PhoneCategory = "TOLL_FREE_NUMBER"
	PhoneCategoryGuestUse         PhoneCategory = "GUEST_USE"
	PhoneCategoryPickupContact    PhoneCategory = "PICKUP_CONTACT"
	PhoneCategoryContact          PhoneCategory = "CONTACT"
)

type DeviceType string

const (
	DeviceTypeFax      DeviceType = "FAX"
	DeviceTypeMobile   DeviceType = "MOBILE"
	DeviceTypeLandline DeviceType = "LANDLINE"
	DeviceTypeVoice    DeviceType = "VOICE"
	DeviceTypeTelex    DeviceType = "TELEX"
)

type EmailCategory string

const (
	Personal          EmailCategory = "PERSONAL"           // Guest's private email
	Business          EmailCategory = "BUSINESS"           // Guest's work email
	Property          EmailCategory = "PROPERTY"           // Direct hotel/resort email
	SalesOffice       EmailCategory = "SALES_OFFICE"       // For group bookings/events
	ReservationOffice EmailCategory = "RESERVATION_OFFICE" // Central reservations (CRO)
	ManagingCompany   EmailCategory = "MANAGING_COMPANY"   // Corporate owner or management group
)

type EmailAddressType string

const (
	EMAIL_ID                         EmailAddressType = "EMAIL_ID"
	EmailAddressTypeDistributionList EmailAddressType = "DISTRIBUTION_LIST"
	EmailAddressTypeAlias            EmailAddressType = "ALIAS"
	EmailAddressTypeGroup            EmailAddressType = "GROUP"
)

type LocationType string

const (
	LocationTypeCentralReservationOffice LocationType = "CENTRAL RESERVATION OFFICE"
	LocationTypeCorporateHeadquarters    LocationType = "CORPORATE HEADQUARTERS"
	LocationTypeCorporateOffice          LocationType = "CORPORATE OFFICE"
	LocationTypeDivisionalOffice         LocationType = "DIVISIONAL OFFICE"
	LocationTypeGlobalSalesOffice        LocationType = "GLOBAL SALES OFFICE"
	LocationTypeHotelDirect              LocationType = "HOTEL DIRECT"
	// Hotel Content returns this unspaced spelling, not "HOTEL DIRECT".
	LocationTypeHotelDirectUnspaced    LocationType = "HOTELDIRECT"
	LocationTypeReservations           LocationType = "RESERVATIONS"
	LocationTypeLocalReservationOffice LocationType = "LOCAL RESERVATION OFFICE"
	LocationTypeSalesOffice            LocationType = "SALES_OFFICE"
	LocationTypeFranchiseCompany       LocationType = "FRANCHISE COMPANY"
	LocationTypeManagementCompany      LocationType = "MANAGEMENT COMPANY"
	LocationTypeOwnershipCompany       LocationType = "OWNERSHIP COMPANY"
	LocationTypeCustomerServiceOffice  LocationType = "CUSTOMER_SERVICE_OFFICE"
	LocationTypeHomeResidence          LocationType = "HOME_RESIDENCE"
	LocationTypeRegionalSalesOffice    LocationType = "REGIONAL SALES OFFICE"
	LocationTypeTechnicalSupportOffice LocationType = "TECHNICAL SUPPORT OFFICE"
)

type (
	// * Contains all the facilities offered by the hotel.
	FacilityResponse struct {
		MeetingRoomInfo MeetingRoomInfoResponse `json:"meetingRoomInfo"` // Indicates the meeting room information
		Amenities       []dto.AmenityResponse   `json:"amenities"`       // Indicates the amenities offered by the facility
		RestaurantInfo  RestaurantInfoResponse  `json:"restaurantInfo"`  // Indicates the restaurant information
	}

	RestaurantInfoResponse struct {
		Quantity    int                  `json:"quantity"`    // Indicates the number of restaurants within the property
		Restaurants []RestaurantResponse `json:"restaurants"` // Indicates the various restaurants in the property
	}

	RestaurantResponse struct {
		Name                  string                    `json:"name"`                  // Indicates the name of the restaurant
		Description           string                    `json:"description"`           // Indicates the description of the restaurant
		Category              RestaurantServiceCategory `json:"category"`              // Restaurant food service category. These enum values are inspired from OTA - "https://opentravel.org/" with code list as - RES. Can contain values such as
		AcceptedCurrencyCodes []string                  `json:"acceptedCurrencyCodes"` // example: EUR
		CuisineTypes          []string                  `json:"cuisineTypes"`          // Indicates the list of cuisines served at the restaurant. These enum values are inspired from OTA - "https://opentravel.org/" with code list as - CUI
		MaxSeatingCapacity    float64                   `json:"maxSeatingCapacity"`    // Indicates the max number of occupancy in the restaurant
		HasBreakfast          bool                      `json:"hasBreakfast"`          // True if breakfast is served in the restaurant. Default value is false
		HasLunch              bool                      `json:"hasLunch"`              // True if lunch is served in the restaurant. Default value is false
		HasBrunch             bool                      `json:"hasBrunch"`             // True if brunch is served in the restaurant. Default value is false
		HasDinner             bool                      `json:"hasDinner"`             // True if dinner is served in the restaurant. Default value is false
		Contact               []ContactResponse         `json:"contact"`               // A contact refers to the information that can be used to reach a person, a company or an organization.
		HonorsAndAwards       []AwardsResponse          `json:"honorsAndAwards"`       // Indicates the honors and awards received by the restaurant
		Media                 []dto.MediaResponse       `json:"media"`                 // Indicates the media of the restaurant
		OperatingHours        CalendarScheduleResponse  `json:"operatingHours"`        // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
		IsReservationRequired bool                      `json:"isReservationRequired"` // True if a reservation is required to dine at the restaurant
	}

	CalendarScheduleResponse struct {
		StartDate string              `json:"startDate"` // The start date or datetime of the item, in ISO 8601 date format: - dates in the form [-]CCYY-MM-DD - datetimes in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm]
		StartTime string              `json:"startTime"` // The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action started to be performed. e.g. John wrote a book from January to December. For media, including audio and video, it's the time offset of the start of a clip within a larger file. It can be represented as a time or a datetime. Datetimes are represented with a sting in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm] (see Chapter 5.4 of ISO 8601). Times are represented with a string in the form hh:mm:ss[Z|(+|-)hh:mm] (see Chapter 5.3 of ISO 8601). If no explicit UTC deviation is provided, then time is intended as local, and the location is given in the scheduleTimezone attribute. If no explicit UTC deviation is given and scheduleTimezone is not present, then time is considered as UTC. If a timezone is provided in scheduleTimezone, but the time or datetime is explicitely providing a UTC deviation, then the timezone is ignored and the UTC deviation given in the time or datetime is to be used instead.
		EndDate   string              `json:"endDate"`   // The end date or datetime of the item, in ISO 8601 date format: - dates in the form [-]CCYY-MM-DD - datetimes in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm]
		EndTime   string              `json:"endTime"`   // The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December. For media, including audio and video, it's the time offset of the end of a clip within a larger file. It can be represented as a time or a datetime. Datetimes are represented with a sting in the form [-]CCYY-MM-DDThh:mm:ss[Z|(+|-)hh:mm] (see Chapter 5.4 of ISO 8601). Times are represented with a string in the form hh:mm:ss[Z|(+|-)hh:mm] (see Chapter 5.3 of ISO 8601) If no explicit UTC deviation is provided, then time is inteded as local, and the location is given in the scheduleTimezone attribute. If no explicit UTC deviation is given and scheduleTimezone is not present, then time is considered as UTC. If a timezone is provided in scheduleTimezone, but the time or datetime is explicitely providing a UTC deviation, then the timezone is ignored and the UTC deviation given in the time or datetime is to be used instead.
		ByDays    []DayOfWeekResponse `json:"byDays"`    //  Defines the day(s) of the week on which a recurring Event takes place.
	}

	DayOfWeekResponse struct {
		Day string `json:"day"` // Indicates the day of the week. It can contain values such as MON, TUE, WED, THU, FRI, SAT, SUN
	}

	ContactResponse struct {
		AddresseeName AddresseeNameResponse `json:"addresseeName"` // Indicates the name of the person, company or organization that the contact is for
		Phones        PhoneResponse         `json:"phone"`         // Indicates the phone numbers of the person, company or organization that the contact is for
		Address       AddressResponse       `json:"address"`       // Indicates the postal address of the person, company or organization that the contact is for
		Email         EmailResponse         `json:"email"`         // Indicates the email addresses of the person, company or organization that the contact is for
		Purpose       []string              `json:"purpose"`       // the purpose for which this contact is to be used
		LocationType  LocationType          `json:"locationType"`  // Describes the locationType of the contact. It can contain values such as
		Website       struct {
			Url  string `json:"url"`  // Indicates the URL of the website
			Href string `json:"href"` // Indicates the URL of the website (key actually returned by Hotel Content)
		} `json:"website"` // Object containing URL and description
	}

	// * A postal address used to locate a person, company or organization.
	AddressResponse struct {
		Category    string   `json:"category"`    // Category of the address
		Lines       []string `json:"lines"`       // Address lines (street, building, etc.)
		PostalCode  string   `json:"postalCode"`  // Postal or ZIP code
		CountryCode string   `json:"countryCode"` // ISO country code
		CityName    string   `json:"cityName"`    // City name
		CountyName  string   `json:"countyName"`  // County name
		CountryName string   `json:"countryName"` // Country name
		StateCode   string   `json:"stateCode"`   // State or province code
		PostalBox   string   `json:"postalBox"`   // Post office box
		Text        string   `json:"text"`        // Free-text representation of the address
		State       string   `json:"state"`       // State or province name
	}

	AddresseeNameResponse struct {
		FirstName  string `json:"firstName"`  // Indicates the first name of the person, company or organization that the contact is for
		LastName   string `json:"lastName"`   // Indicates the last name of the person, company or organization that the contact is for
		MiddleName string `json:"middleName"` // Indicates the middle name of the person, company or organization that the contact is for
		Prefix     string `json:"prefix"`     // Indicates the prefix of the person, company or organization that the contact is for
		Suffix     string `json:"suffix"`     // Indicates the suffix of the person, company or organization that the contact is for
		NameType   string `json:"nameType"`   // the type of the reference name. It can also contain values such as FORMER ,NICKNAME ,ALTERNATE ,MAIDEN

	}

	PhoneResponse struct {
		Category           PhoneCategory `json:"category"`           // Indicates the category of the phone number. It can also contain values such as HOME ,WORK ,MOBILE ,FAX ,PAGER ,OTHER
		DeviceType         DeviceType    `json:"deviceType"`         // Type of the device. It Can also contain values such as
		CountryCode        string        `json:"countryCode"`        // Indicates the country code of the phone number
		CountryCallingCode string        `json:"countryCallingCode"` // Country calling code of the phone number, as defined by the International Communication Union. Examples - "1" for US, "371" for Latvia.
		AreaCode           string        `json:"areaCode"`           // Corresponds to a regional code or a city code. The length of the field varies depending on the area.
		Number             string        `json:"number"`             // Phone number. Composed of digits only. The number of digits depends on the country.
		Extension          string        `json:"extension"`          // Extension of the phone

	}

	EmailResponse struct {
		Category         EmailCategory    `json:"category"`         // Indicates the category of the email address. It can also contain values such as HOME ,WORK ,OTHER
		Email            string           `json:"email"`            // Indicates the email address of the person, company or organization that the contact is for
		Address          string           `json:"address"`          // Email address (e.g. john@smith.com)
		EmailAddressType EmailAddressType `json:"emailAddressType"` // EmailAddressingType defines the format of Email Address. EMAIL_ID (default) - Single Email Address. ex: abc@amadeus.com DISTRIBUTION_LIST - Similar to Email Mailing List - refers to a list of email addresses. Alias - A Nickname to Email Address or list of Emails. Do not follow Email Format. (Refer DLIST ABR Rule). RFC 5321. ex: SUPPAX Group - A group has an email address and whenever an email is sent to that address everyone in the group receives the email. ex: "Help_Desk@amadeus.com"
	}

	MeetingRoomInfoResponse struct {
		Quantity                  int                    `json:"quantity"`                  // Indicates the number of meeting rooms within the property
		SmallestRoomSpace         RoomDimensionsResponse `json:"smallestRoomSpace"`         // Indicates the smallest meeting room space within the property
		LargestRoomSpace          RoomDimensionsResponse `json:"largestRoomSpace"`          // Indicates the largest meeting room space within the property
		SmallestRoomSeatOccupancy int                    `json:"smallestRoomSeatOccupancy"` // Indicates the number of people that can be accomodated in the smallest room meeting room within the property
		LargestRoomSeatOccupancy  int                    `json:"largestRoomSeatOccupancy"`  // Indicates the number of people that can be accomodated in the largest room meeting room within the property
		TotalRoomSeatOccupancy    int                    `json:"totalRoomSeatOccupancy"`    // Indicates the number of people that can be accomodated in the all the meeting rooms combined
		MeetingRooms              []MeetingRoomResponse  `json:"meetingRooms"`              // Indicates the various meeting rooms in the property
	}
	// * A meeting room is a space usually set aside for people to get together, often informally to hold meetings, for issues to be discussed, priorities set and decisions made.
	MeetingRoomResponse struct {
		Name                string                        `json:"name"`                // Indicates the name of the meeting room
		MeetingRoomType     MeetingRoomType               `json:"meetingRoomType"`     // Indicates the type of the meeting room
		Description         string                        `json:"description"`         // Indicates the description of the meeting room
		OccupancyPerLayouts []OccupancyPerLayoutsResponse `json:"occupancyPerLayouts"` // Denotes the occupancy for each type of layout that can be designed in the meeting room. For instance, a U-Shaped meeting room can have a maxOCcupancy of 10
		SortOrder           int                           `json:"sortOrder"`           // Indicates the sort order of the meeting room
		ExhibitDimensions   RoomDimensionsResponse        `json:"exhibitDimensions"`   // Indicates the dimensions of the exhibit
		RoomDimensions      RoomDimensionsResponse        `json:"roomDimensions"`      // Indicates the dimensions of the room
		PriceQuotations     []PriceQuotationsResponse     `json:"priceQuotations"`     // Indicates the price quotations for the meeting room
		Media               []dto.MediaResponse           `json:"media"`               // Indicates the media of the meeting room
	}

	PriceQuotationsResponse struct {
		PricingMethod dto.PricingMethod `json:"pricingMethod"` // Indicates the pricing method used to asses the meeting room's usage cost
		UnitPrice     PriceResponse     `json:"unitPrice"`     // Indicates the price of the meeting room per unit
	}

	OccupancyPerLayoutsResponse struct {
		Layout       Layout `json:"layout"`       // Defines the design layout type of the meeting room
		MaxOccupancy int    `json:"maxOccupancy"` // Denotes the maximum number of people that can be accomodated in the corresponding layout
	}
)

type FrequencyType string

const (
	FrequencyTypeHourly     FrequencyType = "HOURLY"
	FrequencyTypeDaily      FrequencyType = "DAILY"
	FrequencyTypeMonthly    FrequencyType = "MONTHLY"
	FrequencyTypeWeekend    FrequencyType = "WEEKEND"
	FrequencyTypeWeekly     FrequencyType = "WEEKLY"
	FrequencyTypeFullPeriod FrequencyType = "FULL_PERIOD"
)

type (
	BasicResponse struct {
		Season struct {
			OpenCalendar []Period `json:"openCalendar"` // Indicates the opening time and days of the property
		} `json:"season"` // Indicates the season of the point of interest
		HotelID                      string                         `json:"hotelId"`                      // Amadeus Property Code (8 chars). example: ADPAR001
		ChainCode                    string                         `json:"chainCode"`                    // Brand (RT...) or Merchant (AD...)
		BrandCode                    string                         `json:"brandCode"`                    // Brand (RT...) (Amadeus 2 chars Code). Small Properties distributed by Merchants may not have a Brand. Example - AD (Value Hotels) is the Provider/Merchant, and RT (Accor) is the Brand of the Property
		DupeID                       string                         `json:"dupeId"`                       // Unique Property identifier of the physical hotel. One physical hotel can be represented by different Providers, each one having its own hotelID. This attribute allows a client application to group together hotels that are actually the same.
		Name                         string                         `json:"name"`                         // Name of the point of interest
		Rating                       string                         `json:"rating"`                       // Rating of the point of interest
		Description                  dto.QualifiedFreeTextResponse  `json:"description"`                  // Description of the point of interest
		Amenities                    []dto.AmenityResponse          `json:"amenities"`                    // Amenities of the point of interest
		Media                        []dto.MediaResponse            `json:"media"`                        // Media of the point of interest
		DefaultSpokenLanguage        string                         `json:"defaultSpokenLanguage"`        // Describes the default language preferred or used at the property
		ContextProvider              string                         `json:"contextProvider"`              // Describes the provider of the context of the point of interest
		Contact                      []ContactResponse              `json:"contact"`                      // Contact of the point of interest
		Location                     LocationResponse               `json:"location"`                     // Location of the point of interest
		Altitude                     AltitudeResponse               `json:"altitude"`                     // From analytics, Metrics describe the exact numbers that make up the data
		CategoryCode                 CategoryCode                   `json:"categoryCode"`                 // Category code of the point of interest
		Category                     []string                       `json:"category"`                     // Category labels of the property
		Segment                      []Segment                      `json:"segment"`                      // Segments of the point of interest
		Area                         []AreaResponse                 `json:"area"`                         // Geographical zone like City, Region, Country
		ChainName                    string                         `json:"chainName"`                    // Name of the chain to which the hotel belongs to
		BrandName                    string                         `json:"brandName"`                    // Name of the brand to which the hotel or hotel chain belongs to
		Status                       HotelStatus                    `json:"status"`                       // Status of the hotel
		HotelBusinessIdentifications BusinessIdentificationResponse `json:"hotelBusinessIdentifications"` // An business, can be idenfified via business identifiers, those business identifiers are defined by a body of authority thay could be local, national, transnational or supranational (like EU for the EU VAT number).
	}

	BusinessIdentificationResponse struct {
		Identifiers []IdentifierResponse `json:"identifiers"` // Identifiers of the business
	}

	IdentifierResponse struct {
		ID   string `json:"id"`   // Identifier id
		Name string `json:"name"` // Identifier name
	}

	AreaResponse struct {
		HotelAreaType HotelAreaType `json:"hotelAreaType"` // 'Indicates the category of the location. OTA Code Set LOC values are to be considered here. Can contain values such as
		Name          string        `json:"name"`          // Label associated to the location (e.g. Eiffel Tower, Madison Square)
	}

	AltitudeResponse struct {
		Unit  Unit `json:"unit"`  // Indicates the unit of the altitude
		Value int  `json:"value"` // Indicates the value of the altitude
	}

	HotelResponse struct {
		TaxID            string                     `json:"taxId"`            // Describes the unique tax identifier of a hotel property
		CurrencyCode     []string                   `json:"currencyCode"`     // Describes the currency code accepted at the property. Example : [EUR]
		SpokenLanguages  []string                   `json:"spokenLanguages"`  // Describes the list of languages spoken at the property. Follows the standard of ISO 639-1 (Alpha-2). Example : [es]
		TimeZone         TimeZoneResponse           `json:"timeZone"`         // Element defining a time zone
		Climate          string                     `json:"climate"`          // Describes the climate at the location of the property. example: Dry
		Certifications   []AwardsResponse           `json:"certifications"`   // Describes the certifications received by the Hotel
		RelativeLocation []LocationDistanceResponse `json:"relativeLocation"` // To indicate the reference points from the hotel such as the distance to Airport, Bus Stations or Train Station.
		Season           SeasonResponse             `json:"season"`           // Models a period of time between two dates and inclusive only of the days of the week specified.
		Building         BuildingResponse           `json:"building"`         // Indicates the building of the hotel
	}

	BuildingResponse struct {
		ArchitectureCode        ArchitectureCode `json:"architectureCode"`        // Denotes the architecture in which the property was built upon. Can contain values such as
		BuiltDate               string           `json:"builtDate"`               // Denotes the year at which the property was built. Format YYYY-MM-DD (ISO 8601)
		RenovationDate          string           `json:"renovationDate"`          // Denotes the year at which the property was renovated. Format YYYY-MM-DD (ISO 8601)
		NumberOfFloors          int              `json:"numberOfFloors"`          // Indicates the number of floors in the property
		NumberOfRooms           int              `json:"numberOfRooms"`           // Indicates the number of rooms in the property
		NumberOfExecutiveFloors int              `json:"numberOfExecutiveFloors"` // Indicates the number of Executive floors in the property
		NumberOfBuildings       int              `json:"numberOfBuildings"`       // Indicates the number of buildings in the property
		NumberOfElevators       int              `json:"numberOfElevators"`       // Indicates the number of elevators in the property
	}

	SeasonResponse struct {
		ClosedSeasons         []Period               `json:"closedSeasons"`         // Closed seasons of the hotel refers to the season where in the property is shut down
		BlackoutSeasons       []Period               `json:"blackoutSeasons"`       // Blackout dates of the hotel during which the hotel is open but no bookings are available
		OpenCalendar          []Period               `json:"openCalendar"`          // Indicates the opening time and days of the property
		ClosedSeasonsDetail   []SeasonPeriodResponse `json:"closedSeasonsDetail"`   // Detailed closed seasons including recurrence (dow, moy, frequencyType) and excluded periods
		BlackoutSeasonsDetail []SeasonPeriodResponse `json:"blackoutSeasonsDetail"` // Detailed blackout seasons including recurrence (dow, moy, frequencyType) and excluded periods
		OpenCalendarDetail    []SeasonPeriodResponse `json:"openCalendarDetail"`    // Detailed open calendar including recurrence (dow, moy, frequencyType) and excluded periods
	}

	// * Models a period of time between two dates inclusive only of the days of the week and months of the year specified, with optional excluded sub-periods.
	SeasonPeriodResponse struct {
		Start           string                   `json:"start"`           // Start date and time following ISO 8601 format
		End             string                   `json:"end"`             // End date and time following ISO 8601 format
		Dow             string                   `json:"dow"`             // Days of the week the period applies to
		FrequencyType   FrequencyType            `json:"frequencyType"`   // Recurrence frequency of the period
		Moy             string                   `json:"moy"`             // Months of the year the period applies to
		ExcludedPeriods []ExcludedPeriodResponse `json:"excludedPeriods"` // Sub-periods excluded from this period
	}

	// * Models a sub-period excluded from an enclosing season period.
	ExcludedPeriodResponse struct {
		Start string `json:"start"` // Start date and time following ISO 8601 format
		End   string `json:"end"`   // End date and time following ISO 8601 format
	}

	TimeZoneResponse struct {
		ID                     string `json:"id"`                     // Unique id of the time zone. example: Europe/Paris
		Name                   string `json:"name"`                   //Long name of the time zone. example: Central European Summer Time
		Code                   string `json:"code"`                   // Time zone code. example: CEST
		OffSet                 string `json:"offSet"`                 // Total offset from UTC including the Daylight Saving Time (DST) following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: +02:00
		OffSetInSeconds        int    `json:"offSetInSeconds"`        // Total offset from UTC including the Daylight Saving Time (DST) in second. example: 7200
		DstOffset              string `json:"dstOffset"`              // Indicates whether the day light savings is observed at the location. example: True
		DstOffsetInSeconds     int    `json:"dstOffsetInSeconds"`     // Daylight Saving Time (DST) in second. 0 if the zone is not in the Daylight Saving time at specified date. example: -3600
		DstSetInSeconds        int    `json:"dstSetInSeconds"`        // Daylight Saving Time (DST) in second. 0 if the zone is not in the Daylight Saving time.
		ReferenceLocalDateTime string `json:"referenceLocalDateTime"` // Date and time used as reference to determine the time zone name, code, offset, and dstOffset following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: 2022-09-28T19:20:30
	}

	LocationDistanceResponse struct {
		Destination LocationResponse   `json:"destination"` // Indicates the destination of the location distance
		Distances   []DistanceResponse `json:"distances"`   // Indicates the distances from the point of interest to the destination
	}

	DistanceResponse struct {
		Unit         Unit         `json:"unit"`         // Indicates the unit of the distance
		Value        float64      `json:"value"`        // Indicates the value of the distance
		DistanceType DistanceType `json:"distanceType"` // Indicates the type of the distance
	}

	TransportationResponse struct {
		TransportMode         TransportMode             `json:"transportMode"`
		IsReservationRequired bool                      `json:"isReservationRequired"` // True if reservation is required in advance to board the transport
		OperatingHours        CalendarScheduleResponse  `json:"operatingHours"`        // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
		Description           string                    `json:"description"`           // Description of the transportation
		PriceEquation         []PricingEquationResponse `json:"priceEquation"`         // Indicates the price equation of the transportation
		PriceQuotation        []PriceQuotationResponse  `json:"priceQuotation"`        // Indicates the price quotations for the transportation
		Media                 []dto.MediaResponse       `json:"media"`                 // Indicates the media of the transportation
	}

	PricingEquationResponse struct {
		PricingMethod dto.PricingMethod       `json:"pricingMethod"` // Indicates the pricing method of the point of interest
		UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the price of the point of interest per unit
	}

	// * Indicates a price quote for a given pricing method, using an elementary (itemized) price.
	PriceQuotationResponse struct {
		PricingMethod dto.PricingMethod       `json:"pricingMethod"` // Indicates the pricing method used to assess the quoted price
		UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the quoted price per unit
	}

	ElementaryPriceResponse struct {
		Amount              string               `json:"amount"`              // Indicates the amount of the price of the point of interest
		Value               string               `json:"value"`               // Indicates the value of the price of the point of interest
		DecimalPlaces       int                  `json:"decimalPlaces"`       // Indicates the decimal places of the price of the point of interest
		Currency            dto.CurrencyResponse `json:"currency"`            // Indicates the currency of the price of the point of interest
		ElementaryPriceType string               `json:"elementaryPriceType"` // Defines the type of price, eg. for base fare, total, grand total.
	}

	LocationResponse struct {
		SubType  string              `json:"subType"`  // Location sub-type (e.g. airport, port, rail-station, restaurant, atm...)
		Subtype  string              `json:"subtype"`  // Same as SubType; Hotel Content returns this all-lowercase spelling
		Name     string              `json:"name"`     // Name of the location
		IataCode string              `json:"iataCode"` // IATA code of the location
		GeoCode  dto.GeoCodeResponse `json:"geoCode"`  // GeoCode of the location
	}
	Period struct {
		Start *time.Time `json:"start"` // start date and time following ISO 8601 format
		End   *time.Time `json:"end"`   // end date and time following ISO 8601 format
	}
)

type CategoryCode string

const (
	CategoryCodeAirport                                  CategoryCode = "AIRPORT"
	CategoryCodeAmusementPark                            CategoryCode = "AMUSEMENT_PARK"
	CategoryCodeAquarium                                 CategoryCode = "AQUARIUM"
	CategoryCodeBeach                                    CategoryCode = "BEACH"
	CategoryCodeBoatDock                                 CategoryCode = "BOAT_DOCK"
	CategoryCodeBusStation                               CategoryCode = "BUS_STATION"
	CategoryCodeBusinessLocation                         CategoryCode = "BUSINESS_LOCATION"
	CategoryCodeCanal                                    CategoryCode = "CANAL"
	CategoryCodeCarRentalLocation                        CategoryCode = "CAR_RENTAL_LOCATION"
	CategoryCodeCasino                                   CategoryCode = "CASINO"
	CategoryCodeCemetery                                 CategoryCode = "CEMETERY"
	CategoryCodeChurch                                   CategoryCode = "CHURCH"
	CategoryCodeConcertHall                              CategoryCode = "CONCERT_HALL"
	CategoryCodeConferenceCenter                         CategoryCode = "CONFERENCE_CENTER"
	CategoryCodeConventionCenter                         CategoryCode = "CONVENTION_CENTER"
	CategoryCodeFairground                               CategoryCode = "FAIRGROUND"
	CategoryCodeFarm                                     CategoryCode = "FARM"
	CategoryCodeGallery                                  CategoryCode = "GALLERY"
	CategoryCodeHistoricBuilding                         CategoryCode = "HISTORIC_BUILDING"
	CategoryCodeHospital                                 CategoryCode = "HOSPITAL"
	CategoryCodeLake                                     CategoryCode = "LAKE"
	CategoryCodeLandmark                                 CategoryCode = "LANDMARK"
	CategoryCodeMarina                                   CategoryCode = "MARINA"
	CategoryCodeMarket                                   CategoryCode = "MARKET"
	CategoryCodeMonument                                 CategoryCode = "MONUMENT"
	CategoryCodeMountain                                 CategoryCode = "MOUNTAIN"
	CategoryCodeMuseum                                   CategoryCode = "MUSEUM"
	CategoryCodeOcean                                    CategoryCode = "OCEAN"
	CategoryCodePalace                                   CategoryCode = "PALACE"
	CategoryCodePark                                     CategoryCode = "PARK"
	CategoryCodeRecreationCenter                         CategoryCode = "RECREATION_CENTER"
	CategoryCodeRestaurant                               CategoryCode = "RESTAURANT"
	CategoryCodeRiver                                    CategoryCode = "RIVER"
	CategoryCodeShoppingMall                             CategoryCode = "SHOPPING_MALL"
	CategoryCodeSkiArea                                  CategoryCode = "SKI_AREA"
	CategoryCodeStadium                                  CategoryCode = "STADIUM"
	CategoryCodeStore                                    CategoryCode = "STORE"
	CategoryCodeTheaterCinema                            CategoryCode = "THEATER_CINEMA"
	CategoryCodeTrainStation                             CategoryCode = "TRAIN_STATION"
	CategoryCodeUniversity                               CategoryCode = "UNIVERSITY"
	CategoryCodeWinery                                   CategoryCode = "WINERY"
	CategoryCodeZoo                                      CategoryCode = "ZOO"
	CategoryCodeCityEvent                                CategoryCode = "CITY_EVENT"
	CategoryCodeFestival                                 CategoryCode = "FESTIVAL"
	CategoryCodeTour                                     CategoryCode = "TOUR"
	CategoryCodeOther                                    CategoryCode = "OTHER"
	CategoryCodeNightlife                                CategoryCode = "NIGHTLIFE"
	CategoryCodeShopping                                 CategoryCode = "SHOPPING"
	CategoryCodeSports                                   CategoryCode = "SPORTS"
	CategoryCodeCityCenter                               CategoryCode = "CITY_CENTER"
	CategoryCodeCityDowntown                             CategoryCode = "CITY_DOWNTOWN"
	CategoryCodeLivetheater                              CategoryCode = "LIVETHEATER"
	CategoryCodeArena                                    CategoryCode = "ARENA"
	CategoryCodeBar                                      CategoryCode = "BAR"
	CategoryCodeBay                                      CategoryCode = "BAY"
	CategoryCodeCathedral                                CategoryCode = "CATHEDRAL"
	CategoryCodeEducationalInstitution                   CategoryCode = "EDUCATIONAL_INSTITUTION"
	CategoryCodeMedicalFacility                          CategoryCode = "MEDICAL_FACILITY"
	CategoryCodeArmyBase                                 CategoryCode = "ARMY_BASE"
	CategoryCodeCommercialDistrict                       CategoryCode = "COMMERCIAL_DISTRICT"
	CategoryCodeTouristSite                              CategoryCode = "TOURIST_SITE"
	CategoryCodeMiscellaneous                            CategoryCode = "MISCELLANEOUS"
	CategoryCodeAgricultural                             CategoryCode = "AGRICULTURAL"
	CategoryCodeArcheological                            CategoryCode = "ARCHEOLOGICAL"
	CategoryCodeBotanicalGarden                          CategoryCode = "BOTANICAL_GARDEN"
	CategoryCodeBowling                                  CategoryCode = "BOWLING"
	CategoryCodeCulturalCenter                           CategoryCode = "CULTURAL_CENTER"
	CategoryCodeEquestrianCenter                         CategoryCode = "EQUESTRIAN_CENTER"
	CategoryCodeHandicraftCenter                         CategoryCode = "HANDICRAFT_CENTER"
	CategoryCodeNaturalAttraction                        CategoryCode = "NATURAL_ATTRACTION"
	CategoryCodePerformingArtCenter                      CategoryCode = "PERFORMING_ART_CENTER"
	CategoryCodePlanetariumScienceCenter                 CategoryCode = "PLANETARIUM_SCIENCE_CENTER"
	CategoryCodeCableCars                                CategoryCode = "CABLE_CARS"
	CategoryCodeCompany                                  CategoryCode = "COMPANY"
	CategoryCodeFactoryBusinessTour                      CategoryCode = "FACTORY_BUSINESS_TOUR"
	CategoryCodeNighttimeEntertainment                   CategoryCode = "NIGHTTIME_ENTERTAINMENT"
	CategoryCodeArt                                      CategoryCode = "ART"
	CategoryCodeMusic                                    CategoryCode = "MUSIC"
	CategoryCodeStateNationalPark                        CategoryCode = "STATE_NATIONAL_PARK"
	CategoryCodeExhibitionConferenceCenter               CategoryCode = "EXHIBITION_CONFERENCE_CENTER"
	CategoryCodeAirLineDesk                              CategoryCode = "AIR_LINE_DESK"
	CategoryCodeAnimalWatching                           CategoryCode = "ANIMAL_WATCHING"
	CategoryCodeAtmCashMachine                           CategoryCode = "ATM_CASH_MACHINE"
	CategoryCodeBabySitting                              CategoryCode = "BABY_SITTING"
	CategoryCodeBaggageStorage                           CategoryCode = "BAGGAGE_STORAGE"
	CategoryCodeBallroom                                 CategoryCode = "BALLROOM"
	CategoryCodeBeachNearHotel                           CategoryCode = "BEACH_NEAR_HOTEL"
	CategoryCodeHotelWithDirectAccessToABeach            CategoryCode = "HOTEL_WITH_DIRECT_ACCESS_TO_A_BEACH"
	CategoryCodeBirdWatching                             CategoryCode = "BIRD_WATCHING"
	CategoryCodeRoomsWithBalcony                         CategoryCode = "ROOMS_WITH_BALCONY"
	CategoryCodeBoating                                  CategoryCode = "BOATING"
	CategoryCodeBeautyParlour                            CategoryCode = "BEAUTY_PARLOUR"
	CategoryCodeCoachBusParking                          CategoryCode = "COACH_BUS_PARKING"
	CategoryCodeButlerService                            CategoryCode = "BUTLER_SERVICE"
	CategoryCodeCarRental                                CategoryCode = "CAR_RENTAL"
	CategoryCodeChildrenWelcome                          CategoryCode = "CHILDREN_WELCOME"
	CategoryCodeChildrenNotAllowed                       CategoryCode = "CHILDREN_NOT_ALLOWED"
	CategoryCodeConnectingRooms                          CategoryCode = "CONNECTING_ROOMS"
	CategoryCodeConcierge                                CategoryCode = "CONCIERGE"
	CategoryCodeCourtesyCar                              CategoryCode = "COURTESY_CAR"
	CategoryCodeCellularPhoneRental                      CategoryCode = "CELLULAR_PHONE_RENTAL"
	CategoryCodeDutyFreeShop                             CategoryCode = "DUTY_FREE_SHOP"
	CategoryCodeDisco                                    CategoryCode = "DISCO"
	CategoryCodeDrivingRange                             CategoryCode = "DRIVING_RANGE"
	CategoryCodeElevator                                 CategoryCode = "ELEVATOR"
	CategoryCodeLiveEntertainment                        CategoryCode = "LIVE_ENTERTAINMENT"
	CategoryCodeCurrencyExchangeFacilities               CategoryCode = "CURRENCY_EXCHANGE_FACILITIES"
	CategoryCodeExecutiveDesk                            CategoryCode = "EXECUTIVE_DESK"
	CategoryCodeExecutiveFloor                           CategoryCode = "EXECUTIVE_FLOOR"
	CategoryCodeExpressCheckIn                           CategoryCode = "EXPRESS_CHECK_IN"
	CategoryCodeExpressCheckOut                          CategoryCode = "EXPRESS_CHECK_OUT"
	CategoryCodeFrontDeskOpen24HoursADay                 CategoryCode = "FRONT_DESK_OPEN_24_HOURS_A_DAY"
	CategoryCodeFishing                                  CategoryCode = "FISHING"
	CategoryCodeFlorist                                  CategoryCode = "FLORIST"
	CategoryCodeFreeParking                              CategoryCode = "FREE_PARKING"
	CategoryCodeFreeTransportation                       CategoryCode = "FREE_TRANSPORTATION"
	CategoryCodeGamesRoom                                CategoryCode = "GAMES_ROOM"
	CategoryCodeGarageParking                            CategoryCode = "GARAGE_PARKING"
	CategoryCodeGiftShopNewsStand                        CategoryCode = "GIFT_SHOP_NEWS_STAND"
	CategoryCodeGolf                                     CategoryCode = "GOLF"
	CategoryCodeGymNotHealthClub                         CategoryCode = "GYM_NOT_HEALTH_CLUB"
	CategoryCodeHealthClub                               CategoryCode = "HEALTH_CLUB"
	CategoryCodeHorseRiding                              CategoryCode = "HORSE_RIDING"
	CategoryCodeHotspots                                 CategoryCode = "HOTSPOTS"
	CategoryCodeFreeHighSpeedInternetConnection          CategoryCode = "FREE_HIGH_SPEED_INTERNET_CONNECTION"
	CategoryCodeHighSpeedInternetConnection              CategoryCode = "HIGH_SPEED_INTERNET_CONNECTION"
	CategoryCodeInternetServices                         CategoryCode = "INTERNET_SERVICES"
	CategoryCodeJacuzzi                                  CategoryCode = "JACUZZI"
	CategoryCodeJoggingTrack                             CategoryCode = "JOGGING_TRACK"
	CategoryCodeKennels                                  CategoryCode = "KENNELS"
	CategoryCodeLaundryService                           CategoryCode = "LAUNDRY_SERVICE"
	CategoryCodeMassage                                  CategoryCode = "MASSAGE"
	CategoryCodeMiniatureGolf                            CategoryCode = "MINIATURE_GOLF"
	CategoryCodeMultilingualStaff                        CategoryCode = "MULTILINGUAL_STAFF"
	CategoryCodeNightClub                                CategoryCode = "NIGHT_CLUB"
	CategoryCodeHotelDoesNotProvidePornographicFilmsTv   CategoryCode = "HOTEL_DOES_NOT_PROVIDE_PORNOGRAPHIC_FILMS_TV"
	CategoryCodeNursery                                  CategoryCode = "NURSERY"
	CategoryCodeParking                                  CategoryCode = "PARKING"
	CategoryCodePetsAllowed                              CategoryCode = "PETS_ALLOWED"
	CategoryCodePharmacy                                 CategoryCode = "PHARMACY"
	CategoryCodeChildrenSPlayArea                        CategoryCode = "CHILDREN_S_PLAY_AREA"
	CategoryCodePorterBellBoy                            CategoryCode = "PORTER_BELL_BOY"
	CategoryCodePuttingGreen                             CategoryCode = "PUTTING_GREEN"
	CategoryCodeSauna                                    CategoryCode = "SAUNA"
	CategoryCodeScubaDiving                              CategoryCode = "SCUBA_DIVING"
	CategoryCodeFreeAirportShuttle                       CategoryCode = "FREE_AIRPORT_SHUTTLE"
	CategoryCodeIndoorSwimmingPool                       CategoryCode = "INDOOR_SWIMMING_POOL"
	CategoryCodeSightseeing                              CategoryCode = "SIGHTSEEING"
	CategoryCodeSkeetShooting                            CategoryCode = "SKEET_SHOOTING"
	CategoryCodeHotelWithSkiInOutFacilities              CategoryCode = "HOTEL_WITH_SKI_IN_OUT_FACILITIES"
	CategoryCodeSnowSkiing                               CategoryCode = "SNOW_SKIING"
	CategoryCodeSolarium                                 CategoryCode = "SOLARIUM"
	CategoryCodeSpa                                      CategoryCode = "SPA"
	CategoryCodeHeatedSwimmingPool                       CategoryCode = "HEATED_SWIMMING_POOL"
	CategoryCodeSwimmingPool                             CategoryCode = "SWIMMING_POOL"
	CategoryCodeIndoorTennis                             CategoryCode = "INDOOR_TENNIS"
	CategoryCodeTennis                                   CategoryCode = "TENNIS"
	CategoryCodeTennisProfessional                       CategoryCode = "TENNIS_PROFESSIONAL"
	CategoryCodeTheatreDesk                              CategoryCode = "THEATRE_DESK"
	CategoryCodeTourDesk                                 CategoryCode = "TOUR_DESK"
	CategoryCodeTranslationServices                      CategoryCode = "TRANSLATION_SERVICES"
	CategoryCodeTravelAgency                             CategoryCode = "TRAVEL_AGENCY"
	CategoryCodeValetParking                             CategoryCode = "VALET_PARKING"
	CategoryCodeVendingMachines                          CategoryCode = "VENDING_MACHINES"
	CategoryCodeVolleyball                               CategoryCode = "VOLLEYBALL"
	CategoryCodeWaterSports                              CategoryCode = "WATER_SPORTS"
	CategoryCodeWirelessConnectivity                     CategoryCode = "WIRELESS_CONNECTIVITY"
	CategoryCodeWeddingServices                          CategoryCode = "WEDDING_SERVICES"
	CategoryCodeHairDresser                              CategoryCode = "HAIR_DRESSER"
	CategoryCodeBusinessServices                         CategoryCode = "BUSINESS_SERVICES"
	CategoryCodeAccessibleFacilities                     CategoryCode = "ACCESSIBLE_FACILITIES"
	CategoryCodeSecurity                                 CategoryCode = "SECURITY"
	CategoryCodeGroupRates                               CategoryCode = "GROUP_RATES"
	CategoryCode24HourSecurity                           CategoryCode = "24_HOUR_SECURITY"
	CategoryCodePhotocopyCenter                          CategoryCode = "PHOTOCOPY_CENTER"
	CategoryCodeVideoTapes                               CategoryCode = "VIDEO_TAPES"
	CategoryCodeWakeupService                            CategoryCode = "WAKEUP_SERVICE"
	CategoryCodeDirectDialTelephone                      CategoryCode = "DIRECT_DIAL_TELEPHONE"
	CategoryCodeEarlyCheckIn                             CategoryCode = "EARLY_CHECK_IN"
	CategoryCodeBicycleRentals                           CategoryCode = "BICYCLE_RENTALS"
	CategoryCodeLateCheckOutAvailable                    CategoryCode = "LATE_CHECK_OUT_AVAILABLE"
	CategoryCodeBookstore                                CategoryCode = "BOOKSTORE"
	CategoryCodeComplimentarySelfServiceLaundry          CategoryCode = "COMPLIMENTARY_SELF_SERVICE_LAUNDRY"
	CategoryCodeAccessibleParking                        CategoryCode = "ACCESSIBLE_PARKING"
	CategoryCodeBoutiquesStores                          CategoryCode = "BOUTIQUES_STORES"
	CategoryCodeShopsAndCommercialServices               CategoryCode = "SHOPS_AND_COMMERCIAL_SERVICES"
	CategoryCodeSportsBarOpenForLunch                    CategoryCode = "SPORTS_BAR_OPEN_FOR_LUNCH"
	CategoryCodeComplimentaryCoffeeInLobby               CategoryCode = "COMPLIMENTARY_COFFEE_IN_LOBBY"
	CategoryCodeDinnerDeliveryServiceFromLocalRestaurant CategoryCode = "DINNER_DELIVERY_SERVICE_FROM_LOCAL_RESTAURANT"
	CategoryCodeComplimentaryNewspaperInLobby            CategoryCode = "COMPLIMENARY_NEWSPAPER_IN_LOBBY"
	CategoryCodeFrontDesk                                CategoryCode = "FRONT_DESK"
	CategoryCodeGroceryShoppingServiceAvailable          CategoryCode = "GROCERY_SHOPPING_SERVICE_AVAILABLE"
	CategoryCodeManagersReception                        CategoryCode = "MANAGERS_RECEPTION"
	CategoryCodeMedicalFacilitiesService                 CategoryCode = "MEDICAL_FACILITIES_SERVICE"
	CategoryCodeAllInclusiveMealPlan                     CategoryCode = "ALL_INCLUSIVE_MEAL_PLAN"
	CategoryCodeCommunalBarArea                          CategoryCode = "COMMUNAL_BAR_AREA"
	CategoryCodeContinentalBreakfast                     CategoryCode = "CONTINENTAL_BREAKFAST"
	CategoryCodeFullMealPlan                             CategoryCode = "FULL_MEAL_PLAN"
	CategoryCodeOnsiteLaundry                            CategoryCode = "ONSITE_LAUNDRY"
	CategoryCode24HourFoodBeverageKiosk                  CategoryCode = "24_HOUR_FOOD_BEVERAGE_KIOSK"
	CategoryCodeFullServiceHousekeeping                  CategoryCode = "FULL_SERVICE_HOUSEKEEPING"
	CategoryCodeAdditionalServicesAmenitiesFacilities    CategoryCode = "ADDITIONAL_SERVICES_AMENITIES_FACILITIES_ON_PROPERTY"
	CategoryCodeDvdVideoRental                           CategoryCode = "DVD_VIDEO_RENTAL"
	CategoryCodeParkingLot                               CategoryCode = "PARKING_LOT"
	CategoryCodeCocktailLoungeWithEntertainment          CategoryCode = "COCKTAIL_LOUNGE_WITH_ENTERTAINMENT"
	CategoryCodeCocktailLounge                           CategoryCode = "COCKTAIL_LOUNGE"
	CategoryCodePhoneServices                            CategoryCode = "PHONE_SERVICES"
	CategoryCodeAerobicsInstruction                      CategoryCode = "AEROBICS_INSTRUCTION"
	CategoryCodeCoinOperatedLaundry                      CategoryCode = "COIN_OPERATED_LAUNDRY"
	CategoryCodeBankingServices                          CategoryCode = "BANKING_SERVICES"
	CategoryCodeExhibitionConventionFloor                CategoryCode = "EXHIBITION_CONVENTION_FLOOR"
	CategoryCodeCourtyard                                CategoryCode = "COURTYARD"
	CategoryCodeDoorMan                                  CategoryCode = "DOOR_MAN"
	CategoryCodeDrugstorePharmacy                        CategoryCode = "DRUGSTORE_PHARMACY"
	CategoryCodeHousekeepingDaily                        CategoryCode = "HOUSEKEEPING_DAILY"
	CategoryCodeOffSiteParking                           CategoryCode = "OFF_SITE_PARKING"
	CategoryCodeOnSiteParking                            CategoryCode = "ON_SITE_PARKING"
	CategoryCodeOutdoorParking                           CategoryCode = "OUTDOOR_PARKING"
	CategoryCodeRampAccess                               CategoryCode = "RAMP_ACCESS"
	CategoryCodeSportsBar                                CategoryCode = "SPORTS_BAR"
	CategoryCodeValetDryCleaning                         CategoryCode = "VALET_DRY_CLEANING"
	CategoryCodeChildrensProgramOnsite                   CategoryCode = "CHILDRENS_PROGRAM_ONSITE"
	CategoryCodeWindsurfing                              CategoryCode = "WINDSURFING"
	CategoryCodeCamping                                  CategoryCode = "CAMPING"
	CategoryCodeHunting                                  CategoryCode = "HUNTING"
	CategoryCodeIndoorOutdoorConnectingPool              CategoryCode = "INDOOR_OUTDOOR_CONNECTING_POOL"
	CategoryCodeMountainClimbing                         CategoryCode = "MOUNTAIN_CLIMBING"
	CategoryCodeNaturePreserveTrail                      CategoryCode = "NATURE_PRESERVE_TRAIL"
	CategoryCodeBilliards                                CategoryCode = "BILLIARDS"
	CategoryCodeSunTanningBed                            CategoryCode = "SUN_TANNING_BED"
	CategoryCodeSurfing                                  CategoryCode = "SURFING"
	CategoryCodeTableTennis                              CategoryCode = "TABLE_TENNIS"
	CategoryCodeTeenPrograms                             CategoryCode = "TEEN_PROGRAMS"
	CategoryCodeIndoorPool                               CategoryCode = "INDOOR_POOL"
	CategoryCodeOutdoorPool                              CategoryCode = "OUTDOOR_POOL"
	CategoryCodeChildrensProgram                         CategoryCode = "CHILDRENS_PROGRAM"
	CategoryCodeBoxing                                   CategoryCode = "BOXING"
	CategoryCodeChildrensPool                            CategoryCode = "CHILDRENS_POOL"
	CategoryCodeDancing                                  CategoryCode = "DANCING"
	CategoryCodeGarden                                   CategoryCode = "GARDEN"
	CategoryCodeKaraoke                                  CategoryCode = "KARAOKE"
	CategoryCodeMuseumGalleryViewing                     CategoryCode = "MUSEUM_GALLERY_VIEWING"
	CategoryCodeNightclubs                               CategoryCode = "NIGHTCLUBS"
	CategoryCodeSportsEvents                             CategoryCode = "SPORTS_EVENTS"
	CategoryCodeSkydiving                                CategoryCode = "SKYDIVING"
	CategoryCodeSunbathing                               CategoryCode = "SUNBATHING"
	CategoryCodeTheatre                                  CategoryCode = "THEATRE"
	CategoryCodeFitnessCenterOffSite                     CategoryCode = "FITNESS_CENTER_OFF_SITE"
	CategoryCodeFlyFishing                               CategoryCode = "FLY_FISHING"
	CategoryCodeBaseballDiamond                          CategoryCode = "BASEBALL_DIAMOND"
	CategoryCodeGym                                      CategoryCode = "GYM"
	CategoryCodeBasketballCourt                          CategoryCode = "BASKETBALL_COURT"
	CategoryCodeBikeTrail                                CategoryCode = "BIKE_TRAIL"
	CategoryCodeHikingTrail                              CategoryCode = "HIKING_TRAIL"
	CategoryCodeJoggingTrail                             CategoryCode = "JOGGING_TRAIL"
	CategoryCodeKayaking                                 CategoryCode = "KAYAKING"
	CategoryCodeMountainBikingTrail                      CategoryCode = "MOUNTAIN_BIKING_TRAIL"
	CategoryCodeParasailing                              CategoryCode = "PARASAILING"
	CategoryCodePlayground                               CategoryCode = "PLAYGROUND"
	CategoryCodePool                                     CategoryCode = "POOL"
	CategoryCodeRiverRafting                             CategoryCode = "RIVER_RAFTING"
	CategoryCodeSailing                                  CategoryCode = "SAILING"
	CategoryCodeSnorkeling                               CategoryCode = "SNORKELING"
	CategoryCodeTennisCourt                              CategoryCode = "TENNIS_COURT"
	CategoryCodeWaterSkiing                              CategoryCode = "WATER_SKIING"
	CategoryCodeFineDining                               CategoryCode = "FINE_DINING"
	CategoryCodeGolfLocation                             CategoryCode = "GOLF_LOCATION"
	CategoryCodeBilingualStaff                           CategoryCode = "BILINGUAL_STAFF"
	CategoryCodeAirConditioning                          CategoryCode = "AIR_CONDITIONING"
	CategoryCodeNonSmokingRooms                          CategoryCode = "NON_SMOKING_ROOMS"
	CategoryCodeInternetAccess                           CategoryCode = "INTERNET_ACCESS"
	CategoryCodeSundryConvenienceStore                   CategoryCode = "SUNDRY_CONVENIENCE_STORE"
	CategoryCodeTransportation                           CategoryCode = "TRANSPORTATION"
	CategoryCodeComplimentaryBreakfast                   CategoryCode = "COMPLIMENTARY_BREAKFAST"
	CategoryCodeHighSpeedInternetAccess                  CategoryCode = "HIGH_SPEED_INTERNET_ACCESS"
	CategoryCodeLobby                                    CategoryCode = "LOBBY"
	CategoryCode24HourCoffeeShop                         CategoryCode = "24_HOUR_COFFEE_SHOP"
	CategoryCodeAirportShuttleService                    CategoryCode = "AIRPORT_SHUTTLE_SERVICE"
	CategoryCodeLuggageService                           CategoryCode = "LUGGAGE_SERVICE"
	CategoryCodePianoBar                                 CategoryCode = "PIANO_BAR"
	CategoryCodeVipSecurity                              CategoryCode = "VIP_SECURITY"
	CategoryCodeWheelChairAccess                         CategoryCode = "WHEEL_CHAIR_ACCESS"
	CategoryCodeBusinessCenter                           CategoryCode = "BUSINESS_CENTER"
	CategoryCodeChildPrograms                            CategoryCode = "CHILD_PROGRAMS"
	CategoryCodeSeaside                                  CategoryCode = "SEASIDE"
	CategoryCodePrivateDiningForGroups                   CategoryCode = "PRIVATE_DINING_FOR_GROUPS"
	CategoryCodeHighSpeedWireless                        CategoryCode = "HIGH_SPEED_WIRELESS"
	CategoryCodePrinter                                  CategoryCode = "PRINTER"
	CategoryCodeIfGuestRoomsHaveMoreThanOnePhoneLine     CategoryCode = "IF_GUEST_ROOMS_HAVE_MORE_THAN_ONE_PHONE_LINE"
	CategoryCodeComplimentaryWirelessInternet            CategoryCode = "COMPLIMENTARY_WIRELESS_INTERNET"
	CategoryCodeSameGenderFloor                          CategoryCode = "SAME_GENDER_FLOOR"
	CategoryCodeChildrenPrograms                         CategoryCode = "CHILDREN_PROGRAMS"
	CategoryCodeBuildingMeetsLocal                       CategoryCode = "BUILDING_MEETS_LOCAL"
	CategoryCodeInternetBrowserOnTv                      CategoryCode = "INTERNET_BROWSER_ON_TV"
	CategoryCodeNewspaper                                CategoryCode = "NEWSPAPER"
	CategoryCodeParkingControlledAccessGates             CategoryCode = "PARKING_CONTROLLED_ACCESS_GATES_TO_ENTER_PARKING_AREA"
	CategoryCodeHotelSafeDepositBoxNotRoomSafeBox        CategoryCode = "HOTEL_SAFE_DEPOSIT_BOX__NOT_ROOM_SAFE_BOX"
	CategoryCodeStorageSpaceAvailableForFee              CategoryCode = "STORAGE_SPACE_AVAILABLE_FOR_FEE"
	CategoryCodeTypeOfEntranceToGuestRoom                CategoryCode = "TYPE_OF_ENTRANCE_TO_GUEST_ROOM"
	CategoryCodeBeverageCocktail                         CategoryCode = "BEVERAGE_COCKTAIL"
	CategoryCodeCellPhoneRental                          CategoryCode = "CELL_PHONE_RENTAL"
	CategoryCodeCoffeeTea                                CategoryCode = "COFFEE_TEA"
	CategoryCodeEarlyCheckInGuarantee                    CategoryCode = "EARLY_CHECK_IN_GUARANTEE"
	CategoryCodeFoodAndBeverageDiscount                  CategoryCode = "FOOD_AND_BEVERAGE_DISCOUNT"
	CategoryCodeLateCheckOutGuarantee                    CategoryCode = "LATE_CHECK_OUT_GUARANTEE"
	CategoryCodeRoomUpgradeConfirmed                     CategoryCode = "ROOM_UPGRADE_CONFIRMED"
	CategoryCodeRoomUpgradeOnAvailability                CategoryCode = "ROOM_UPGRADE_ON_AVAILABILITY"
	CategoryCodeShuttleToLocalBusinesses                 CategoryCode = "SHUTTLE_TO_LOCAL_BUSINESSES"
	CategoryCodeShuttleToLocalAttractions                CategoryCode = "SHUTTLE_TO_LOCAL_ATTRACTIONS"
	CategoryCodeSocialHour                               CategoryCode = "SOCIAL_HOUR"
	CategoryCodeVideoBilling                             CategoryCode = "VIDEO_BILLING"
	CategoryCodeWelcomeGift                              CategoryCode = "WELCOME_GIFT"
	CategoryCodeHypoallergenicRooms                      CategoryCode = "HYPOALLERGENIC_ROOMS"
	CategoryCodeRoomAirFiltration                        CategoryCode = "ROOM_AIR_FILTRATION"
	CategoryCodeSmokeFreeProperty                        CategoryCode = "SMOKE_FREE_PROPERTY"
	CategoryCodeWaterPurificationSystemInUse             CategoryCode = "WATER_PURIFICATION_SYSTEM_IN_USE"
	CategoryCodePoolsideService                          CategoryCode = "POOLSIDE_SERVICE"
	CategoryCodeClothingStore                            CategoryCode = "CLOTHING_STORE"
	CategoryCodeEvElectricVehicleChargingLocation        CategoryCode = "EV_ELECTRIC_VEHICLE_CHARGING_LOCATION"
	CategoryCodeOfficeRental                             CategoryCode = "OFFICE_RENTAL"
	CategoryCodeIncomingFax                              CategoryCode = "INCOMING_FAX"
	CategoryCodeOutgoingFax                              CategoryCode = "OUTGOING_FAX"
	CategoryCodeBabyKit                                  CategoryCode = "BABY_KIT"
	CategoryCodeChildrenSBreakfast                       CategoryCode = "CHILDREN_S_BREAKFAST"
	CategoryCodeCloakroomService                         CategoryCode = "CLOAKROOM_SERVICE"
	CategoryCodeCoffeeLounge                             CategoryCode = "COFFEE_LOUNGE"
	CategoryCodeEventsTicketService                      CategoryCode = "EVENTS_TICKET_SERVICE"
	CategoryCodeLateCheckIn                              CategoryCode = "LATE_CHECK_IN"
	CategoryCodeLimitedParking                           CategoryCode = "LIMITED_PARKING"
	CategoryCodeOutdoorSummerBarCafe                     CategoryCode = "OUTDOOR_SUMMER_BAR_CAFE"
	CategoryCodeNoParkingAvailable                       CategoryCode = "NO_PARKING_AVAILABLE"
	CategoryCodeBeerGarden                               CategoryCode = "BEER_GARDEN"
	CategoryCodeGardenLoungeBar                          CategoryCode = "GARDEN_LOUNGE_BAR"
	CategoryCodeSummerTerrace                            CategoryCode = "SUMMER_TERRACE"
	CategoryCodeWinterTerrace                            CategoryCode = "WINTER_TERRACE"
	CategoryCodeRoofTerrace                              CategoryCode = "ROOF_TERRACE"
	CategoryCodeBeachBar                                 CategoryCode = "BEACH_BAR"
	CategoryCodeHelicopterService                        CategoryCode = "HELICOPTER_SERVICE"
	CategoryCodeFerry                                    CategoryCode = "FERRY"
	CategoryCodeTapasBar                                 CategoryCode = "TAPAS_BAR"
	CategoryCodeCafeBar                                  CategoryCode = "CAFE_BAR"
	CategoryCodeSnackBar                                 CategoryCode = "SNACK_BAR"
	CategoryCodeEnhancedSafetyProtocol                   CategoryCode = "ENHANCED_SAFETY_PROTOCOL"
	CategoryCodeBusinessLibrary                          CategoryCode = "BUSINESS_LIBRARY"
	CategoryCodeCheckInKioskAvailable                    CategoryCode = "CHECK_IN_KIOSK_AVAILABLE"
	CategoryCodeConciergeFloor                           CategoryCode = "CONCIERGE_FLOOR"
	CategoryCodeHousekeepingWeekly                       CategoryCode = "HOUSEKEEPING_WEEKLY"
	CategoryCodePackageReceiving                         CategoryCode = "PACKAGE_RECEIVING"
	CategoryCodePublicAddressSystem                      CategoryCode = "PUBLIC_ADDRESS_SYSTEM"
	CategoryCodeShoeShine                                CategoryCode = "SHOE_SHINE"
	CategoryCodeStorageSpaceAvailable                    CategoryCode = "STORAGE_SPACE_AVAILABLE"
	CategoryCodeTechnicalConciergeAvailable              CategoryCode = "TECHNICAL_CONCIERGE_AVAILABLE"
	CategoryCodeTruckParking                             CategoryCode = "TRUCK_PARKING"
	CategoryCodeWakeUpCalls                              CategoryCode = "WAKE_UP_CALLS"
	CategoryCodeVideoGames                               CategoryCode = "VIDEO_GAMES"
	CategoryCodeRoomServiceLimitedHours                  CategoryCode = "ROOM_SERVICE_LIMITED_HOURS"
	CategoryCodePublicAreasAirConditioned                CategoryCode = "PUBLIC_AREAS_AIR_CONDITIONED"
	CategoryCodeComplimentaryInRoomCoffeeOrTea           CategoryCode = "COMPLIMENTARY_IN_ROOM_COFFEE_OR_TEA"
	CategoryCodeComplimentaryBuffetBreakfast             CategoryCode = "COMPLIMENTARY_BUFFET_BREAKFAST"
	CategoryCodeComplimentaryContinentalBreakfast        CategoryCode = "COMPLIMENTARY_CONTINENTAL_BREAKFAST"
	CategoryCodeLimousineService                         CategoryCode = "LIMOUSINE_SERVICE"
	CategoryCodeTelephoneJackAdaptorAvailable            CategoryCode = "TELEPHONE_JACK_ADAPTOR_AVAILABLE"
	CategoryCodeBreakfastFull                            CategoryCode = "BREAKFAST_FULL"
	CategoryCodeVipLounge                                CategoryCode = "VIP_LOUNGE"
	CategoryCodeParkingFeeManagedByTheHotel              CategoryCode = "PARKING_FEE_MANAGED_BY_THE_HOTEL"
	CategoryCodeHousekeepingLimited                      CategoryCode = "HOUSEKEEPING_LIMITED"
	CategoryCodeTransportationServicesLocalArea          CategoryCode = "TRANSPORTATION_SERVICES_LOCAL_AREA"
	CategoryCodeTransportationServicesLocalOffice        CategoryCode = "TRANSPORTATION_SERVICES_LOCAL_OFFICE"
	CategoryCodeParkingDeck                              CategoryCode = "PARKING_DECK"
	CategoryCodeParkingSideStreet                        CategoryCode = "PARKING_SIDE_STREET"
	CategoryCodeCocktailLoungeWithLightFare              CategoryCode = "COCKTAIL_LOUNGE_WITH_LIGHT_FARE"
	CategoryCodeMotorcycleParking                        CategoryCode = "MOTORCYCLE_PARKING"
	CategoryCodePersonalTrainer                          CategoryCode = "PERSONAL_TRAINER"
	CategoryCodeJetskiing                                CategoryCode = "JETSKIING"
	CategoryCodeRacquetballCourt                         CategoryCode = "RACQUETBALLCOURT"
	CategoryCodeSquashCourts                             CategoryCode = "SQUASHCOURTS"
	CategoryCodeSteamBath                                CategoryCode = "STEAM_BATH"
	CategoryCodeWhirlpool                                CategoryCode = "WHIRLPOOL"
	CategoryCodeSafari                                   CategoryCode = "SAFARI"
	CategoryCodeRecreationSportsCourt                    CategoryCode = "RECREATION_SPORTS_COURT"
	CategoryCodeSnowmobiling                             CategoryCode = "SNOWMOBILING"
	CategoryCodePolo                                     CategoryCode = "POLO"
	CategoryCodeWeightliftingEquipment                   CategoryCode = "WEIGHTLIFTINGEQUIPMENT"
	CategoryCodeCardiovascularEquipment                  CategoryCode = "CARDIOVASCULAREQUIPMENT"
	CategoryCodeExtensiveHealthClub                      CategoryCode = "EXTENSIVEHEALTHCLUB"
	CategoryCodeLimitedHealthClub                        CategoryCode = "LIMITEDHEALTHCLUB"
	CategoryCodeDiving                                   CategoryCode = "DIVING"
	CategoryCodeWalkingTrack                             CategoryCode = "WALKING_TRACK"
	CategoryCodePaddleCourt                              CategoryCode = "PADDLE_COURT"
	CategoryCodeBoatTours                                CategoryCode = "BOAT_TOURS"
	CategoryCodeKidsGolfAcademy                          CategoryCode = "KIDS_GOLF_ACADEMY"
	CategoryCodeKidsBeachClub                            CategoryCode = "KIDS_BEACH_CLUB"
	CategoryCodeKidsEquestrianClub                       CategoryCode = "KIDS_EQUESTRIAN_CLUB"
	CategoryCodeLounge                                   CategoryCode = "LOUNGE"

	// Values observed from Hotel Content that the original list omitted. Note
	// that CategoryCode is not a closed set: for airport points of interest the
	// API returns the IATA airport code itself (CDG, ORY, BVA, LBG, ...), so
	// callers must tolerate values with no matching constant here.
	CategoryCode24HourFoodOrBeverageKiosk CategoryCode = "24_HOUR_FOOD_OR_BEVERAGE_KIOSK"
	CategoryCodeEntertainmentDistrict     CategoryCode = "ENTERTAINMENT_DISTRICT"
	CategoryCodeTheaterOrCinema           CategoryCode = "THEATER_OR_CINEMA"
)

type TransportMode string

// Enum values for TransportMode
const (
	Bicycle          TransportMode = "BICYCLE"
	Boat             TransportMode = "BOAT"
	Bus              TransportMode = "BUS"
	CableCar         TransportMode = "CABLE_CAR"
	CourtesyCar      TransportMode = "COURTESY_CAR"
	Car              TransportMode = "CAR"
	Carriage         TransportMode = "CARRIAGE"
	Helicopter       TransportMode = "HELICOPTER"
	Limousine        TransportMode = "LIMOUSINE"
	Metro            TransportMode = "METRO"
	Monorail         TransportMode = "MONORAIL"
	Motorbike        TransportMode = "MOTORBIKE"
	PackAnimal       TransportMode = "PACK_ANIMAL"
	Plane            TransportMode = "PLANE"
	Rickshaw         TransportMode = "RICKSHAW"
	Shuttle          TransportMode = "SHUTTLE"
	SedanChair       TransportMode = "SEDAN_CHAIR"
	Subway           TransportMode = "SUBWAY"
	Taxi             TransportMode = "TAXI"
	Train            TransportMode = "TRAIN"
	Walk             TransportMode = "WALK"
	WaterTaxi        TransportMode = "WATER_TAXI"
	OtherOrAlternate TransportMode = "OTHER_OR_ALTERNATE"
	ExpressTrain     TransportMode = "EXPRESS_TRAIN"
	Alternate        TransportMode = "ALTERNATE"
	Ferry            TransportMode = "FERRY"
)

type DistanceType string

// Enum values for DistanceType
const (
	Airways   DistanceType = "AIRWAYS"
	Roadways  DistanceType = "ROADWAYS"
	Railways  DistanceType = "RAILWAYS"
	Waterways DistanceType = "WATERWAYS"
	Birdseye  DistanceType = "BIRDSEYE"
)

type Segment string

// Enum values for Segment
const (
	SegmentAllSuite                   Segment = "ALL_SUITE"
	SegmentBudget                     Segment = "BUDGET"
	SegmentCorporateBusinessTransient Segment = "CORPORATE_BUSINESS_TRANSIENT"
	SegmentDeluxe                     Segment = "DELUXE"
	SegmentEconomy                    Segment = "ECONOMY"
	SegmentExtendedStay               Segment = "EXTENDED_STAY"
	SegmentFirstClass                 Segment = "FIRST_CLASS"
	SegmentLuxury                     Segment = "LUXURY"
	SegmentMeetingOrConvention        Segment = "MEETING_OR_CONVENTION"
	SegmentModerate                   Segment = "MODERATE"
	SegmentResidentialApartment       Segment = "RESIDENTIAL_APARTMENT"
	SegmentResort                     Segment = "RESORT"
	SegmentTourist                    Segment = "TOURIST"
	SegmentUpscale                    Segment = "UPSCALE"
	SegmentEfficiency                 Segment = "EFFICIENCY"
	SegmentStandard                   Segment = "STANDARD"
	SegmentMidscale                   Segment = "MIDSCALE"
	SegmentQuality                    Segment = "QUALITY"
	SegmentUnknown                    Segment = "UNKNOWN"
	SegmentMidscaleWithoutFAndB       Segment = "MIDSCALE_WITHOUT_F_AND_B"
	SegmentUpperUpscale               Segment = "UPPER_UPSCALE"
)

type HotelAreaType string

const (
	AreaAirport               HotelAreaType = "AIRPORT"
	AreaBeach                 HotelAreaType = "BEACH"
	AreaCity                  HotelAreaType = "CITY"
	AreaDowntown              HotelAreaType = "DOWNTOWN"
	AreaEast                  HotelAreaType = "EAST"
	AreaExpressway            HotelAreaType = "EXPRESSWAY"
	AreaLake                  HotelAreaType = "LAKE"
	AreaMountain              HotelAreaType = "MOUNTAIN"
	AreaNorth                 HotelAreaType = "NORTH"
	AreaResort                HotelAreaType = "RESORT"
	AreaRural                 HotelAreaType = "RURAL"
	AreaSouth                 HotelAreaType = "SOUTH"
	AreaSuburban              HotelAreaType = "SUBURBAN"
	AreaWest                  HotelAreaType = "WEST"
	AreaBeachfront            HotelAreaType = "BEACHFRONT"
	AreaOceanfront            HotelAreaType = "OCEANFRONT"
	AreaGulf                  HotelAreaType = "GULF"
	AreaBusinessDistrict      HotelAreaType = "BUSINESS_DISTRICT"
	AreaEntertainmentDistrict HotelAreaType = "ENTERTAINMENT_DISTRICT"
	AreaFinancialDistrict     HotelAreaType = "FINANCIAL_DISTRICT"
	AreaShoppingDistrict      HotelAreaType = "SHOPPING_DISTRICT"
	AreaTheatreDistrict       HotelAreaType = "THEATRE_DISTRICT"
	AreaCountryside           HotelAreaType = "COUNTRYSIDE"
	AreaBay                   HotelAreaType = "BAY"
	AreaMarina                HotelAreaType = "MARINA"
	AreaPark                  HotelAreaType = "PARK"
	AreaRiver                 HotelAreaType = "RIVER"
	AreaTouristSite           HotelAreaType = "TOURIST_SITE"
	AreaNorthSuburb           HotelAreaType = "NORTH_SUBURB"
	AreaSouthSuburb           HotelAreaType = "SOUTH_SUBURB"
	AreaEastSuburb            HotelAreaType = "EAST_SUBURB"
	AreaWestSuburb            HotelAreaType = "WEST_SUBURB"
	AreaWaterfront            HotelAreaType = "WATERFRONT"
	AreaSkiResort             HotelAreaType = "SKI_RESORT"
)

type HotelStatus string

const (
	StatusOpen                        HotelStatus = "OPEN"
	StatusClosed                      HotelStatus = "CLOSED"
	StatusPreOpening                  HotelStatus = "PRE_OPENING"
	StatusTest                        HotelStatus = "TEST"
	StatusPropertySuitableForChildren HotelStatus = "PROPERTY_SUITABLE_FOR_CHILDREN"
	StatusDeleted                     HotelStatus = "DELETED"
	StatusLocked                      HotelStatus = "LOCKED"
	StatusUnlocked                    HotelStatus = "UNLOCKED"
)

type (
	// * A point of interest is a specific location that someone may find useful or interesting that tourists visit, typically for its inherent or an exhibited natural or cultural value, historical significance, natural or built beauty, offering leisure and amusement.
	PointOfInterestResponse struct {
		Location            LocationResponse                `json:"location"`            // Indicates the location of the point of interest
		CategoryCode        CategoryCode                    `json:"categoryCode"`        // Indicates the category to which the points of interest belongs to. It can contain values such as
		Description         string                          `json:"description"`         // Description of the point of interest
		Season              Period                          `json:"season"`              // Models a period of time between two dates and inclusive only of the days of the week specified.
		Contact             ContactResponse                 `json:"contact"`             // A contact refers to the information that can be used to reach a person, a company or an organization.
		EligibilityForEntry []dto.QualifiedFreeTextResponse `json:"eligibilityForEntry"` // Indicates the eligibility for entry to the point of interest
		OfficialWebsite     struct {
			Url string `json:"url"` // Indicates the URL of the website
		} `json:"officialWebsite"` // Indicates the official website of the point of interest
		OperatingHours   CalendarScheduleResponse  `json:"operatingHours"`   // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
		PriceEquation    []PricingEquationResponse `json:"priceEquation"`    // Indicates the price equation of the point of interest
		PriceQuotation   []PriceQuotationResponse  `json:"priceQuotation"`   // Indicates the price quotations to access the point of interest
		Transportations  []TransportationResponse  `json:"transportations"`  // Indicates the transportation info to reach the point of interest via various transportation modes
		LocationDistance LocationDistanceResponse  `json:"locationDistance"` // Indicates the location distance of the point of interest
		Media            []dto.MediaResponse       `json:"media"`            // Indicates the media of the point of interest
		Hotel            HotelResponse             `json:"hotel"`            // Provides Information related to Hotel Calendar, Climate and Spoken Language.
		Basic            BasicResponse             `json:"basic"`            // By default, this Model would be returned in all successful cases. This information provides Information related to Hotel Name, Chain name and Hotel Id.
	}

	// BasicResponse struct {
	// 	Season struct {
	// 		OpenCalendar []Period `json:"openCalendar"` // Indicates the opening time and days of the property
	// 	} `json:"season"` // Indicates the season of the point of interest
	// 	HotelID                      string                         `json:"hotelId"`                      // Amadeus Property Code (8 chars). example: ADPAR001
	// 	ChainCode                    string                         `json:"chainCode"`                    // Brand (RT...) or Merchant (AD...)
	// 	BrandCode                    string                         `json:"brandCode"`                    // Brand (RT...) (Amadeus 2 chars Code). Small Properties distributed by Merchants may not have a Brand. Example - AD (Value Hotels) is the Provider/Merchant, and RT (Accor) is the Brand of the Property
	// 	DupeID                       string                         `json:"dupeId"`                       // Unique Property identifier of the physical hotel. One physical hotel can be represented by different Providers, each one having its own hotelID. This attribute allows a client application to group together hotels that are actually the same.
	// 	Name                         string                         `json:"name"`                         // Name of the point of interest
	// 	Rating                       string                         `json:"rating"`                       // Rating of the point of interest
	// 	Description                  QualifiedFreeTextResponse      `json:"description"`                  // Description of the point of interest
	// 	Amenities                    []AmenityResponse              `json:"amenities"`                    // Amenities of the point of interest
	// 	Media                        []MediaResponse                `json:"media"`                        // Media of the point of interest
	// 	DefaultSpokenLanguage        string                         `json:"defaultSpokenLanguage"`        // Describes the default language preferred or used at the property
	// 	ContextProvider              string                         `json:"contextProvider"`              // Describes the provider of the context of the point of interest
	// 	Contact                      []ContactResponse              `json:"contact"`                      // Contact of the point of interest
	// 	Location                     LocationResponse               `json:"location"`                     // Location of the point of interest
	// 	Altitude                     AltitudeResponse               `json:"altitude"`                     // From analytics, Metrics describe the exact numbers that make up the data
	// 	CategoryCode                 CategoryCode                   `json:"categoryCode"`                 // Category code of the point of interest
	// 	Segment                      Segment                        `json:"segment"`                      // Segment of the point of interest
	// 	Area                         []AreaResponse                 `json:"area"`                         // Geographical zone like City, Region, Country
	// 	ChainName                    string                         `json:"chainName"`                    // Name of the chain to which the hotel belongs to
	// 	BrandName                    string                         `json:"brandName"`                    // Name of the brand to which the hotel or hotel chain belongs to
	// 	Status                       HotelStatus                    `json:"status"`                       // Status of the hotel
	// 	HotelBusinessIdentifications BusinessIdentificationResponse `json:"hotelBusinessIdentifications"` // An business, can be idenfified via business identifiers, those business identifiers are defined by a body of authority thay could be local, national, transnational or supranational (like EU for the EU VAT number).
	// }

	// BusinessIdentificationResponse struct {
	// 	Identifiers []IdentifierResponse `json:"identifiers"` // Identifiers of the business
	// }

	// IdentifierResponse struct {
	// 	ID   string `json:"id"`   // Identifier id
	// 	Name string `json:"name"` // Identifier name
	// }

	// AreaResponse struct {
	// 	HotelAreaType HotelAreaType `json:"hotelAreaType"` // 'Indicates the category of the location. OTA Code Set LOC values are to be considered here. Can contain values such as
	// 	Name          string        `json:"name"`          // Label associated to the location (e.g. Eiffel Tower, Madison Square)
	// }

	// AltitudeResponse struct {
	// 	Unit  Unit `json:"unit"`  // Indicates the unit of the altitude
	// 	Value int  `json:"value"` // Indicates the value of the altitude
	// }

	// HotelResponse struct {
	// 	TaxID            string                     `json:"taxId"`        // Describes the unique tax identifier of a hotel property
	// 	CurrencyCode     []string                   `json:"currencyCode"` // Describes the currency code accepted at the property. Example : [EUR]
	// 	SpokenLanguages  []string                   // Describes the list of languages spoken at the property. Follows the standard of ISO 639-1 (Alpha-2). Example : [es]
	// 	TimeZone         TimeZoneResponse           `json:"timeZone"`         // Element defining a time zone
	// 	Climate          string                     `json:"climate"`          // Describes the climate at the location of the property. example: Dry
	// 	Certifications   []AwardsResponse           `json:"certifications"`   // Describes the certifications received by the Hotel
	// 	RelativeLocation []LocationDistanceResponse `json:"relativeLocation"` // To indicate the reference points from the hotel such as the distance to Airport, Bus Stations or Train Station.
	// 	Season           SeasonResponse             `json:"season"`           // Models a period of time between two dates and inclusive only of the days of the week specified.
	// 	Building         BuildingResponse           `json:"building"`         // Indicates the building of the hotel
	// }

	// BuildingResponse struct {
	// 	ArchitectureCode        ArchitectureCode `json:"architectureCode"`        // Denotes the architecture in which the property was built upon. Can contain values such as
	// 	BuildDate               string           `json:"buildDate"`               // Denotes the year at which the property was built. Format YYYY-MM-DD (ISO 8601)
	// 	RenovationDate          string           `json:"renovationDate"`          // Denotes the year at which the property was renovated. Format YYYY-MM-DD (ISO 8601)
	// 	NumberOfFloors          int              `json:"numberOfFloors"`          // Indicates the number of floors in the property
	// 	NumberOfRooms           int              `json:"numberOfRooms"`           // Indicates the number of rooms in the property
	// 	NumberOfExecutiveFloors int              `json:"numberOfExecutiveFloors"` // Indicates the number of Executive floors in the property
	// 	NumberOfBuildings       int              `json:"numberOfBuildings"`       // Indicates the number of buildings in the property
	// 	NumberOfElevators       int              `json:"numberOfElevators"`       // Indicates the number of elevators in the property
	// }

	// SeasonResponse struct {
	// 	ClosedSeasons   []Period `json:"closedSeasons"`   // Closed seasons of the hotel refers to the season where in the property is shut down
	// 	BlackoutSeasons []Period `json:"blackoutSeasons"` // Blackout dates of the hotel during which the hotel is open but no bookings are available
	// 	OpenCalendar    []Period `json:"openCalendar"`    // Indicates the opening time and days of the property
	// }

	// TimeZoneResponse struct {
	// 	ID                     string `json:"id"`                     // Unique id of the time zone. example: Europe/Paris
	// 	Name                   string `json:"name"`                   //Long name of the time zone. example: Central European Summer Time
	// 	Code                   string `json:"code"`                   // Time zone code. example: CEST
	// 	OffSet                 string `json:"offSet"`                 // Total offset from UTC including the Daylight Saving Time (DST) following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: +02:00
	// 	OffSetInSeconds        int    `json:"offSetInSeconds"`        // Total offset from UTC including the Daylight Saving Time (DST) in second. example: 7200
	// 	DstOffset              string `json:"dstOffset"`              // Indicates whether the day light savings is observed at the location. example: True
	// 	DstOffsetInSeconds     int    `json:"dstOffsetInSeconds"`     // Daylight Saving Time (DST) in second. 0 if the zone is not in the Daylight Saving time at specified date. example: -3600
	// 	ReferenceLocalDateTime string `json:"referenceLocalDateTime"` // Date and time used as reference to determine the time zone name, code, offset, and dstOffset following ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601) standard. example: 2022-09-28T19:20:30
	// }

	// LocationDistanceResponse struct {
	// 	Destination LocationResponse   `json:"destination"` // Indicates the destination of the location distance
	// 	Distances   []DistanceResponse `json:"distances"`   // Indicates the distances from the point of interest to the destination
	// }

	// DistanceResponse struct {
	// 	Unit         Unit         `json:"unit"`         // Indicates the unit of the distance
	// 	Value        int          `json:"value"`        // Indicates the value of the distance
	// 	DistanceType DistanceType `json:"distanceType"` // Indicates the type of the distance
	// }

	// TransportationResponse struct {
	// 	TransportMode         TransportMode             `json:"transportMode"`
	// 	IsReservationRequired bool                      `json:"isReservationRequired"` // True if reservation is required in advance to board the transport
	// 	OperatingHours        CalendarScheduleResponse  `json:"operatingHours"`        // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
	// 	Description           string                    `json:"description"`           // Description of the transportation
	// 	PriceEquation         []PricingEquationResponse `json:"priceEquation"`         // Indicates the price equation of the transportation
	// 	Media                 []MediaResponse           `json:"media"`                 // Indicates the media of the transportation
	// }

	// PricingEquationResponse struct {
	// 	PricingMethod PricingMethod           `json:"pricingMethod"` // Indicates the pricing method of the point of interest
	// 	UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the price of the point of interest per unit
	// }

	// ElementaryPriceResponse struct {
	// 	Amount              string           `json:"amount"`              // Indicates the amount of the price of the point of interest
	// 	Value               string           `json:"value"`               // Indicates the value of the price of the point of interest
	// 	DecimalPlaces       int              `json:"decimalPlaces"`       // Indicates the decimal places of the price of the point of interest
	// 	Currency            CurrencyResponse `json:"currency"`            // Indicates the currency of the price of the point of interest
	// 	ElementaryPriceType string           `json:"elementaryPriceType"` // Defines the type of price, eg. for base fare, total, grand total.
	// }

	// LocationResponse struct {
	// 	SubType  string                            `json:"subType"`  // Location sub-type (e.g. airport, port, rail-station, restaurant, atm...)
	// 	Name     string                            `json:"name"`     // Name of the location
	// 	IataCode string                            `json:"iataCode"` // IATA code of the location
	// 	GeoCode  dto.GeoCodeResponse `json:"geoCode"`  // GeoCode of the location
	// }
	// Period struct {
	// 	Start *time.Time `json:"start"` // start date and time following ISO 8601 format
	// 	End   *time.Time `json:"end"`   // end date and time following ISO 8601 format
	// }
)

type PaymentType string

const (
	DEPOSIT   PaymentType = "DEPOSIT"
	GUARANTEE PaymentType = "GUARANTEE"
	PREPAY    PaymentType = "PREPAY"
	HoldTime  PaymentType = "HOLDTIME"
)

type PaymentMethod string

const (
	Cash                     PaymentMethod = "CASH"
	DirectBill               PaymentMethod = "DIRECT_BILL"
	Voucher                  PaymentMethod = "VOUCHER"
	CreditCard               PaymentMethod = "CREDIT_CARD"
	DebitCard                PaymentMethod = "DEBIT_CARD"
	Check                    PaymentMethod = "CHECK"
	Deposit                  PaymentMethod = "DEPOSIT"
	Coupon                   PaymentMethod = "COUPON"
	BusinessCheck            PaymentMethod = "BUSINESS_CHECK"
	PersonalCheck            PaymentMethod = "PERSONAL_CHECK"
	MoneyOrder               PaymentMethod = "MONEY_ORDER"
	CertificatesAwards       PaymentMethod = "CERTIFICATES_AWARDS"
	MiscellaneousChargeOrder PaymentMethod = "MISCELLANEOUS_CHARGE_ORDER"
	TravelAgencyNameAddress  PaymentMethod = "TRAVEL_AGENCY_NAME_ADDRESS"
	TravelAgencyIataNumber   PaymentMethod = "TRAVEL_AGENCY_IATA_NUMBER"
	CertifiedCheck           PaymentMethod = "CERTIFIED_CHECK"
	ClubMembershipId         PaymentMethod = "CLUB_MEMBERSHIP_ID"
	FrequentGuestNumber      PaymentMethod = "FREQUENT_GUEST_NUMBER"
	FrequentTravelerNumber   PaymentMethod = "FREQUENT TRAVELER NUMBER"
	GuestNameAddress         PaymentMethod = "GUEST_NAME_ADDRESS"
	SpecialIndustryProgram   PaymentMethod = "SPECIAL_INDUSTRY_PROGRAM"
	TourOrder                PaymentMethod = "TOUR_ORDER"
	TravelersCheck           PaymentMethod = "TRAVELERS_CHECK"
	WirePayment              PaymentMethod = "WIRE_PAYMENT"
	CompanyNameAddress       PaymentMethod = "COMPANY_NAME_ADDRESS"
	CorporateIdCdNumber      PaymentMethod = "CORPORTE_ID_CD_NUMBER"
	Guarantee                PaymentMethod = "GUARANTEE"
	VirtualCreditCard        PaymentMethod = "VIRTUAL_CREDIT_CARD"
)

type PricingMethod string

const (
	Daily                 PricingMethod = "DAILY"
	Hourly                PricingMethod = "HOURLY"
	HalfDay               PricingMethod = "HALF_DAY"
	AdditionsPerStay      PricingMethod = "ADDITIONS_PER_STAY"
	PerOccurrence         PricingMethod = "PER_OCCURRENCE"
	PerEvent              PricingMethod = "PER_EVENT"
	PerPerson             PricingMethod = "PER_PERSON"
	FirstUse              PricingMethod = "FIRST_USE"
	OneTimeUse            PricingMethod = "ONE_TIME_USE"
	PerMinute             PricingMethod = "PER_MINUTE"
	PerFunction           PricingMethod = "PER_FUNCTION"
	PerStay               PricingMethod = "PER_STAY"
	Complimentary         PricingMethod = "COMPLIMENTARY"
	Other                 PricingMethod = "OTHER"
	MaximumCharge         PricingMethod = "MAXIMUM_CHARGE"
	OverMinuteCharge      PricingMethod = "OVER_MINUTE_CHARGE"
	Weekly                PricingMethod = "WEEKLY"
	PerRoomPerStay        PricingMethod = "PER_ROOM_PER_STAY"
	PerRoomPerNight       PricingMethod = "PER_ROOM_PER_NIGHT"
	PerPersonPerStay      PricingMethod = "PER_PERSON_PER_STAY"
	PerPersonPerNight     PricingMethod = "PER_PERSON_PER_NIGHT"
	MinimumCharge         PricingMethod = "MINIMUM_CHARGE"
	PerRental             PricingMethod = "PER_RENTAL"
	PerItem               PricingMethod = "PER_ITEM"
	PerRoom               PricingMethod = "PER_ROOM"
	PerReservationBooking PricingMethod = "PER_RESERVATION_BOOKING"
	PerGallon             PricingMethod = "PER_GALLON"
	PerDozen              PricingMethod = "PER_DOZEN"
	PerTray               PricingMethod = "PER_TRAY"
	PerOrder              PricingMethod = "PER_ORDER"
	PerUnit               PricingMethod = "PER_UNIT"
	OneWay                PricingMethod = "ONE_WAY"
	RoundTrip             PricingMethod = "ROUND_TRIP"
)

type (
	PolicyResponse struct {
		PaymentPolicies      []PaymentPolicyResponse         `json:"paymentPolicies"`
		CheckInOutPolicies   []CheckInOutPolicyResponse      `json:"checkInOutPolicies"`
		PetsPolicies         []PetsPolicyResponse            `json:"petsPolicies"`
		CancellationPolicies []CancellationPolicyResponse    `json:"cancellationPolicies"` // Describes the cancellation policies applicable to the property
		TaxPolicies          []TaxPolicyResponse             `json:"taxPolicies"`          // Describes the taxes applicable at the property
		CommissionPolicies   []CommissionPolicyResponse      `json:"commissionPolicies"`   // Describes the commission policies applicable to the property
		StayRequirements     []dto.QualifiedFreeTextResponse `json:"stayRequirements"`     // Describes the stay requirements such as minimum/maximum length of stay
		GuestPolicies        []GuestPolicyResponse           `json:"guestPolicies"`        // Describes the guest policies applicable to the property
		LoyaltyPolicies      []LoyaltyBenefitResponse        `json:"loyaltyPolicies"`      // Describes the loyalty benefits applicable to the property
	}

	// * Describes the conditions under which a booking can be cancelled and the resulting charges.
	CancellationPolicyResponse struct {
		Amount         string                        `json:"amount"`         // Cancellation charge amount applicable when the policy is triggered
		NumberOfNights int                           `json:"numberOfNights"` // Number of nights charged as a cancellation penalty
		Percentage     string                        `json:"percentage"`     // Cancellation charge expressed as a percentage
		Deadline       string                        `json:"deadline"`       // Deadline before which the booking can be cancelled free of charge
		Description    dto.QualifiedFreeTextResponse `json:"description"`    // Free-text description of the cancellation policy
		PolicyType     string                        `json:"policyType"`     // Type of the cancellation policy
	}

	// * Describes a tax or fee applicable at the property.
	TaxPolicyResponse struct {
		Currency         string `json:"currency"`         // ISO currency code of the tax amount
		Amount           string `json:"amount"`           // Tax amount
		Code             string `json:"code"`             // Code identifying the tax
		Percentage       string `json:"percentage"`       // Tax expressed as a percentage
		Included         bool   `json:"included"`         // True if the tax is already included in the price
		Description      string `json:"description"`      // Description of the tax
		PricingFrequency string `json:"pricingFrequency"` // Frequency at which the tax is applied
		PricingMode      string `json:"pricingMode"`      // Mode in which the tax is priced
	}

	// * Describes the commission applicable to a booking.
	CommissionPolicyResponse struct {
		Percentage  string                        `json:"percentage"`  // Commission expressed as a percentage
		Amount      string                        `json:"amount"`      // Commission amount
		Description dto.QualifiedFreeTextResponse `json:"description"` // Free-text description of the commission policy
	}

	// * Describes the policies applicable to guests such as age restrictions and child sharing rules.
	GuestPolicyResponse struct {
		MinGuestAge              int  `json:"minGuestAge"`              // Minimum age required for a guest to stay at the property
		MaxChildAgeforBedSharing int  `json:"maxChildAgeforBedSharing"` // Maximum age of a child allowed to share a bed
		ChildStayFreeCutoffAge   int  `json:"childStayFreeCutoffAge"`   // Maximum age up to which a child stays free of charge
		ChildStayFree            bool `json:"childStayFree"`            // True if children stay free of charge
	}

	// * Describes the loyalty benefits a member can accrue or redeem at the property.
	LoyaltyBenefitResponse struct {
		Eligibility      string                    `json:"eligibility"`      // Describes the eligibility for the loyalty benefit
		BenefitsAccruals []BenefitAccrualResponse  `json:"benefitsAccruals"` // Describes the benefits accrued through the loyalty program
		Discount         LoyaltyDiscountResponse   `json:"discount"`         // Describes the discount associated with the loyalty benefit
		Membership       LoyaltyMembershipResponse `json:"membership"`       // Describes the membership tied to the loyalty benefit
	}

	// * Describes a benefit accrued through a loyalty program.
	BenefitAccrualResponse struct {
		LoyaltyAwardType string `json:"loyaltyAwardType"` // Type of the loyalty award
		Amount           string `json:"amount"`           // Amount accrued for the benefit
		Category         string `json:"category"`         // Category of the benefit
		Code             string `json:"code"`             // Code identifying the benefit
		CodeDescription  string `json:"codeDescription"`  // Description of the benefit code
	}

	// * Describes a discount applicable through a loyalty program.
	LoyaltyDiscountResponse struct {
		Percentage string `json:"percentage"` // Discount expressed as a percentage
	}

	// * Describes the membership tied to a loyalty benefit.
	LoyaltyMembershipResponse struct {
		ActiveTier LoyaltyTierResponse    `json:"activeTier"` // Describes the active tier of the membership
		Program    LoyaltyProgramResponse `json:"program"`    // Describes the loyalty program of the membership
	}

	// * Describes the active tier of a loyalty membership.
	LoyaltyTierResponse struct {
		Level string `json:"level"` // Level of the active tier
	}

	// * Describes a loyalty program.
	LoyaltyProgramResponse struct {
		Name  string              `json:"name"`  // Name of the loyalty program
		Owner string              `json:"owner"` // Owner of the loyalty program
		Media []dto.MediaResponse `json:"media"` // Media associated to the loyalty program
	}

	// * Pets policies
	PetsPolicyResponse struct {
		Code          string            `json:"code"`          // example: 119 describes the pets policy code of the property
		Description   string            `json:"description"`   // example: Only dogs are allowed
		PricingMethod dto.PricingMethod `json:"pricingMethod"` // example: PER_NIGHT. Pricing method for the pets policy
	}

	// * Check-in and Check-out policies
	CheckInOutPolicyResponse struct {
		CheckIn             string                        `json:"checkIn"`             // example: 13:00:00. Check-in From time limit in ISO-8601 format [http://www.w3.org/TR/xmlschema-2/#time]
		CheckInDescription  dto.QualifiedFreeTextResponse `json:"checkInDescription"`  // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
		CheckOut            string                        `json:"checkOut"`            // example: 12:00:00. Check-out To time limit in ISO-8601 format [http://www.w3.org/TR/xmlschema-2/#time]
		CheckOutDescription dto.QualifiedFreeTextResponse `json:"checkOutDescription"` // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
	}

	// * Booking Rules
	PaymentPolicyResponse struct {
		PaymentType       PaymentType                     `json:"paymentType"` // example: DEPOSIT payment type. Guarantee means Pay at Check Out. Check the methods in guarantee or deposit or prepay.
		Guarantee         GuaranteeResponse               `json:"guarantee"`
		AdditionalDetails []dto.QualifiedFreeTextResponse `json:"additionalDetails"`
	}

	// * the guarantee policy information applicable to the offer. It includes accepted payments
	GuaranteeResponse struct {
		Description      dto.QualifiedFreeTextResponse `json:"description"`      // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
		AcceptedPayments AcceptedPaymentResponse       `json:"acceptedPayments"` // Accepted Payment Methods and Card Types. Several Payment Methods and Card Types may be available.
	}

	AcceptedPaymentResponse struct {
		CreditCards []string        `json:"creditCards"` // example: VI .CA - MasterCard (warning - use it instead of MC/IK/EC/MD/XS) VI - Visa AX - American Express DC - Diners Club AU - Carte Aurore CG - Cofinoga DS - Discover GK - Lufthansa GK Card JC - Japanese Credit Bureau TC - Torch Club TP - Universal Air Travel Card BC - Bank Card DL - Delta MA - Maestro UP - China UnionPay
		Methods     []PaymentMethod `json:"methods"`     // example: CREDIT_CARD. CREDIT_CARD (CC) - Payment Cards in creditCards are accepted AGENCY_ACCOUNT - Agency Account (Credit Line) is accepted. Agency is Charged at CheckOut TRAVEL_AGENT_ID - Agency IATA/ARC Number is accepted to Guarantee the booking CORPORATE_ID (COR-ID) - Corporate Account is accepted to Guarantee the booking HOTEL_GUEST_ID - Hotel Chain Rewards Card Number is accepted to Guarantee the booking CHECK - Checks are accepted MISC_CHARGE_ORDER - Miscellaneous Charge Order is accepted ADVANCE_DEPOSIT - Cash is accepted for Deposit/PrePay COMPANY_ADDRESS - Company Billing Address is accepted to Guarantee the booking
	}
)

type PromotionCategory string

const (
	Aaa                    PromotionCategory = "AAA"
	Aarp                   PromotionCategory = "AARP"
	Convention             PromotionCategory = "CONVENTION"
	Corporate              PromotionCategory = "CORPORATE"
	Family                 PromotionCategory = "FAMILY"
	Government             PromotionCategory = "GOVERNMENT"
	Group                  PromotionCategory = "GROUP"
	Military               PromotionCategory = "MILITARY"
	Promotional            PromotionCategory = "PROMOTIONAL"
	SeniorCitizen          PromotionCategory = "SENIOR_CITIZEN"
	Tour                   PromotionCategory = "TOUR"
	Weekday                PromotionCategory = "WEEKDAY"
	Weekend                PromotionCategory = "WEEKEND"
	FullPrice              PromotionCategory = "FULL_PRICE"
	FreeOfCharge           PromotionCategory = "FREE_OF_CHARGE"
	HomeVisitingProof      PromotionCategory = "HOME_VISITING_PROOF"
	RedTour                PromotionCategory = "RED_TOUR"
	SingleBedForChild      PromotionCategory = "SINGLE_BED_FOR_CHILD"
	ChildSharingSeat       PromotionCategory = "CHILD_SHARING_SEAT"
	Anniversary            PromotionCategory = "ANNIVERSARY"
	AdventurePackage       PromotionCategory = "ADVENTURE_PACKAGE"
	BedAndBreakfastPackage PromotionCategory = "BED_AND_BREAKFAST_PACKAGE"
	Dinner                 PromotionCategory = "DINNER"
	FishingPackage         PromotionCategory = "FISHING_PACKAGE"
	GolfPackage            PromotionCategory = "GOLF_PACKAGE"
	Getaway                PromotionCategory = "GETAWAY"
	HolidayPackage         PromotionCategory = "HOLIDAY_PACKAGE"
	HoneymoonPackage       PromotionCategory = "HONEYMOON_PACKAGE"
	InternetPackage        PromotionCategory = "INTERNET_PACKAGE"
	ParkAndFlyPackage      PromotionCategory = "PARK_AND_FLY_PACKAGE"
	Park                   PromotionCategory = "PARK"
	Romance                PromotionCategory = "ROMANCE"
	RecreationPackage      PromotionCategory = "RECREATION_PACKAGE"
	ShoppingPackage        PromotionCategory = "SHOPPING_PACKAGE"
	SkiPackage             PromotionCategory = "SKI_PACKAGE"
	SpaPackage             PromotionCategory = "SPA_PACKAGE"
	TravelAgentRates       PromotionCategory = "TRAVEL_AGENT_RATES"
	TheaterPackage         PromotionCategory = "THEATER_PACKAGE"
	TennisPackage          PromotionCategory = "TENNIS_PACKAGE"
	Travel                 PromotionCategory = "TRAVEL"
)

type (
	// * Promotion refers to the list of things that is done to increase the sales or to advertise a product
	PromotionResponse struct {
		Name               string                   `json:"name"`               // example: Year End Sale . Name of the promotion
		Description        string                   `json:"description"`        // example: Get 20% off by using this promotion code. Enjoy extra benefits while applying this promotion Description of the promotion
		Category           PromotionCategory        `json:"category"`           // Category of the promotion. These enum values are inspired from OTA - "https://opentravel.org/" with code list as - DIS
		Code               string                   `json:"code"`               // example: XYZ123 Promotion code to be used at the time of booking
		TermsAndConditions TermsOfConditionResponse `json:"termsAndConditions"` // example: Terms and conditions of the promotion
		Media              []dto.MediaResponse      `json:"media"`              // example: Media of the promotion
	}

	TermsOfConditionResponse struct {
		Language        string `json:"language"`        // example: fr-FR. see RFC 5646
		DescriptionType string `json:"descriptionType"` // example: TEXT. Type of the description
		Text            string `json:"text"`            // example: Terms and conditions of the promotion
	}

	// * Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
	QualifiedFreeTextResponse = dto.QualifiedFreeTextResponse
)

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
		Name                     dto.QualifiedFreeTextResponse    `json:"name"`
		HotelRoomClassification  HotelRoomClassification          `json:"hotelRoomClassification"`  // Indicates the type of the hotel room. Enum values are taken from the OTA Code list GRI
		ProviderContentReference ProviderContentReferenceResponse `json:"providerContentReference"` // Identifies the provider content reference for the room.
		HotelRoomCategory        HotelRoomCategory                `json:"hotelRoomCategory"`        //
		Beds                     int                              `json:"beds"`                     // example: 2. Number of beds in the room
		BedType                  BedType                          `json:"bedType"`                  // Type of the bed.Enum values are taken from the OTA code list BED Here are the list of values
		Description              dto.QualifiedFreeTextResponse    `json:"description"`              //
		Quantity                 int                              `json:"quantity"`                 // Indicates the number of rooms under the given type
		BedRoomsPerRoom          int                              `json:"bedRoomsPerRoom"`          // Indicates the number of bed rooms under the room
		BathroomsPerRoom         int                              `json:"bathroomsPerRoom"`         // Indicates the number of bathrooms under the room
		PolicyDescriptions       []string                         `json:"policyDescriptions"`       // Lists out the set of policies for the room
		ViewCode                 RoomViewType                     `json:"viewCode"`                 // Indicates the view of the room. Enum values are taken from the OTA Code list VIE
		ArchitectureCode         ArchitectureCode                 `json:"architectureCode"`         // Indicates the architecture of the room. Enum values are taken from the OTA Code list ARC
		HotelRoomLocation        string                           `json:"hotelRoomLocation"`        // Indicates the location of the room within the hotel property
		SortOrder                int                              `json:"sortOrder"`                // Absolute order in which to display the room
		Media                    []dto.MediaResponse              `json:"media"`                    // List of media associated to the room
		Amenities                []dto.AmenityResponse            `json:"amenities"`                // Amenities available in the room
		IsNonSmoking             bool                             `json:"isNonSmoking"`             // Indicates whether the room is a smoking room or not
		StandardPersonCapacity   int                              `json:"standardPersonCapacity"`   // Indicates the typical capacity for the room
		MaxPersonCapacity        dto.MaxPersonCapacityResponse    `json:"maxPersonCapacity"`        // Capacity of hotel room ( total of number of adult, children)
		MaxSleepFurnishings      MaxSleepFurnishingResponse       `json:"maxSleepFurnishings"`      // Indicates the maximum number of extra beds, cribs that can be accomodated in the room
		Dimensions               dto.DimensionsResponse           `json:"dimensions"`               // A dimension is a measurement such as length, width, or height. Dimensions of a place refers to its size and proportions. The value of height, length and width can be collectively used to calculate multiple values such as total surface area and weight, volume and density.
	}

	RoomDimensionsResponse struct {
		Length        int     `json:"length"`        // Indicates the length of the room
		Width         int     `json:"width"`         // Indicates the width of the room
		Height        int     `json:"height"`        // Indicates the height of the room
		DecimalPlaces int     `json:"decimalPlaces"` // Indicates the number of decimal places for the length, width and height
		Unit          Unit    `json:"unit"`          // Indicates the unit of the length, width and height
		Area          float64 `json:"area"`          // Indicates the area of the room
		AreaUnit      Unit    `json:"areaUnit"`      // Indicates the unit of the area
	}

	MaxSleepFurnishingResponse = dto.MaxSleepFurnishingsResponse

	MaxPersonCapacityResponse = dto.MaxPersonCapacityResponse

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
		Media            []dto.MediaResponse     `json:"media"`            // List of media associated to the amenity
		Quantity         int                     `json:"quantity"`         // Indicates the quantity of the amenity, i.e; how many counts are available for a particular amenity type
		PricingMethod    dto.PricingMethod       `json:"pricingMethod"`    // indicates the pricing method used to asses the amenity's usage cost
	}

	AmenityProviderResponse = dto.AmenityProvider

	PriceResponse struct {
		Currency          dto.CurrencyResponse  `json:"currency"`          // Indicates the currency of the price.
		SellingTotal      string                `json:"sellingTotal"`      // sellingTotal = Total + margins + markup + totalFees - discounts
		Total             string                `json:"total"`             // total = base + totalTaxes
		PricingTimeWindow dto.PricingTimeWindow `json:"pricingTimeWindow"` // Unit timing Price for which the rate plan pricing applies. It can contain values such as
	}

	CurrencyResponse = dto.CurrencyResponse
)
