package codes

// RateCode is a special-rate code sent in the `rateCodes` query parameter on
// Hotel Search.
//
// Unlike the other types in this package this set is open-ended: any Amadeus
// 3-character code is legal, including corporate codes negotiated per account
// (IBM, ACME...). The constants below are the widely-available public and
// qualified rates only. IsValid therefore checks the code's shape, not
// membership of a list, and AllRateCodes returns the documented subset rather
// than an exhaustive set.
//
// Sending a corporate code makes Amadeus check which chains may see that rate,
// and only authorised chains are queried. The response may still include public
// rates alongside the corporate ones.
type RateCode string

const (
	// RateCodePublic is the standard public rate.
	RateCodePublic RateCode = "PRO"
	// RateCodeGovernment is a qualified rate requiring government ID.
	RateCodeGovernment RateCode = "GOV"
	// RateCodeAAA is a qualified rate requiring AAA membership.
	RateCodeAAA RateCode = "AAA"
	// RateCodeMilitary is a qualified rate requiring military or veteran ID.
	RateCodeMilitary RateCode = "MIL"
	// RateCodeSenior is a qualified rate with a minimum-age requirement.
	RateCodeSenior RateCode = "SNR"
	// RateCodeCorporate is the generic corporate rate.
	RateCodeCorporate RateCode = "COR"
	// RateCodeRack is the undiscounted rack rate.
	RateCodeRack RateCode = "RAC"
)

var rateCodeCatalog = []entry[RateCode]{
	{RateCodePublic, "Promotional rate"},
	{RateCodeGovernment, "Government rate"},
	{RateCodeAAA, "AAA rate"},
	{RateCodeMilitary, "Military / veteran rate"},
	{RateCodeSenior, "Senior rate"},
	{RateCodeCorporate, "Corporate rate"},
	{RateCodeRack, "Rack rate"},
}

// AllRateCodes returns the documented public and qualified rate codes.
//
// This is NOT the complete set of codes Amadeus accepts: corporate codes are
// negotiated per account and cannot be enumerated. Treat it as a starting list
// for a filter UI, not as a whitelist.
func AllRateCodes() []RateCode { return allOf(rateCodeCatalog) }

// Label returns a human-readable name for c, or "" when c is not one of the
// documented codes. An empty label does not mean c is invalid: a valid
// corporate code has no label here.
func (c RateCode) Label() string { return labelOf(rateCodeCatalog, c) }

// IsValid reports whether c has the shape Amadeus requires: exactly three
// characters, uppercase letters or digits. It cannot verify that a corporate
// code exists or that the account may use it; only Amadeus can.
func (c RateCode) IsValid() bool {
	if len(c) != 3 {
		return false
	}
	for _, r := range c {
		if (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}
