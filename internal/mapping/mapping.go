// Package mapping holds the wire-to-domain conversions shared by more than one
// bounded context.
//
// Each context owns its own mapper; this is only the primitive layer beneath
// them - parsing Amadeus's money strings, its date formats and its quoted
// booleans, and translating the media block that offers and content both carry.
// Putting them here keeps one definition of "how Amadeus spells a timestamp"
// rather than four subtly different ones.
//
// Nothing here returns an error. A malformed field in one corner of a response
// must not discard an otherwise usable offer, so unparseable input becomes the
// zero value and the caller keeps everything else.
package mapping

import (
	"strings"
	"time"

	"github.com/techpartners-asia/amadeus-hotel-integration/datetime"
	"github.com/techpartners-asia/amadeus-hotel-integration/internal/amadeus/dto"
	"github.com/techpartners-asia/amadeus-hotel-integration/media"
	"github.com/techpartners-asia/amadeus-hotel-integration/money"
)

// Money parses one of Amadeus's decimal price strings against a currency code.
// An empty or malformed amount yields a zero Money that still carries the
// currency, so a caller can tell "no price" from "price in EUR".
func Money(amount, currency string) money.Money {
	parsed, err := money.Parse(amount, currency)
	if err != nil {
		return money.New(money.Amount{}, money.Currency(currency))
	}
	return parsed
}

// Date parses Amadeus's "2006-01-02" calendar dates, yielding the zero Date for
// anything it cannot read.
func Date(s string) datetime.Date {
	parsed, err := datetime.ParseDate(s)
	if err != nil {
		return datetime.Date{}
	}
	return parsed
}

// timestampLayouts are the forms Amadeus uses for an instant, in the order they
// are tried.
//
// The list is this long because Amadeus is genuinely inconsistent: cancellation
// deadlines arrive with a zone, hold times usually without, and some sources
// send a bare date. A parser accepting only one of them silently drops most
// cancellation deadlines, which is the field callers most need.
var timestampLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02T15:04",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

// Timestamp parses an Amadeus instant, returning nil when the value is absent
// or unreadable. A nil result means "no deadline", which callers must not
// confuse with "the deadline has passed".
func Timestamp(s string) *time.Time {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	for _, layout := range timestampLayouts {
		if parsed, err := time.Parse(layout, s); err == nil {
			return &parsed
		}
	}
	return nil
}

// Bool reads the quoted booleans Amadeus sends for isLoyaltyRate and
// isOptional, which arrive as the JSON strings "true" and "false" rather than
// as JSON booleans. Anything unrecognised reads as false.
func Bool(s string) bool {
	return strings.EqualFold(strings.TrimSpace(s), "true")
}

// Text translates a description block, returning nil when it carries no text so
// callers can distinguish "no description" from "empty description".
func Text(d *dto.DescriptionResponse) *media.Text {
	if d == nil || d.Text == "" {
		return nil
	}
	return &media.Text{Value: d.Text, Lang: d.Lang}
}

// TextContent translates the richer text block, which adds encoding metadata.
func TextContent(t *dto.TextContentResponse) *media.Text {
	if t == nil || t.Text == "" {
		return nil
	}
	return &media.Text{
		Value:       t.Text,
		Lang:        t.Lang,
		Status:      t.Status,
		CharSet:     t.CharSet,
		Encoding:    t.Encoding,
		ContentType: t.IanaContentType,
	}
}

// QualifiedText translates the qualified free-text block, which additionally
// classifies what the text describes.
func QualifiedText(q *dto.QualifiedFreeTextResponse) *media.Text {
	if q == nil || q.Text == "" {
		return nil
	}
	return &media.Text{
		Value:       q.Text,
		Type:        string(q.Type),
		Lang:        q.Lang,
		Status:      q.Status,
		CharSet:     q.CharSet,
		Encoding:    q.Encoding,
		ContentType: q.IamaContentType,
	}
}

// Dimensions translates a measurement block.
func Dimensions(d *dto.DimensionsResponse) *media.Dimensions {
	if d == nil {
		return nil
	}
	return &media.Dimensions{
		Width:         d.Width,
		Height:        d.Height,
		Length:        d.Length,
		Unit:          d.Unit,
		Area:          d.Area,
		AreaUnit:      d.AreaUnit,
		DecimalPlaces: d.DecimalPlaces,
	}
}

// MediaAssets translates a list of media items.
func MediaAssets(wire []dto.MediaResponse) []media.Asset {
	if wire == nil {
		return nil
	}
	out := make([]media.Asset, len(wire))
	for i, m := range wire {
		out[i] = MediaAsset(m)
	}
	return out
}

// MediaAsset translates one media item.
func MediaAsset(m dto.MediaResponse) media.Asset {
	asset := media.Asset{
		ID:          m.Id,
		Kind:        mediaKind(m),
		Name:        m.Name,
		Title:       m.Title,
		Caption:     m.Caption,
		Hint:        m.Hint,
		Alt:         m.Alt,
		URL:         m.Href,
		Category:    m.Category,
		Tags:        m.Tags,
		Description: QualifiedText(m.Description),
		Metadata:    mediaMetadata(m.MediaMetaData),
	}

	for _, scale := range m.MediaScales {
		asset.Scales = append(asset.Scales, media.Scale{
			URL:        scale.Href,
			Size:       mediaSize(scale.Size),
			Dimensions: Dimensions(scale.Dimensions),
			Duration:   scale.Duration,
		})
	}

	return asset
}

// mediaKind normalises the fields Amadeus uses to say what an asset is.
//
// The declared type is preferred, but in practice it is usually absent: the
// Hotel Content API leaves "type" and "mediaType" unset on every entry and
// expects you to infer the kind from whether the entry actually carries an
// image. An entry with renditions or a URL is one; an entry with only text is
// a prose block that happens to travel in the same array.
func mediaKind(m dto.MediaResponse) media.Kind {
	for _, candidate := range []string{m.Type, m.MediaType} {
		switch strings.ToUpper(strings.TrimSpace(candidate)) {
		case "IMAGE":
			return media.KindImage
		case "ICON":
			return media.KindIcon
		case "FILE":
			return media.KindFile
		}
	}

	if m.Href != "" || len(m.MediaScales) > 0 {
		return media.KindImage
	}
	return ""
}

func mediaSize(s *dto.MediaSizeResponse) *media.Size {
	if s == nil {
		return nil
	}
	return &media.Size{Unit: s.Unit, Value: s.Value}
}

func mediaMetadata(m *dto.MediaMetaDataResponse) *media.Metadata {
	if m == nil {
		return nil
	}

	metadata := &media.Metadata{
		MediaType:   m.MediaType,
		SubType:     m.SubType,
		Encoding:    m.Encoding,
		ETag:        m.Etag,
		Duration:    m.Duration,
		Application: m.Application,
		Size:        mediaSize(m.Size),
		Dimensions:  Dimensions(m.Dimensions),
	}
	if m.MediaSource != nil {
		metadata.Source = &media.Source{
			Code:      m.MediaSource.Code,
			Copyright: m.MediaSource.Copyright,
			Filename:  m.MediaSource.Filename,
			Symbology: m.MediaSource.Symbology,
			Version:   m.MediaSource.Version,
		}
	}
	if m.ClickToAction != nil {
		metadata.ClickToAction = &media.ClickToAction{
			Text: m.ClickToAction.PlainText,
			URL:  m.ClickToAction.Href,
		}
	}
	return metadata
}

// Or returns value when it is non-empty, and fallback otherwise. It is the
// common shape of "this block has its own currency, or inherits the parent's".
func Or(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
