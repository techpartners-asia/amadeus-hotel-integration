// Command capture records live Amadeus responses into the testdata fixtures
// the offline mapper tests run against.
//
// Run it when Amadeus changes a schema, or to replace hand-written fixtures
// with real payloads:
//
//	AMADEUS_CLIENT_ID=... AMADEUS_CLIENT_SECRET=... go run ./internal/capture
//
// It writes into each context's testdata directory. Re-running the mapper tests
// afterwards is the point of the exercise: a test that now fails is a field
// Amadeus changed, and a diff in the fixture shows exactly what.
//
// It never captures booking responses. Creating an order on the sandbox is
// still a booking, and this tool must not make one as a side effect of running
// it; booking/testdata/order.json is maintained by hand.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus"
)

func main() {
	var (
		cityCode = flag.String("city", "PAR", "IATA city code to capture against")
		host     = flag.String("host", "https://test.travel.api.amadeus.com", "Amadeus host root")
		outDir   = flag.String("out", ".", "repository root to write testdata into")
	)
	flag.Parse()

	id, secret := os.Getenv("AMADEUS_CLIENT_ID"), os.Getenv("AMADEUS_CLIENT_SECRET")
	if id == "" || secret == "" {
		fmt.Fprintln(os.Stderr, "set AMADEUS_CLIENT_ID and AMADEUS_CLIENT_SECRET")
		os.Exit(1)
	}

	client := amadeus.NewClient(amadeus.Options{
		ClientID:     id,
		ClientSecret: secret,
		Host:         *host,
		Timeout:      90 * time.Second,
		UserAgent:    "amadeus-hotel-integration-capture",
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err := run(ctx, client, *cityCode, *outDir); err != nil {
		fmt.Fprintf(os.Stderr, "capture failed: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, client *amadeus.Client, cityCode, outDir string) error {
	// Hotel list first: everything downstream needs real property codes, and
	// hard-coding them would make this tool go stale.
	hotels, err := capture(ctx, client, outDir, "inventory", "hotels-by-city", amadeus.Request{
		Path:  "/v1/reference-data/locations/hotels/by-city",
		Query: url.Values{"cityCode": {cityCode}, "radius": {"5"}, "radiusUnit": {"KM"}},
	})
	if err != nil {
		return fmt.Errorf("hotel list: %w", err)
	}

	ids, err := hotelIDs(hotels, 15)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		return fmt.Errorf("no hotels returned for %s; try another city", cityCode)
	}
	fmt.Printf("captured %d hotels for %s\n", len(ids), cityCode)

	// Content, for the first property that has any.
	if _, err := capture(ctx, client, outDir, "content", "hotel", amadeus.Request{
		Path:  "/v3/reference-data/locations/by-hotel",
		Query: url.Values{"hotelID": {ids[0]}, "view": {"FULL"}},
	}); err != nil {
		return fmt.Errorf("hotel content: %w", err)
	}

	// Offers. bestRateOnly=false is what makes the fixture useful: with the
	// default, every hotel collapses to one offer and the grouping tests have
	// nothing to group.
	checkIn := time.Now().AddDate(0, 1, 0).Format("2006-01-02")
	checkOut := time.Now().AddDate(0, 1, 3).Format("2006-01-02")

	search, err := capture(ctx, client, outDir, "offers", "search", amadeus.Request{
		Path: "/v3/shopping/hotel-offers",
		Query: url.Values{
			"hotelIds":     {strings.Join(ids, ",")},
			"checkInDate":  {checkIn},
			"checkOutDate": {checkOut},
			"adults":       {"2"},
			"bestRateOnly": {"false"},
		},
	})
	if err != nil {
		return fmt.Errorf("hotel offers: %w", err)
	}

	offerID, err := firstOfferID(search)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: no bookable offer found, skipping offer-by-id: %v\n", err)
		return nil
	}

	if _, err := capture(ctx, client, outDir, "offers", "offer-by-id", amadeus.Request{
		Path: "/v3/shopping/hotel-offers/" + offerID,
	}); err != nil {
		return fmt.Errorf("offer by id: %w", err)
	}

	fmt.Println("\nRe-run the tests: a failure now is a field Amadeus changed.")
	fmt.Println("\tgo test ./...")
	return nil
}

// capture performs one request and writes its indented body to
// <outDir>/<context>/testdata/<name>.json.
func capture(ctx context.Context, client *amadeus.Client, outDir, contextDir, name string, req amadeus.Request) ([]byte, error) {
	body, err := amadeus.DoRaw(ctx, client, req)
	if err != nil {
		return nil, err
	}

	var indented any
	if err := json.Unmarshal(body, &indented); err != nil {
		return nil, fmt.Errorf("response was not JSON: %w", err)
	}
	pretty, err := json.MarshalIndent(indented, "", "  ")
	if err != nil {
		return nil, err
	}

	path := filepath.Join(outDir, contextDir, "testdata", name+".json")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, append(pretty, '\n'), 0o644); err != nil {
		return nil, err
	}

	fmt.Printf("wrote %s (%d bytes)\n", path, len(pretty))
	return body, nil
}

// hotelIDs pulls up to limit property codes out of a hotel-list response.
func hotelIDs(body []byte, limit int) ([]string, error) {
	var envelope struct {
		Data []struct {
			HotelID string `json:"hotelId"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}

	ids := make([]string, 0, limit)
	for _, hotel := range envelope.Data {
		if hotel.HotelID == "" {
			continue
		}
		if ids = append(ids, hotel.HotelID); len(ids) == limit {
			break
		}
	}
	return ids, nil
}

// firstOfferID finds an offer to re-fetch by ID.
func firstOfferID(body []byte) (string, error) {
	var envelope struct {
		Data []struct {
			Offers []struct {
				ID string `json:"id"`
			} `json:"offers"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return "", err
	}

	for _, hotel := range envelope.Data {
		for _, offer := range hotel.Offers {
			if offer.ID != "" {
				return offer.ID, nil
			}
		}
	}
	return "", fmt.Errorf("the search returned no offers")
}
