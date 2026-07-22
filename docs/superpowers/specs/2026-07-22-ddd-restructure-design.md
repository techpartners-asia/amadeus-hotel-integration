# DDD Restructure — Amadeus Hotel SDK

**Date:** 2026-07-22
**Status:** Implemented

## Problem

The SDK works but its structure fights the reader. Concretely:

1. **Non-idiomatic package names.** `usecasesHotelOffers`, `requestHotelOffersDTO`,
   `sharedResponseDTO`. Every import needs an alias; every call site stutters.
2. **Global mutable singleton.** `amadeusIntegration.manager` is package-level state
   set by `Init()`. Two clients (test + prod, or two tenants) cannot coexist in one
   process, and nothing is testable without network access.
3. **No `context.Context`.** No cancellation, no deadlines, no tracing propagation.
4. **Environment hardcoded.** `test.travel.api.amadeus.com` is compiled into
   `constants/url.go`; going to production means editing the source.
5. **Duplicated transport.** The same result/error block appears in 11 methods.
6. **Inconsistent shape.** `offers/usecases` (plural) vs `booking/usecase`; only
   `list` splits requests into sub-packages; grouping logic sits in a DTO package.
7. **Untestable tests.** Everything in `tests/` hits the live sandbox, with
   credentials committed to the repository.
8. **Wire format is the public API.** Callers consume Amadeus's JSON shapes
   directly: prices as `string`, dates as `string`, enums as `string` with the
   permitted values listed only in a doc comment.

## Decisions

| Decision | Choice |
|---|---|
| DDD depth | Pragmatic — bounded contexts, value objects, ACL, ports. No aggregates/repositories where they earn nothing. |
| Compatibility | Break freely. Pre-1.0, no consumers to protect. |
| ACL scope | **Full translation.** Every wire DTO gets a domain counterpart and a mapper. |
| Testing | Golden fixtures captured from live responses + offline unit tests. Live suite kept behind a build tag. |
| Layout | Context-per-package, flat. Wire DTOs under `internal/`. |
| Module path | **Unchanged**: `github.com/techpartners-asia/amadeus-hotel-integration`. |

## Bounded contexts

| Context | Package | Responsibility | Language |
|---|---|---|---|
| Offers | `offers/` | Priced, bookable availability for a stay | Offer, Stay, Guests, Rate, BoardType, CancellationPolicy |
| Booking | `booking/` | The commercial transaction and its lifecycle | Order, Booking, Guest, Payment, Reference |
| Content | `content/` | Descriptive, non-priced hotel information | Hotel, Room, Facility, Policy, Media, Award |
| Inventory | `inventory/` | Which hotels exist, and where | Hotel, Location, SearchArea |

`modules/list` becomes `inventory`: "list" names a mechanism, not a domain concept,
and it collides with the verb on every other service.

### Context relationships

```
        codes/  money/  geo/          <- shared kernel (value objects only)
              ^    ^    ^
   +----------+----+----+-------------+
   |                                  |
inventory --> content            offers --> booking
 (find a       (describe        (price it)  (buy it)
  hotel)        it)
        -------- Customer/Supplier -------->
```

Contexts never import each other's services, only published domain types. `booking`
accepting an `offers.OfferID` is a value crossing a boundary, not a dependency on
offers' behaviour.

## Shared kernel

Deliberately small. Value objects only; no behaviour beyond validation and formatting.

- `money/` — `Money{Amount, Currency}`. Prices are `string` on the wire today, so
  arithmetic is currently the caller's problem.
- `geo/` — `Coordinates`, `Distance{Value, Unit}`.
- `codes/` — the current `searchcriteria` types (Amenity, Rating, BoardType, ...),
  already value objects in all but name.

`Stay{CheckIn, CheckOut}` lives in `offers/`, not the kernel: booking receives a stay
that offers produced. That is Customer/Supplier, not shared ownership.

## Anti-corruption layer

Every Amadeus wire DTO moves to `internal/amadeus/dto/`. Go forbids importing
`internal/` from outside the module, so **a caller cannot depend on the wire format** —
the ACL is a build error, not a convention. Each context's `mapper.go` is the only
code that sees both sides.

Translation rules applied uniformly:

| Wire | Domain |
|---|---|
| `"120.50"` + `"EUR"` | `money.Money` (decimal, not float) |
| `"2026-07-22"` | `datetime.Date` (date-only type, no spurious timezone) |
| `"2026-07-22T14:00:00"` | `time.Time` |
| `string` enum + doc comment | typed constant in `codes/` |
| `latitude`/`longitude` pair | `geo.Coordinates` |
| absent vs. zero | pointer on the wire, `Optional`/zero-value semantics documented in domain |

Unmapped wire fields are a test failure, not a silent drop — see Testing.

## Package layout

```
sdk.go                  Client: New(cfg), context wiring
config.go               Environment, credentials, timeouts, *http.Client

offers/     offer.go  search.go  service.go  mapper.go
booking/    order.go  service.go  mapper.go
content/    hotel.go  service.go  mapper.go
inventory/  hotel.go  service.go  mapper.go

money/  geo/  codes/          shared kernel

internal/
  amadeus/
    client.go  auth.go  errors.go   transport, token manager (no globals)
    dto/                            wire structs, unreachable by callers
  testdata/*.json                   golden fixtures
```

## Transport

`internal/amadeus.Client` replaces the package-level singleton:

- Constructed per SDK instance, holding its own token manager. Multiple clients coexist.
- Every method takes `ctx context.Context` as its first parameter.
- One generic `do[T]` helper absorbs the result/error block duplicated 11 times.
- The token manager keeps its mutex and refresh window; it becomes a field, not a global.
- `Environment` (`Test`/`Production`) selects the host at construction. No source edits.

## Error handling

`APIError` keeps its shape (status, structured Amadeus errors, raw body) and moves to
the root package so callers need no internal import. Additions:

- Sentinel errors for the common cases: `ErrUnauthorized`, `ErrNotFound`,
  `ErrRateLimited`, `ErrInvalidRequest`, matched via `errors.Is`.
- `APIError` wraps them, so `errors.Is(err, sdk.ErrNotFound)` works alongside
  `errors.As(err, &apiErr)` for full detail.
- Validation failures return before any network call, as `ValidationError` naming the
  offending field.

## Testing

Three layers:

1. **Fixture capture** (`-tags capture`, run on demand): hits the sandbox and writes
   `internal/testdata/*.json` for every endpoint. This is the existing probe tests
   promoted from "print to stdout" to "write to disk".
2. **Mapper tests** (default, offline): decode each fixture, map to domain, assert
   field-by-field. Plus a **coverage assertion** that walks the fixture JSON and fails
   on any field the domain model drops — this is what makes full translation safe.
3. **Live suite** (`-tags live`): the current end-to-end tests, kept but no longer the
   default path, and reading credentials from the environment only.

Committed sandbox credentials are removed; the live suite skips when the environment
does not supply them.

## Out of scope

- Retry/backoff policy (worth doing; a separate change).
- Pagination helpers/iterators (the `meta.links.next` plumbing exists but is unused).
- Any second supplier behind the same domain interface.

---

## Implementation notes

Where the built thing differs from the design above, and why.

### Wire DTOs are split per context, not one `dto` package

The design put every wire structure in `internal/amadeus/dto`. That does not
work: Amadeus reuses `Guest`, `Payment`, `Address` and `RoomAssociation` across
schemas with **different shapes**, and booking uses different shapes again for
requests and responses (a request `Guest` carries a caller-assigned `tid`, a
response `Guest` an Amadeus-assigned `id`).

The first attempt renamed the collisions. That silently mangled struct *field*
names as well as type names, producing `BookingGuestReference.BookingGuestReference`
and a `RoomInformation.MediaResponse` field. Renaming was the wrong tool.

Final layout:

```
internal/amadeus/dto/            shared: media, amenity, markup, dimensions
internal/amadeus/dto/offersdto/  Hotel Search
internal/amadeus/dto/contentdto/ Hotel Content
internal/amadeus/dto/bookingreq/ Hotel Booking requests
internal/amadeus/dto/bookingres/ Hotel Booking responses
```

Every struct stays faithful to its own schema; nothing is renamed to accommodate
another.

### Domain types follow the wire, not the plan

Two domain types were drafted from the design's assumptions and then corrected
against what Amadeus actually sends:

- **`content.PetPolicy`** was drafted with an `Allowed bool` and a maximum
  weight. The schema has neither: it sends a code, prose and a pricing method.
  The domain now exposes what arrives.
- **`booking.Price`** had a `PayableAtProperty` mirroring the offers one. The
  booking tax block carries no `collectionPoint`, so the method could only ever
  have returned zero — which reads as "nothing due on arrival" rather than "not
  known". It was removed rather than shipped; `TaxesTotal` remains.

The general rule this settled: do not add a method the wire cannot answer.

### `money` gained division

`Price.PerNight` needs to divide a stay total. `money.Money.Split` and
`Amount.DivMod` return the undivided remainder rather than rounding, because
deciding which night absorbs the odd cent is a presentation choice the SDK has
no basis to make silently. There is deliberately no plain `Div`.

### `offers.Service` collided with the domain's `Service`

Amadeus calls a chargeable add-on a "service", which collided with the service
interface. The domain type is `offers.Extra` - the standard hospitality term,
and unambiguous.

### Fixtures are captured from the live sandbox

`internal/capture` has been run; every `testdata/*.json` outside booking is a
real Amadeus response. The live suite passes end-to-end.

Running it against real data was worth it: three mapper bugs survived the
hand-built fixtures, because a fixture written from the same misunderstanding
as the mapper agrees with it.

1. **The property description was dropped entirely.** Amadeus leaves
   `basic.description` null and ships prose as entries in the *media* array
   carrying text and a tag but no image. The mapper read the empty field and
   put all seven prose blocks into `Media`. Now `splitMedia` separates them,
   `Hotel.Description` picks the long description, and `Hotel.Descriptions`
   keeps the rest keyed by tag.

2. **`PrimaryPhoto()` returned a text block.** Amadeus never populates `type`
   on a media entry, so the kind inference always yielded `""`, which the
   "image or unset" check accepted. It returned a description with no URL - an
   `<img>` with an empty `src`. Kind is now inferred from whether the entry
   actually carries imagery, and `media.Asset.IsVisual` makes the check
   available to callers.

3. **`Room.Dimensions` was never nil.** The wire sends dimensions as an
   embedded struct, and the mapper took its address unconditionally, so
   `if room.Dimensions != nil` was always true and told callers nothing.

Two API behaviours the capture also documented:

- **Hotel Search fails the whole request when any one property is sold out**,
  returning a 400 naming the offending codes rather than omitting them. An
  application searching twenty hotels where two have no rooms gets nothing, not
  eighteen results. `captureOffers` narrows its set on this error, and
  `TestNoAvailabilityFailsTheWholeSearch` pins the behaviour.
- **Live rate codes are not in the documented list.** The sandbox returns
  `1KD`, `EAM`, `D20`. `RateCode.IsValid` checks shape rather than membership,
  which is what lets these round-trip.

The context tests were rewritten to assert on invariants rather than fixture
positions, so a re-capture with different hotels does not fail them spuriously.
Cases the sandbox does not exercise skip explicitly rather than passing
vacuously.

### Not done

- Retry/backoff beyond the single 401-refresh retry.
- Pagination helpers. `meta.links.next` is decoded but unused.
