# iTrip Hotel SDK for Go

A Go SDK that wraps the [Amadeus Hotel APIs](https://developers.amadeus.com/) to provide type-safe access to hotel search, content, offers, and booking functionality.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Architecture](#architecture)
- [Modules](#modules)
  - [List](#list-module)
  - [Offers](#offers-module)
  - [Content](#content-module)
  - [Booking](#booking-module)
- [Search Criteria](#search-criteria)
- [Authentication](#authentication)
- [Error Handling](#error-handling)
- [API Reference](#api-reference)
- [Testing](#testing)
- [Project Structure](#project-structure)

---

## Requirements

- **Go** 1.25.4 or later
- **Amadeus API credentials** (client ID and client secret) from [Amadeus for Developers](https://developers.amadeus.com/)

## Installation

```bash
go get test
```

### Dependencies

| Package            | Version       | Purpose                      |
| ------------------ | ------------- | ---------------------------- |
| `resty.dev/v3`     | v3.0.0-beta.6 | HTTP client for API requests |
| `golang.org/x/net` | v0.43.0       | Network utilities (indirect) |

---

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    sdk "test"
    requestList "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
    requestOffers "github.com/techpartners-asia/amadeus-hotel-integration/modules/offers/dto/request"
    requestContent "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/request"
    requestBooking "github.com/techpartners-asia/amadeus-hotel-integration/modules/booking/dto/request"
)

func main() {
    // 1. Initialize the SDK with your Amadeus credentials
    client, err := sdk.New("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
    if err != nil {
        log.Fatal(err)
    }

    // 2. Search for hotels in Paris
    hotels, err := client.List.HotelListByCityCode(requestList.HotelListByCityCodeRequest{
        CityCode: "PAR",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d hotels\n", len(hotels))

    // 3. Get offers for the first hotel
    if len(hotels) > 0 {
        offers, err := client.Offers.List(requestOffers.HotelOffersListRequest{
            HotelIDs:     []string{hotels[0].HotelId},
            Adults:       2,
            CheckInDate:  "2026-06-01",
            CheckOutDate: "2026-06-05",
        })
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Found %d offer groups\n", len(offers))
    }

    // 4. Get detailed content for a hotel
    content, err := client.Content.GetByID(requestContent.ContentByIDRequest{
        ID: hotels[0].HotelId,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Hotel: %s\n", content.Hotel.Name)

    // 5. Create a booking
    order, err := client.Booking.Create(requestBooking.HotelBookingRequest{
        Data: requestBooking.BookingData{
            Type: "hotel-order",
            Guests: []requestBooking.Guest{
                {
                    Tid:       1,
                    Title:     "MR",
                    FirstName: "JOHN",
                    LastName:  "DOE",
                    Phone:     "+33679278416",
                    Email:     "john.doe@example.com",
                },
            },
            RoomAssociations: []requestBooking.RoomAssociation{
                {
                    HotelOfferId: "OFFER_ID_FROM_STEP_3",
                    GuestReferences: []requestBooking.GuestReference{
                        {GuestReference: "1"},
                    },
                },
            },
            Payment: requestBooking.Payment{
                Method: "CREDIT_CARD",
                PaymentCard: requestBooking.PaymentCard{
                    PaymentCardInfo: requestBooking.PaymentCardInfo{
                        VendorCode: "VI",
                        CardNumber: "4111111111111111",
                        ExpiryDate: "1226",
                        HolderName: "JOHN DOE",
                    },
                },
            },
            TravelAgent: requestBooking.TravelAgent{
                Contact: requestBooking.Contact{
                    Email: "agent@travelagency.com",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Booking confirmed! Order ID: %s\n", order.Id)
}
```

---

## Architecture

The SDK follows a modular, use-case-driven architecture:

```
SDK.New(clientID, clientSecret)
  |
  |-- Authenticates via OAuth2 (client_credentials grant)
  |-- Starts a token manager that auto-refreshes the Bearer token
  |
  +-- SDK
       |-- List    (HotelListUsecase)     --> Amadeus Hotel List API v1
       |-- Offers  (HotelOffersUsecase)   --> Amadeus Hotel Shopping API v3
       |-- Content (ContentUsecase)       --> Amadeus Hotel Content API v3
       |-- Booking (BookingUsecase)       --> Amadeus Hotel Booking API v2 (travel host)
```

**Key design decisions:**

- **Auto-refreshing auth**: a token manager fetches the OAuth2 token once and refreshes it before expiry; each module gets its own HTTP client (via `amadeus.NewClient`) bound to its own base URL, so modules never clobber each other's configuration.
- **Use case interfaces**: Each module exposes an interface, making it easy to mock for testing.
- **Separated DTOs**: Request and response data structures are in dedicated packages, fully typed with JSON tags modeled from the Amadeus swagger schemas.
- **Shared models**: Types that are identical across the APIs (media, dimensions, geocode, amenity, currency, and other value objects) live once in `shared/dto/response` and are aliased by each module. Types that genuinely differ per API (`Hotel`, `Price`, `Room`, `PaymentPolicy`) are kept per-module to preserve full schema fidelity.
- **Generic base response**: A shared `BaseResponse[T]` generic struct handles the standard Amadeus response envelope (data, errors, meta).
- **Typed errors**: Every failure is a `*sharedResponseDTO.APIError` (see [Error Handling](#error-handling)) carrying the HTTP status and structured Amadeus error objects.

---

## Modules

### List Module

Search for hotels by city code or geographic coordinates.

**Amadeus API**: [Hotel List API v1](https://developers.amadeus.com/self-service/category/hotels/api-doc/hotel-list)

#### Search by City Code

```go
hotels, err := client.List.HotelListByCityCode(requestList.HotelListByCityCodeRequest{
    CityCode:    "PAR",                        // Required: IATA 3-letter city code
    Radius:      searchcriteria.Ptr(10),       // Optional: search radius (default: 5)
    RadiusUnit:  searchcriteria.Ptr(searchcriteria.RadiusUnitKM),
    Ratings:     []searchcriteria.Rating{searchcriteria.Rating4, searchcriteria.Rating5},
    Amenities:   []searchcriteria.Amenity{searchcriteria.AmenityWifi, searchcriteria.AmenitySpa},
    HotelSource: searchcriteria.Ptr(searchcriteria.HotelSourceAll),
})
```

Filter values come from the `searchcriteria` package rather than raw strings, so a
typo is a compile error instead of an Amadeus 400. See [Search Criteria](#search-criteria).

**Returns**: `[]GeneralInfoResponse` -- array of hotels with ID, name, IATA code, geo coordinates, address, and distance.

#### Search by Geocode

```go
hotels, err := client.List.HotelListByGeocode(requestGeocode.HotelListByGeocodeRequest{
    Latitude:    48.85,  // Required
    Longitude:   2.29,   // Required
    Radius:      5,
    RadiusUnit:  searchcriteria.RadiusUnitKM,
    Ratings:     []searchcriteria.Rating{searchcriteria.Rating3, searchcriteria.Rating4, searchcriteria.Rating5},
    HotelSource: searchcriteria.HotelSourceAll,
})
```

**Returns**: `[]HotelListResponse` -- array of hotels matching the geographic criteria.

#### Search by Hotel IDs

```go
hotels, err := client.List.HotelListByHotelIds(requestHotels.HotelListByHotelsRequest{
    HotelIds: []string{"MCLONGHM", "ACPAR419"},
})
```

**Returns**: `[]GeneralInfoResponse` -- the hotels matching the supplied Amadeus
property codes.

---

### Offers Module

Search hotel availability and retrieve specific offer details.

**Amadeus API**: [Hotel Search API v3](https://developers.amadeus.com/self-service/category/hotels/api-doc/hotel-search)

#### List Offers

```go
offers, err := client.Offers.List(requestOffers.HotelOffersListRequest{
    HotelIDs:     []string{"MCLONGHM"},  // Required: Amadeus property codes (8 chars)
    Adults:       2,                       // Guests per room (1-9, default: 1)
    CheckInDate:  "2026-06-01",           // Format: YYYY-MM-DD
    CheckOutDate: "2026-06-05",           // Format: YYYY-MM-DD
    RoomQuantity: 1,                       // Number of rooms (1-9, default: 1)
    Currency:     "EUR",                   // ISO currency code
    BoardType:    searchcriteria.BoardTypeBreakfast,
    BestRateOnly: requestOffers.Bool(true), // Only cheapest offer per hotel (*bool; nil = API default true)
    Lang:         "EN",                    // Language for descriptions
})
```

**Returns**: `[]OffersResponse` -- array of hotel offer groups, each containing the hotel info, availability flag, and an array of individual offers with pricing, room details, and policies.

#### Get Offer by ID

```go
offer, err := client.Offers.GetByID(requestOffers.HotelOffersByIDRequest{
    OfferID: "63A93695B58821ABB0EC2B33FE9FAB24D72BF34B1BD7D707293763D8D9378FC3",
    Lang:    "EN",
})
```

**Returns**: `*OffersResponse` -- a single offer group with full details.

---

### Content Module

Retrieve detailed hotel content including rooms, facilities, policies, awards, and points of interest.

**Amadeus API**: [Hotel Content API v3](https://developers.amadeus.com/self-service/category/hotels)

#### Get Hotel Content by ID

```go
content, err := client.Content.GetByID(requestContent.ContentByIDRequest{
    ID: "ADPAR001",  // Amadeus hotel property code
})
```

**Returns**: `*HotelContentResponse` with these sections:

| Section           | Description                                 |
| ----------------- | ------------------------------------------- |
| `Hotel`           | Basic hotel info (name, chain, category)    |
| `Basic`           | Building details, timezone, season info     |
| `Rooms`           | Room classifications, bed types, dimensions |
| `Facilities`      | Meeting rooms, restaurants, amenities       |
| `Policies`        | Payment, check-in/out, pets, cancellation   |
| `Awards`          | Certifications and awards                   |
| `Promotions`      | Current promotional offers                  |
| `PointOfInterest` | Nearby points of interest                   |

---

### Booking Module

Create hotel bookings and retrieve existing orders.

**Amadeus API**: [Hotel Booking API v2](https://developers.amadeus.com/self-service/category/hotels/api-doc/hotel-booking)

#### Create a Booking

```go
order, err := client.Booking.Create(requestBooking.HotelBookingRequest{
    Data: requestBooking.BookingData{
        Type: "hotel-order",
        Guests: []requestBooking.Guest{
            {
                Tid:       1,
                Title:     "MR",
                FirstName: "JOHN",
                LastName:  "DOE",
                Phone:     "+33679278416",
                Email:     "john@example.com",
            },
        },
        RoomAssociations: []requestBooking.RoomAssociation{
            {
                HotelOfferId: "OFFER_ID",
                GuestReferences: []requestBooking.GuestReference{
                    {GuestReference: "1"},
                },
            },
        },
        Payment: requestBooking.Payment{
            Method: "CREDIT_CARD",
            PaymentCard: requestBooking.PaymentCard{
                PaymentCardInfo: requestBooking.PaymentCardInfo{
                    VendorCode:   "VI",
                    CardNumber:   "4111111111111111",
                    ExpiryDate:   "1226",
                    HolderName:   "JOHN DOE",
                    SecurityCode: "123",
                },
            },
        },
        TravelAgent: requestBooking.TravelAgent{
            Contact: requestBooking.Contact{
                Email: "agent@agency.com",
            },
        },
    },
})
```

**Returns**: `*HotelOrder` with:

| Field               | Description                                     |
| ------------------- | ----------------------------------------------- |
| `Id`                | Hotel order ID (store this for cancel/retrieve) |
| `Type`              | Always `"hotel-order"`                          |
| `HotelBookings`     | Array of bookings with confirmation numbers     |
| `Guests`            | Guests with Amadeus-assigned IDs                |
| `AssociatedRecords` | PNR record locators                             |
| `Self`              | URL to retrieve this order                      |

**Booking statuses**: `CONFIRMED`, `PENDING`, `CANCELLED`, `ON_HOLD`, `PAST`, `UNCONFIRMED`, `DENIED`, `GHOST`, `DELETED`

**Payment methods**: `CREDIT_CARD`, `CREDIT_CARD_AGENCY`, `CREDIT_CARD_TRAVELER`, `AGENCY_ACCOUNT`, `VCC_BILLBACK`, `VCC_B2B_WALLET`, `TRAVEL_AGENT_ID`

#### Retrieve a Booking by Reference

```go
order, err := client.Booking.GetByReference("ABCDEF")  // 6-char PNR locator
```

**Returns**: `*HotelOrder` -- the full hotel order with all bookings, guests, and associated records.

#### Retrieve a Booking by Order ID

```go
order, err := client.Booking.GetByID("HOTEL_ORDER_ID")
```

**Returns**: `*HotelOrder` -- the full hotel order. Uses Hotel Booking Retrieve (v2.1).

#### Cancel a Booking

```go
order, err := client.Booking.Cancel("HOTEL_ORDER_ID", "HOTEL_BOOKING_ID")
```

**Returns**: `*HotelOrder` -- the order with the cancelled booking's `bookingStatus`
set to `CANCELLED`. Uses Hotel Booking Manage (v2.2).

#### Modify a Booking

```go
result, err := client.Booking.Modify("HOTEL_ORDER_ID", "HOTEL_BOOKING_ID",
    requestBooking.UpdateHotelBookingRequest{
        Data: requestBooking.UpdateHotelBooking{
            HotelBooking: requestBooking.UpdateHotelBookingData{
                RoomAssociation: &requestBooking.UpdateRoomAssociation{
                    SpecialRequest: "Late check-in",
                },
            },
        },
    })
```

**Returns**: `*HotelBookingUpdateResponse` -- `Included` carries the full updated
order. Send only the fields you want to change. Uses Hotel Booking Manage (v2.2).

#### Delete a Booking

```go
result, err := client.Booking.Delete("HOTEL_ORDER_ID", "HOTEL_BOOKING_ID")
```

**Returns**: `*DeleteBookingResult` -- contains the provider `CancellationNumber`.
Uses Hotel Booking Manage (v2.2).

---

## Search Criteria

Search filters take typed values from the `searchcriteria` package instead of raw
strings, so an invalid code fails to compile rather than returning an Amadeus 400.

```go
import "github.com/techpartners-asia/amadeus-hotel-integration/searchcriteria"

req := requestList.HotelListByCityCodeRequest{
    CityCode:  "PAR",
    Amenities: []searchcriteria.Amenity{
        searchcriteria.AmenitySwimmingPool,
        searchcriteria.AmenityWifi,
    },
}
```

### Available lists

| Type | Values | Used by |
|------|--------|---------|
| `Amenity` | 34 codes | List (by-city, by-geocode) |
| `Rating` | `1`-`5` (max 4 per request) | List (by-city, by-geocode) |
| `HotelSource` | `BEDBANK`, `DIRECTCHAIN`, `ALL` | List (by-city, by-geocode) |
| `RadiusUnit` | `KM`, `MILE` | List (by-city, by-geocode) |
| `BoardType` | `ROOM_ONLY`, `BREAKFAST`, `HALF_BOARD`, `FULL_BOARD`, `ALL_INCLUSIVE` | Offers |
| `PaymentPolicy` | `GUARANTEE`, `DEPOSIT`, `NONE` | Offers |
| `RateCode` | `PRO`, `GOV`, `AAA`, `MIL`, `SNR`, `COR`, `RAC`, plus corporate codes | Offers |
| `ContentView` | `FULL`, `LIGHT` | Content |

### Enumerating a list

Every type exposes `All<Type>s()` for building a filter UI, plus `Label()` and
`IsValid()`:

```go
for _, a := range searchcriteria.AllAmenities() {
    fmt.Printf("%s -- %s\n", a, a.Label())  // SWIMMING_POOL -- Swimming Pool
}

searchcriteria.Amenity("SWIMMING_POOL").IsValid()  // true
searchcriteria.Amenity("POOL").IsValid()           // false
```

The same lists hang off the SDK value when that is easier to pass around:

```go
client.SearchCriteria.Amenities()
client.SearchCriteria.BoardTypes()
```

Both forms return the same static data compiled into the SDK -- no network call,
no credentials, and nothing that can fail. The package-level `All*` functions work
without an SDK value at all.

### Two codes differ from the Amadeus docs

Amadeus documents two amenities with spellings its own API rejects
(`7211 INVALID FACILITY CODE`). This SDK uses the values that actually work:

| Amadeus docs | Accepted by the API | Constant |
|--------------|---------------------|----------|
| `BABY-SITTING` | `BABY_SITTING` | `AmenityBabySitting` |
| `BAR or LOUNGE` | `BAR_LOUNGE` | `AmenityBarOrLounge` |

`TestAmenityCodesAcceptedByAmadeus` probes every amenity against the live endpoint
to catch further drift:

```bash
AMADEUS_PROBE_AMENITIES=1 go test ./tests/ -run TestAmenityCodesAcceptedByAmadeus
```

### Rate codes are not a closed set

`RateCode.IsValid()` checks only that a code is three uppercase alphanumerics,
because corporate rate codes are negotiated per account and cannot be enumerated.
`AllRateCodes()` returns the documented public and qualified rates -- treat it as a
starting list for a UI, not a whitelist.

---

## Authentication

The SDK uses **OAuth2 Client Credentials** flow to authenticate with the Amadeus API.

```
POST https://test.travel.api.amadeus.com/v1/security/oauth2/token
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_id=YOUR_CLIENT_ID
&client_secret=YOUR_CLIENT_SECRET
```

Authentication happens automatically when you call `sdk.New()`, which returns an
`error` if the credentials are rejected. The resulting Bearer token is attached to
all subsequent API requests and is **automatically refreshed** before it expires, so
a single SDK instance can be reused for the lifetime of your process.

> **Important — Enterprise credentials required.** Every endpoint targets the Amadeus
> Enterprise ("travel") host, matching the swagger specifications. Self-service
> credentials (`api.amadeus.com`) are **rejected** by this host — `sdk.New()` will
> return an `invalid_client` error. Use the Enterprise credentials issued for your
> Amadeus travel-host contract.

**Environments:**

| Environment | Auth + API host                          |
| ----------- | ---------------------------------------- |
| Test        | `https://test.travel.api.amadeus.com`    |
| Production  | `https://travel.api.amadeus.com`         |

The SDK uses the **test environment** by default. To switch to production, change the
`test.travel.api.amadeus.com` host prefixes to `travel.api.amadeus.com` in
`constants/url.go` (a single set of `TRAVEL_BASE_*` constants drives both auth and all
modules).

---

## Error Handling

All SDK methods return `(result, error)`. Errors come from:

1. **Network errors** -- connection failures, timeouts (returned verbatim from the HTTP client).
2. **API errors** -- any non-2xx response, or a `2xx` response that carries an `errors` array. These are returned as a typed `*APIError`.

### The `APIError` type

Every API failure is a `*sharedResponseDTO.APIError`, which carries the HTTP status and the structured Amadeus error objects:

```go
import (
    "errors"
    sharedResponseDTO "github.com/techpartners-asia/amadeus-hotel-integration/shared/dto/response"
)

offers, err := client.Offers.List(request)
if err != nil {
    var apiErr *sharedResponseDTO.APIError
    if errors.As(err, &apiErr) {
        // Structured API error -- inspect it.
        log.Printf("HTTP %d", apiErr.StatusCode)
        for _, e := range apiErr.Errors {
            log.Printf("  [%d] %s - %s", e.Code, e.Title, e.Detail)
            // e.Source.Parameter / e.Source.Pointer / e.Documentation also available
        }
        switch apiErr.StatusCode {
        case 401:
            // token rejected -- re-create the SDK with valid credentials
        case 404:
            // resource not found (e.g. unknown offer id / hotel order)
        case 429:
            // rate limited -- back off and retry
        }
        return
    }
    // Network / transport error.
    log.Printf("request failed: %v", err)
    return
}
```

`err.Error()` renders a readable summary, e.g.
`amadeus: [38196] Resource not found - The targeted resource doesn't exist (status 404)`.

`APIError` fields:

| Field        | Type              | Description                                              |
| ------------ | ----------------- | ------------------------------------------------------- |
| `StatusCode` | `int`             | HTTP status of the failed response.                     |
| `Errors`     | `[]ErrorResponse` | Structured Amadeus errors (`Status`, `Code`, `Title`, `Detail`, `Source`, `Documentation`). |
| `Raw`        | `string`          | Raw response body (fallback for non-standard errors).   |

The Amadeus error envelope looks like:

```json
{
  "errors": [
    {
      "status": 400,
      "code": 477,
      "title": "INVALID FORMAT",
      "detail": "invalid query parameter format",
      "source": { "parameter": "cityCode", "pointer": "", "example": "PAR" },
      "documentation": "https://developers.amadeus.com"
    }
  ]
}
```

### Possible errors

**HTTP status codes**

| Status | Meaning | Typical cause |
| ------ | ------- | ------------- |
| `400`  | Bad Request | Malformed/missing parameter or body (see error `code`). |
| `401`  | Unauthorized | Expired/invalid token, or **self-service credentials used against the Enterprise host** (`invalid_client` at `New()`). |
| `403`  | Forbidden | Credentials lack access to the endpoint. |
| `404`  | Not Found | Unknown hotel id, offer id, or hotel order; or endpoint not enabled for your account. |
| `429`  | Too Many Requests | Rate limit exceeded -- back off and retry. |
| `500`  | Server Error | Amadeus-side failure (`code` 141 "SYSTEM ERROR HAS OCCURRED"); safe to retry idempotent calls. |

**Common Amadeus error codes** (the `code` field; not exhaustive):

| Code  | Title | Where it shows up |
| ----- | ----- | ----------------- |
| 477   | INVALID FORMAT | bad query parameter / body format |
| 1797  | NOT FOUND | no result for the query |
| 141   | SYSTEM ERROR HAS OCCURRED | transient Amadeus error |
| 38196 | Resource not found | retrieving a non-existent hotel order |
| 1257  | INVALID PROPERTY CODE | bad `hotelId` |
| 3237  | PROPERTY CODE NOT FOUND IN SYSTEM | unknown property in List/Content |
| 1157  | INVALID CITY CODE | bad `cityCode` in List |
| 36801 | INVALID OFFER ID | bad `offerId` in Offers/Booking |
| 38420 | OFFER NOT FOUND | offer expired between search and booking |
| 33555 | NUMBER OF ROOMS MISMATCH BETWEEN SHOPPING AND BOOKING | booking payload inconsistent with offer |
| 27706 | PRICING CONDITIONS HAVE CHANGED | re-price before booking |
| 38592 | INVALID METHOD OF PAYMENT | unsupported `payment.method` |
| 1146 / 1427 | DEPOSIT / GUARANTEE REQUIRED | payment guarantee missing |
| 8622  | MODIFY/CANCEL NOT ALLOWED | booking can no longer be modified/cancelled |
| 1694  | ITEM ALREADY CANCELLED | cancelling an already-cancelled booking |
| 25859 | MODIFICATION NOT ALLOWED FOR THIS CHAIN | chain does not support `Modify` |

> Amadeus error codes are returned as strings/ints depending on the endpoint; the SDK
> exposes them as `int` in `ErrorResponse.Code`. Always branch on `code` (machine-readable)
> rather than `title` (which may be localized).

---

## API Reference

### SDK Initialization

```go
func New(id, secret string) (*SDK, error)
```

Creates and returns a new SDK instance. Authenticates with Amadeus using the provided credentials and returns an `error` if authentication fails.

### List Module

| Method                | Signature                                                     | Description                             |
| --------------------- | ------------------------------------------------------------- | --------------------------------------- |
| `HotelListByCityCode` | `(HotelListByCityCodeRequest) ([]GeneralInfoResponse, error)` | Search hotels by IATA city code         |
| `HotelListByGeocode`  | `(HotelListByGeocodeRequest) ([]HotelListResponse, error)`    | Search hotels by geographic coordinates |
| `HotelListByHotelIds` | `(HotelListByHotelsRequest) ([]GeneralInfoResponse, error)`   | Look up hotels by Amadeus property codes |

### Offers Module

| Method    | Signature                                            | Description                      |
| --------- | ---------------------------------------------------- | -------------------------------- |
| `List`    | `(HotelOffersListRequest) ([]OffersResponse, error)` | Search hotel offers/availability |
| `GetByID` | `(HotelOffersByIDRequest) (*OffersResponse, error)`  | Retrieve a specific offer by ID  |

### Content Module

| Method    | Signature                                             | Description                |
| --------- | ----------------------------------------------------- | -------------------------- |
| `GetByID` | `(ContentByIDRequest) (*HotelContentResponse, error)` | Get detailed hotel content |

### Booking Module

| Method           | Signature                                                                       | Description                              |
| ---------------- | ------------------------------------------------------------------------------- | ---------------------------------------- |
| `Create`         | `(HotelBookingRequest) (*HotelOrder, error)`                                    | Create a hotel booking                   |
| `GetByReference` | `(string) (*HotelOrder, error)`                                                 | Retrieve a booking by PNR reference      |
| `GetByID`        | `(string) (*HotelOrder, error)`                                                 | Retrieve a hotel order by order id       |
| `Cancel`         | `(orderID, bookingID string) (*HotelOrder, error)`                              | Cancel a booking within an order         |
| `Modify`         | `(orderID, bookingID string, UpdateHotelBookingRequest) (*HotelBookingUpdateResponse, error)` | Modify a booking within an order |
| `Delete`         | `(orderID, bookingID string) (*DeleteBookingResult, error)`                     | Delete a booking within an order         |

### Search Criteria (`client.SearchCriteria`)

Static lists compiled into the SDK. No method performs a network call or returns
an error. See [Search Criteria](#search-criteria).

| Method             | Signature            | Description                          |
| ------------------ | -------------------- | ------------------------------------ |
| `Amenities`        | `() []Amenity`       | 34 amenity filter codes              |
| `Ratings`          | `() []Rating`        | Star ratings 1-5                     |
| `HotelSources`     | `() []HotelSource`   | BEDBANK, DIRECTCHAIN, ALL            |
| `RadiusUnits`      | `() []RadiusUnit`    | KM, MILE                             |
| `BoardTypes`       | `() []BoardType`     | Meal plans                           |
| `PaymentPolicies`  | `() []PaymentPolicy` | GUARANTEE, DEPOSIT, NONE             |
| `ContentViews`     | `() []ContentView`   | FULL, LIGHT                          |
| `RateCodes`        | `() []RateCode`      | Documented rate codes (not a whitelist) |

Each returned type also has `Label() string` and `IsValid() bool`.

---

## Testing

Run the integration tests:

```bash
go test ./tests/ -v
```

The test suite demonstrates the typical SDK workflow:

1. Initialize the SDK with credentials
2. Search hotels by city code
3. Iterate through results and fetch content for each hotel

> **Note**: Tests use the Amadeus test environment and require valid API credentials.

---

## Project Structure

```
sdk/
├── sdk.go                                  # SDK entry point and initialization
├── go.mod                                  # Go module definition
├── README.md                               # This file
│
├── constants/
│   └── url.go                              # API base URLs (v1, v3)
│
├── integrations/
│   └── amadeus/
│       ├── base.go                         # OAuth2 authentication and HTTP client setup
│       └── type.go                         # Auth response type definition
│
├── modules/
│   ├── list/                               # Hotel List Module
│   │   ├── usecase/
│   │   │   └── base.go                     # HotelListByGeocode, HotelListByCityCode
│   │   └── dto/
│   │       ├── request/
│   │       │   ├── city/request.go         # City code search request DTO
│   │       │   └── geocode/request.go      # Geocode search request DTO
│   │       └── response/
│   │           └── list-response.go        # Hotel list response DTOs
│   │
│   ├── offers/                             # Hotel Offers Module
│   │   ├── usecases/
│   │   │   └── base.go                     # List, GetByID
│   │   └── dto/
│   │       ├── request/
│   │       │   ├── list-request.go         # Offer list search request DTO
│   │       │   └── by-id-request.go        # Get offer by ID request DTO
│   │       └── response/
│   │           └── offer-response.go       # Hotel offer response DTOs
│   │
│   ├── content/                            # Hotel Content Module
│   │   ├── usecase/
│   │   │   └── base.go                     # GetByID
│   │   └── dto/
│   │       ├── request/
│   │       │   └── request.go              # Content request DTO
│   │       └── response/
│   │           ├── base.go                 # HotelContentResponse (top-level)
│   │           ├── hotel-general-response.go
│   │           ├── room-response.go
│   │           ├── policy-response.go
│   │           ├── facility-response.go
│   │           ├── awards-response.go
│   │           ├── promotion-response.go
│   │           └── point-of-interest-response.go
│   │
│   └── booking/                            # Hotel Booking Module
│       ├── usecase/
│       │   └── base.go                     # Create, GetByReference
│       └── dto/
│           ├── request/
│           │   └── base.go                 # Full booking request DTOs (22 structs)
│           └── response/
│               └── base.go                 # Full booking response DTOs (50+ structs)
│
├── searchcriteria/                         # Typed values accepted in search filters
│   ├── searchcriteria.go                   # entry[T], codes/labelOf/isValid, Join
│   ├── catalog.go                          # Catalog interface behind sdk.SearchCriteria
│   ├── amenity.go                          # Amenity (34 codes)
│   ├── rating.go                           # Rating (1-5)
│   ├── hotel-source.go                     # HotelSource
│   ├── radius-unit.go                      # RadiusUnit
│   ├── board-type.go                       # BoardType
│   ├── payment-policy.go                   # PaymentPolicy
│   ├── content-view.go                     # ContentView
│   └── rate-code.go                        # RateCode (open-ended)
│
├── shared/
│   └── dto/
│       ├── request/
│       │   └── hotel-general-info.go       # Shared request types
│       └── response/
│           ├── base.go                     # BaseResponse[T], ErrorResponse, MetaResponse
│           ├── amenity.go                  # Shared amenity response types
│           └── ...                         # Other shared response types
│
└── tests/
    ├── search_test.go                      # Integration test: search + content
    ├── searchcriteria_test.go              # Search-criteria lists (no network)
    └── amenity_probe_test.go               # Probes each amenity live (opt-in)
```

---

## Typical Workflow

The standard hotel booking flow using this SDK follows these steps:

```
1. List Hotels          -->  sdk.List.HotelListByCityCode()
   Find hotels in a          or sdk.List.HotelListByGeocode()
   city or location

2. Get Offers           -->  sdk.Offers.List()
   Check availability        Pass hotel IDs from step 1,
   and pricing               with dates and guest count

3. Get Content          -->  sdk.Content.GetByID()
   (Optional) Fetch          Rooms, facilities, policies,
   rich hotel details        photos, awards

4. Book                 -->  sdk.Booking.Create()
   Create the order          Pass offer ID from step 2,
                             guest details, and payment

5. Retrieve             -->  sdk.Booking.GetByReference()
   (Optional) Check          Pass PNR reference from step 4
   booking status
```

---

## License

Private / Internal use.
