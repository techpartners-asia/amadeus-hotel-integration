package sharedResponseDTO

type PricingTimeWindow string

const (
	PricingTimeWindowHourly     PricingTimeWindow = "HOURLY"
	PricingTimeWindowDaily      PricingTimeWindow = "DAILY"
	PricingTimeWindowMonthly    PricingTimeWindow = "MONTHLY"
	PricingTimeWindowWeekend    PricingTimeWindow = "WEEKEND"
	PricingTimeWindowWeekly     PricingTimeWindow = "WEEKLY"
	PricingTimeWindowFullPeriod PricingTimeWindow = "FULL_PERIOD"
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
	PriceResponse struct {
		// Base - base price of the amenity.
		Base string `json:"base,omitempty"`
		// Currency - currency code applied to the price.
		Currency CurrencyResponse `json:"currency,omitempty"`
		// Markups - markups applied to the amenity price.
		Markups []MarkupResponse `json:"markups,omitempty"`
		// SellingTotal - selling total = total + margins + markup + totalFees - discounts.
		SellingTotal string `json:"sellingTotal,omitempty"`
		// Total - total = base + totalTaxes.
		Total string `json:"total,omitempty"`
	}

	CurrencyResponse struct {
		Code string `json:"code"` // ISO currency code (http://www.iso.org/iso/home/standards/currency_codes.htm). For miles the code associated is MIL.
		Name string `json:"name"` // Indicates the name of the currency.
	}
)
