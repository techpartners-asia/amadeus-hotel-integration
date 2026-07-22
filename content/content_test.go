package content_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/apierr"
	"github.com/techpartners-asia/amadeus-hotel-integration/codes"
	"github.com/techpartners-asia/amadeus-hotel-integration/content"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeustest"
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
)

const contentPath = "/v3/reference-data/locations/by-hotel"

func newService(t *testing.T) (content.Service, *amadeustest.Server) {
	t.Helper()
	server := amadeustest.New(t)
	server.Fixture(t, http.MethodGet, contentPath, "hotel")
	return content.NewService(server.Client()), server
}

func fetch(t *testing.T) *content.Hotel {
	t.Helper()
	service, _ := newService(t)

	hotel, err := service.Get(context.Background(), content.Query{HotelID: "HLPAR266"})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	return hotel
}

func TestHotelIdentityIsMapped(t *testing.T) {
	hotel := fetch(t)

	if hotel.ID != "HLPAR266" || hotel.Name != "Hilton Paris Opera" {
		t.Errorf("hotel = %s / %s", hotel.ID, hotel.Name)
	}
	if hotel.Rating != codes.Rating4 {
		t.Errorf("Rating = %q, want a typed rating", hotel.Rating)
	}
	if hotel.ChainName != "Hilton" || hotel.BrandName != "Hilton Hotels & Resorts" {
		t.Errorf("chain/brand = %q / %q", hotel.ChainName, hotel.BrandName)
	}
	if hotel.Status != "OPEN" || hotel.DefaultLanguage != "FR" {
		t.Errorf("status/language = %q / %q", hotel.Status, hotel.DefaultLanguage)
	}
	if hotel.Description == nil || hotel.Description.Lang != "EN" {
		t.Errorf("description = %+v", hotel.Description)
	}
	if len(hotel.Categories) != 2 || len(hotel.Segments) != 2 {
		t.Errorf("categories = %v, segments = %v", hotel.Categories, hotel.Segments)
	}
	if len(hotel.BusinessIdentifiers) != 1 || hotel.BusinessIdentifiers[0].Name != "VAT" {
		t.Errorf("business identifiers = %+v", hotel.BusinessIdentifiers)
	}
}

func TestPositionIsExposedThroughLocation(t *testing.T) {
	hotel := fetch(t)

	position, ok := hotel.Position()
	if !ok {
		t.Fatal("expected coordinates")
	}
	if position.Latitude != 48.87626 || position.Longitude != 2.32603 {
		t.Errorf("position = %v", position)
	}
	if hotel.Location.IATACode != "PAR" || hotel.Location.SubType != "CITY" {
		t.Errorf("location = %+v", hotel.Location)
	}
}

func TestPhoneNumberIsAssembledFromItsParts(t *testing.T) {
	// Amadeus splits a number across countryCallingCode, areaCode and number,
	// and frequently leaves the bare number unusable on its own.
	hotel := fetch(t)

	if len(hotel.Contacts) != 1 {
		t.Fatalf("got %d contacts", len(hotel.Contacts))
	}
	contact := hotel.Contacts[0]
	if len(contact.Phones) != 1 || contact.Phones[0].Number != "+33140088844" {
		t.Errorf("phone = %+v, want the parts joined into a dialable number", contact.Phones)
	}
	if len(contact.Emails) != 1 || contact.Emails[0] != "reservations@example.invalid" {
		t.Errorf("emails = %v", contact.Emails)
	}
	if contact.AddresseeName != "Reservations Desk" {
		t.Errorf("addressee = %q, want the name parts joined without the empty prefix", contact.AddresseeName)
	}
	if contact.Address == nil || contact.Address.CityName != "Paris" {
		t.Errorf("address = %+v", contact.Address)
	}
	if contact.Website != "https://example.invalid/hilton-paris-opera" {
		t.Errorf("website = %q", contact.Website)
	}
}

func TestMediaScalesLetCallersAvoidDownloadingOriginals(t *testing.T) {
	hotel := fetch(t)

	photo, ok := hotel.PrimaryPhoto()
	if !ok {
		t.Fatal("expected a photograph")
	}
	if photo.Alt == "" {
		t.Error("the screen-reader description was dropped; it is the only accessible text Amadeus sends")
	}
	if len(photo.Scales) != 2 {
		t.Fatalf("got %d scales", len(photo.Scales))
	}

	// Asking for a thumbnail should not hand back the full-size original.
	if got := photo.Best(400); got != "https://example.invalid/exterior-320.jpg" {
		t.Errorf("Best(400) = %q, want the 320px rendition", got)
	}
	if got := photo.Best(2000); got != "https://example.invalid/exterior-1024.jpg" {
		t.Errorf("Best(2000) = %q, want the largest that fits", got)
	}
	// Below every rendition, fall back to the smallest rather than the original.
	if got := photo.Best(100); got != "https://example.invalid/exterior-320.jpg" {
		t.Errorf("Best(100) = %q, want the smallest rendition", got)
	}
	if photo.Kind != media.KindImage {
		t.Errorf("Kind = %q, want IMAGE (Amadeus spells it \"Image\")", photo.Kind)
	}
}

func TestAmenitiesCarryTheirCharges(t *testing.T) {
	hotel := fetch(t)

	if len(hotel.Amenities) != 2 {
		t.Fatalf("got %d amenities", len(hotel.Amenities))
	}
	if !hotel.HasAmenity("SWIMMING_POOL") || hotel.HasAmenity("CASINO") {
		t.Error("HasAmenity gave the wrong answer")
	}

	parking := hotel.Amenities[1]
	if !parking.IsChargeable || parking.PricingMethod != "PER_ROOM_PER_NIGHT" {
		t.Errorf("parking = %+v", parking)
	}
	if parking.Provider != "ATPCO" {
		t.Errorf("provider = %q", parking.Provider)
	}
	if parking.Quantity != 40 {
		t.Errorf("quantity = %d", parking.Quantity)
	}
}

func TestRoomsAreMapped(t *testing.T) {
	hotel := fetch(t)

	if len(hotel.Rooms) != 1 {
		t.Fatalf("got %d rooms", len(hotel.Rooms))
	}
	room := hotel.Rooms[0]

	if room.Name == nil || room.Name.Value != "Deluxe King Room" {
		t.Errorf("name = %+v", room.Name)
	}
	if room.Category != "DELUXE" || room.Classification != "ROOM" {
		t.Errorf("category/classification = %q / %q", room.Category, room.Classification)
	}
	if room.Beds != 1 || room.BedType != "KING" {
		t.Errorf("beds = %d %s", room.Beds, room.BedType)
	}
	if !room.IsNonSmoking || room.StandardOccupancy != 2 {
		t.Errorf("room = %+v", room)
	}
	if room.MaxOccupancy == nil || room.MaxOccupancy.Total != 3 {
		t.Errorf("max occupancy = %+v", room.MaxOccupancy)
	}
	if room.SleepFurnishings == nil || room.SleepFurnishings.Cribs != 1 {
		t.Errorf("sleep furnishings = %+v", room.SleepFurnishings)
	}
	if room.Dimensions == nil || room.Dimensions.Area != 28 {
		t.Errorf("dimensions = %+v", room.Dimensions)
	}
	if len(room.Amenities) != 1 || len(room.Media) != 1 {
		t.Errorf("room amenities = %v, media = %v", room.Amenities, room.Media)
	}
	if room.ProviderReference == nil || room.ProviderReference.ID != "RT-DLX-01" {
		t.Errorf("provider reference = %+v", room.ProviderReference)
	}
}

func TestPoliciesAreMapped(t *testing.T) {
	hotel := fetch(t)

	policies := hotel.Policies
	if policies == nil {
		t.Fatal("policies were dropped")
	}

	if len(policies.CheckInOut) != 1 || policies.CheckInOut[0].CheckIn != "15:00" {
		t.Errorf("check-in/out = %+v", policies.CheckInOut)
	}
	if len(policies.Payment) != 1 {
		t.Fatalf("got %d payment policies", len(policies.Payment))
	}
	payment := policies.Payment[0]
	if payment.Type != "GUARANTEE" || payment.Guarantee == nil {
		t.Fatalf("payment = %+v", payment)
	}
	if len(payment.Guarantee.AcceptedCards) != 3 {
		t.Errorf("accepted cards = %v", payment.Guarantee.AcceptedCards)
	}
	if len(payment.Details) != 1 {
		t.Errorf("additional details = %v", payment.Details)
	}

	if len(policies.Cancellation) != 1 {
		t.Fatalf("got %d cancellation policies", len(policies.Cancellation))
	}
	if got := policies.Cancellation[0].Amount.Amount().String(); got != "150" {
		t.Errorf("cancellation fee = %q", got)
	}

	if len(policies.Pets) != 1 || policies.Pets[0].Code != "PETS_ALLOWED" {
		t.Errorf("pet policies = %+v", policies.Pets)
	}
	if len(policies.Commission) != 1 || policies.Commission[0].Percentage != "10" {
		t.Errorf("commission = %+v", policies.Commission)
	}
	if len(policies.StayRequirements) != 1 {
		t.Errorf("stay requirements = %v", policies.StayRequirements)
	}
}

func TestUnincludedTaxIsDistinguishable(t *testing.T) {
	// A tax not included in quoted rates is what makes the final bill bigger
	// than the booking total, so the flag has to survive.
	hotel := fetch(t)

	taxes := hotel.Policies.Tax
	if len(taxes) != 2 {
		t.Fatalf("got %d tax policies", len(taxes))
	}

	cityTax, vat := taxes[0], taxes[1]
	if cityTax.Included {
		t.Error("the city tax is not included in quoted rates")
	}
	if got := cityTax.Amount.String(); got != "4.4 EUR" {
		t.Errorf("city tax = %q", got)
	}
	if cityTax.Frequency != "PER_NIGHT" || cityTax.Mode != "PER_PERSON" {
		t.Errorf("city tax assessment = %q / %q", cityTax.Frequency, cityTax.Mode)
	}
	if !vat.Included || vat.Percentage != "10" {
		t.Errorf("VAT = %+v", vat)
	}
}

func TestGuestPolicyCarriesTheAgeRules(t *testing.T) {
	hotel := fetch(t)

	if len(hotel.Policies.Guest) != 1 {
		t.Fatalf("got %d guest policies", len(hotel.Policies.Guest))
	}
	guest := hotel.Policies.Guest[0]
	if guest.MinimumGuestAge != 18 {
		t.Errorf("minimum age = %d", guest.MinimumGuestAge)
	}
	if !guest.ChildStayFree || guest.ChildStayFreeCutoffAge != 12 {
		t.Errorf("child policy = %+v", guest)
	}
	if guest.MaxChildAgeForBedSharing != 12 {
		t.Errorf("bed sharing age = %d", guest.MaxChildAgeForBedSharing)
	}
}

func TestFacilitiesAreSummarised(t *testing.T) {
	hotel := fetch(t)

	facilities := hotel.Facilities
	if facilities == nil {
		t.Fatal("facilities were dropped")
	}
	if len(facilities.Amenities) != 1 {
		t.Errorf("facility amenities = %v", facilities.Amenities)
	}
	if facilities.MeetingRooms == nil || facilities.MeetingRooms.Count != 6 {
		t.Fatalf("meeting rooms = %+v", facilities.MeetingRooms)
	}
	if facilities.MeetingRooms.LargestCapacity != 220 {
		t.Errorf("largest capacity = %d", facilities.MeetingRooms.LargestCapacity)
	}
	if facilities.Restaurants == nil || facilities.Restaurants.Count != 2 {
		t.Fatalf("restaurants = %+v", facilities.Restaurants)
	}
	// FRENCH appears on both restaurants and must not be listed twice.
	if got := facilities.Restaurants.Cuisines; len(got) != 3 {
		t.Errorf("cuisines = %v, want 3 distinct (FRENCH deduplicated)", got)
	}
}

func TestNearestDistanceComparesAcrossUnits(t *testing.T) {
	// The fixture quotes 0.4 MILE and 0.5 KM for the same landmark. Comparing
	// the raw numbers would pick the mile as "nearer"; it is 644m against 500m.
	hotel := fetch(t)

	if len(hotel.NearbyLandmarks) != 1 {
		t.Fatalf("got %d landmarks", len(hotel.NearbyLandmarks))
	}
	landmark := hotel.NearbyLandmarks[0]
	if landmark.Name != "Opera Garnier" {
		t.Errorf("landmark = %q", landmark.Name)
	}
	if landmark.Distance == nil {
		t.Fatal("distance was dropped")
	}
	if landmark.Distance.Value != 0.5 || landmark.Distance.Unit != "KM" {
		t.Errorf("nearest distance = %s, want 0.5 KM (shorter than 0.4 MILE)", landmark.Distance)
	}
}

func TestAwardsAndPromotions(t *testing.T) {
	hotel := fetch(t)

	if len(hotel.Awards) != 1 || hotel.Awards[0].Name != "AAA Diamond" {
		t.Errorf("awards = %+v", hotel.Awards)
	}
	if len(hotel.Certifications) != 1 || hotel.Certifications[0].Name != "Green Key" {
		t.Errorf("certifications = %+v", hotel.Certifications)
	}
	if len(hotel.Promotions) != 1 {
		t.Fatalf("got %d promotions", len(hotel.Promotions))
	}
	promotion := hotel.Promotions[0]
	if promotion.Code != "S3P2" || promotion.TermsAndConditions == "" {
		t.Errorf("promotion = %+v", promotion)
	}
}

func TestPointsOfInterest(t *testing.T) {
	hotel := fetch(t)

	if len(hotel.PointsOfInterest) != 1 {
		t.Fatalf("got %d points of interest", len(hotel.PointsOfInterest))
	}
	poi := hotel.PointsOfInterest[0]
	if poi.Name != "Palais Garnier" || poi.CategoryCode != "SIGHTSEEING" {
		t.Errorf("point of interest = %+v", poi)
	}
	if poi.Location == nil || poi.Location.Position == nil {
		t.Errorf("location = %+v", poi.Location)
	}
	if poi.Distance == nil || poi.Distance.Value != 0.6 {
		t.Errorf("distance = %+v", poi.Distance)
	}
	if poi.Website == "" {
		t.Error("official website was dropped")
	}
}

func TestBuildingAndTimeZone(t *testing.T) {
	hotel := fetch(t)

	if hotel.Building == nil || hotel.Building.Floors != 7 || hotel.Building.TotalRooms != 268 {
		t.Errorf("building = %+v", hotel.Building)
	}
	if hotel.Building.YearBuilt != "1889" || hotel.Building.YearRenovated != "2015" {
		t.Errorf("building dates = %+v", hotel.Building)
	}
	if hotel.TimeZone == nil || hotel.TimeZone.Name != "Central European Time" {
		t.Errorf("timezone = %+v", hotel.TimeZone)
	}
	if !hotel.TimeZone.DaylightSaving {
		t.Error("Europe/Paris observes daylight saving")
	}
	if hotel.Altitude == nil || hotel.Altitude.Value != 35 {
		t.Errorf("altitude = %+v", hotel.Altitude)
	}
}

func TestViewDefaultsToFull(t *testing.T) {
	// Amadeus's own default returns the basic block alone, which is almost
	// never what a caller asking for content wants.
	service, server := newService(t)

	if _, err := service.Get(context.Background(), content.Query{HotelID: "HLPAR266"}); err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got := server.LastRequest(t).Query.Get("view"); got != "FULL" {
		t.Errorf("view = %q, want FULL by default", got)
	}
}

func TestQueryParameters(t *testing.T) {
	service, server := newService(t)

	_, err := service.Get(context.Background(), content.Query{
		HotelID: "HLPAR266",
		View:    codes.ContentViewLight,
		Fields:  []string{"hotel", "rooms"},
		Lang:    "FR",
	})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	query := server.LastRequest(t).Query
	if query.Get("hotelID") != "HLPAR266" {
		t.Errorf("hotelID = %q", query.Get("hotelID"))
	}
	if query.Get("view") != "LIGHT" {
		t.Errorf("view = %q, an explicit view must not be overridden", query.Get("view"))
	}
	if query.Get("fields") != "hotel,rooms" {
		t.Errorf("fields = %q", query.Get("fields"))
	}
	if query.Get("lang") != "FR" {
		t.Errorf("lang = %q", query.Get("lang"))
	}
}

func TestValidation(t *testing.T) {
	service, server := newService(t)

	cases := []content.Query{
		{},
		{HotelID: "TOOSHORT1234"},
		{HotelID: "HLPAR266", View: "MEDIUM"},
	}

	before := len(server.Requests())
	for _, query := range cases {
		if _, err := service.Get(context.Background(), query); !errors.Is(err, apierr.ErrValidation) {
			t.Errorf("%+v: err = %v, want ErrValidation", query, err)
		}
	}
	if len(server.Requests()) != before {
		t.Error("an invalid query reached the network")
	}
}

func TestSparseContentDoesNotProduceEmptyStructs(t *testing.T) {
	// Most properties publish far less than the schema allows. An absent block
	// must be nil, so a caller can tell "not published" from "published empty".
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, contentPath, http.StatusOK,
		`{"data":{"basic":{"hotelId":"XXPAR999","name":"SPARSE PROPERTY"}}}`)
	service := content.NewService(server.Client())

	hotel, err := service.Get(context.Background(), content.Query{HotelID: "XXPAR999"})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if hotel.Name != "SPARSE PROPERTY" {
		t.Errorf("name = %q", hotel.Name)
	}
	if hotel.Policies != nil {
		t.Errorf("Policies = %+v, want nil when none were published", hotel.Policies)
	}
	if hotel.Facilities != nil {
		t.Errorf("Facilities = %+v, want nil", hotel.Facilities)
	}
	if hotel.Location != nil {
		t.Errorf("Location = %+v, want nil", hotel.Location)
	}
	if hotel.Building != nil || hotel.TimeZone != nil || hotel.Altitude != nil {
		t.Error("absent blocks should stay nil")
	}
	if len(hotel.Rooms) != 0 || len(hotel.Contacts) != 0 {
		t.Error("absent lists should stay empty")
	}
	if _, ok := hotel.Position(); ok {
		t.Error("a property with no geoCode has no position")
	}
	if _, ok := hotel.PrimaryPhoto(); ok {
		t.Error("a property with no media has no primary photo")
	}
}

func TestNotFoundSurfacesTyped(t *testing.T) {
	server := amadeustest.New(t)
	server.JSON(http.MethodGet, contentPath, http.StatusNotFound,
		`{"errors":[{"status":404,"code":1797,"title":"NOT FOUND"}]}`)
	service := content.NewService(server.Client())

	_, err := service.Get(context.Background(), content.Query{HotelID: "XXPAR999"})
	if !errors.Is(err, apierr.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}
