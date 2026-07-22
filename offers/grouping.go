package offers

import (
	"sort"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/money"
)

// Grouping offers by room is domain logic, not presentation. Amadeus returns a
// flat list of offers per hotel, where an offer is a bookable rate rather than
// a room: one room (Room.Type, e.g. "C3S") appears in many offers differing by
// rate code, board type, cancellation policy and price. A room-picker shows the
// inverse - each room once, with its rate options underneath - and computing
// that inversion correctly, including which offer is genuinely cheapest, is the
// kind of decision that belongs in the domain rather than in each caller's UI
// layer.
//
// Grouping is only meaningful when the search set BestRateOnly to false. With
// Amadeus's default of true the API returns one cheapest offer per hotel, so
// every hotel collapses to a single group of a single offer.

// GroupedHotel is one hotel with its offers regrouped by room.
type GroupedHotel struct {
	// Hotel is the property the offers belong to.
	Hotel Hotel
	// Available reports whether the property has bookable inventory.
	Available bool
	// Rooms are the room groups, cheapest first.
	Rooms []RoomGroup
}

// RoomGroup collects every offer for one room and precomputes what a room card
// shows before it is expanded.
type RoomGroup struct {
	// RoomType is the Room.Type shared by every offer in the group. Offers
	// whose room type is empty are collected under "".
	RoomType string
	// Room is copied from the cheapest offer, so a caller can render the
	// description, bed type and category without reaching into Offers.
	Room Room
	// PriceFrom is the cheapest offer's total, and is the "from" price a room
	// card displays. It is zero only when no offer in the group had a price.
	PriceFrom money.Money
	// Cheapest is the lowest-priced offer, which is Offers[0]. It is nil only
	// for an empty group, which cannot occur through GroupByRoom.
	Cheapest *Offer
	// Offers are the room's rate options, cheapest first, with the offer ID as
	// a tie-breaker so the order is stable across identical responses.
	Offers []Offer
}

// GroupByRoom regroups one hotel's offers by room type.
//
// The result is deterministic: groups are ordered by their cheapest price then
// room type, and each group's offers by price then ID. An offer with no usable
// price sorts last and is never chosen as the cheapest unless it is the only
// offer in its group, so a missing price cannot mask a real one.
func (h HotelOffers) GroupByRoom() []RoomGroup {
	// First-seen order is recorded so grouping stays deterministic even when
	// every offer lacks a price and the sort has nothing to order on.
	order := make([]string, 0, len(h.Offers))
	byRoom := make(map[string][]Offer, len(h.Offers))

	for _, offer := range h.Offers {
		key := offer.Room.Type
		if _, seen := byRoom[key]; !seen {
			order = append(order, key)
		}
		byRoom[key] = append(byRoom[key], offer)
	}

	groups := make([]RoomGroup, 0, len(order))
	for _, key := range order {
		offers := byRoom[key]
		sort.SliceStable(offers, func(i, j int) bool {
			if cmp := comparePrices(offers[i].Price.Total, offers[j].Price.Total); cmp != 0 {
				return cmp < 0
			}
			return offers[i].ID < offers[j].ID
		})

		group := RoomGroup{RoomType: key, Offers: offers}
		if len(offers) > 0 {
			group.Cheapest = &group.Offers[0]
			group.Room = group.Cheapest.Room
			group.PriceFrom = group.Cheapest.Price.Total
		}
		groups = append(groups, group)
	}

	sort.SliceStable(groups, func(i, j int) bool {
		if cmp := comparePrices(groups[i].PriceFrom, groups[j].PriceFrom); cmp != 0 {
			return cmp < 0
		}
		return groups[i].RoomType < groups[j].RoomType
	})

	return groups
}

// GroupByRoom regroups every hotel in a search result, preserving the order in
// which Amadeus returned the hotels.
func GroupByRoom(hotels []HotelOffers) []GroupedHotel {
	out := make([]GroupedHotel, 0, len(hotels))
	for _, hotel := range hotels {
		out = append(out, GroupedHotel{
			Hotel:     hotel.Hotel,
			Available: hotel.Available,
			Rooms:     hotel.GroupByRoom(),
		})
	}
	return out
}

// comparePrices orders two prices, sorting a missing price last so it never
// displaces a genuine cheapest offer.
//
// Prices in different currencies cannot be ordered; those compare equal, which
// leaves the ID tie-breaker to produce a stable result rather than an arbitrary
// one. A single Amadeus response quotes one currency per hotel, so this is the
// degenerate case rather than the common one.
func comparePrices(a, b money.Money) int {
	aMissing, bMissing := a.Amount().IsZero(), b.Amount().IsZero()
	switch {
	case aMissing && bMissing:
		return 0
	case aMissing:
		return 1
	case bMissing:
		return -1
	}

	cmp, err := a.Compare(b)
	if err != nil {
		return 0
	}
	return cmp
}
