package offers

import (
	"regexp"
	"strconv"
)

// CapacitySource says where a capacity figure came from.
//
// It exists because the two sources are not equally trustworthy, and a caller
// deciding whether four people can book a room deserves to know which one
// answered.
type CapacitySource string

const (
	// CapacityStructured means Amadeus supplied maxPersonCapacity as a field.
	// Trust it.
	CapacityStructured CapacitySource = "STRUCTURED"
	// CapacityFromDescription means the figure was read out of the room's
	// free-text description, e.g. "Apartment with 1 bedroom for 4 persons".
	// It is a best-effort reading of prose a human wrote, so treat it as a
	// strong hint rather than a guarantee.
	CapacityFromDescription CapacitySource = "DESCRIPTION"
)

// MaxPersonCapacity returns the most people the room can hold, and how that was
// established.
//
// Amadeus does not reliably populate maxPersonCapacity on Hotel Search: across
// a captured sandbox search of 44 offers it was absent from every one, while
// the room description said "for 4 persons" in plain English. So this checks
// the structured field first and falls back to reading the description, and
// reports which via CapacitySource.
//
// It returns ok=false when neither source says anything. Do not infer capacity
// from Room.Beds - a room with two beds may sleep two, three or four - and do
// not infer it from Guests, which is the occupancy the price was quoted for
// rather than what the room holds.
//
// When the answer matters commercially, the authoritative check is to search
// with the occupancy you actually want: Amadeus filtering on Guests is more
// dependable than anything it publishes about capacity.
func (o Offer) MaxPersonCapacity() (people int, source CapacitySource, ok bool) {
	if o.RoomDetails != nil && o.RoomDetails.MaxOccupancy != nil {
		if total := o.RoomDetails.MaxOccupancy.Total; total > 0 {
			return total, CapacityStructured, true
		}
		// Some sources give the breakdown without the total.
		if occupancy := o.RoomDetails.MaxOccupancy; occupancy.Adults > 0 {
			return occupancy.Adults + occupancy.Children, CapacityStructured, true
		}
	}

	if o.StandardizedRoom != nil && o.StandardizedRoom.MaxOccupancy != nil {
		if total := o.StandardizedRoom.MaxOccupancy.Total; total > 0 {
			return total, CapacityStructured, true
		}
	}

	// Fall back to the prose, checking the richer description block first.
	for _, text := range []string{o.roomDetailsDescription(), o.roomDescription()} {
		if capacity, found := capacityFromText(text); found {
			return capacity, CapacityFromDescription, true
		}
	}

	return 0, "", false
}

func (o Offer) roomDescription() string {
	if o.Room.Description == nil {
		return ""
	}
	return o.Room.Description.Value
}

func (o Offer) roomDetailsDescription() string {
	if o.RoomDetails == nil {
		return ""
	}
	return o.RoomDetails.Description
}

// capacityPatterns match the ways Amadeus's room descriptions state an
// occupancy. They are ordered by how explicit they are.
//
// The person-noun is required, so "1 bedroom" and "2 beds" cannot be mistaken
// for a capacity - which matters, because "Apartment with 1 bedroom for 4
// persons" must yield 4 and never 1.
var capacityPatterns = []*regexp.Regexp{
	// "for 1 or 2 persons", "for 2 to 4 people", "for 4 persons"
	regexp.MustCompile(`(?i)\b(\d{1,2})\s*(?:or|to|-|–)\s*(\d{1,2})\s*(?:persons?|people|pax|adults?|guests?|occupants?)\b`),
	regexp.MustCompile(`(?i)\b(\d{1,2})\s*(?:persons?|people|pax|adults?|guests?|occupants?)\b`),
	// "sleeps 4", "sleeping 4"
	regexp.MustCompile(`(?i)\bsleep(?:s|ing)?\s+(?:up\s+to\s+)?(\d{1,2})\b`),
	// "max occupancy 3", "maximum occupancy: 3"
	regexp.MustCompile(`(?i)\bmax(?:imum)?\s+occupancy\s*:?\s*(\d{1,2})\b`),
}

// maxPlausibleCapacity guards against reading a room number, a floor or a year
// as an occupancy. No hotel room in this inventory sleeps more than this.
const maxPlausibleCapacity = 20

// capacityFromText reads an occupancy out of a room description.
//
// A range takes its upper bound: "Room for 1 or 2 persons" holds two. Where a
// description states several figures, the largest plausible one wins, since
// descriptions tend to qualify upward ("for 2 persons, extra bed for 1 more").
func capacityFromText(text string) (int, bool) {
	if text == "" {
		return 0, false
	}

	best := 0
	for _, pattern := range capacityPatterns {
		for _, match := range pattern.FindAllStringSubmatch(text, -1) {
			for _, group := range match[1:] {
				if group == "" {
					continue
				}
				value, err := strconv.Atoi(group)
				if err != nil || value <= 0 || value > maxPlausibleCapacity {
					continue
				}
				if value > best {
					best = value
				}
			}
		}
		if best > 0 {
			// An earlier, more explicit pattern matched; do not let a looser
			// one override it.
			break
		}
	}

	return best, best > 0
}

// Accommodates reports whether the room can hold the given number of people.
//
// It returns certain=false when capacity is unknown or was only inferred from
// prose. Booking a party of four into a room the SDK merely guessed holds four
// is how a family arrives to find one bed, so check certain before relying on
// a true.
func (o Offer) Accommodates(people int) (fits bool, certain bool) {
	capacity, source, ok := o.MaxPersonCapacity()
	if !ok {
		return false, false
	}
	return capacity >= people, source == CapacityStructured
}
