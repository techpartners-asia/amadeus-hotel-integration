# Search: available hotels with offers

**Date:** 2026-07-21
**Status:** approved, not yet implemented

## Problem

`List.HotelListByCityCode` returns hotels and nothing else. Callers building a search
results page then need price and availability, and today must orchestrate that
themselves. The obvious request — "make the list call return content and offers too" —
does not survive contact with the API's cost profile.

## Measurements

Taken against the live Amadeus test environment on 2026-07-21, `cityCode=PAR`. These
drove every decision below and are recorded so nobody has to re-derive them.

| Operation | Measured |
| --------- | -------- |
| `HotelListByCityCode("PAR")` | **3,508 hotels in 2.7s** |
| `Content.GetByID` | **1.18s per hotel, no batch endpoint** |
| Content extrapolated to all 3,508 hotels | **~69 minutes** |
| `Offers.List`, 10 ids | 1.8s |
| `Offers.List`, 20 ids | 11.8s |
| `Offers.List`, 50 ids | 12.0s |
| `Offers.List`, 51 ids | rejected: `477 INVALID FORMAT - Exceeding max items for: hotelIds` |

Two facts established by probing, both previously undocumented:

1. **The `hotelIds` cap is exactly 50.** Found by binary search: 50 accepted, 51
   rejected. Encoded as `MaxHotelIDsPerOffersRequest`.
2. **A bad hotel id does not poison its batch.** `hotels[0]` alone returned
   `3237 PROPERTY CODE NOT FOUND IN SYSTEM`, but the same id inside a batch of 10
   returned results normally. Batches tolerate unknown properties, so we do not need to
   pre-filter ids.

A third observation shaped the design more than either: in the Paris sandbox only 3-4
hotels out of 50 had any availability. Filtering on availability removes roughly 90% of
candidates before any per-hotel work.

## Decisions

**Content is not fetched.** `Offers.List` already returns hotel name, star rating,
chain/brand code, coordinates, address, contact and amenity codes alongside pricing.
That covers a search results row. Content's unique payload — media, long descriptions,
room detail, facilities, policies, awards, points of interest — is detail-page material,
so it stays a lazy per-hotel call at 1.18s, which is fine for one hotel. This single
decision removes the ~69-minute cost entirely.

**Availability filtering is the point, not a side effect.** The method returns only
hotels that came back with offers. Hotels with no inventory for the dates are dropped.

**Work is bounded by the caller.** `MaxHotels` defaults to 100 (2 batches). A full city
sweep is ~71 batches and belongs in a background job, not a request path.

**Partial batch failures are reported, not hidden.** For a booking product, silently
omitting a hotel that has inventory is worse than admitting the check was incomplete.

## Placement

The capability composes `List` and `Offers`, so it belongs to neither module. New
`modules/search/usecase`, constructed from both, exposed as `client.Search` — consistent
with how every other capability hangs off the SDK.

```go
client.Search.AvailableHotels(req)
```

## API

```go
type AvailableHotelsRequest struct {
    // Which hotels to consider - forwarded to List.
    CityCode    string
    Amenities   []searchcriteria.Amenity
    Ratings     []searchcriteria.Rating
    HotelSource *searchcriteria.HotelSource

    // The stay - forwarded to Offers.
    CheckInDate  string
    CheckOutDate string
    Adults       int
    RoomQuantity int
    BoardType    searchcriteria.BoardType
    Currency     string

    // Bounding.
    MaxHotels   int // default 100
    Concurrency int // default 4
}

// AvailabilityResult reports what was found and what could not be checked. Hotels
// alone would make "no availability" and "we failed to look" indistinguishable.
type AvailabilityResult struct {
    Hotels        []responseOffers.OffersResponse // only hotels with offers, in List order
    HotelsChecked int                             // ids actually submitted
    FailedBatches int                             // batches that errored
    Errors        []error                         // one per failed batch
}

func (s *searchUsecase) AvailableHotels(req AvailableHotelsRequest) (*AvailabilityResult, error)
```

`OffersResponse` is reused rather than wrapped: it already carries `{Hotel, Available,
Offers}`, which is exactly the result shape.

The returned `error` is non-nil only when the search could not proceed at all (the List
call failed, or every batch failed). Otherwise partial results come back with
`FailedBatches > 0`.

## Algorithm

1. `List.HotelListByCityCode` with the property filters.
2. Truncate to `MaxHotels`.
3. Chunk ids into groups of `MaxHotelIDsPerOffersRequest` (50).
4. Run `Offers.List` per chunk, concurrency-bounded by `Concurrency`.
5. Collect results, preserving List's ordering; record per-batch errors.
6. Return hotels that have offers.

Ordering is preserved by writing each batch's results into a pre-sized slot rather than
appending as they complete, so output order does not depend on completion order.

## Testing

- **Chunking** is a pure function (`chunkIDs`) unit-tested with no network: exact
  multiples of 50, remainders, empty input, single id, and the 50/51 boundary.
- **Ordering** is verified with a faked offers port returning out of order.
- **Partial failure** is verified with a faked port that fails one batch: results from
  the surviving batch are returned and `FailedBatches == 1`.
- **Live test** mirrors the existing `search_offers_test.go` style, skipped in `-short`.

Tests go under `tests/`, matching the SDK's existing layout.

## Known gaps

- **Latency.** A 100-hotel search is ~12s. Parallelism keeps wall-clock near the slowest
  batch but cannot go below it. If this ends up behind a user-facing HTTP request, 12s
  likely exceeds the timeout budget and the capability should be an async job with
  polling instead. Not resolved here.
- **Rate limits.** `Concurrency: 4` is a conservative guess. Amadeus' actual limit for
  these Enterprise endpoints has not been measured. Worth probing the way the batch cap
  was, then tuning the default.
- **Sandbox availability is not production availability.** The 90% reduction figure comes
  from a test environment with thin inventory. Production will return far more hotels
  with offers, which makes `MaxHotels` matter more, not less.
