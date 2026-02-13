package responseContentDTO

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
)

type DeviceType string

const (
	DeviceTypeFax      DeviceType = "FAX"
	DeviceTypeMobile   DeviceType = "MOBILE"
	DeviceTypeLandline DeviceType = "LANDLINE"
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
	LocationTypeLocalReservationOffice   LocationType = "LOCAL RESERVATION OFFICE"
	LocationTypeSalesOffice              LocationType = "SALES_OFFICE"
	LocationTypeFranchiseCompany         LocationType = "FRANCHISE COMPANY"
	LocationTypeManagementCompany        LocationType = "MANAGEMENT COMPANY"
	LocationTypeOwnershipCompany         LocationType = "OWNERSHIP COMPANY"
	LocationTypeCustomerServiceOffice    LocationType = "CUSTOMER_SERVICE_OFFICE"
	LocationTypeHomeResidence            LocationType = "HOME_RESIDENCE"
	LocationTypeRegionalSalesOffice      LocationType = "REGIONAL SALES OFFICE"
	LocationTypeTechnicalSupportOffice   LocationType = "TECHNICAL SUPPORT OFFICE"
)

type (
	// * Contains all the facilities offered by the hotel.
	FacilityResponse struct {
		MeetingRoomInfo MeetingRoomInfoResponse `json:"meetingRoomInfo"` // Indicates the meeting room information
		Amenities       []AmenityResponse       `json:"amenities"`       // Indicates the amenities offered by the facility
		RestaurantInfo  RestaurantInfoResponse  `json:"restaurantInfo"`  // Indicates the restaurant information
	}

	RestaurantInfoResponse struct {
		Quantity   int                  `json:"quantity"`   // Indicates the number of restaurants within the property
		Restaurant []RestaurantResponse `json:"restaurant"` // Indicates the various restaurants in the property
	}

	RestaurantResponse struct {
		Name                  string                    `json:"name"`                  // Indicates the name of the restaurant
		Description           string                    `json:"description"`           // Indicates the description of the restaurant
		Category              RestaurantServiceCategory `json:"category"`              // Restaurant food service category. These enum values are inspired from OTA - "https://opentravel.org/" with code list as - RES. Can contain values such as
		AcceptedCurrencyCodes []string                  `json:"acceptedCurrencyCodes"` // example: EUR
		CuisineTypes          []string                  `json:"cuisineTypes"`          // Indicates the list of cuisines served at the restaurant. These enum values are inspired from OTA - "https://opentravel.org/" with code list as - CUI
		MaxSeatingCapacity    int                       `json:"maxSeatingCapacity"`    // Indicates the max number of occupancy in the restaurant
		HasBreakfast          bool                      `json:"hasBreakfast"`          // True if breakfast is served in the restaurant. Default value is false
		HasLunch              bool                      `json:"hasLunch"`              // True if lunch is served in the restaurant. Default value is false
		HasBrunch             bool                      `json:"hasBrunch"`             // True if brunch is served in the restaurant. Default value is false
		HasDinner             bool                      `json:"hasDinner"`             // True if dinner is served in the restaurant. Default value is false
		Contact               []ContactResponse         `json:"contact"`               // A contact refers to the information that can be used to reach a person, a company or an organization.
		HonorsAndAwards       []AwardsResponse          `json:"honorsAndAwards"`       // Indicates the honors and awards received by the restaurant
		Media                 []MediaResponse           `json:"media"`                 // Indicates the media of the restaurant
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
		Email         EmailResponse         `json:"email"`         // Indicates the email addresses of the person, company or organization that the contact is for
		Purpose       []string              `json:"purpose"`       // the purpose for which this contact is to be used
		LocationType  LocationType          `json:"locationType"`  // Describes the locationType of the contact. It can contain values such as
		Website       struct {
			Url string `json:"url"` // Indicates the URL of the website
		} `json:"website"` // Object containing URL and description
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
		MeetingRooms              MeetingRoomResponse    `json:"meetingRooms"`              // Indicates the various meeting rooms in the property
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
		Media               []MediaResponse               `json:"media"`               // Indicates the media of the meeting room
	}

	PriceQuotationsResponse struct {
		PricingMethod PricingMethod `json:"pricingMethod"` // Indicates the pricing method used to asses the meeting room's usage cost
		UnitPrice     PriceResponse `json:"unitPrice"`     // Indicates the price of the meeting room per unit
	}

	OccupancyPerLayoutsResponse struct {
		Layout       Layout `json:"layout"`       // Defines the design layout type of the meeting room
		MaxOccupancy int    `json:"maxOccupancy"` // Denotes the maximum number of people that can be accomodated in the corresponding layout
	}
)
