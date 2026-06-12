# Amadeus Hotel SDK — Full Integration Design

**Date:** 2026-06-12
**Status:** Approved (user: "cook it")

## Goal

Bring the Go SDK (`github.com/techpartners-asia/amadeus-hotel-integration`) to full
coverage of the six provided Amadeus swagger specifications, with full-fidelity DTOs
generated from each swagger's schemas.

## Swagger coverage map

| Swagger | Endpoint | Status before | Action |
|---|---|---|---|
| Hotel Search v3.5 | `GET /shopping/hotel-offers`, `/{offerId}` | ✅ covered | reconcile DTOs |
| Hotel Booking v2.3 | `POST /booking/hotel-orders` | ✅ covered | keep |
| Hotel Booking Retrieve v2.1 | `GET /booking/hotel-orders/{hotelOrderId}` | ❌ only by-reference | **add** |
| Hotel Booking Manage v2.2 | `POST .../hotel-bookings/{id}/cancel` | ❌ missing | **add** |
| Hotel Booking Manage v2.2 | `PATCH .../hotel-bookings/{id}` | ❌ missing | **add** |
| Hotel Booking Manage v2.2 | `DELETE .../hotel-bookings/{id}` | ❌ missing | **add** |
| Hotel List v1.2 | `GET .../hotels/by-hotels` | ❌ missing | **add** |
| Hotel List v1.2 | `by-city`, `by-geocode` | ✅ covered | keep |
| Hotel Content v3.1 | `GET /reference-data/locations/by-hotel` | ⚠️ wrong path/param | **align** |

## Architecture changes

### 1. Per-module HTTP client (fixes shared-client bug)

Today every use-case mutates the single global `amadeusIntegration.Client` via
`SetBaseURL()`. Because they share one pointer, the last constructor wins, so
`Content` and `Booking` (which never set a URL) run against whatever base URL the
List/Offers constructor set last. Fix:

```go
// integrations/amadeus
func NewClient(baseURL string) *resty.Client
```

Each use-case gets its own client cloned from the authenticated base (auth header +
JSON accept), with its own base URL. The global mutable `Client` is removed.

### 2. Token manager with refresh

Auth currently runs once in `Init` and the token silently expires (~30 min). Add a
token manager that caches the token + `expires_in` and refreshes on demand before each
request (or when within a safety window). `Init` returns an `error` instead of
`log.Fatalf`.

### 3. Per-module host constants

```
test.api.amadeus.com/v3   → Offers (search), Content
test.api.amadeus.com/v1   → List
test.travel.api.amadeus.com/v2 → Booking (create + retrieve + manage)
```

**Decision (resolves the open question):** the entire Booking module moves to the
`travel` host, because the Manage/Retrieve swaggers and the Create swagger all specify
that host, and manage calls reference orders created there. Hosts are constants so
test↔prod is a one-line change.

## New / changed API surface

### Booking (`modules/booking`)
- `GetByID(hotelOrderId string) (*HotelOrder, error)` — `GET /booking/hotel-orders/{id}`
- `Cancel(hotelOrderId, hotelBookingId string) (*HotelOrder, error)` — `POST .../cancel`
- `Modify(hotelOrderId, hotelBookingId string, req UpdateHotelBookingRequest) (*HotelOrderUpdateResponse, error)` — `PATCH`
- `Delete(hotelOrderId, hotelBookingId string) (*DeleteBookingResult, error)` — `DELETE`

Response reuse: Retrieve and Cancel return the existing `HotelOrder` type (already
modeled). Modify returns `{data:{type,id}, included:HotelOrder, warnings}`. Delete
returns `{included:{cancellationNumber}, warnings}`.

New request DTO `UpdateHotelBookingRequest` from `UpdateHotelBooking` schema
(roomAssociation / hotelOffer.product / payment.paymentCard with 3DS + address).

### List (`modules/list`)
- `HotelListByHotelIds(req HotelListByHotelsRequest) ([]GeneralInfoResponse, error)` —
  `GET .../by-hotels?hotelIds=...`. Reuses existing `GeneralInfoResponse`
  (same `HotelSearchResponse` schema as by-city).

### Content (`modules/content`)
- Align `GetByID` to `GET /reference-data/locations/by-hotel` with query `hotelID`
  (capital ID) plus optional `fields`, `lang`, `view` (FULL|LIGHT). Existing
  `HotelContentResponse` already matches the data schema.

## Testing

Extend `tests/` with one test per new use-case method, matching `tests/search_test.go`,
credential-gated (skip when env creds absent).

## Out of scope

- Production host auto-detection beyond a constant switch.
- Caching/rate-limit handling beyond token refresh.
