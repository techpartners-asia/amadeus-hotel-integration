package sharedResponseDTO

type (
	AmenityResponse struct {
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
		PricingMethod PricingMethod `json:"pricingMethod,omitempty"`
		// Quantity - how many counts are available for this amenity type.
		Quantity int `json:"quantity,omitempty"`
		// Media - list of media (images/videos) associated with the amenity.
		Media []MediaResponse `json:"media,omitempty"`
	}

	// AmenityProvider contains the source of the amenity content.
	AmenityProvider struct {
		// Name - name of the amenity content source. Example: "ATPCO".
		Name string `json:"name,omitempty"`
	}

	// AmenityPrice contains price information for an amenity.
	AmenityPrice struct {
		// Base - base price of the amenity.
		Base string `json:"base,omitempty"`
		// Currency - currency code applied to the price.
		Currency string `json:"currency,omitempty"`
		// Markups - markups applied to the amenity price.
		Markups []MarkupResponse `json:"markups,omitempty"`
		// SellingTotal - selling total = total + margins + markup + totalFees - discounts.
		SellingTotal string `json:"sellingTotal,omitempty"`
		// Total - total = base + totalTaxes.
		Total string `json:"total,omitempty"`
	}

	// Markup represents a markup applied by a stakeholder (travel agent, merchant mode, etc.).
	MarkupResponse struct {
		// Amount - the monetary value of the markup as a string with decimal.
		Amount string `json:"amount,omitempty"`
	}
)
