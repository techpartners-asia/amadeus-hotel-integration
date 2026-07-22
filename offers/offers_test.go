package offers_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeustest"
	"github.com/techpartners-asia/amadeus-hotel-integration/offers"
)

const searchPath = "/v3/shopping/hotel-offers"

func newService(t *testing.T) (offers.Service, *amadeustest.Server) {
	t.Helper()
	server := amadeustest.New(t)
	server.Fixture(t, http.MethodGet, searchPath, "search")
	return offers.NewService(server.Client()), server
}

// search runs the fixture-backed search and returns the first hotel.
func searchFirst(t *testing.T) offers.HotelOffers {
	t.Helper()
	service, _ := newService(t)

	results, err := service.Search(context.Background(), offers.SearchQuery{
		HotelIDs: []string{"HLPAR266"},
	})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("got %d hotels, want 2", len(results))
	}
	return results[0]
}

func TestHotelIsMapped(t *testing.T) {
	hotel := searchFirst(t).Hotel

	if hotel.ID != "HLPAR266" || hotel.Name != "HILTON PARIS OPERA" {
		t.Errorf("hotel = %+v", hotel)
	}
	if hotel.Rating != codes.Rating4 {
		t.Errorf("Rating = %q, want a typed codes.Rating", hotel.Rating)
	}
	if hotel.ChainCode != "HL" || hotel.BrandCode != "HI" || hotel.DupeID != "700034421" {
		t.Errorf("codes = %q/%q/%q", hotel.ChainCode, hotel.BrandCode, hotel.DupeID)
	}
	if hotel.Position == nil || hotel.Position.Latitude != 48.87626 {
		t.Errorf("Position = %+v", hotel.Position)
	}
	if hotel.Contact == nil || hotel.Contact.Email != "reservations@example.invalid" {
		t.Errorf("Contact = %+v", hotel.Contact)
	}
	if hotel.Address == nil || hotel.Address.StateCode != "IDF" {
		t.Errorf("Address = %+v", hotel.Address)
	}
	if len(hotel.AmenityCodes) != 3 {
		t.Errorf("AmenityCodes = %v", hotel.AmenityCodes)
	}
	if hotel.TermsAndConditions == "" {
		t.Error("TermsAndConditions was dropped")
	}
}

func TestPricesBecomeMoneyNotStrings(t *testing.T) {
	// The headline improvement of the restructure: a caller never parses a
	// price string, and cannot add two currencies together by accident.
	offer := searchFirst(t).Offers[0]

	if got := offer.Price.Total.String(); got != "600 EUR" {
		t.Errorf("Total = %q, want %q", got, "600 EUR")
	}
	if got := offer.Price.Base.String(); got != "540 EUR" {
		t.Errorf("Base = %q, want %q", got, "540 EUR")
	}
	if got := offer.Price.SellingTotal.String(); got != "612 EUR" {
		t.Errorf("SellingTotal = %q, want %q", got, "612 EUR")
	}
	if offer.Price.Currency != "EUR" {
		t.Errorf("Currency = %q", offer.Price.Currency)
	}
}

func TestTaxesCarryTheirCollectionPoint(t *testing.T) {
	// Whether a tax is already in the base, and whether the guest pays it at
	// the property, changes what you display. Both must survive mapping.
	price := searchFirst(t).Offers[0].Price

	if len(price.Taxes) != 2 {
		t.Fatalf("got %d taxes, want 2", len(price.Taxes))
	}

	vat, tourism := price.Taxes[0], price.Taxes[1]
	if !vat.Included {
		t.Error("VAT should be marked as included in the base")
	}
	if vat.CollectedAtProperty() {
		t.Error("VAT is collected at booking, not at the property")
	}
	if tourism.Included {
		t.Error("the tourist tax is not included in the base")
	}
	if !tourism.CollectedAtProperty() {
		t.Error("the tourist tax is collected at the property")
	}
	if tourism.Applicable == nil || tourism.Applicable.Start.String() != "2026-08-10" {
		t.Errorf("Applicable = %+v", tourism.Applicable)
	}

	// Only the non-included tax should be added to the base.
	total, err := price.TaxesTotal()
	if err != nil {
		t.Fatalf("TaxesTotal: %v", err)
	}
	if total.String() != "12 EUR" {
		t.Errorf("TaxesTotal = %q, want 12 EUR (the included VAT must not be counted)", total)
	}

	atProperty, err := price.PayableAtProperty()
	if err != nil {
		t.Fatalf("PayableAtProperty: %v", err)
	}
	if atProperty.String() != "12 EUR" {
		t.Errorf("PayableAtProperty = %q, want 12 EUR", atProperty)
	}
}

func TestQuotedBooleansAreNormalised(t *testing.T) {
	// Amadeus sends isLoyaltyRate and isOptional as the strings "true"/"false"
	// rather than JSON booleans. Left as strings they are a trap: the string
	// "false" is truthy to any caller doing a non-empty check.
	hotel := searchFirst(t)

	if !hotel.Offers[0].IsLoyaltyRate {
		t.Error(`isLoyaltyRate "true" should map to true`)
	}
	if hotel.Offers[1].IsLoyaltyRate {
		t.Error(`isLoyaltyRate "false" should map to false`)
	}

	guarantee := hotel.Offers[0].Policies.Guarantee
	if guarantee == nil || len(guarantee.AcceptedPayments.CardPolicies) != 1 {
		t.Fatalf("guarantee = %+v", guarantee)
	}
	inputs := guarantee.AcceptedPayments.CardPolicies[0].Inputs
	if len(inputs) != 2 {
		t.Fatalf("got %d inputs, want 2", len(inputs))
	}
	if inputs[0].Optional {
		t.Errorf(`%q has isOptional "false" and must not be optional`, inputs[0].Label)
	}
	if !inputs[1].Optional {
		t.Errorf(`%q has isOptional "true" and must be optional`, inputs[1].Label)
	}
}

func TestDatesBecomeCalendarDates(t *testing.T) {
	offer := searchFirst(t).Offers[0]

	if got := offer.Stay.CheckIn.String(); got != "2026-08-10" {
		t.Errorf("CheckIn = %q", got)
	}
	if got := offer.Stay.Nights(); got != 3 {
		t.Errorf("Nights = %d, want 3", got)
	}
}

func TestGuestsCountChildrenByAge(t *testing.T) {
	guests := searchFirst(t).Offers[0].Guests

	if guests.Adults != 2 {
		t.Errorf("Adults = %d", guests.Adults)
	}
	if guests.Children() != 2 || guests.Total() != 4 {
		t.Errorf("Children = %d, Total = %d, want 2 and 4", guests.Children(), guests.Total())
	}
	if len(guests.ChildAges) != 2 || guests.ChildAges[0] != 7 {
		t.Errorf("ChildAges = %v", guests.ChildAges)
	}
}

func TestCancellationPoliciesAreMergedAndDeduplicated(t *testing.T) {
	// Amadeus sends the same policy in both "cancellation" and
	// "cancellations". A caller reading both gets it twice; a caller reading
	// one gets nothing when the other was populated instead.
	policies := searchFirst(t).Offers[0].Policies

	if len(policies.Cancellation) != 2 {
		t.Fatalf("got %d cancellation policies, want 2 (the duplicate merged away): %+v",
			len(policies.Cancellation), policies.Cancellation)
	}

	if policies.Cancellation[0].PolicyType != "CANCELLATION" {
		t.Errorf("first policy = %+v", policies.Cancellation[0])
	}
	noShow := policies.Cancellation[1]
	if noShow.PolicyType != "NO_SHOW" || noShow.NumberOfNights != 1 {
		t.Errorf("no-show policy = %+v", noShow)
	}
	if noShow.IsFree() {
		t.Error("a policy charging one night is not free cancellation")
	}
}

func TestRefundabilityIsReportedConservatively(t *testing.T) {
	hotel := searchFirst(t)

	refundable, certain := hotel.Offers[0].Policies.IsRefundable()
	if !refundable || !certain {
		t.Errorf("flexible offer: refundable=%v certain=%v, want true/true", refundable, certain)
	}

	refundable, certain = hotel.Offers[1].Policies.IsRefundable()
	if refundable || !certain {
		t.Errorf("advance-purchase offer: refundable=%v certain=%v, want false/true", refundable, certain)
	}

	// An offer with no refundability block and no cancellation terms must
	// report uncertainty, not a cheerful "refundable".
	refundable, certain = hotel.Offers[3].Policies.IsRefundable()
	if refundable || certain {
		t.Errorf("offer with no policy: refundable=%v certain=%v, want false/false", refundable, certain)
	}
}

func TestFreeCancellationDeadline(t *testing.T) {
	policies := searchFirst(t).Offers[0].Policies

	deadline, ok := policies.FreeCancellationUntil()
	if !ok {
		t.Fatal("expected a free-cancellation deadline")
	}
	if got := deadline.Format("2006-01-02T15:04:05"); got != "2026-08-08T18:00:00" {
		t.Errorf("deadline = %q", got)
	}

	// The no-refund offer has no free-cancellation deadline at all.
	if _, ok := searchFirst(t).Offers[1].Policies.FreeCancellationUntil(); ok {
		t.Error("a non-refundable offer must not report a free-cancellation deadline")
	}
}

func TestPriceVariationsAreMapped(t *testing.T) {
	variations := searchFirst(t).Offers[0].Price.Variations

	if got := variations.Average.Total.String(); got != "200 EUR" {
		t.Errorf("average total = %q", got)
	}
	if len(variations.Changes) != 2 {
		t.Fatalf("got %d changes, want 2", len(variations.Changes))
	}
	if got := variations.Changes[0].Start.String(); got != "2026-08-10" {
		t.Errorf("first change start = %q", got)
	}
	if got := variations.Changes[1].Total.String(); got != "220 EUR" {
		t.Errorf("second change total = %q", got)
	}
}

func TestCommissionsFromBothWireShapes(t *testing.T) {
	// Amadeus sends commission as a legacy nested array on some sources and a
	// flat one on others. Both fold into one list.
	offer := searchFirst(t).Offers[0]

	if offer.Commission == nil || offer.Commission.Amount.String() != "18.75 EUR" {
		t.Errorf("offer commission = %+v", offer.Commission)
	}
	if len(offer.Price.Commissions) != 1 {
		t.Fatalf("got %d price commissions, want 1", len(offer.Price.Commissions))
	}
	if got := offer.Price.Commissions[0]; got.Percentage != 10 || got.DecimalPlaces != 2 {
		t.Errorf("commission = %+v", got)
	}
}

func TestExtrasAreMapped(t *testing.T) {
	extras := searchFirst(t).Offers[0].Extras

	if len(extras) != 2 {
		t.Fatalf("got %d extras, want 2", len(extras))
	}
	if extras[0].IsChargeable {
		t.Error("breakfast is complimentary in the fixture")
	}
	if !extras[1].IsChargeable || extras[1].Price == nil {
		t.Errorf("parking = %+v", extras[1])
	}
	if got := extras[1].Price.Total.String(); got != "42 EUR" {
		t.Errorf("parking price = %q", got)
	}
}

func TestPerNightSplitsExactlyWithRemainder(t *testing.T) {
	price := searchFirst(t).Offers[0].Price

	perNight, remainder, ok := price.PerNight(3)
	if !ok {
		t.Fatal("PerNight failed")
	}
	if perNight.String() != "200 EUR" || !remainder.Amount().IsZero() {
		t.Errorf("600/3 = %s remainder %s, want 200 EUR remainder 0", perNight, remainder)
	}

	// A total that does not divide evenly must report the leftover rather than
	// rounding it away silently.
	perNight, remainder, ok = searchFirst(t).Offers[2].Price.PerNight(7)
	if !ok {
		t.Fatal("PerNight failed")
	}
	recombined := perNight.Mul(7)
	sum, err := recombined.Add(remainder)
	if err != nil {
		t.Fatalf("recombining: %v", err)
	}
	if sum.String() != "360 EUR" {
		t.Errorf("part*7 + remainder = %s, want the original 360 EUR", sum)
	}
}

func TestCheapestSkipsPricelessOffers(t *testing.T) {
	hotel := searchFirst(t)

	cheapest, ok := hotel.Cheapest()
	if !ok {
		t.Fatal("expected a cheapest offer")
	}
	if cheapest.ID != "OFFER_STANDARD" {
		t.Errorf("cheapest = %s (%s), want OFFER_STANDARD at 360 EUR",
			cheapest.ID, cheapest.Price.Total)
	}
}

func TestGroupByRoomInvertsTheOfferList(t *testing.T) {
	// Amadeus returns offers flat; a room picker needs each room once with its
	// rates underneath.
	rooms := searchFirst(t).GroupByRoom()

	if len(rooms) != 2 {
		t.Fatalf("got %d room groups, want 2 (DLX and STD)", len(rooms))
	}

	// Groups are ordered by their cheapest price: STD at 360 beats DLX at 495.
	if rooms[0].RoomType != "STD" || rooms[1].RoomType != "DLX" {
		t.Errorf("group order = %q, %q; want STD then DLX", rooms[0].RoomType, rooms[1].RoomType)
	}
	if got := rooms[0].PriceFrom.String(); got != "360 EUR" {
		t.Errorf("STD PriceFrom = %q, want 360 EUR", got)
	}
	if rooms[0].Room.Category != "STANDARD_ROOM" {
		t.Errorf("group room was not copied from the cheapest offer: %+v", rooms[0].Room)
	}

	// The DLX group holds both deluxe rates, cheapest first.
	deluxe := rooms[1]
	if len(deluxe.Offers) != 2 {
		t.Fatalf("DLX has %d offers, want 2", len(deluxe.Offers))
	}
	if deluxe.Offers[0].ID != "OFFER_DELUXE_SAVER" {
		t.Errorf("DLX cheapest = %s, want the 495 EUR saver", deluxe.Offers[0].ID)
	}
	if deluxe.Cheapest == nil || deluxe.Cheapest.ID != deluxe.Offers[0].ID {
		t.Error("Cheapest should point at Offers[0]")
	}
}

func TestPricelessOfferSortsLastAndNeverWins(t *testing.T) {
	// A missing price must not read as zero and become the cheapest.
	rooms := searchFirst(t).GroupByRoom()

	standard := rooms[0]
	if len(standard.Offers) != 2 {
		t.Fatalf("STD has %d offers, want 2", len(standard.Offers))
	}
	if standard.Offers[0].ID != "OFFER_STANDARD" {
		t.Errorf("STD cheapest = %s, want the priced offer", standard.Offers[0].ID)
	}
	if standard.Offers[1].ID != "OFFER_NO_PRICE" {
		t.Errorf("the priceless offer should sort last, got order %s, %s",
			standard.Offers[0].ID, standard.Offers[1].ID)
	}
}

func TestGroupingIsStableAcrossRuns(t *testing.T) {
	first := searchFirst(t).GroupByRoom()
	second := searchFirst(t).GroupByRoom()

	if len(first) != len(second) {
		t.Fatal("grouping produced different group counts")
	}
	for i := range first {
		if first[i].RoomType != second[i].RoomType {
			t.Errorf("group %d: %q then %q", i, first[i].RoomType, second[i].RoomType)
		}
		for j := range first[i].Offers {
			if first[i].Offers[j].ID != second[i].Offers[j].ID {
				t.Errorf("group %q offer %d: %s then %s",
					first[i].RoomType, j, first[i].Offers[j].ID, second[i].Offers[j].ID)
			}
		}
	}
}

func TestSoldOutHotelIsReturnedWithoutOffers(t *testing.T) {
	service, _ := newService(t)
	results, err := service.Search(context.Background(), offers.SearchQuery{HotelIDs: []string{"XXPAR999"}})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}

	soldOut := results[1]
	if soldOut.Available {
		t.Error("the sold-out property should report Available=false")
	}
	if len(soldOut.Offers) != 0 {
		t.Errorf("sold-out property has %d offers", len(soldOut.Offers))
	}
	if _, ok := soldOut.Cheapest(); ok {
		t.Error("a hotel with no offers has no cheapest offer")
	}
}

func TestSearchQueryParameters(t *testing.T) {
	service, server := newService(t)

	_, err := service.Search(context.Background(), offers.SearchQuery{
		HotelIDs:           []string{"HLPAR266", "MCPARC12"},
		Stay:               offers.NewStay(datetime.MustParseDate("2026-08-10"), datetime.MustParseDate("2026-08-13")),
		Guests:             offers.Guests{Adults: 2, ChildAges: []int{7, 12}},
		Rooms:              2,
		BoardType:          codes.BoardTypeBreakfast,
		RateCodes:          []codes.RateCode{codes.RateCodePublic},
		PaymentPolicy:      codes.PaymentPolicyGuarantee,
		Currency:           "EUR",
		PriceRange:         "100-400",
		CountryOfResidence: "FR",
		BestRateOnly:       codes.Ptr(false),
		Lang:               "FR",
	})
	if err != nil {
		t.Fatalf("Search: %v", err)
	}

	query := server.LastRequest(t).Query
	want := map[string]string{
		"hotelIds":           "HLPAR266,MCPARC12",
		"checkInDate":        "2026-08-10",
		"checkOutDate":       "2026-08-13",
		"adults":             "2",
		"childAges":          "7,12",
		"roomQuantity":       "2",
		"boardType":          "BREAKFAST",
		"rateCodes":          "PRO",
		"paymentPolicy":      "GUARANTEE",
		"currency":           "EUR",
		"priceRange":         "100-400",
		"countryOfResidence": "FR",
		"bestRateOnly":       "false",
		"lang":               "FR",
	}
	for key, expected := range want {
		if got := query.Get(key); got != expected {
			t.Errorf("query[%s] = %q, want %q", key, got, expected)
		}
	}
}

func TestUnsetParametersAreOmitted(t *testing.T) {
	service, server := newService(t)

	if _, err := service.Search(context.Background(), offers.SearchQuery{
		HotelIDs: []string{"HLPAR266"},
	}); err != nil {
		t.Fatalf("Search: %v", err)
	}

	query := server.LastRequest(t).Query
	for _, key := range []string{"adults", "currency", "boardType", "roomQuantity", "bestRateOnly"} {
		if _, present := query[key]; present {
			t.Errorf("unset %q was sent as %q", key, query.Get(key))
		}
	}
}

func TestSearchValidation(t *testing.T) {
	service, server := newService(t)

	cases := []struct {
		name  string
		query offers.SearchQuery
	}{
		{"no hotel IDs", offers.SearchQuery{}},
		{"backwards stay", offers.SearchQuery{
			HotelIDs: []string{"HLPAR266"},
			Stay: offers.NewStay(
				datetime.MustParseDate("2026-08-13"),
				datetime.MustParseDate("2026-08-10")),
		}},
		{"zero-night stay", offers.SearchQuery{
			HotelIDs: []string{"HLPAR266"},
			Stay: offers.NewStay(
				datetime.MustParseDate("2026-08-10"),
				datetime.MustParseDate("2026-08-10")),
		}},
		{"too many adults", offers.SearchQuery{
			HotelIDs: []string{"HLPAR266"},
			Guests:   offers.Guests{Adults: 20},
		}},
		{"adult age given as a child", offers.SearchQuery{
			HotelIDs: []string{"HLPAR266"},
			Guests:   offers.Guests{Adults: 1, ChildAges: []int{45}},
		}},
		{"price range without a currency", offers.SearchQuery{
			HotelIDs:   []string{"HLPAR266"},
			PriceRange: "100-400",
		}},
		{"unknown board type", offers.SearchQuery{
			HotelIDs:  []string{"HLPAR266"},
			BoardType: "BRUNCH",
		}},
		{"malformed rate code", offers.SearchQuery{
			HotelIDs:  []string{"HLPAR266"},
			RateCodes: []codes.RateCode{"TOOLONG"},
		}},
	}

	before := len(server.Requests())
	for _, c := range cases {
		if _, err := service.Search(context.Background(), c.query); !errors.Is(err, apierr.ErrValidation) {
			t.Errorf("%s: err = %v, want ErrValidation", c.name, err)
		}
	}
	if after := len(server.Requests()); after != before {
		t.Errorf("%d invalid searches reached the network", after-before)
	}
}

func TestGetReturnsOfferWithItsHotel(t *testing.T) {
	server := amadeustest.New(t)
	server.Fixture(t, http.MethodGet, searchPath+"/OFFER_DELUXE_FLEX", "offer-by-id")
	service := offers.NewService(server.Client())

	detail, err := service.Get(context.Background(), offers.GetQuery{OfferID: "OFFER_DELUXE_FLEX"})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if detail.Offer.ID != "OFFER_DELUXE_FLEX" {
		t.Errorf("offer = %s", detail.Offer.ID)
	}
	// The hotel must come back too: a price with no property cannot be shown.
	if detail.Hotel.ID != "HLPAR266" {
		t.Errorf("hotel = %+v", detail.Hotel)
	}
	if !detail.Available {
		t.Error("Available was dropped")
	}
}

func TestGetOnAnEmptyResponseReportsNotFound(t *testing.T) {
	// An expired offer ID can return 200 with nothing in it. Returning a nil
	// offer and a nil error would hand the caller a panic.
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, searchPath+"/EXPIRED", http.StatusOK, `{"data":{}}`)
	service := offers.NewService(server.Client())

	detail, err := service.Get(context.Background(), offers.GetQuery{OfferID: "EXPIRED"})
	if detail != nil {
		t.Errorf("detail = %+v, want nil", detail)
	}
	if !errors.Is(err, apierr.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestGetValidatesOfferID(t *testing.T) {
	service, _ := newService(t)
	if _, err := service.Get(context.Background(), offers.GetQuery{}); !errors.Is(err, apierr.ErrValidation) {
		t.Errorf("err = %v, want ErrValidation", err)
	}
}

func TestExpiredOfferSurfacesAsNotFound(t *testing.T) {
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, searchPath+"/STALE", http.StatusNotFound,
		`{"errors":[{"status":404,"code":1797,"title":"NOT FOUND","detail":"offer no longer available"}]}`)
	service := offers.NewService(server.Client())

	_, err := service.Get(context.Background(), offers.GetQuery{OfferID: "STALE"})
	if !errors.Is(err, apierr.ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestMalformedPriceDoesNotDiscardTheOffer(t *testing.T) {
	// One bad decimal must not cost the caller the whole response.
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, searchPath, http.StatusOK, `{
	  "data": [{
	    "hotel": {"hotelId":"HLPAR266","name":"TEST"},
	    "available": true,
	    "offers": [{
	      "id":"BAD_PRICE",
	      "checkInDate":"2026-08-10",
	      "checkOutDate":"2026-08-13",
	      "room":{"type":"STD"},
	      "price":{"currency":"EUR","total":"not-a-number","base":"100.00"},
	      "guests":{"adults":2}
	    }]
	  }]
	}`)
	service := offers.NewService(server.Client())

	results, err := service.Search(context.Background(), offers.SearchQuery{HotelIDs: []string{"HLPAR266"}})
	if err != nil {
		t.Fatalf("a malformed price should not fail the search: %v", err)
	}

	offer := results[0].Offers[0]
	if offer.ID != "BAD_PRICE" {
		t.Fatalf("offer was dropped: %+v", results)
	}
	if !offer.Price.Total.Amount().IsZero() {
		t.Errorf("unparseable total = %s, want zero", offer.Price.Total)
	}
	if offer.Price.Base.String() != "100 EUR" {
		t.Errorf("the readable base should survive, got %s", offer.Price.Base)
	}
	if offer.Stay.Nights() != 3 {
		t.Error("the rest of the offer should be intact")
	}
}
