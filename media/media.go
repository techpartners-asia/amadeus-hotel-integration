// Package media holds the image and asset value objects shared by the offers
// and content contexts.
//
// Amadeus returns the same media block for a room photo attached to an offer
// and for a property photo attached to hotel content, so it belongs to neither
// context alone.
package media

// Kind is the sort of asset an Asset points at.
type Kind string

const (
	KindImage Kind = "IMAGE"
	KindIcon  Kind = "ICON"
	KindFile  Kind = "FILE"
)

// Asset is a single media item: a photograph, an icon or a file.
type Asset struct {
	// ID is Amadeus's identifier for the asset.
	ID string
	// Kind is the asset type as Amadeus classified it.
	Kind Kind
	// Name is the asset's file-level name, e.g. "guest_room".
	Name string
	// Title is a display title.
	Title string
	// Caption is display text accompanying the asset.
	Caption string
	// Hint is supplementary display text.
	Hint string
	// Alt is the description intended for screen readers. Carry it through to
	// any img tag you render: it is the only accessible description Amadeus
	// supplies.
	Alt string
	// URL points at the original, full-size asset.
	URL string
	// Category groups the asset, e.g. "EXTERIOR", "GUEST_ROOM".
	Category string
	// Tags are Amadeus's free-form labels for the asset.
	Tags []string
	// Description is the asset's descriptive text with its language.
	Description *Text
	// Scales are alternative renditions at different sizes. Prefer one of
	// these over URL when you know the display size; see Best.
	Scales []Scale
	// Metadata describes the encoding, dimensions and rights of the asset.
	Metadata *Metadata
}

// Scale is one rendition of an asset at a particular size.
type Scale struct {
	// URL points at this rendition.
	URL string
	// Size is the file size, when Amadeus reports it.
	Size *Size
	// Dimensions are the pixel dimensions, when Amadeus reports them.
	Dimensions *Dimensions
	// Duration is an ISO 8601 duration for time-based media.
	Duration string
}

// Size is a file size with its unit.
type Size struct {
	Unit  string
	Value int
}

// Dimensions are the physical or pixel measurements of an asset or a room.
type Dimensions struct {
	Width         int
	Height        int
	Length        int
	Unit          string
	Area          float64
	AreaUnit      string
	DecimalPlaces int
}

// Metadata describes an asset's encoding and provenance.
type Metadata struct {
	MediaType     string
	SubType       string
	Encoding      string
	ETag          string
	Duration      string
	Application   string
	Size          *Size
	Dimensions    *Dimensions
	Source        *Source
	ClickToAction *ClickToAction
}

// Source identifies who owns an asset and under what terms.
type Source struct {
	Code      string
	Copyright string
	Filename  string
	Symbology string
	Version   string
}

// ClickToAction is a hyperlink Amadeus attaches to an asset.
type ClickToAction struct {
	Text string
	URL  string
}

// Text is a string with the language and encoding Amadeus supplied for it.
type Text struct {
	// Value is the text itself.
	Value string
	// Type classifies what the text describes, e.g. "PROPERTY_DESCRIPTION".
	Type string
	// Lang is the language tag, e.g. "fr-FR". Amadeus falls back to English
	// when it holds no text in the language you asked for.
	Lang string
	// Status is Amadeus's lifecycle marker for the text, e.g. "ACTIVE".
	Status string
	// CharSet and Encoding describe the representation. Amadeus occasionally
	// sends base-64, in which case Value is the encoded form.
	CharSet  string
	Encoding string
	// ContentType is the IANA media type of the text.
	ContentType string
}

// String returns the text value, so a Text prints as its content.
func (t Text) String() string { return t.Value }

// IsEmpty reports whether the text carries no content.
func (t Text) IsEmpty() bool { return t.Value == "" }

// Best returns the scale whose width is closest to targetWidth without
// exceeding it, falling back to the smallest available and finally to the
// asset's own URL.
//
// It exists because Amadeus returns up to a dozen renditions per photograph,
// and rendering a thumbnail grid from the originals downloads tens of megabytes
// the user never sees.
func (a Asset) Best(targetWidth int) string {
	best := ""
	bestWidth := 0
	smallest := ""
	smallestWidth := 0

	for _, scale := range a.Scales {
		if scale.URL == "" || scale.Dimensions == nil {
			continue
		}
		width := scale.Dimensions.Width
		if width <= 0 {
			continue
		}

		if smallest == "" || width < smallestWidth {
			smallest, smallestWidth = scale.URL, width
		}
		if width <= targetWidth && width > bestWidth {
			best, bestWidth = scale.URL, width
		}
	}

	switch {
	case best != "":
		return best
	case smallest != "":
		return smallest
	default:
		return a.URL
	}
}
