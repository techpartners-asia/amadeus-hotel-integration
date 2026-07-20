package requestBookingDTO

import "time"

// Int returns a pointer to v, a convenience for optional numeric request fields
// such as Guest.ChildAge where 0 is a meaningful value (an infant) and must be
// distinguishable from "not set".
func Int(v int) *int { return &v }

type (
	// HotelBookingRequest allows the creation of a hotel order (PNR) and its hotel segments (HHL).
	// It can also be used to add a hotel booking (HHL) in an already existing hotel order (PNR).
	HotelBookingRequest struct {
		// Data - the booking data payload (required)
		Data BookingData `json:"data,omitempty"`
	}

	// BookingData contains all the information needed to create a hotel order.
	BookingData struct {
		// Type - "hotel-order" to create a hotel order (required)
		Type string `json:"type,omitempty"`
		// Guests - list of all guests with detailed information (required)
		Guests []Guest `json:"guests,omitempty"`
		// RoomAssociations - correlates a room to a guest and an offer. Min 1, max 9 rooms.
		// The application only supports multi-identical rooms (same hotel, same dates, same supplier).
		RoomAssociations []RoomAssociation `json:"roomAssociations,omitempty"`
		// Payment - payment information for the booking (required)
		Payment Payment `json:"payment,omitempty"`
		// TravelAgent - travel agent details including contact info (required)
		TravelAgent TravelAgent `json:"travelAgent,omitempty"`
		// ArrivalInformation - optional information on how the guest is arriving to the hotel
		ArrivalInformation *ArrivalInformation `json:"arrivalInformation,omitempty"`
		// AssociatedRecord - reference to an existing Amadeus PNR to add the hotel booking into
		AssociatedRecord *AssociatedRecord `json:"associatedRecord,omitempty"`
	}

	// Guest represents a guest with their personal details, loyalty programs, and age when child.
	Guest struct {
		// Tid - temporary id of a guest. Correlates a given guest with a room in the roomAssociation section.
		// It is arbitrarily chosen by the user and must be unique. (required)
		// No omitempty: tid is required, and a caller numbering guests from 0
		// would otherwise have the field silently dropped.
		Tid int `json:"tid"`
		// Title - title/gender of the guest. Enum: MRS, MR, MS, CHILD, DR, MADAM, MESSRS, MISS, SIR.
		// Only English alphas [A-Z] and spaces are supported. Sum of title + firstName + lastName <= 62 chars.
		Title string `json:"title,omitempty"`
		// FirstName - first name (and middle name) of the guest. Mandatory when creating a hotel order.
		// Only English alphas [A-Z] and spaces. Pattern: ^[A-Za-z ]*$. MinLength: 1, MaxLength: 56.
		FirstName string `json:"firstName,omitempty"`
		// LastName - last name of the guest. Mandatory when creating a hotel order.
		// Only English alphas [A-Z] and spaces. Pattern: ^[A-Za-z ]*$. MinLength: 2, MaxLength: 57.
		LastName string `json:"lastName,omitempty"`
		// Phone - phone number of the guest. Recommended to use standard E.123 format.
		// MinLength: 2, MaxLength: 199. Example: "+33679278416"
		Phone string `json:"phone,omitempty"`
		// Email - email address of the guest.
		// Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$. MinLength: 3, MaxLength: 90.
		Email string `json:"email,omitempty"`
		// ChildAge - mandatory if the guest is a child. Otherwise, the system considers them as an adult.
		// Pointer so that nil means "adult" while 0 still means an infant under
		// one year; a plain int with omitempty would silently turn an infant
		// into an adult.
		ChildAge *int `json:"childAge,omitempty"`
		// FrequentTraveler - airline frequent flyer info of the guest.
		// Only the first element is transmitted to the hotel provider during creation. Provide only one.
		FrequentTraveler []FrequentTraveler `json:"frequentTraveler,omitempty"`
	}

	// FrequentTraveler represents an airline frequent flyer program membership.
	FrequentTraveler struct {
		// AirlineCode - code of the airline. MinLength: 2, MaxLength: 3. Example: "AF" (required)
		AirlineCode string `json:"airlineCode,omitempty"`
		// FrequentTravelerId - the frequent traveler membership ID. Example: "32546971326" (required)
		FrequentTravelerId string `json:"frequentTravelerId,omitempty"`
	}

	// RoomAssociation correlates one single room to guest(s), a payment, and a hotel offer.
	// Min 1, max 9 rooms. Multi-room bookings must have same hotel, same dates, same supplier.
	RoomAssociation struct {
		// HotelOfferId - hotel offer ID from availability response, identifying the product to book.
		// Pattern: ^[A-Z0-9]*$. MinLength: 2, MaxLength: 100. (required)
		HotelOfferId string `json:"hotelOfferId,omitempty"`
		// GuestReferences - array of guest references listing all guests occupying the room.
		// First guest is the main guest holding the reservation and form of payment.
		// Following references are accompagnants. (required)
		GuestReferences []GuestReference `json:"guestReferences,omitempty"`
		// SpecialRequest - special request to send to the reception (optional).
		// MinLength: 2, MaxLength: 120.
		SpecialRequest string `json:"specialRequest,omitempty"`
		// TravelAgentManualMarkup - overrides the amount computed by the Margin Manager
		// when Hotel Markup is activated for the travel agency.
		TravelAgentManualMarkup *TravelAgentManualMarkup `json:"travelAgentManualMarkup,omitempty"`
	}

	// GuestReference links a guest to a room with optional hotel loyalty program.
	GuestReference struct {
		// GuestReference - reference to the guest id (tid at creation time). (required)
		GuestReference string `json:"guestReference,omitempty"`
		// HotelLoyaltyId - Hotel Chain Rewards Program Membership ID of the guest.
		// Used for Rewards Points, online check-in, fast check-out.
		// An error is returned by the Chain if the number is invalid.
		// Pattern: ^[A-Z0-9-]{1,21}$. MinLength: 1, MaxLength: 21.
		HotelLoyaltyId string `json:"hotelLoyaltyId,omitempty"`
	}

	// TravelAgentManualMarkup overrides the margin computed by Margin Manager when Hotel Markup is activated.
	TravelAgentManualMarkup struct {
		// Amount - the markup amount. Pattern: ^\-?[0-9]+(\.[0-9]+)?$ (required)
		Amount string `json:"amount,omitempty"`
		// Currency - 3-letter currency code. Pattern: ^[A-Z0-9*]{3}$. Example: "EUR" (required)
		Currency string `json:"currency,omitempty"`
	}

	// Payment contains the hotel payment information.
	Payment struct {
		// Method - indicates the method of payment. (required)
		// Enum: CREDIT_CARD, CREDIT_CARD_AGENCY, CREDIT_CARD_TRAVELER, AGENCY_ACCOUNT,
		// VCC_BILLBACK, VCC_B2B_WALLET, TRAVEL_AGENT_ID.
		// - CREDIT_CARD: payment through a credit card (provide creditCard info)
		// - AGENCY_ACCOUNT: payment through agency credit line
		// - VCC_BILLBACK: direct payment via billback provider (e.g. Conferma)
		// - VCC_B2B_WALLET: payment between travel agency and Amadeus Merchant using VCC
		// - TRAVEL_AGENT_ID: payment with IATA booking source
		// - CREDIT_CARD_AGENCY: exclusively for Amadeus Value Hotel (agency card)
		// - CREDIT_CARD_TRAVELER: exclusively for Amadeus Value Hotel (guest card)
		Method string `json:"method,omitempty"`
		// PaymentInstructions - optional free text specifying payment instructions sent to the hotelier.
		PaymentInstructions string `json:"paymentInstructions,omitempty"`
		// PayerCode - optional free text specifying the corporation payerCode for VCC generation.
		// Applicable only for VCC_B2B_WALLET. Pattern: ^[A-Z0-9_]*$. MinLength: 1, MaxLength: 40.
		PayerCode string `json:"payerCode,omitempty"`
		// HotelSupplierInformation - contact details of the hotel supplier.
		HotelSupplierInformation *HotelSupplierInformation `json:"hotelSupplierInformation,omitempty"`
		// IataTravelAgency - agency IATA/ARC Number used to guarantee the booking.
		// If not provided, taken from the Amadeus Office profile.
		IataTravelAgency *IataTravelAgency `json:"iataTravelAgency,omitempty"`
		// BillBack - used when the booking is paid with a virtual credit card via an external provider (e.g. Conferma).
		BillBack *BillBack `json:"billBack,omitempty"`
		// PaymentCard - credit card information for CREDIT_CARD payment method.
		PaymentCard *PaymentCard `json:"paymentCard,omitempty"`
	}

	// HotelSupplierInformation contains hotel supplier contact details.
	HotelSupplierInformation struct {
		// Phone - phone number. Recommended E.123 format. MinLength: 2, MaxLength: 90.
		Phone string `json:"phone,omitempty"`
		// Fax - fax number. Recommended E.123 format. MinLength: 2, MaxLength: 90. (required)
		Fax string `json:"fax,omitempty"`
		// Email - email address. Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$. MinLength: 3, MaxLength: 90.
		Email string `json:"email,omitempty"`
	}

	// IataTravelAgency holds the IATA/ARC Number used to guarantee the booking.
	IataTravelAgency struct {
		// IataNumber - the IATA/ARC number. If not provided, taken from Amadeus Office profile. (required)
		IataNumber string `json:"iataNumber,omitempty"`
	}

	// BillBack is used for VCC_BILLBACK payment. Requires a contract with Conferma and a banking partner.
	// CAI (TravelAgencyId) and CBI (BookerId) can be provided in input or taken from Amadeus Office Profile.
	BillBack struct {
		// TravelAgencyId - Travel Agency Conferma account (CAI).
		TravelAgencyId string `json:"travelAgencyId,omitempty"`
		// BookerId - Travel Agent Conferma ID (CBI).
		BookerId string `json:"bookerId,omitempty"`
		// PaymentInstructions - DEPRECATED. Use the same field under Payment. Specifies instructions sent to the hotelier.
		PaymentInstructions string `json:"paymentInstructions,omitempty"`
		// BillbackProviderCode - billback provider code. For Conferma provider, it will be "CN". (required)
		BillbackProviderCode string `json:"billbackProviderCode,omitempty"`
		// BillbackProviderAccountNumber - Conferma account number. (required)
		BillbackProviderAccountNumber string `json:"billbackProviderAccountNumber,omitempty"`
		// HotelSupplierInformation - hotel supplier contact details.
		HotelSupplierInformation *HotelSupplierInformation `json:"hotelSupplierInformation,omitempty"`
	}

	// PaymentCard contains credit card details including optional 3DS authentication and billing address.
	PaymentCard struct {
		// PaymentCardInfo - the credit card details (vendorCode, cardNumber, expiryDate are required).
		PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo,omitempty"`
		// ThreeDomainSecure - 3D Secure authentication data (version, eci, cryptogramValue are required).
		ThreeDomainSecure *ThreeDomainSecure `json:"threeDomainSecure,omitempty"`
		// Address - billing address of the credit card holder.
		Address *Address `json:"address,omitempty"`
	}

	// PaymentCardInfo contains the credit card details.
	PaymentCardInfo struct {
		// VendorCode - two-letter card type code. E.g. VI=VISA, MA=MasterCard, AX=American Express.
		// Pattern: ^[A-Z]{2}$. (required)
		VendorCode string `json:"vendorCode,omitempty"`
		// CardNumber - credit card number. Pattern: ^[0-9]{14,19}$. MinLength: 14, MaxLength: 19. (required)
		CardNumber string `json:"cardNumber,omitempty"`
		// ExpiryDate - expiration date. Format: MMYY or YYYY-MM. Example: "1224" for December 2024. (required)
		ExpiryDate string `json:"expiryDate,omitempty"`
		// SecurityCode - card verification code (CVV/CVC). Pattern: ^[0-9]{3,4}$. MinLength: 3, MaxLength: 4.
		// Strongly recommended especially for Aggregators.
		SecurityCode string `json:"securityCode,omitempty"`
		// HolderName - name of the credit card holder. Pattern: ^[A-Za-z ]*$. MinLength: 1, MaxLength: 99.
		HolderName string `json:"holderName,omitempty"`
	}

	// ThreeDomainSecure contains the 3D Secure (3DS) authentication transaction summary.
	ThreeDomainSecure struct {
		// Version - 3DS protocol version. Examples: "1.0.2", "2.1.0", "2.2". MaxLength: 5. (required)
		Version string `json:"version,omitempty"`
		// DsTransactionId - unique transaction identifier assigned by the Directory Server (3DS V2). MaxLength: 36.
		DsTransactionId string `json:"dsTransactionId,omitempty"`
		// Eci - Electronic Commerce Indicator. Required for Version 1.0.2 and 2.1.0.
		// Enum: "00" (Failed-Visa/MC), "01" (Incomplete-MC), "02" (Successful-MC),
		// "05" (Successful-Visa), "06" (Attempted-Visa), "07" (Unable-Visa). (required)
		Eci string `json:"eci,omitempty"`
		// Xid - 3DS transaction identifier for version < 2.0. Must be Base64 encoded. MaxLength: 28.
		Xid string `json:"xid,omitempty"`
		// CryptogramValue - authentication verification code (CAVV for Visa, AEVV for AmEx).
		// Base64 encoded, 20-byte giving 28-byte result. MaxLength: 28. (required)
		CryptogramValue string `json:"cryptogramValue,omitempty"`
		// ParesStatus - authentication status for 3DS Version 1.x.
		// Enum: Y (Successful), N (Failed), U (Unable), A (Attempts), E (Error).
		ParesStatus string `json:"paresStatus,omitempty"`
		// ParesStatusLabel - human-readable label for ParesStatus.
		// Enum: SUCCESSFUL, FAILED, UNABLE_TO_AUTHENTICATE, ATTEMPTS_PROCESSING, ERROR.
		ParesStatusLabel string `json:"paresStatusLabel,omitempty"`
		// VeresStatus - indicates whether the cardholder is enrolled. Only for Version 1.x.
		// Enum: Y (Enrolled), N (Not Enrolled), U (Unable to Verify), E (Error), A (Attempts).
		VeresStatus string `json:"veresStatus,omitempty"`
		// VeresStatusLabel - human-readable label for VeresStatus.
		// Enum: ENROLLED, NOT_ENROLLED, UNABLE_TO_VERIFY, ERROR, ATTEMPTS_PROCESSING.
		VeresStatusLabel string `json:"veresStatusLabel,omitempty"`
		// TransStatus - overall 3DS transaction state. Only for Version >= 2.x.
		// Enum: Y, N, U, A, E, C (Challenge), D (Decoupled), R (Rejected), I (Info Only).
		TransStatus string `json:"transStatus,omitempty"`
		// TransStatusLabel - human-readable label for TransStatus.
		// Enum: SUCCESSFUL, FAILED, UNABLE_TO_AUTHENTICATE, ATTEMPTS_PROCESSING, ERROR,
		// CHALLENGE_REQUESTED, DECOUPLED_CHALLENGE_REQUESTED, AUTHENTICATION_REJECTED, INFORMATION_ONLY.
		TransStatusLabel string `json:"transStatusLabel,omitempty"`
	}

	// Address contains the billing or postal address details.
	Address struct {
		// Lines - unformatted address lines (street, apartment, suite, etc.).
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

	// TravelAgent contains travel agent details. The contact email is required.
	TravelAgent struct {
		// Contact - travel agent contact information (email is required).
		Contact Contact `json:"contact,omitempty"`
		// TravelAgentId - Travel Agent ID / Booking source / IATA number.
		// When provided, it overrides the booking source receiving the commission.
		// If not provided, defaults to the IATA Number of the connected office profile.
		TravelAgentId string `json:"travelAgentId,omitempty"`
	}

	// Contact contains the travel agent's contact information.
	Contact struct {
		// Email - travel agency email. (required)
		// Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$. MinLength: 3, MaxLength: 90.
		Email string `json:"email,omitempty"`
		// Fax - travel agency fax. Recommended to fill. Taken from Office Profile if not provided.
		// Pattern: ^[+][1-9][0-9]{4,18}$. MinLength: 2, MaxLength: 90.
		Fax string `json:"fax,omitempty"`
		// Phone - travel agency phone. Recommended to fill. Taken from Office Profile if not provided.
		// MinLength: 2, MaxLength: 199.
		Phone string `json:"phone,omitempty"`
	}

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
		// Arrival - arrival airport info with terminal and time.
		Arrival *Arrival `json:"arrival,omitempty"`
	}

	// Departure contains the departure airport information.
	Departure struct {
		// IataCode - IATA airport code. Example: "JFK". (required)
		IataCode string `json:"iataCode,omitempty"`
	}

	// Arrival contains the arrival airport information including terminal and local arrival time.
	Arrival struct {
		// IataCode - IATA airport code. Example: "JFK". (required)
		IataCode string `json:"iataCode,omitempty"`
		// Terminal - terminal name/number. Example: "T2". (required)
		Terminal string `json:"terminal,omitempty"`
		// At - local date and time of the flight arrival.
		// Format: YYYY-MM-DDTHH:mm:ss (e.g. 2017-02-10T20:40:00). (required)
		At time.Time `json:"at,omitempty"`
	}

	// AssociatedRecord describes the association with an existing Amadeus PNR.
	AssociatedRecord struct {
		// Reference - record locator of the PNR in Amadeus GDS. Example: "JKL789". (required)
		Reference string `json:"reference,omitempty"`
		// OriginSystemCode - always set to "GDS" for Amadeus PNR record locators. (required)
		OriginSystemCode string `json:"originSystemCode,omitempty"`
	}
)
