package responseContentDTO

import (
	"time"

	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"
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
		HotelID                      string                              `json:"hotelId"`                      // Amadeus Property Code (8 chars). example: ADPAR001
		ChainCode                    string                              `json:"chainCode"`                    // Brand (RT...) or Merchant (AD...)
		BrandCode                    string                              `json:"brandCode"`                    // Brand (RT...) (Amadeus 2 chars Code). Small Properties distributed by Merchants may not have a Brand. Example - AD (Value Hotels) is the Provider/Merchant, and RT (Accor) is the Brand of the Property
		DupeID                       string                              `json:"dupeId"`                       // Unique Property identifier of the physical hotel. One physical hotel can be represented by different Providers, each one having its own hotelID. This attribute allows a client application to group together hotels that are actually the same.
		Name                         string                              `json:"name"`                         // Name of the point of interest
		Rating                       string                              `json:"rating"`                       // Rating of the point of interest
		Description                  QualifiedFreeTextResponse           `json:"description"`                  // Description of the point of interest
		Amenities                    []sharedResponseDTO.AmenityResponse `json:"amenities"`                    // Amenities of the point of interest
		Media                        []MediaResponse                     `json:"media"`                        // Media of the point of interest
		DefaultSpokenLanguage        string                              `json:"defaultSpokenLanguage"`        // Describes the default language preferred or used at the property
		ContextProvider              string                              `json:"contextProvider"`              // Describes the provider of the context of the point of interest
		Contact                      []ContactResponse                   `json:"contact"`                      // Contact of the point of interest
		Location                     LocationResponse                    `json:"location"`                     // Location of the point of interest
		Altitude                     AltitudeResponse                    `json:"altitude"`                     // From analytics, Metrics describe the exact numbers that make up the data
		CategoryCode                 CategoryCode                        `json:"categoryCode"`                 // Category code of the point of interest
		Segment                      Segment                             `json:"segment"`                      // Segment of the point of interest
		Area                         []AreaResponse                      `json:"area"`                         // Geographical zone like City, Region, Country
		ChainName                    string                              `json:"chainName"`                    // Name of the chain to which the hotel belongs to
		BrandName                    string                              `json:"brandName"`                    // Name of the brand to which the hotel or hotel chain belongs to
		Status                       HotelStatus                         `json:"status"`                       // Status of the hotel
		HotelBusinessIdentifications BusinessIdentificationResponse      `json:"hotelBusinessIdentifications"` // An business, can be idenfified via business identifiers, those business identifiers are defined by a body of authority thay could be local, national, transnational or supranational (like EU for the EU VAT number).
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
		Value        int          `json:"value"`        // Indicates the value of the distance
		DistanceType DistanceType `json:"distanceType"` // Indicates the type of the distance
	}

	TransportationResponse struct {
		TransportMode         TransportMode             `json:"transportMode"`
		IsReservationRequired bool                      `json:"isReservationRequired"` // True if reservation is required in advance to board the transport
		OperatingHours        CalendarScheduleResponse  `json:"operatingHours"`        // As defined in: https://schema.org/Schedule A schedule defines a repeating time period used to describe a regularly occurring Event. At a minimum a schedule will specify repeatFrequency which describes the interval between occurences of the event. Additional information can be provided to specify the schedule more precisely. This includes identifying the day(s) of the week or month when the recurring event will take place, in addition to its start and end time. Schedules may also have start and end dates to indicate when they are active, e.g. to define a limited calendar of events.
		Description           string                    `json:"description"`           // Description of the transportation
		PriceEquation         []PricingEquationResponse `json:"priceEquation"`         // Indicates the price equation of the transportation
		PriceQuotation        []PriceQuotationResponse  `json:"priceQuotation"`        // Indicates the price quotations for the transportation
		Media                 []MediaResponse           `json:"media"`                 // Indicates the media of the transportation
	}

	PricingEquationResponse struct {
		PricingMethod PricingMethod           `json:"pricingMethod"` // Indicates the pricing method of the point of interest
		UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the price of the point of interest per unit
	}

	// * Indicates a price quote for a given pricing method, using an elementary (itemized) price.
	PriceQuotationResponse struct {
		PricingMethod PricingMethod           `json:"pricingMethod"` // Indicates the pricing method used to assess the quoted price
		UnitPrice     ElementaryPriceResponse `json:"unitPrice"`     // Indicates the quoted price per unit
	}

	ElementaryPriceResponse struct {
		Amount              string           `json:"amount"`              // Indicates the amount of the price of the point of interest
		Value               string           `json:"value"`               // Indicates the value of the price of the point of interest
		DecimalPlaces       int              `json:"decimalPlaces"`       // Indicates the decimal places of the price of the point of interest
		Currency            CurrencyResponse `json:"currency"`            // Indicates the currency of the price of the point of interest
		ElementaryPriceType string           `json:"elementaryPriceType"` // Defines the type of price, eg. for base fare, total, grand total.
	}

	LocationResponse struct {
		SubType  string                            `json:"subType"`  // Location sub-type (e.g. airport, port, rail-station, restaurant, atm...)
		Name     string                            `json:"name"`     // Name of the location
		IataCode string                            `json:"iataCode"` // IATA code of the location
		GeoCode  sharedResponseDTO.GeoCodeResponse `json:"geoCode"`  // GeoCode of the location
	}
	Period struct {
		Start *time.Time `json:"start"` // start date and time following ISO 8601 format
		End   *time.Time `json:"end"`   // end date and time following ISO 8601 format
	}
)
