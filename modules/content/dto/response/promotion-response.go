package responseContentDTO

type QualifiedFreeTextType string

const (
	PropertyDescription               QualifiedFreeTextType = "PROPERTY_DESCRIPTION"
	AmenityInformation                QualifiedFreeTextType = "AMENITY_INFORMATION"
	PromotionalInformation            QualifiedFreeTextType = "PROMOTIONAL_INFORMATION"
	Dining                            QualifiedFreeTextType = "DINING"
	GeneralMeetingPlanningInformation QualifiedFreeTextType = "GENERAL_MEETING_PLANNING_INFORMATION"
	Services                          QualifiedFreeTextType = "SERVICES"
	Marketing                         QualifiedFreeTextType = "MARKETING"
	TypicalDescription                QualifiedFreeTextType = "TYPICAL_DESCRIPTION"
	SellMarketing                     QualifiedFreeTextType = "SELL_MARKETING"
	TopSellingFeature                 QualifiedFreeTextType = "TOP_SELLING_FEATURE"
	AreasServed                       QualifiedFreeTextType = "AREAS_SERVED"
	CategoryDescription               QualifiedFreeTextType = "CATEGORY_DESCRIPTION"
	OtherDescription                  QualifiedFreeTextType = "OTHER_DESCRIPTION"
	HotelShortDescription             QualifiedFreeTextType = "HOTEL_SHORT_DESCRIPTION"
	HotelLongDescription              QualifiedFreeTextType = "HOTEL_LONG_DESCRIPTION"
	LongLocationDescription           QualifiedFreeTextType = "LONG_LOCATION_DESCRIPTION"
	ShortLocationDescription          QualifiedFreeTextType = "SHORT_LOCATION_DESCRIPTION"
	DefaultRoomTypeDescription        QualifiedFreeTextType = "DEFAULT_ROOM_TYPE_DESCRIPTION"
	MeetingFacilitiesDescription      QualifiedFreeTextType = "MEETING_FACILITIES_DESCRIPTION"
	GroupMeetingDescription           QualifiedFreeTextType = "GROUP_MEETING_DESCRIPTION"
	FacilityDescription               QualifiedFreeTextType = "FACILITY_DESCRIPTION"
	OnsiteFacilities                  QualifiedFreeTextType = "ONSITE_FACILITIES"
	OffsiteFacilities                 QualifiedFreeTextType = "OFFSITE_FACILITIES"
	OnsiteServices                    QualifiedFreeTextType = "ONSITE_SERVICES"
	OffsiteServices                   QualifiedFreeTextType = "OFFSITE_SERVICES"
	OnsiteRecreationalActivities      QualifiedFreeTextType = "ONSITE_RECREATIONAL_ACTIVITIES"
	OffsiteRecreationalActivities     QualifiedFreeTextType = "OFFSITE_RECREATIONAL_ACTIVITIES"
	SecurityInformation               QualifiedFreeTextType = "SECURITY_INFORMATION"
	AdditionalOccupant                QualifiedFreeTextType = "ADDITIONAL_OCCUPANT"
	RateDisclaimer                    QualifiedFreeTextType = "RATE_DISCLAIMER"
	TaxAndFeeDescription              QualifiedFreeTextType = "TAX_AND_FEE_DESCRIPTION"
	GeneralPolicyDescription          QualifiedFreeTextType = "GENERAL_POLICY_DECRIPTION"
	CommissionPolicyDescription       QualifiedFreeTextType = "COMMISSION_POLICY_DESCRIPTION"
	CommissionException               QualifiedFreeTextType = "COMMISSION_EXCEPTION"
	VisaTravelRequirements            QualifiedFreeTextType = "VISA_TRAVEL_REQUIREMENTS"
	ExtraChargesDescription           QualifiedFreeTextType = "EXTRA_CHARGES_DESCRIPTION"
	ExtendedStayDescription           QualifiedFreeTextType = "EXTENDED_STAY_DESCRIPTION"
	BookingPolicyDescription          QualifiedFreeTextType = "BOOKING_POLICY_DESCRIPTION"
	ServiceChargeDescription          QualifiedFreeTextType = "SERVICE_CHARGE_DESCRIPTION"
	GroupConditions                   QualifiedFreeTextType = "GROUP_CONDITIONS"
	EarlyCheckoutDescription          QualifiedFreeTextType = "EARLY_CHECKOUT_DESCRIPTION"
	LateCheckoutDescription           QualifiedFreeTextType = "LATE_CHECKOUT_DESCRIPTION"
	LastRoomDescription               QualifiedFreeTextType = "LAST_ROOM_DESCRIPTION"
	RoomTypeGuaranteed                QualifiedFreeTextType = "ROOM_TYPE_GUARANTEED"
	DiningDescription                 QualifiedFreeTextType = "DINING_DECRIPTION"
	HotelRoomDescription              QualifiedFreeTextType = "HOTEL_ROOM_DESCRIPTION"
	RoomAmenityDescription            QualifiedFreeTextType = "ROOM_AMENITY_DESCRIPTION"
	StandardRoomCategory              QualifiedFreeTextType = "STANDARD_ROOM_CATEGORY"
	RoomCategoryName                  QualifiedFreeTextType = "ROOM_CATEGORY_NAME"
	DefaultRoomName                   QualifiedFreeTextType = "DEFAULT_ROOM_NAME"
	RoomCategory                      QualifiedFreeTextType = "ROOM_CATEGORY"
	RestaurantImages                  QualifiedFreeTextType = "RESTAURANT_IMAGES"
	SpecialOffersDescription          QualifiedFreeTextType = "SPECIAL_OFFERS_DESCRIPTION"
	CateringDescription               QualifiedFreeTextType = "CATERING_DESCRIPTION"
	CuisineDescription                QualifiedFreeTextType = "CUISINE_DESCRIPTION"
	RestaurantService                 QualifiedFreeTextType = "RESTAURANT_SERVICE"
	TransportationDescription         QualifiedFreeTextType = "TRANSPORTATION_DESCRIPTION"
	CheckoutInstructions              QualifiedFreeTextType = "CHECKOUT_INSTRUCTIONS"
	DamageDeposit                     QualifiedFreeTextType = "DAMAGE_DEPOSIT"
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
		Media              []MediaResponse          `json:"media"`              // example: Media of the promotion
	}

	TermsOfConditionResponse struct {
		Language        string `json:"language"`        // example: fr-FR. see RFC 5646
		DescriptionType string `json:"descriptionType"` // example: TEXT. Type of the description
		Text            string `json:"text"`            // example: Terms and conditions of the promotion
	}

	// * Media is a digital content like image, video with associated text and description, several scales and some metadata can be provided also.
	MediaResponse struct {
		Id            string                    `json:"id"`            // example: 69810B23CB8644A18AF760DC66BE41A6. Image Id
		Title         string                    `json:"title"`         // example: My image title. media title
		Caption       string                    `json:"caption"`       // example: Hotel exterior view. Caption of the media
		Href          string                    `json:"href"`          // example: http:pdt.multimediarepository.testing.amadeus.com/cmr/retrieve/hotel/69810B23CB8644A18AF760DC66BE41A6. href to display the original media. href for scaled versions of that media are provided at MediaScale level
		Description   QualifiedFreeTextResponse `json:"description"`   // example: Description of the media
		Tags          []string                  `json:"tags"`          // example: ["hotel", "promotion", "sale"]. Tags of the media
		Category      string                    `json:"category"`      // example: EXTERIOR. media category
		MediaScales   []MediaScaleResponse      `json:"mediaScales"`   // example: Media scales of the media
		MediaMetaData MediaMetaDataResponse     `json:"mediaMetaData"` // example: Media meta data of the media
	}

	MediaMetaDataResponse struct {
		SubType    string             `json:"subType"`    // example: PNG, MKV media subtype / file format
		ETag       string             `json:"etag"`       // example: 2010-08-14T13:00:00 The date and time of the last update.
		Dimensions DimensionsResponse `json:"dimensions"` // example: Dimensions of the media
	}

	// * Media Scale is a version in the media with different size and dimension.
	MediaScaleResponse struct {
		Href       string             `json:"href"`       // example: http:pdt.multimediarepository.testing.amadeus.com/cmr/retrieve/hotel/69810B23CB8644A18AF760DC66BE41A6. href to display the original media. href for scaled versions of that media are provided at MediaScale level
		Dimensions DimensionsResponse `json:"dimensions"` // example: Dimensions of the scaled media
	}

	DimensionsResponse struct {
		Height int `json:"height"` // example: 100. Height of the scaled media
		Width  int `json:"width"`  // example: 100. Width of the scaled media
	}

	// * Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
	QualifiedFreeTextResponse struct {
		Text            string                `json:"text"`            // example: Text of the qualified free text
		Type            QualifiedFreeTextType `json:"type"`            // example: PROPERTY_DESCRIPTION. Type of the qualified free text
		Lang            string                `json:"lang"`            // example: fr-FR. Language of the qualified free text
		Status          string                `json:"status"`          // example: ACTIVE. Status of the qualified free text
		CharSet         string                `json:"charSet"`         // example: UTF-8. Character set of the qualified free text
		Encoding        string                `json:"encoding"`        // example: Base-64. Encoding of the qualified free text
		IamaContentType string                `json:"iamaContentType"` // example: text/plain Follow the RFC define by http://www.iana.org/assignments/media-types/media-types.xhtml
	}
)
