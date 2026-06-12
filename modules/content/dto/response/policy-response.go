package responseContentDTO

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
		PaymentPolicies      []PaymentPolicyResponse      `json:"paymentPolicies"`
		CheckInOutPolicies   []CheckInOutPolicyResponse   `json:"checkInOutPolicies"`
		PetsPolicies         []PetsPolicyResponse         `json:"petsPolicies"`
		CancellationPolicies []CancellationPolicyResponse `json:"cancellationPolicies"` // Describes the cancellation policies applicable to the property
		TaxPolicies          []TaxPolicyResponse          `json:"taxPolicies"`          // Describes the taxes applicable at the property
		CommissionPolicies   []CommissionPolicyResponse   `json:"commissionPolicies"`   // Describes the commission policies applicable to the property
		StayRequirements     []QualifiedFreeTextResponse  `json:"stayRequirements"`     // Describes the stay requirements such as minimum/maximum length of stay
		GuestPolicies        []GuestPolicyResponse        `json:"guestPolicies"`        // Describes the guest policies applicable to the property
		LoyaltyPolicies      []LoyaltyBenefitResponse     `json:"loyaltyPolicies"`      // Describes the loyalty benefits applicable to the property
	}

	// * Describes the conditions under which a booking can be cancelled and the resulting charges.
	CancellationPolicyResponse struct {
		Amount         string                    `json:"amount"`         // Cancellation charge amount applicable when the policy is triggered
		NumberOfNights int                       `json:"numberOfNights"` // Number of nights charged as a cancellation penalty
		Percentage     string                    `json:"percentage"`     // Cancellation charge expressed as a percentage
		Deadline       string                    `json:"deadline"`       // Deadline before which the booking can be cancelled free of charge
		Description    QualifiedFreeTextResponse `json:"description"`    // Free-text description of the cancellation policy
		PolicyType     string                    `json:"policyType"`     // Type of the cancellation policy
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
		Percentage  string                    `json:"percentage"`  // Commission expressed as a percentage
		Amount      string                    `json:"amount"`      // Commission amount
		Description QualifiedFreeTextResponse `json:"description"` // Free-text description of the commission policy
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
		Name  string          `json:"name"`  // Name of the loyalty program
		Owner string          `json:"owner"` // Owner of the loyalty program
		Media []MediaResponse `json:"media"` // Media associated to the loyalty program
	}

	// * Pets policies
	PetsPolicyResponse struct {
		Code          string        `json:"code"`          // example: 119 describes the pets policy code of the property
		Description   string        `json:"description"`   // example: Only dogs are allowed
		PricingMethod PricingMethod `json:"pricingMethod"` // example: PER_NIGHT. Pricing method for the pets policy
	}

	// * Check-in and Check-out policies
	CheckInOutPolicyResponse struct {
		CheckIn             string                    `json:"checkIn"`             // example: 13:00:00. Check-in From time limit in ISO-8601 format [http://www.w3.org/TR/xmlschema-2/#time]
		CheckInDescription  QualifiedFreeTextResponse `json:"checkInDescription"`  // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
		CheckOut            string                    `json:"checkOut"`            // example: 12:00:00. Check-out To time limit in ISO-8601 format [http://www.w3.org/TR/xmlschema-2/#time]
		CheckOutDescription QualifiedFreeTextResponse `json:"checkOutDescription"` // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
	}

	// * Booking Rules
	PaymentPolicyResponse struct {
		PaymentType       PaymentType                 `json:"paymentType"` // example: DEPOSIT payment type. Guarantee means Pay at Check Out. Check the methods in guarantee or deposit or prepay.
		Guarantee         GuaranteeResponse           `json:"guarantee"`
		AdditionalDetails []QualifiedFreeTextResponse `json:"additionalDetails"`
	}

	// * the guarantee policy information applicable to the offer. It includes accepted payments
	GuaranteeResponse struct {
		Description      QualifiedFreeTextResponse `json:"description"`      // Specific type to convey a list of string for specific information type ( via qualifier) in specific character set, or language
		AcceptedPayments []AcceptedPaymentResponse `json:"acceptedPayments"` // Accepted Payment Methods and Card Types. Several Payment Methods and Card Types may be available.
	}

	AcceptedPaymentResponse struct {
		CreditCards []string      `json:"creditCards"` // example: VI .CA - MasterCard (warning - use it instead of MC/IK/EC/MD/XS) VI - Visa AX - American Express DC - Diners Club AU - Carte Aurore CG - Cofinoga DS - Discover GK - Lufthansa GK Card JC - Japanese Credit Bureau TC - Torch Club TP - Universal Air Travel Card BC - Bank Card DL - Delta MA - Maestro UP - China UnionPay
		Method      PaymentMethod `json:"method"`      // example: CREDIT_CARD. CREDIT_CARD (CC) - Payment Cards in creditCards are accepted AGENCY_ACCOUNT - Agency Account (Credit Line) is accepted. Agency is Charged at CheckOut TRAVEL_AGENT_ID - Agency IATA/ARC Number is accepted to Guarantee the booking CORPORATE_ID (COR-ID) - Corporate Account is accepted to Guarantee the booking HOTEL_GUEST_ID - Hotel Chain Rewards Card Number is accepted to Guarantee the booking CHECK - Checks are accepted MISC_CHARGE_ORDER - Miscellaneous Charge Order is accepted ADVANCE_DEPOSIT - Cash is accepted for Deposit/PrePay COMPANY_ADDRESS - Company Billing Address is accepted to Guarantee the booking
	}
)
