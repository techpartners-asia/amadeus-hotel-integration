package searchcriteria

// ContentView sets how much detail the Hotel Content API returns, via the
// `view` query parameter.
type ContentView string

const (
	// ContentViewFull returns every content block: rooms, facilities, policies,
	// awards and points of interest. The SDK sends this when View is unset,
	// because the Amadeus default returns only the basic block.
	ContentViewFull ContentView = "FULL"
	// ContentViewLight returns the basic block only.
	ContentViewLight ContentView = "LIGHT"
)

var contentViewCatalog = []entry[ContentView]{
	{ContentViewFull, "Full"},
	{ContentViewLight, "Light"},
}

// AllContentViews returns every view Amadeus accepts.
func AllContentViews() []ContentView { return codes(contentViewCatalog) }

// Label returns a human-readable name for v, or "" when v is not a known code.
func (v ContentView) Label() string { return labelOf(contentViewCatalog, v) }

// IsValid reports whether v is a code Amadeus accepts.
func (v ContentView) IsValid() bool { return isValid(contentViewCatalog, v) }
