package sharedResponseDTO

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

type (
	MediaResponse struct {
		// Id - unique media identifier. Example: "69810B23CB8644A18AF760DC66BE41A6".
		Id string `json:"id,omitempty"`
		// Type - media data type. Enum: file, Image, Icon.
		Type string `json:"type,omitempty"`
		// Name - name of the media file. Example: "guest_room".
		Name string `json:"name,omitempty"`
		// Title - media title. Example: "My image title".
		Title string `json:"title,omitempty"`
		// Caption - media caption text.
		Caption string `json:"caption,omitempty"`
		// Hint - additional hint for the media.
		Hint string `json:"hint,omitempty"`
		// Alt - media description for visually impaired (screen reader text). See W3C WAI guidelines.
		Alt string `json:"alt,omitempty"`
		// Href - URL to display the original media.
		Href string `json:"href,omitempty"`
		// Description - free text description of the media with language info.
		Description *QualifiedFreeTextResponse `json:"description,omitempty"`
		// Category - media category. Example: "EXTERIOR".
		Category string `json:"category,omitempty"`
		// Tags - tags associated with the media.
		Tags []string `json:"tags,omitempty"`
		// MediaType - MIME type of the media. Example: "IMAGE".
		MediaType string `json:"mediaType,omitempty"`
		// MediaScales - scaled versions of the media with different sizes and dimensions.
		MediaScales []MediaScaleResponse `json:"mediaScales,omitempty"`
		// MediaMetaData - metadata about the media (encoding, dimensions, source, etc.).
		MediaMetaData *MediaMetaDataResponse `json:"mediaMetaData,omitempty"`
	}

	// MediaScale represents a scaled version of media with different size and dimension.
	MediaScaleResponse struct {
		// Href - URL to display the scaled version of the media.
		Href string `json:"href,omitempty"`
		// Size - file size of the scaled media.
		Size *MediaSizeResponse `json:"size,omitempty"`
		// Dimensions - physical dimensions (width, height, etc.) of the scaled media.
		Dimensions *DimensionsResponse `json:"dimensions,omitempty"`
		// Duration - duration of the media per ISO 8601. Example: "P1Y2M3DT4H5M6S".
		Duration string `json:"duration,omitempty"`
	}

	// MediaSize represents the size of a media file.
	MediaSizeResponse struct {
		// Unit - unit type for the size value.
		// Enum: NIGHT, PIXELS, KILOGRAMS, POUNDS, CENTIMETERS, INCHES, BYTES, KILOBYTES, etc.
		Unit string `json:"unit,omitempty"`
		// Value - numeric size value. Example: 200.
		Value int `json:"value,omitempty"`
	}

	// Dimensions represents measurements (width, height, length, area) of a media or object.
	DimensionsResponse struct {
		// Area - total surface area. Example: 445.
		Area float64 `json:"area,omitempty"`
		// AreaUnit - unit for area measurement.
		// Enum: SQUARE_FEET, SQUARE_METERS, SQUARE_INCHES, SQUARE_YARDS, etc.
		AreaUnit string `json:"areaUnit,omitempty"`
		// DecimalPlaces - number of decimal places for values.
		DecimalPlaces int `json:"decimalPlaces,omitempty"`
		// Height - height of the object in specified unit.
		Height int `json:"height,omitempty"`
		// Length - length of the object in specified unit.
		Length int `json:"length,omitempty"`
		// Unit - unit type for height/width/length.
		// Enum: PIXELS, CENTIMETERS, INCHES, etc.
		Unit string `json:"unit,omitempty"`
		// Width - width of the object in specified unit.
		Width int `json:"width,omitempty"`
	}

	// MediaMetaData contains metadata about a media file (type, encoding, source, etc.).
	MediaMetaDataResponse struct {
		// MediaType - IANA media type. Enum: application, audio, font, example, image, message, model, multipart, text, video.
		MediaType string `json:"mediaType,omitempty"`
		// SubType - media subtype / file format. Example: "PNG", "MKV".
		SubType string `json:"subType,omitempty"`
		// Encoding - media encoding format. Example: "PNG", "H265".
		Encoding string `json:"encoding,omitempty"`
		// Etag - date and time of the last update in ISO 8601 format.
		Etag string `json:"etag,omitempty"`
		// Size - file size of the media.
		Size *MediaSizeResponse `json:"size,omitempty"`
		// Dimensions - physical dimensions of the media.
		Dimensions *DimensionsResponse `json:"dimensions,omitempty"`
		// Duration - duration per ISO 8601. Example: "P1Y2M3DT4H5M6S".
		Duration string `json:"duration,omitempty"`
		// Application - application name for viewing or editing the media.
		Application string `json:"application,omitempty"`
		// MediaSource - source and copyright information of the media owner.
		MediaSource *MediaSourceResponse `json:"mediaSource,omitempty"`
		// ClickToAction - hyperlink action associated with the media.
		ClickToAction *ClickToActionResponse `json:"clickToAction,omitempty"`
	}

	// MediaSource contains source and copyright information for the media owner.
	MediaSourceResponse struct {
		// Code - owner code of the media.
		Code string `json:"code,omitempty"`
		// Copyright - copyright text related to the media owner.
		Copyright string `json:"copyright,omitempty"`
		// Filename - file name of the media.
		Filename string `json:"filename,omitempty"`
		// Symbology - logo or icon designation.
		Symbology string `json:"symbology,omitempty"`
		// Version - version of the file.
		Version string `json:"version,omitempty"`
	}

	ClickToActionResponse struct {
		// PlainText - hyperlink text content.
		PlainText string `json:"plainText,omitempty"`
		// Href - URL associated with the action text.
		Href string `json:"href,omitempty"`
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
