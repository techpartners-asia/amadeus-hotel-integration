package responseBookingDTO

import (
	"time"

	sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"
)

type (
	// ==================== Aliases to shared response DTOs ====================

	Media               = sharedResponseDTO.MediaResponse
	MediaMetaData       = sharedResponseDTO.MediaMetaDataResponse
	MediaScale          = sharedResponseDTO.MediaScaleResponse
	MediaSize           = sharedResponseDTO.MediaSizeResponse
	MediaSource         = sharedResponseDTO.MediaSourceResponse
	Dimensions          = sharedResponseDTO.DimensionsResponse
	ClickToAction       = sharedResponseDTO.ClickToActionResponse
	Markup              = sharedResponseDTO.MarkupResponse
	AmenityProvider     = sharedResponseDTO.AmenityProvider
	AmenityPrice        = sharedResponseDTO.AmenityPrice
	MaxPersonCapacity   = sharedResponseDTO.MaxPersonCapacityResponse
	MaxSleepFurnishings = sharedResponseDTO.MaxSleepFurnishingsResponse
	RateFamilyEstimated = sharedResponseDTO.RateFamilyEstimatedResponse
	WarningSource       = sharedResponseDTO.WarningSourceResponse

	// ==================== Top-Level Response ====================

	// HotelBookingResponse is the 201 Created response from POST /v2/booking/hotel-orders.
	HotelBookingResponse struct {
		// Data - the hotel order data returned upon successful booking.
		Data HotelOrder `json:"data"`
		// Warnings - optional array of warnings returned alongside a successful response.
		Warnings []Warning `json:"warnings,omitempty"`
	}

	// ErrorResponse is the error response (400, 500) from the API.
	ErrorResponse struct {
		// Errors - array of error objects describing what went wrong.
		Errors []Error `json:"errors"`
	}

	// ==================== Hotel Order ====================

	// HotelOrder represents one or several hotel bookings done for a set of guests.
	// It corresponds to one PNR in Amadeus GDS.
	HotelOrder struct {
		// Type - resource name, always set to "hotel-order".
		Type string `json:"type"`
		// Id - unique identifier of the hotel order. Must be stored in the client system
		// as it is mandatory for further cancel or retrieve operations.
		Id string `json:"id"`
		// HotelBookings - array of hotel bookings within this order. MinItems: 1.
		HotelBookings []HotelBooking `json:"hotelBookings"`
		// Guests - array of guests sharing this hotel order.
		Guests []ResponseGuest `json:"guests"`
		// AssociatedRecords - references and origin of the hotel order record (PNR).
		AssociatedRecords []AssociatedRecord `json:"associatedRecords"`
		// Self - URL for retrieving the Hotel Order.
		Self string `json:"self"`
	}

	// ==================== Hotel Booking ====================

	// HotelBooking represents one or several rooms booked in the same physical hotel.
	// There is always a reference to this booking provided by the hotel provider.
	HotelBooking struct {
		// Type - data type, always set to "hotel-booking".
		Type string `json:"type"`
		// Id - unique ID of the hotel booking. Computed by Amadeus based on technical data.
		Id string `json:"id"`
		// BookingStatus - status of the booking.
		// Enum: CONFIRMED (HK), PENDING (HN, on-request), CANCELLED, ON_HOLD (HO, deferred payment),
		// PAST (confirmed with check-out in past), UNCONFIRMED (UC), DENIED (NO), GHOST (GK), DELETED.
		BookingStatus string `json:"bookingStatus"`
		// HotelProviderInformation - references and origin of the hotel booking records
		// including confirmation/cancellation numbers from the provider. MinItems: 1.
		HotelProviderInformation []HotelProviderInfo `json:"hotelProviderInformation"`
		// RoomAssociations - array of room associations. Each correlates one room to guest(s),
		// a payment and a hotel offer. One roomAssociation per hotelBooking for multi-room.
		RoomAssociations []RoomAssociation `json:"roomAssociations"`
		// HotelOffer - the full details of the hotel offer that was booked.
		HotelOffer HotelOffer `json:"hotelOffer"`
		// Hotel - hotel content information including name, chain code, and terms.
		Hotel Hotel `json:"hotel"`
		// Payment - payment information used for this booking.
		Payment *PaymentOutput `json:"payment,omitempty"`
		// TravelAgentId - Travel Agent ID / Booking source / IATA number.
		// If not provided in request, set to the IATA Number of the booking office profile.
		TravelAgentId string `json:"travelAgentId"`
		// ArrivalInformation - optional information on how the guest is arriving to the hotel.
		// Displayed if provided at booking creation time.
		ArrivalInformation *ArrivalInformation `json:"arrivalInformation,omitempty"`
	}

	// HotelProviderInfo contains references from the hotel provider for the booking.
	HotelProviderInfo struct {
		// HotelProviderCode - 2-letter hotel provider code. Example: "RT" (for Accor). (required)
		HotelProviderCode string `json:"hotelProviderCode"`
		// ConfirmationNumber - provider confirmation number. Never empty; for on-request it can be "......".
		// If calling the provider, this reference may be asked. Pattern: ^[A-Z0-9]*$. MaxLength: 16. (required)
		ConfirmationNumber string `json:"confirmationNumber"`
		// CancellationNumber - provider cancellation number. Filled for cancelled bookings.
		// If not returned by the hotel provider, it will be "NONE". Pattern: ^[A-Z0-9]*$. MaxLength: 16.
		CancellationNumber string `json:"cancellationNumber,omitempty"`
		// OnRequestNumber - on-request identifier. Pattern: ^[A-Z0-9]*$. MaxLength: 16.
		OnRequestNumber string `json:"onRequestNumber,omitempty"`
	}

	// ==================== Room Association ====================

	// RoomAssociation correlates one single room to guest(s), a payment, and a hotel offer.
	RoomAssociation struct {
		// HotelOfferId - hotel offer ID received in availability response, identifying the product booked.
		// Pattern: ^[A-Z0-9]*$. MinLength: 2, MaxLength: 100.
		HotelOfferId string `json:"hotelOfferId,omitempty"`
		// GuestReferences - array of guest references listing all guests occupying the room.
		// First reference is the main guest holding the reservation and form of payment.
		GuestReferences []GuestReference `json:"guestReferences"`
		// SpecialRequest - special request sent to the reception. MinLength: 2, MaxLength: 120.
		SpecialRequest string `json:"specialRequest,omitempty"`
		// TravelAgentManualMarkup - override amount computed by Margin Manager for Hotel Markup.
		TravelAgentManualMarkup *TravelAgentManualMarkup `json:"travelAgentManualMarkup,omitempty"`
	}

	// GuestReference links a guest to a room with an optional hotel loyalty program.
	GuestReference struct {
		// GuestReference - reference to the guest id. At creation time this is the temporary id (tid). (required)
		GuestReference string `json:"guestReference"`
		// HotelLoyaltyId - Hotel Chain Rewards Program Membership ID.
		// Used for Rewards Points, online check-in, fast check-out.
		// Pattern: ^[A-Z0-9-]{1,21}$. MaxLength: 21. Example: "3081031320523260".
		HotelLoyaltyId string `json:"hotelLoyaltyId,omitempty"`
	}

	// TravelAgentManualMarkup overrides the margin computed by Margin Manager.
	TravelAgentManualMarkup struct {
		// Amount - the markup amount. Pattern: ^\-?[0-9]+(\.[0-9]+)?$ (required)
		Amount string `json:"amount"`
		// Currency - 3-letter currency code. Pattern: ^[A-Z0-9*]{3}$. Example: "EUR". (required)
		Currency string `json:"currency"`
	}

	// ==================== Hotel Offer ====================

	// HotelOffer contains the full details of the hotel offer that was booked.
	HotelOffer struct {
		// Id - unique identifier of the offer. Pattern: ^[A-Z0-9]*$. MaxLength: 100.
		Id string `json:"id"`
		// Type - data type, always "hotel-offer".
		Type string `json:"type"`
		// BoardType - the included breakfast/meals.
		// Enum: ROOM_ONLY, BREAKFAST, HALF_BOARD, ALL_INCLUSIVE, BUFFET_BREAKFAST,
		// CARIBBEAN_BREAKFAST, CONTINENTAL_BREAKFAST, ENGLISH_BREAKFAST, FULL_BREAKFAST,
		// LUNCH, DINNER, FAMILY, AS_BROCHURED, SELF_CATERING, BERMUDA, FULL_BOARD,
		// FAMILY_AMERICAN, MODIFIED, BREAKFAST_AND_LUNCH, LUNCH_AND_DINNER.
		BoardType string `json:"boardType,omitempty"`
		// Category - special rate category. Examples: ASSOCIATION, FAMILY_PLAN.
		Category string `json:"category,omitempty"`
		// CheckInDate - check-in date (hotel local date). Format: YYYY-MM-DD (ISO 8601).
		CheckInDate string `json:"checkInDate"`
		// CheckOutDate - check-out date (hotel local date). Format: YYYY-MM-DD (ISO 8601).
		CheckOutDate string `json:"checkOutDate"`
		// Commission - commission paid to the travel seller, including amount and/or percentage.
		Commission *Commission `json:"commission,omitempty"`
		// Guests - number of adult guests and children with their ages.
		Guests *HotelProductGuests `json:"guests,omitempty"`
		// Policies - booking rules including cancellation, deposit, guarantee, hold time, etc.
		Policies *PolicyDetails `json:"policies,omitempty"`
		// Price - price information including base, taxes, total, markups, and daily variations.
		Price *HotelPrice `json:"price,omitempty"`
		// RateCode - special rate provider response code (3 chars).
		// Examples: RAC (Rack), BAR (Best Available), PRO (Promotional), COR (Corporate),
		// GOV (Government), AAA, BNB (Bed & Breakfast), PKG (Package), WKD (Weekend).
		RateCode string `json:"rateCode,omitempty"`
		// RateFamilyEstimated - estimated rate code family grouping various rate plan codes.
		RateFamilyEstimated *RateFamilyEstimated `json:"rateFamilyEstimated,omitempty"`
		// Room - DEPRECATED: please refer to RoomInformation instead.
		Room *RoomDetails `json:"room,omitempty"`
		// RoomInformation - detailed hotel room information including amenities and room type.
		RoomInformation *RoomInformation `json:"roomInformation,omitempty"`
		// RoomQuantity - number of rooms booked under this offer.
		RoomQuantity int `json:"roomQuantity,omitempty"`
		// Services - list of additional services attached to the offer (e.g. extra beds, parking).
		Services []OfferService `json:"services,omitempty"`
	}

	// OfferService represents an additional service offered with the hotel offer.
	OfferService struct {
		// Code - unique code representing the service.
		Code string `json:"code,omitempty"`
		// Description - free text description of the service.
		Description string `json:"description,omitempty"`
		// IsChargeable - true if the service is chargeable. Default: false.
		IsChargeable bool `json:"isChargeable,omitempty"`
		// Price - price information for the service.
		Price *HotelPrice `json:"price,omitempty"`
		// PricingMethod - how the service cost is assessed.
		// Enum: DAILY, HOURLY, HALF_DAY, PER_OCCURRENCE, PER_EVENT, PER_PERSON, FIRST_USE,
		// PER_MINUTE, COMPLIMENTARY, WEEKLY, PER_STAY, PER_FUNCTION, PER_ROOM_PER_STAY,
		// PER_ROOM_PER_NIGHT, PER_PERSON_PER_STAY, PER_PERSON_PER_NIGHT, PER_RESERVATION_BOOKNG, PER_USE.
		PricingMethod string `json:"pricingMethod,omitempty"`
		// Quantity - how many counts are available for this service.
		Quantity int `json:"quantity,omitempty"`
		// ServiceAttribute - attribute related to the service. Example: "Parking" attribute.
		ServiceAttribute string `json:"serviceAttribute,omitempty"`
	}

	// Commission represents commission paid to the travel seller.
	Commission struct {
		// Amount - amount of the commission. Linked to the currency code of the offer. Pattern: ^\d+(\.\d+)?$
		Amount string `json:"amount,omitempty"`
		// Description - free text description of the commission with language info.
		Description *QualifiedFreeText `json:"description,omitempty"`
		// Percentage - percentage of the commission (0-100). Pattern: ^\d+(\.\d+)?$
		Percentage string `json:"percentage,omitempty"`
	}

	// QualifiedFreeText conveys free text content with language information.
	QualifiedFreeText struct {
		// Lang - language tag per RFC 5646. Example: "fr-FR".
		Lang string `json:"lang,omitempty"`
		// Text - the free text content.
		Text string `json:"text,omitempty"`
	}

	// HotelProductGuests contains guest count information for the offer.
	HotelProductGuests struct {
		// Adults - number of adult guests per room (1-9).
		Adults int `json:"adults,omitempty"`
		// ChildAges - list of ages of each child at the time of check-out (0-20).
		// If several children have the same age, the ages are repeated.
		ChildAges []int `json:"childAges,omitempty"`
	}

	// ==================== Policies ====================

	// PolicyDetails contains all booking rules and policies for the offer.
	PolicyDetails struct {
		// AdditionalDetails - additional policy descriptions.
		AdditionalDetails []AdditionalDetail `json:"additionalDetails,omitempty"`
		// Cancellations - list of cancellation policies. The deadline indicates when a policy starts applying.
		Cancellations []CancellationPolicy `json:"cancellations,omitempty"`
		// CheckInOut - check-in and check-out time policies.
		CheckInOut *CheckInOutPolicy `json:"checkInOut,omitempty"`
		// Deposit - deposit/prepay policy including accepted payments, deadline, and amount due.
		Deposit *DepositPolicy `json:"deposit,omitempty"`
		// Guarantee - guarantee policy including accepted payments.
		Guarantee *GuaranteePolicy `json:"guarantee,omitempty"`
		// HoldTime - hold policy with a deadline for deferred payment.
		HoldTime *HoldPolicy `json:"holdTime,omitempty"`
		// LengthOfStay - minimum and maximum number of nights for the hotel stay.
		LengthOfStay *LengthOfStayPolicy `json:"lengthOfStay,omitempty"`
		// MaximumLengthOfStay - maximum number of nights accepted for the hotel stay.
		MaximumLengthOfStay int `json:"maximumLengthOfStay,omitempty"`
		// MinimumLengthOfStay - minimum number of nights needed for the hotel stay.
		// Note: the Swagger spec uses the typo "minimu" in the JSON key.
		MinimumLengthOfStay int `json:"minimuLengthOfStay,omitempty"`
		// PaymentType - payment type. Guarantee means Pay at Check Out.
		// Enum: GUARANTEE, DEPOSIT, PREPAY, HOLDTIME.
		PaymentType string `json:"paymentType,omitempty"`
		// Prepay - prepay policy (same structure as deposit).
		Prepay *DepositPolicy `json:"prepay,omitempty"`
		// Refundable - the refund/cancellation refund policy.
		Refundable *RefundablePolicy `json:"refundable,omitempty"`
	}

	// AdditionalDetail contains additional policy descriptions.
	AdditionalDetail struct {
		// Description - array of free text descriptions with language info.
		Description []QualifiedFreeText `json:"description,omitempty"`
	}

	// CancellationPolicy describes a cancellation penalty.
	CancellationPolicy struct {
		// Amount - amount of the cancellation fee. Pattern: ^\d+(\.\d+)?$
		Amount string `json:"amount,omitempty"`
		// Deadline - deadline after which the penalty applies. ISO 8601 date-time in hotel local time.
		// Example: "2010-08-14T12:00:00+01:00".
		Deadline string `json:"deadline,omitempty"`
		// Description - free text description of the cancellation policy.
		Description *QualifiedFreeText `json:"description,omitempty"`
		// NumberOfNights - number of nights due as fee in case of cancellation. Minimum: 0.
		NumberOfNights int `json:"numberOfNights,omitempty"`
		// Percentage - percentage of total stay amount to pay on cancellation (0-100). Pattern: ^\d+(\.\d+)?$
		Percentage string `json:"percentage,omitempty"`
		// PolicyType - type of cancellation policy.
		// Enum: CANCELLATION (full booking cancelled), EARLY_CHECKOUT (partial cancel),
		// NO_SHOW (guest did not show up).
		PolicyType string `json:"policyType,omitempty"`
		// Type - DEPRECATED. FULL_STAY means penalty equals the total price.
		Type string `json:"type,omitempty"`
	}

	// CheckInOutPolicy describes check-in and check-out time limits.
	CheckInOutPolicy struct {
		// CheckIn - check-in from time limit in ISO-8601 time format. Example: "13:00:00".
		CheckIn string `json:"checkIn,omitempty"`
		// CheckInDescription - free text description of the check-in policy.
		CheckInDescription *QualifiedFreeText `json:"checkInDescription,omitempty"`
		// CheckOut - check-out until time limit in ISO-8601 time format. Example: "11:00:00".
		CheckOut string `json:"checkOut,omitempty"`
		// CheckOutDescription - free text description of the check-out policy.
		CheckOutDescription *QualifiedFreeText `json:"checkOutDescription,omitempty"`
	}

	// DepositPolicy describes deposit/prepay policy (used for both deposit and prepay sections).
	DepositPolicy struct {
		// AcceptedPayments - accepted payment methods and card types for deposit/prepay.
		AcceptedPayments *PaymentPolicy `json:"acceptedPayments,omitempty"`
		// Amount - deposit/prepay amount. Pattern: ^\d+(\.\d+)?$
		Amount string `json:"amount,omitempty"`
		// Deadline - deadline for deposit/prepay in ISO 8601 date-time format (hotel local time).
		Deadline string `json:"deadline,omitempty"`
		// Description - free text description of the deposit/prepay policy.
		Description *QualifiedFreeText `json:"description,omitempty"`
	}

	// GuaranteePolicy describes the guarantee payment policy.
	GuaranteePolicy struct {
		// AcceptedPayments - accepted payment methods and card types for the guarantee.
		AcceptedPayments *PaymentPolicy `json:"acceptedPayments,omitempty"`
		// Description - free text description of the guarantee policy.
		Description *QualifiedFreeText `json:"description,omitempty"`
	}

	// PaymentPolicy describes accepted payment methods and card types.
	PaymentPolicy struct {
		// CreditCardPolicies - credit card policy information with input parameter details.
		CreditCardPolicies []CreditCardPolicy `json:"creditCardPolicies,omitempty"`
		// CreditCards - DEPRECATED. Use creditCardPolicies. Accepted card type codes (e.g. VI, MA, AX).
		CreditCards []string `json:"creditCards,omitempty"`
		// Methods - accepted payment methods.
		// Enum: CREDIT_CARD, AGENCY_ACCOUNT, TRAVEL_AGENT_ID, CORPORATE_ID, HOTEL_GUEST_ID,
		// CHECK, MISC_CHARGE_ORDER, ADVANCE_DEPOSIT, COMPANY_ADDRESS, VCC_BILLBACK,
		// VCC_B2B_WALLET, DEFERED_PAYMENT, VCC_EXTERNAL_PROVIDER,
		// TRAVEL_AGENT_IMMEDIATE (deprecated), CREDIT_CARD_AGENCY, CREDIT_CARD_TRAVELER.
		Methods []string `json:"methods,omitempty"`
	}

	// CreditCardPolicy contains credit card policy information including input parameter requirements.
	CreditCardPolicy struct {
		// InputParameters - array of input parameter requirements for the credit card (e.g. CVV).
		InputParameters []CreditCardInputParam `json:"inputParameters,omitempty"`
		// VendorCode - card type code. Pattern: ^[A-Z]{2}$.
		// Examples: CA (MasterCard), VI (Visa), AX (AmEx), DC (Diners Club), MA (Maestro), UP (UnionPay).
		VendorCode string `json:"vendorCode,omitempty"`
	}

	// CreditCardInputParam describes a required/optional input parameter for credit card processing.
	CreditCardInputParam struct {
		// ContextDescription - description of the input parameter.
		// Example: "CVV2, card verification value also known as security code, present on the card".
		ContextDescription string `json:"contextDescription,omitempty"`
		// InputRegularExpression - regular expression pattern for the input. Example: "^[0-9]{3,4}$".
		InputRegularExpression string `json:"inputRegularExpression,omitempty"`
		// IsOptional - indicates if the parameter is optional. Example: false.
		IsOptional bool `json:"isOptional,omitempty"`
		// Label - label of the input parameter. Example: "CVV".
		Label string `json:"label,omitempty"`
		// MustBeConcealed - indicates if the parameter needs to be concealed/masked. Example: true.
		MustBeConcealed bool `json:"mustBeConcealed,omitempty"`
		// ParameterFormat - format of the input parameter. Example: "String".
		ParameterFormat string `json:"parameterFormat,omitempty"`
		// ResourceLocatorKey - JSON path key to link to the parameter in the booking API.
		// Example: "$.data.payments[*].cvv".
		ResourceLocatorKey string `json:"resourceLocatorKey,omitempty"`
	}

	// HoldPolicy describes the hold time deadline for deferred payment.
	HoldPolicy struct {
		// Deadline - deadline for hold time in ISO 8601 date-time (hotel local time). (required)
		Deadline string `json:"deadline"`
	}

	// LengthOfStayPolicy describes minimum and maximum length of stay requirements.
	LengthOfStayPolicy struct {
		// MaximumLengthOfStay - maximum number of nights allowed in a single reservation.
		MaximumLengthOfStay int `json:"maximumLengthOfStay,omitempty"`
		// MaximumLengthOfStayDescription - free text description of the maximum stay policy.
		MaximumLengthOfStayDescription *QualifiedFreeText `json:"maximumLengthOfStayDescription,omitempty"`
		// MinimumLengthOfStay - minimum number of nights required by hotel booking conditions.
		MinimumLengthOfStay int `json:"minimumLengthOfStay,omitempty"`
		// MinimumLengthOfStayDescription - free text description of the minimum stay policy.
		MinimumLengthOfStayDescription *QualifiedFreeText `json:"minimumLengthOfStayDescription,omitempty"`
	}

	// RefundablePolicy describes the refund policy for cancellations.
	RefundablePolicy struct {
		// CancellationRefund - the cancellation refund policy.
		// Enum: NON_REFUNDABLE, REFUNDABLE, REFUNDABLE_UP_TO_DEADLINE, UNKNOWN.
		CancellationRefund string `json:"cancellationRefund,omitempty"`
	}

	// ==================== Price ====================

	// HotelPrice contains price information for the hotel offer.
	HotelPrice struct {
		// Base - base price amount (before taxes).
		Base string `json:"base,omitempty"`
		// Currency - currency code applied to all price elements.
		Currency string `json:"currency,omitempty"`
		// Markups - array of markups applied by any stakeholder (travel agent, merchant mode, etc.).
		Markups []Markup `json:"markups,omitempty"`
		// SellingTotal - selling total = total + margins + markup + totalFees - discounts.
		SellingTotal string `json:"sellingTotal,omitempty"`
		// Taxes - array of taxes applied to the price.
		Taxes []HotelTax `json:"taxes,omitempty"`
		// Total - total price = base + totalTaxes.
		Total string `json:"total,omitempty"`
		// Variations - daily price variations and average daily price when available.
		Variations *PriceVariations `json:"variations,omitempty"`
	}

	// HotelTax represents an IATA tax definition applied to the hotel price.
	HotelTax struct {
		// Amount - amount of the tax.
		Amount string `json:"amount,omitempty"`
		// Code - tax code identifying the tax. Examples: 1=BED_TAX, 2=CITY_TAX.
		Code string `json:"code,omitempty"`
		// Currency - currency code of the tax. MaxLength: 3.
		Currency string `json:"currency,omitempty"`
		// Description - textual description of the tax. Example: "Government tax".
		Description string `json:"description,omitempty"`
		// Included - whether the tax is included in the base amount.
		Included bool `json:"included,omitempty"`
		// Percentage - percentage of the tax. Use with PricingFrequency and PricingMode.
		Percentage string `json:"percentage,omitempty"`
		// ApplicableDate - the applicable period for the tax amount provided by the hotel provider.
		ApplicableDate *TaxApplicableDate `json:"applicableDate,omitempty"`
		// PricingFrequency - specifies if the tax applies per stay or per night.
		// Values: PER_STAY, PER_NIGHT.
		PricingFrequency string `json:"pricingFrequency,omitempty"`
		// PricingMode - specifies if the tax applies per occupant or per room.
		// Values: PER_OCCUPANT, PER_PRODUCT.
		PricingMode string `json:"pricingMode,omitempty"`
	}

	// TaxApplicableDate defines the applicable period for a tax.
	TaxApplicableDate struct {
		// End - end date and time in ISO 8601 format. Example: "2019-12-20T00:00:00Z".
		End string `json:"end,omitempty"`
		// Start - start date and time in ISO 8601 format. Example: "2019-11-22T00:00:00Z".
		Start string `json:"start,omitempty"`
	}

	// PriceVariations contains daily price variations during a stay.
	PriceVariations struct {
		// Changes - collection of price periods when the daily price changes during the stay.
		Changes []PriceVariation `json:"changes,omitempty"`
	}

	// PriceVariation represents a single price period during a stay.
	PriceVariation struct {
		// EndDate - end date of the price period. Format: YYYY-MM-DD. (required)
		EndDate string `json:"endDate"`
		// StartDate - begin date of the price period. Format: YYYY-MM-DD. (required)
		StartDate string `json:"startDate"`
		// Base - base price for this period.
		Base string `json:"base,omitempty"`
		// Currency - currency code for this period.
		Currency string `json:"currency,omitempty"`
		// Markups - markups applied during this period.
		Markups []Markup `json:"markups,omitempty"`
		// SellingTotal - selling total for this period = total + margins + markup + totalFees - discounts.
		SellingTotal string `json:"sellingTotal,omitempty"`
		// Total - total for this period = base + totalTaxes.
		Total string `json:"total,omitempty"`
	}

	// ==================== Room ====================

	// RoomDetails contains room information. DEPRECATED: use RoomInformation instead.
	RoomDetails struct {
		// Description - free text description of the room.
		Description *QualifiedFreeText `json:"description,omitempty"`
		// Type - room type code (3 chars). First char = room type category,
		// second numeric char = number of beds, third char = bed type.
		// Special case "ROH" = Run Of House. Pattern: ^[A-Z0-9*]{3}$.
		Type string `json:"type,omitempty"`
		// TypeEstimated - estimated room category, bed type and number of beds.
		// Parsed from room description, provided for informational purposes only.
		TypeEstimated *EstimatedRoomType `json:"typeEstimated,omitempty"`
	}

	// EstimatedRoomType contains estimated room category, bed type, and number of beds.
	// This information is parsed from the room description (informational only).
	EstimatedRoomType struct {
		// BedType - type of the bed (e.g. SINGLE, DOUBLE, KING, QUEEN).
		BedType string `json:"bedType,omitempty"`
		// Beds - number of beds in the room (1-9).
		Beds int `json:"beds,omitempty"`
		// Category - room category code.
		Category string `json:"category,omitempty"`
	}

	// RoomInformation contains detailed hotel room information including amenities.
	RoomInformation struct {
		// Id - unique identifier of the room information record.
		Id string `json:"id,omitempty"`
		// Amenities - list of room amenities (e.g. WIFI, minibar, etc.).
		Amenities []Amenity `json:"amenities,omitempty"`
		// ArchitectureCode - architectural style of the room.
		// Enum: ART_DECO, BRAZILIAN, CONTEMPORARY, HIGH_RISE, HISTORIC, MEDITERRANEAN,
		// MODERN, ORIENTAL, SOUTHWEST, TRADITIONAL, VICTORIAN, WESTERN, ANCIENT, THEMED.
		ArchitectureCode string `json:"architectureCode,omitempty"`
		// BathroomsPerRoom - number of bathrooms in the room.
		BathroomsPerRoom int `json:"bathroomsPerRoom,omitempty"`
		// BedroomsPerRoom - number of bedrooms in the room.
		BedroomsPerRoom int `json:"bedroomsPerRoom,omitempty"`
		// BedType - type of the bed.
		// Enum: DOUBLE, FUTON, KING, MURPHY_BED, QUEEN, SOFA_BED, TATAMI_MATS, TWIN, SINGLE,
		// FULL, RUN_OF_THE_HOUSE, DORM_BED, WATER_BED.
		BedType string `json:"bedType,omitempty"`
		// Beds - number of beds in the room.
		Beds int `json:"beds,omitempty"`
		// Description - free text description of the room.
		Description *QualifiedFreeText `json:"description,omitempty"`
		// Dimensions - physical dimensions (area, height, width, etc.) of the room.
		Dimensions *Dimensions `json:"dimensions,omitempty"`
		// HotelRoomCategory - marketing/pricing category of the room.
		// Enum: SUITE, BUDGET, CORPORATE_BUSINESS_TRANSIENT, DELUXE, ECONOMY, EXTENDED_STAY,
		// FIRST_CLASS, LUXURY, MEETING_CONVENTION, MODERATE, RESIDENTIAL_APARTMENT, RESORT,
		// UPSCALE, EFFICIENCY, STANDARD, MIDSCALE, OTHER, MIDSCALE_WITHOUT_FOOD_AND_BEVERAGES,
		// UPPER_UPSCALE.
		HotelRoomCategory string `json:"hotelRoomCategory,omitempty"`
		// HotelRoomClassification - structural classification of the room.
		// Enum: ROOM, VILLA, SUITES, APARTMENTS, PENTHOUSES, LOFTS, ACCESSIBLE_ROOMS,
		// NONSMOKING_ROOMS, BUNGALOWS_AND_VILLAS, EXECUTIVE_FLOOR, DOUBLE_BEDROOMS,
		// KING_BEDROOMS, QUEEN_BEDROOMS, STUDIOS, SMOKING_ROOMS, TWIN_BEDROOMS, etc.
		HotelRoomClassification string `json:"hotelRoomClassification,omitempty"`
		// HotelRoomLocation - location of the room within the hotel.
		HotelRoomLocation string `json:"hotelRoomLocation,omitempty"`
		// MaxPersonCapacity - maximum occupancy of the room (adults, children, total).
		MaxPersonCapacity *MaxPersonCapacity `json:"maxPersonCapacity,omitempty"`
		// MaxSleepFurnishings - extra sleeping furnishings available in the room (cribs, extra beds).
		MaxSleepFurnishings *MaxSleepFurnishings `json:"maxSleepFurnishings,omitempty"`
		// Media - list of media (images/videos) associated with the room.
		Media []Media `json:"media,omitempty"`
		// Name - room name as free text with language info.
		Name *QualifiedFreeText `json:"name,omitempty"`
		// PolicyDescriptions - free text descriptions of the room-level policies with language info.
		PolicyDescriptions []QualifiedFreeText `json:"policyDescriptions,omitempty"`
		// Quantity - number of rooms of this type available.
		Quantity int `json:"quantity,omitempty"`
		// Type - room type code (3 chars).
		Type string `json:"type,omitempty"`
		// TypeEstimated - estimated room category, bed type, and number of beds.
		TypeEstimated *EstimatedRoomType `json:"typeEstimated,omitempty"`
		// ViewCode - the view from the room.
		// Enum: AIRPORT, BAY, CITY, COURTYARD, GOLF, HARBOR, INTERCOASTAL, LAKE, MARINA,
		// MOUNTAIN, OCEAN, POOL, RIVER, WATER, BEACH, GARDEN, PARK, FOREST, RAIN_FOREST,
		// VARIOUS, SLOPE, STRIP, COUNTRYSIDE, SEA.
		ViewCode string `json:"viewCode,omitempty"`
	}

	// Amenity represents a room amenity with its type, pricing, and associated media.
	Amenity struct {
		// Code - unique code representing the amenity.
		Code string `json:"code,omitempty"`
		// Description - description of the amenity.
		Description string `json:"description,omitempty"`
		// AmenityType - type/category of the amenity. Example: "WIFI", "Internet".
		AmenityType string `json:"amenityType,omitempty"`
		// AmenityAttribute - attribute related to the amenity type. Example: "USB Outlet" for Power amenity.
		AmenityAttribute string `json:"amenityAttribute,omitempty"`
		// AmenityQualityAssessment - ranking indicator for the amenity. Example: "Standard".
		AmenityQualityAssessment string `json:"amenityQualityAssessment,omitempty"`
		// AmenityPerformanceAssessment - performance indicator. Example: "Low", "High" (e.g. Wifi bandwidth).
		AmenityPerformanceAssessment string `json:"amenityPerformanceAssessment,omitempty"`
		// AmenityProvider - source of the amenity content. Example: ATPCO.
		AmenityProvider *AmenityProvider `json:"amenityProvider,omitempty"`
		// IsChargeable - true if usage of the amenity is chargeable. Default: false.
		IsChargeable bool `json:"isChargeable,omitempty"`
		// Price - price information for using the amenity.
		Price *AmenityPrice `json:"price,omitempty"`
		// PricingMethod - how the amenity usage cost is assessed.
		// Enum: DAILY, HOURLY, HALF_DAY, PER_OCCURRENCE, PER_EVENT, PER_PERSON, FIRST_USE,
		// PER_MINUTE, COMPLIMENTARY, WEEKLY, PER_STAY, PER_FUNCTION, PER_ROOM_PER_STAY,
		// PER_ROOM_PER_NIGHT, PER_PERSON_PER_STAY, PER_PERSON_PER_NIGHT, PER_RESERVATION_BOOKNG.
		PricingMethod string `json:"pricingMethod,omitempty"`
		// Quantity - how many counts are available for this amenity type.
		Quantity int `json:"quantity,omitempty"`
		// Medias - list of media (images/videos) associated with the amenity.
		Medias []Media `json:"Medias,omitempty"`
	}

	// ==================== Hotel ====================

	// Hotel contains hotel content information.
	Hotel struct {
		// HotelId - Amadeus Property Code (8 chars). Pattern: ^[A-Z0-9]{8}$. Example: "ADPAR001".
		HotelId string `json:"hotelId,omitempty"`
		// ChainCode - brand code (e.g. "RT") or merchant code (e.g. "AD"). Example: "AD".
		ChainCode string `json:"chainCode,omitempty"`
		// Name - hotel name. Example: "Hotel de Paris".
		Name string `json:"name,omitempty"`
		// TermsAndConditions - link to the terms and conditions the guest must approve to book.
		TermsAndConditions string `json:"termsAndConditions,omitempty"`
		// Self - link to retrieve the hotel details.
		Self string `json:"self,omitempty"`
	}

	// ==================== Payment Output ====================

	// PaymentOutput contains payment information in the booking response.
	PaymentOutput struct {
		// Method - indicates the method of payment used. (required)
		// Enum: CREDIT_CARD, CREDIT_CARD_AGENCY, CREDIT_CARD_TRAVELER,
		// AGENCY_ACCOUNT, VCC_BILLBACK, VCC_B2B_WALLET, TRAVEL_AGENT_ID.
		Method string `json:"method"`
		// PaymentInstructions - optional free text with payment instructions sent to the hotelier.
		PaymentInstructions string `json:"paymentInstructions,omitempty"`
		// HotelSupplierInformation - hotel supplier contact details (phone, fax, email).
		HotelSupplierInformation *HotelSupplierInformation `json:"hotelSupplierInformation,omitempty"`
		// IataTravelAgency - agency IATA/ARC Number used to guarantee the booking.
		IataTravelAgency *IataTravelAgency `json:"iataTravelAgency,omitempty"`
		// BillBack - billback payment details for VCC_BILLBACK method, including Conferma deployment info.
		BillBack *BillBackOutput `json:"billBack,omitempty"`
		// B2bWallet - VCC B2B wallet information with the generated virtual credit card reference.
		B2bWallet *B2bWallet `json:"b2bWallet,omitempty"`
		// PaymentCard - credit card information used for payment (no securityCode in response).
		PaymentCard *PaymentCardOutput `json:"paymentCard,omitempty"`
	}

	// HotelSupplierInformation contains hotel supplier contact details.
	HotelSupplierInformation struct {
		// Phone - phone number. Recommended E.123 format. MinLength: 2, MaxLength: 90.
		Phone string `json:"phone,omitempty"`
		// Fax - fax number. Recommended E.123 format. MinLength: 2, MaxLength: 90.
		Fax string `json:"fax,omitempty"`
		// Email - email address. Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$. MaxLength: 90.
		Email string `json:"email,omitempty"`
	}

	// IataTravelAgency holds the IATA/ARC Number used to guarantee the booking.
	IataTravelAgency struct {
		// IataNumber - the agency IATA/ARC number. (required)
		IataNumber string `json:"iataNumber"`
	}

	// BillBackOutput contains billback payment information in the response.
	// Includes the deployment ID returned by Amadeus Payment server.
	BillBackOutput struct {
		// TravelAgencyId - Travel Agency Conferma account (CAI).
		TravelAgencyId string `json:"travelAgencyId,omitempty"`
		// BookerId - Travel Agent Conferma ID (CBI).
		BookerId string `json:"bookerId,omitempty"`
		// PaymentInstructions - optional free text specifying payment instructions to the hotelier.
		PaymentInstructions string `json:"paymentInstructions,omitempty"`
		// BillbackProviderDeploymentId - returned by Amadeus Payment server, used as the payment reference. (required)
		BillbackProviderDeploymentId string `json:"billbackProviderDeploymentId"`
		// BillbackProviderCode - billback provider code. For Conferma, it is "CN". (required)
		BillbackProviderCode string `json:"billbackProviderCode"`
		// BillbackProviderAccountNumber - Conferma account number. (required)
		BillbackProviderAccountNumber string `json:"billbackProviderAccountNumber"`
		// HotelSupplierInformation - hotel supplier contact details.
		HotelSupplierInformation *HotelSupplierInformation `json:"hotelSupplierInformation,omitempty"`
	}

	// B2bWallet contains VCC B2B wallet information generated for the booking.
	// Only used for VCC_B2B_WALLET payment method (Amadeus Value Hotels).
	B2bWallet struct {
		// VirtualCreditCardId - Amadeus Payment Reference for the generated virtual credit card. (required)
		VirtualCreditCardId string `json:"virtualCreditCardId"`
		// PaymentProvider - payment provider name. Example: "IXARIS". (required)
		PaymentProvider string `json:"paymentProvider"`
	}

	// PaymentCardOutput contains credit card information in the response.
	// Note: securityCode is NOT included in the response for security reasons.
	PaymentCardOutput struct {
		// PaymentCardInfo - credit card details (vendorCode, cardNumber, expiryDate). (required)
		PaymentCardInfo PaymentCardInfoOutput `json:"paymentCardInfo"`
		// Address - billing address of the credit card holder.
		Address *Address `json:"address,omitempty"`
	}

	// PaymentCardInfoOutput contains credit card details in the response.
	// securityCode is intentionally omitted for security.
	PaymentCardInfoOutput struct {
		// VendorCode - two-letter card type code. Example: VI (Visa), MA (MasterCard). MaxLength: 30. (required)
		VendorCode string `json:"vendorCode"`
		// CardNumber - the credit card number. (required)
		CardNumber string `json:"cardNumber"`
		// ExpiryDate - expiration date in MMYY format. (required)
		ExpiryDate string `json:"expiryDate"`
		// HolderName - name of the credit card holder. MaxLength: 99.
		HolderName string `json:"holderName,omitempty"`
	}

	// Address contains postal/billing address details.
	Address struct {
		// Lines - unformatted address lines (street, apartment, suite, building, floor, etc.).
		Lines []string `json:"lines,omitempty"`
		// PostalCode - post office code number.
		PostalCode string `json:"postalCode,omitempty"`
		// CityName - city name.
		CityName string `json:"cityName,omitempty"`
		// PostalBox - postal box. Example: "BP 220".
		PostalBox string `json:"postalBox,omitempty"`
		// StateCode - ISO 3166-2 subdivision code (province/state).
		StateCode string `json:"stateCode,omitempty"`
		// CountryCode - ISO 3166-1 country code. Pattern: ^[A-Z]{2}$. Example: "FR".
		CountryCode string `json:"countryCode,omitempty"`
	}

	// ==================== Guests ====================

	// ResponseGuest contains guest information in the booking response.
	// Includes both the temporary id (tid) and the Amadeus-assigned id.
	ResponseGuest struct {
		// Tid - temporary unique id of the guest, arbitrarily chosen by the user until Amadeus provides one.
		Tid int `json:"tid,omitempty"`
		// Id - unique id of the guest provided by Amadeus application. (required)
		Id int `json:"id"`
		// Title - title/gender of the guest.
		// Enum: MRS, MR, MS, CHILD, DR, MADAM, MESSRS, MISS, SIR.
		Title string `json:"title,omitempty"`
		// FirstName - first name (and middle name) of the guest.
		// Pattern: ^[A-Za-z ]*$. MaxLength: 56.
		FirstName string `json:"firstName,omitempty"`
		// LastName - last name of the guest.
		// Pattern: ^[A-Za-z ]*$. MaxLength: 57.
		LastName string `json:"lastName,omitempty"`
		// Phone - phone number. Recommended E.123 format. MaxLength: 199.
		Phone string `json:"phone,omitempty"`
		// Email - email address. MaxLength: 90.
		Email string `json:"email,omitempty"`
		// ChildAge - age of the child guest. If not provided, guest is treated as an adult.
		ChildAge int `json:"childAge,omitempty"`
		// FrequentTraveler - airline frequent flyer info.
		// In retrieve, when a guest has several frequent flyer numbers used in different bookings, all are listed.
		FrequentTraveler []FrequentTraveler `json:"frequentTraveler,omitempty"`
	}

	// FrequentTraveler represents an airline frequent flyer program membership.
	FrequentTraveler struct {
		// AirlineCode - code of the airline. MinLength: 2, MaxLength: 3. Example: "AF". (required)
		AirlineCode string `json:"airlineCode"`
		// FrequentTravelerId - the frequent traveler membership ID. Example: "32546971326". (required)
		FrequentTravelerId string `json:"frequentTravelerId"`
	}

	// ==================== Arrival Information ====================

	// ArrivalInformation contains optional information on how the guest is arriving to the hotel.
	ArrivalInformation struct {
		// ArrivalFlightDetails - flight details of the guest's arriving flight.
		ArrivalFlightDetails *ArrivalFlightDetails `json:"arrivalFlightDetails,omitempty"`
	}

	// ArrivalFlightDetails contains the arriving flight segment details.
	ArrivalFlightDetails struct {
		// CarrierCode - airline carrier code. Example: "LH". (required)
		CarrierCode string `json:"carrierCode,omitempty"`
		// Number - flight segment number. Example: "1050". (required)
		Number string `json:"number,omitempty"`
		// Departure - departure airport info. (required)
		Departure *Departure `json:"departure,omitempty"`
		// Arrival - arrival airport info with terminal and local time. (required)
		Arrival *Arrival `json:"arrival,omitempty"`
	}

	// Departure contains departure airport information.
	Departure struct {
		// IataCode - IATA airport code. Example: "JFK". (required)
		IataCode string `json:"iataCode"`
	}

	// Arrival contains arrival airport information including terminal and local arrival time.
	Arrival struct {
		// IataCode - IATA airport code. Example: "JFK". (required)
		IataCode string `json:"iataCode"`
		// Terminal - terminal name/number. Example: "T2". (required)
		Terminal string `json:"terminal"`
		// At - local date and time of the flight arrival.
		// Format: YYYY-MM-DDTHH:mm:ss (e.g. 2017-10-23T20:00:00+02:00). (required)
		At time.Time `json:"at"`
	}

	// ==================== Associated Records ====================

	// AssociatedRecord contains reference and origin of the hotel order record in Amadeus GDS.
	AssociatedRecord struct {
		// Reference - record locator of the PNR in Amadeus GDS.
		// Pattern: ^[A-Z0-9]{6}$. Example: "ABCDEF". (required)
		Reference string `json:"reference"`
		// OriginSystemCode - always set to "GDS" for Amadeus PNR record locators. (required)
		OriginSystemCode string `json:"originSystemCode"`
		// Direction - indicates whether the reference results from a split operation.
		// Values: "CHILD" (resulted from split), "PARENT" (originated a split).
		Direction string `json:"direction,omitempty"`
		// AssociationType - specifies the nature of association between two orders.
		// Example: "SPLIT".
		AssociationType string `json:"associationType,omitempty"`
	}

	// ==================== Warning & Error ====================

	// Warning contains warning information returned alongside a successful response.
	Warning struct {
		// Code - machine-readable error code from the Canned Messages table. (required)
		Code int `json:"code"`
		// Title - error title with 1:1 correspondence to the error code. May be localized. (required)
		Title string `json:"title"`
		// Detail - human-readable explanation specific to this occurrence.
		// Gives the consumer an idea of what went wrong and how to recover.
		Detail string `json:"detail,omitempty"`
		// Documentation - link to a web page or file with further documentation.
		Documentation string `json:"documentation,omitempty"`
		// Sources - array of source objects identifying the parameter or field that caused the warning.
		Sources []WarningSource `json:"sources,omitempty"`
		// Relationships - relationships from one entity to other entities (e.g. passenger to flight segments).
		Relationships *Relationships `json:"relationships,omitempty"`
	}

	// Relationships indicates relationships from one entity to many other entities.
	Relationships struct {
		// Href - URL of the related resource.
		Href string `json:"href,omitempty"`
		// Methods - accepted HTTP methods. Enum: GET, POST, PUT, PATCH, DELETE.
		Methods []string `json:"methods,omitempty"`
		// Id - id of the related resource.
		Id string `json:"id,omitempty"`
		// Rel - type of relation between entities per IANA link-relations.
		Rel string `json:"rel,omitempty"`
		// Collection - details of the related items.
		Collection []Relationship `json:"collection,omitempty"`
	}

	// Relationship allows cross-referencing two entities via link and/or id.
	Relationship struct {
		// Id - id of the related resource.
		Id string `json:"id,omitempty"`
		// Type - type of the related resource. Example: "processed-dcs-passenger".
		Type string `json:"type,omitempty"`
		// Ref - local reference of the related resource (URI reference).
		Ref string `json:"ref,omitempty"`
		// TargetSchema - schema definition for building the associated API request operation.
		TargetSchema string `json:"targetSchema,omitempty"`
		// TargetMediaType - media type for the referenced API request (if different from default).
		TargetMediaType string `json:"targetMediaType,omitempty"`
		// HrefSchema - type/format/pattern/enum definitions for each URI parameter.
		HrefSchema string `json:"hrefSchema,omitempty"`
		// ExpirationDate - expiration date and time of the related resource. ISO 8601 format.
		ExpirationDate string `json:"expirationDate,omitempty"`
	}

	// Error contains error information for 400/500 responses.
	Error struct {
		// Status - HTTP status code of this response. Present only in terminal errors.
		// In case of multiple errors, they all have the same status.
		Status int `json:"status,omitempty"`
		// Code - machine-readable error code to enable consumer code to handle this error type. (required)
		Code int `json:"code"`
		// Title - error title with 1:1 correspondence to the error code. May be localized. (required)
		Title string `json:"title"`
		// Detail - human-readable explanation specific to this occurrence.
		// Gives the consumer an idea of what went wrong and how to recover.
		Detail string `json:"detail,omitempty"`
		// Source - identifies the parameter or field in the request that caused the error.
		Source *ErrorSource `json:"source,omitempty"`
		// Documentation - link to a web page or file with further documentation.
		Documentation string `json:"documentation,omitempty"`
	}

	// ErrorSource identifies the source of an error in the request.
	ErrorSource struct {
		// Parameter - the key of the URI path or query parameter that caused the error.
		Parameter string `json:"parameter,omitempty"`
		// Pointer - a JSON Pointer [RFC6901] to the associated entity in the request body.
		Pointer string `json:"pointer,omitempty"`
		// Example - a sample input to guide the user when resolving the issue.
		Example string `json:"example,omitempty"`
	}
)
