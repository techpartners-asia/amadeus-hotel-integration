package offers_test

import (
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/offers"
)

// offerWithDescription builds an offer carrying only a room description, which
// is the shape Hotel Search actually returns: 44 of 44 captured offers had a
// description and none had maxPersonCapacity.
func offerWithDescription(text string) offers.Offer {
	return offers.Offer{
		Room: offers.Room{Description: &media.Text{Value: text, Lang: "EN"}},
	}
}

func TestCapacityIsReadFromRealDescriptions(t *testing.T) {
	// Every string here is either taken verbatim from the captured sandbox
	// search or is a close variant of one.
	cases := []struct {
		description string
		want        int
	}{
		{"FLEX - RO B2C-Room only\nApartment with 1 bedroom for 4 persons", 4},
		{"ADVANCE SAVER-Room only\nRoom for 1 or 2 persons", 2},
		{"Mercure Breaks-Early / RO-Room only\nStandard Room 1 dble bed GB Test GB", 0},
		{"Studio for 2 people", 2},
		{"Suite sleeps 6", 6},
		{"Family room, maximum occupancy 5", 5},
		{"Twin room for 2 adults", 2},
		{"Apartment for 2 to 4 guests", 4},
		{"Deluxe room, max occupancy: 3", 3},
	}

	for _, c := range cases {
		got, source, ok := offerWithDescription(c.description).MaxPersonCapacity()

		if c.want == 0 {
			if ok {
				t.Errorf("%.40q: read a capacity of %d where none is stated", c.description, got)
			}
			continue
		}
		if !ok {
			t.Errorf("%.40q: no capacity found, want %d", c.description, c.want)
			continue
		}
		if got != c.want {
			t.Errorf("%.40q: capacity = %d, want %d", c.description, got, c.want)
		}
		if source != offers.CapacityFromDescription {
			t.Errorf("%.40q: source = %q, want DESCRIPTION", c.description, source)
		}
	}
}

func TestBedroomCountIsNotMistakenForCapacity(t *testing.T) {
	// "Apartment with 1 bedroom for 4 persons" must yield 4, never 1. The
	// person-noun requirement in the patterns is what prevents it, and getting
	// this wrong would understate every apartment in the inventory.
	got, _, ok := offerWithDescription("Apartment with 1 bedroom for 4 persons").MaxPersonCapacity()
	if !ok || got != 4 {
		t.Errorf("capacity = %d (found=%v), want 4", got, ok)
	}

	for _, description := range []string{
		"Room with 2 beds",
		"Apartment with 3 bedrooms",
		"Suite on floor 12",
		"Room 204, renovated 2019",
	} {
		if capacity, _, ok := offerWithDescription(description).MaxPersonCapacity(); ok {
			t.Errorf("%q: read %d as a capacity, want none", description, capacity)
		}
	}
}

func TestRangeTakesItsUpperBound(t *testing.T) {
	// "for 1 or 2 persons" holds two, not one.
	got, _, ok := offerWithDescription("Room for 1 or 2 persons").MaxPersonCapacity()
	if !ok || got != 2 {
		t.Errorf("capacity = %d, want the upper bound 2", got)
	}
}

func TestImplausibleFiguresAreIgnored(t *testing.T) {
	// A room number or a year must not become an occupancy.
	for _, description := range []string{
		"Room for 200 persons",
		"Conference space for 5000 people",
	} {
		if capacity, _, ok := offerWithDescription(description).MaxPersonCapacity(); ok {
			t.Errorf("%q: accepted an implausible capacity of %d", description, capacity)
		}
	}
}

func TestStructuredCapacityWinsOverProse(t *testing.T) {
	// When Amadeus does supply the field, it is authoritative and the prose is
	// not consulted.
	offer := offerWithDescription("Room for 2 persons")
	offer.RoomDetails = &offers.RoomDetails{
		Description:  "Room for 2 persons",
		MaxOccupancy: &offers.Occupancy{Total: 3},
	}

	got, source, ok := offer.MaxPersonCapacity()
	if !ok || got != 3 {
		t.Errorf("capacity = %d, want the structured 3", got)
	}
	if source != offers.CapacityStructured {
		t.Errorf("source = %q, want STRUCTURED", source)
	}
}

func TestOccupancyBreakdownWithoutATotal(t *testing.T) {
	offer := offers.Offer{
		RoomDetails: &offers.RoomDetails{
			MaxOccupancy: &offers.Occupancy{Adults: 2, Children: 2},
		},
	}

	got, source, ok := offer.MaxPersonCapacity()
	if !ok || got != 4 {
		t.Errorf("capacity = %d, want adults+children = 4", got)
	}
	if source != offers.CapacityStructured {
		t.Errorf("source = %q, want STRUCTURED", source)
	}
}

func TestUnknownCapacityIsReportedAsUnknown(t *testing.T) {
	// The honest answer when Amadeus says nothing. Silently returning 0, or
	// guessing from Room.Beds, would be worse than admitting ignorance.
	var offer offers.Offer

	if capacity, _, ok := offer.MaxPersonCapacity(); ok {
		t.Errorf("capacity = %d for an empty offer, want unknown", capacity)
	}

	fits, certain := offer.Accommodates(2)
	if fits || certain {
		t.Errorf("Accommodates(2) = %v/%v on an empty offer, want false/false", fits, certain)
	}
}

func TestAccommodatesFlagsAnInferredAnswer(t *testing.T) {
	// A party booked into a room whose capacity was merely guessed is how a
	// family arrives to find one bed. The caller must be able to tell.
	prose := offerWithDescription("Apartment with 1 bedroom for 4 persons")

	fits, certain := prose.Accommodates(4)
	if !fits {
		t.Error("a room described as holding 4 should accommodate 4")
	}
	if certain {
		t.Error("a capacity read from prose must not report as certain")
	}

	if fits, _ := prose.Accommodates(6); fits {
		t.Error("a room described as holding 4 should not accommodate 6")
	}

	structured := offers.Offer{
		RoomDetails: &offers.RoomDetails{MaxOccupancy: &offers.Occupancy{Total: 4}},
	}
	if fits, certain := structured.Accommodates(4); !fits || !certain {
		t.Errorf("structured capacity: %v/%v, want true/true", fits, certain)
	}
}

// The live sandbox supplies no structured capacity at all, so this documents
// what callers can actually expect from a real search today.
func TestCapacityAvailabilityInTheCapturedFixture(t *testing.T) {
	var structured, fromProse, unknown int
	for _, offer := range allOffers(t) {
		_, source, ok := offer.MaxPersonCapacity()
		switch {
		case !ok:
			unknown++
		case source == offers.CapacityStructured:
			structured++
		default:
			fromProse++
		}
	}

	total := structured + fromProse + unknown
	t.Logf("of %d captured offers: %d structured, %d from prose, %d unknown",
		total, structured, fromProse, unknown)

	if structured+fromProse == 0 {
		t.Error("no offer yielded a capacity by any route; the parser may have regressed")
	}
}
