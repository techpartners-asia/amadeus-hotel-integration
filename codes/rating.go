package codes

// Rating is a hotel star rating accepted by the Hotel List API in the `ratings`
// query parameter. Amadeus matches both the property's official rating and its
// own self-rating. Up to four values may be sent at once.
type Rating string

const (
	Rating1 Rating = "1"
	Rating2 Rating = "2"
	Rating3 Rating = "3"
	Rating4 Rating = "4"
	Rating5 Rating = "5"
)

var ratingCatalog = []entry[Rating]{
	{Rating1, "1 star"},
	{Rating2, "2 stars"},
	{Rating3, "3 stars"},
	{Rating4, "4 stars"},
	{Rating5, "5 stars"},
}

// MaxRatings is the number of star ratings Amadeus accepts in one request.
const MaxRatings = 4

// AllRatings returns every star rating, ascending.
func AllRatings() []Rating { return allOf(ratingCatalog) }

// Label returns a human-readable name for r, or "" when r is not 1-5.
func (r Rating) Label() string { return labelOf(ratingCatalog, r) }

// IsValid reports whether r is a rating Amadeus accepts.
func (r Rating) IsValid() bool { return isValid(ratingCatalog, r) }
