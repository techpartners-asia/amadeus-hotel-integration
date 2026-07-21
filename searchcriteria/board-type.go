package searchcriteria

// BoardType filters hotel offers by the meals included in the rate, via the
// `boardType` query parameter on Hotel Search.
type BoardType string

const (
	BoardTypeRoomOnly  BoardType = "ROOM_ONLY"
	BoardTypeBreakfast BoardType = "BREAKFAST"
	// BoardTypeHalfBoard covers dinner and breakfast. Aggregator inventory only.
	BoardTypeHalfBoard BoardType = "HALF_BOARD"
	// BoardTypeFullBoard is aggregator inventory only.
	BoardTypeFullBoard BoardType = "FULL_BOARD"
	// BoardTypeAllInclusive is aggregator inventory only.
	BoardTypeAllInclusive BoardType = "ALL_INCLUSIVE"
)

var boardTypeCatalog = []entry[BoardType]{
	{BoardTypeRoomOnly, "Room Only"},
	{BoardTypeBreakfast, "Breakfast"},
	{BoardTypeHalfBoard, "Half Board"},
	{BoardTypeFullBoard, "Full Board"},
	{BoardTypeAllInclusive, "All Inclusive"},
}

// AllBoardTypes returns every board type Amadeus accepts.
//
// HALF_BOARD, FULL_BOARD and ALL_INCLUSIVE are only honoured against
// aggregator inventory, so a filter UI built from this list may return no
// offers on GDS/Distribution properties.
func AllBoardTypes() []BoardType { return codes(boardTypeCatalog) }

// Label returns a human-readable name for b, or "" when b is not a known code.
func (b BoardType) Label() string { return labelOf(boardTypeCatalog, b) }

// IsValid reports whether b is a code Amadeus accepts.
func (b BoardType) IsValid() bool { return isValid(boardTypeCatalog, b) }
