package offers

import (
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/mapping"
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Thin aliases over internal/mapping, so mapper.go reads as domain translation
// rather than as a series of package-qualified calls.

func parseMoney(amount, currency string) money.Money { return mapping.Money(amount, currency) }
func parseDate(s string) datetime.Date               { return mapping.Date(s) }
func parseTimestamp(s string) *time.Time             { return mapping.Timestamp(s) }
func parseWireBool(s string) bool                    { return mapping.Bool(s) }
func orDefault(value, fallback string) string        { return mapping.Or(value, fallback) }

func mapDescription(d *dto.DescriptionResponse) *media.Text     { return mapping.Text(d) }
func mapTextContent(t *dto.TextContentResponse) *media.Text     { return mapping.TextContent(t) }
func mapDimensions(d *dto.DimensionsResponse) *media.Dimensions { return mapping.Dimensions(d) }
func mapMediaAssets(m []dto.MediaResponse) []media.Asset        { return mapping.MediaAssets(m) }

// currencyOr returns the block's own currency, or the parent's when it has none.
func currencyOr(value, fallback string) money.Currency {
	return money.Currency(mapping.Or(value, fallback))
}
