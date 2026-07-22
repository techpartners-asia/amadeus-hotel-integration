package responseHotelOffersDTO

import (
	"math"
	"sort"
	"strconv"
)

// The Hotel Offers API returns a flat list of offers per hotel. An "offer" is a
// bookable rate, not a room: one room (room.type, e.g. "C3S") appears in many
// offers that differ by rate code, board type, cancellation policy and price.
// A room-picker UI wants the inverse shape - each room once, with its rate
// options underneath - which is what the types and helpers below produce.
//
// Grouping is only meaningful when the request set bestRateOnly=false. With the
// default (true) the API returns a single cheapest offer per hotel, so every
// hotel collapses to one group of one offer.

type (
	// GroupedHotelOffers is one hotel with its offers regrouped by room. It is
	// the element type returned by GroupByRoom over a List result.
	GroupedHotelOffers struct {
		Hotel     HotelResponse    `json:"hotel"`
		Available bool             `json:"available"`
		Rooms     []RoomOfferGroup `json:"rooms"`
	}

	// RoomOfferGroup collects every offer for a single room (keyed by
	// room.type) and pre-computes what a room card renders before it is
	// expanded: the headline price and a representative room description.
	RoomOfferGroup struct {
		// RoomType is the room.type code shared by every offer in the group.
		// Offers whose room.type is empty are collected under "".
		RoomType string `json:"roomType"`
		// Room is copied from the cheapest offer, so callers can render the
		// description, bed type and category without reaching into Offers.
		Room RoomResponse `json:"room"`
		// PriceFrom is the price.total of the cheapest offer, verbatim (a
		// string, matching the API). Empty only when no offer had a price.
		PriceFrom string `json:"priceFrom"`
		// PriceFromCurrency is the currency of the cheapest offer. Prices in a
		// single response share a currency, so this labels PriceFrom.
		PriceFromCurrency string `json:"priceFromCurrency"`
		// Cheapest points at the lowest-priced offer in Offers (Offers[0]).
		Cheapest *OfferResponse `json:"-"`
		// Offers are the room's rate options, sorted by price ascending with
		// the offer id as a tie-breaker for deterministic output.
		Offers []OfferResponse `json:"offers"`
	}
)

// GroupByRoom regroups one hotel's offers by room.type. The result is stable:
// groups are ordered by their cheapest price (then room type), and each group's
// offers are ordered by price (then id). An offer with an unparseable or missing
// price.total sorts last and is never chosen as the cheapest unless it is the
// only offer in its group, so a bad price cannot mask a real one.
func (o OffersResponse) GroupByRoom() []RoomOfferGroup {
	// Preserve first-seen order of room types before sorting, so grouping is
	// deterministic even when every price is missing.
	order := make([]string, 0)
	byRoom := make(map[string][]OfferResponse)
	for _, offer := range o.Offers {
		key := offer.Room.Type
		if _, seen := byRoom[key]; !seen {
			order = append(order, key)
		}
		byRoom[key] = append(byRoom[key], offer)
	}

	groups := make([]RoomOfferGroup, 0, len(order))
	for _, key := range order {
		offers := byRoom[key]
		sort.SliceStable(offers, func(i, j int) bool {
			pi, pj := offerPrice(offers[i]), offerPrice(offers[j])
			if pi != pj {
				return pi < pj
			}
			return offers[i].ID < offers[j].ID
		})

		g := RoomOfferGroup{
			RoomType: key,
			Offers:   offers,
		}
		if len(offers) > 0 {
			cheapest := &g.Offers[0]
			g.Cheapest = cheapest
			g.Room = cheapest.Room
			g.PriceFrom = cheapest.Price.Total
			g.PriceFromCurrency = cheapest.Price.Currency
		}
		groups = append(groups, g)
	}

	sort.SliceStable(groups, func(i, j int) bool {
		pi, pj := groupPrice(groups[i]), groupPrice(groups[j])
		if pi != pj {
			return pi < pj
		}
		return groups[i].RoomType < groups[j].RoomType
	})

	return groups
}

// GroupByRoom regroups every hotel in a List result by room, preserving the
// order in which the hotels were returned.
func GroupByRoom(hotels []OffersResponse) []GroupedHotelOffers {
	out := make([]GroupedHotelOffers, 0, len(hotels))
	for _, h := range hotels {
		out = append(out, GroupedHotelOffers{
			Hotel:     h.Hotel,
			Available: h.Available,
			Rooms:     h.GroupByRoom(),
		})
	}
	return out
}

// offerPrice parses an offer's price.total into a comparable value. A missing
// or malformed price sorts last (+Inf) rather than 0, so it never displaces a
// genuine cheapest offer.
func offerPrice(o OfferResponse) float64 {
	f, err := strconv.ParseFloat(o.Price.Total, 64)
	if err != nil {
		return math.Inf(1)
	}
	return f
}

// groupPrice is the sort key for a room group: its cheapest offer's price, or
// +Inf when the group has no priced offer.
func groupPrice(g RoomOfferGroup) float64 {
	if g.Cheapest == nil {
		return math.Inf(1)
	}
	return offerPrice(*g.Cheapest)
}
