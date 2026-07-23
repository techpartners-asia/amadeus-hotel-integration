package money

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// Amount is an exact decimal number, stored as an integer of units scaled by a
// power of ten: the value is units / 10^scale. "120.50" is units=12050,
// scale=2.
//
// It exists so the SDK can carry Amadeus prices without a third-party decimal
// dependency and without float64. Hotel prices are decimal strings with at most
// a handful of fraction digits, and the only arithmetic the SDK performs on
// them is addition, subtraction and comparison, all of which are exact on this
// representation. It is deliberately not a general-purpose decimal type: there
// is no division, because nothing here needs to define a rounding policy.
//
// The zero Amount is exactly zero.
type Amount struct {
	units int64
	scale uint8
}

// maxScale caps the fraction digits Amount will accept. Currencies use at most
// four; the extra room absorbs the conversion rates Amadeus quotes at higher
// precision, while keeping 10^scale far inside int64.
const maxScale = 9

// ErrMalformedAmount is returned by ParseAmount for input that is not a plain
// decimal number.
var ErrMalformedAmount = errors.New("malformed decimal amount")

// ErrAmountOverflow is returned when a value or an operation exceeds the range
// Amount can represent exactly.
var ErrAmountOverflow = errors.New("decimal amount out of range")

// ParseAmount parses a plain decimal string such as "120.50", "-3", or ".75".
// It rejects exponent notation, thousands separators and embedded spaces, none
// of which appear in Amadeus price fields.
func ParseAmount(s string) (Amount, error) {
	if s == "" {
		return Amount{}, fmt.Errorf("%w: empty string", ErrMalformedAmount)
	}

	negative := false
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		negative = true
		s = s[1:]
	}

	whole, fraction, hasFraction := strings.Cut(s, ".")
	if whole == "" && fraction == "" {
		return Amount{}, fmt.Errorf("%w: no digits", ErrMalformedAmount)
	}
	if hasFraction && strings.Contains(fraction, ".") {
		return Amount{}, fmt.Errorf("%w: more than one decimal point", ErrMalformedAmount)
	}
	if !isDigits(whole) || !isDigits(fraction) {
		return Amount{}, fmt.Errorf("%w: %q contains non-digits", ErrMalformedAmount, s)
	}

	// Trailing zeros carry no value and only push the scale toward its limit.
	fraction = strings.TrimRight(fraction, "0")
	if len(fraction) > maxScale {
		return Amount{}, fmt.Errorf("%w: more than %d fraction digits", ErrAmountOverflow, maxScale)
	}

	digits := whole + fraction
	digits = strings.TrimLeft(digits, "0")
	if digits == "" {
		return Amount{}, nil // the value is zero, however it was spelled
	}

	units, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return Amount{}, fmt.Errorf("%w: %q", ErrAmountOverflow, s)
	}
	if negative {
		units = -units
	}

	return Amount{units: units, scale: uint8(len(fraction))}, nil
}

// isDigits reports whether s consists solely of ASCII digits. The empty string
// qualifies, so "5." and ".5" are both accepted.
func isDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// ParseRate parses a decimal that may carry more precision than an Amount can
// hold, rounding the excess away instead of refusing it.
//
// It exists for exchange rates. Amadeus quotes them to sixteen decimal places
// - "4099.1909999999998035" for EUR to MNT - which is float noise dressed up
// as precision: the true rate is 4099.191. ParseAmount rightly rejects that as
// beyond what it can represent exactly, but rejecting a currency conversion
// over digits nobody means is not useful, so this rounds at the ninth place.
//
// Use ParseAmount for prices, where excess precision means the input is wrong
// and should be reported rather than quietly adjusted.
func ParseRate(s string) (Amount, error) {
	amount, err := ParseAmount(s)
	if err == nil {
		return amount, nil
	}
	if !errors.Is(err, ErrAmountOverflow) {
		return Amount{}, err
	}

	text := strings.TrimSpace(s)
	negative := strings.HasPrefix(text, "-")
	text = strings.TrimLeft(text, "+-")

	whole, fraction, _ := strings.Cut(text, ".")
	if len(fraction) <= maxScale {
		return Amount{}, err // overflowed on magnitude, not precision
	}
	if !isDigits(whole) || !isDigits(fraction) {
		return Amount{}, err
	}

	// Build the units from the digits that fit, then round using the first
	// digit that does not. Rounding by integer addition carries correctly into
	// the whole part, so 0.9999999999 becomes 1 rather than 0.999999999.
	units, parseErr := strconv.ParseInt(whole+fraction[:maxScale], 10, 64)
	if parseErr != nil {
		return Amount{}, err
	}
	if fraction[maxScale] >= '5' {
		units++
	}
	if negative {
		units = -units
	}

	return Amount{units: units, scale: maxScale}, nil
}

// MustParseAmount is ParseAmount for values known to be well-formed. It panics
// on malformed input.
func MustParseAmount(s string) Amount {
	a, err := ParseAmount(s)
	if err != nil {
		panic(err)
	}
	return a
}

// IsZero reports whether a is exactly zero.
func (a Amount) IsZero() bool { return a.units == 0 }

// IsNegative reports whether a is less than zero.
func (a Amount) IsNegative() bool { return a.units < 0 }

// String renders a in canonical plain decimal notation, round-tripping through
// ParseAmount exactly.
//
// Canonical means no insignificant trailing zeros, so the rendering depends
// only on the value and not on how it was arrived at: "120.50" parses and
// renders as "120.5", and 1.999+0.001 renders as "2" rather than "2.000".
// Equal values therefore always produce equal strings.
func (a Amount) String() string {
	if a.scale == 0 || a.units == 0 {
		return strconv.FormatInt(a.units, 10)
	}

	units := a.units
	sign := ""
	if units < 0 {
		sign = "-"
		units = -units
	}

	digits := strconv.FormatInt(units, 10)
	if len(digits) <= int(a.scale) {
		digits = strings.Repeat("0", int(a.scale)-len(digits)+1) + digits
	}
	split := len(digits) - int(a.scale)

	fraction := strings.TrimRight(digits[split:], "0")
	if fraction == "" {
		return sign + digits[:split]
	}
	return sign + digits[:split] + "." + fraction
}

// Float64 returns a as a float64, for callers feeding it to code that demands
// one - a chart, a sort key, a JSON number. It is lossy by nature and must not
// be used to compute a total that anyone will be charged.
func (a Amount) Float64() float64 {
	return float64(a.units) / math.Pow10(int(a.scale))
}

// Add returns a+other. Overflow saturates rather than wrapping, on the
// principle that a visibly absurd total is safer than a silently negative one;
// values anywhere near the limit do not occur in hotel pricing.
func (a Amount) Add(other Amount) Amount {
	x, y, scale := align(a, other)
	sum := x + y
	if (x > 0 && y > 0 && sum < 0) || (x < 0 && y < 0 && sum > 0) {
		return saturated(x > 0)
	}
	return Amount{units: sum, scale: scale}
}

// Sub returns a-other.
func (a Amount) Sub(other Amount) Amount {
	return a.Add(Amount{units: -other.units, scale: other.scale})
}

// DivMod splits a into n equal parts, returning the part and whatever could not
// be divided evenly. part*n + remainder == a exactly.
//
// There is no plain Div, because dividing money requires deciding who gets the
// leftover cent and the domain has no basis to decide that silently. Returning
// the remainder makes the choice the caller's, and visible.
//
// It reports ok=false for n <= 0.
func (a Amount) DivMod(n int) (part Amount, remainder Amount, ok bool) {
	if n <= 0 {
		return Amount{}, Amount{}, false
	}

	units := int64(n)
	return Amount{units: a.units / units, scale: a.scale},
		Amount{units: a.units % units, scale: a.scale},
		true
}

// Scale returns the number of decimal places the amount is held at.
func (a Amount) Scale() int { return int(a.scale) }

// MulAmount returns a*other, exactly.
//
// It reports ok=false when the product cannot be represented - which needs
// operands far larger than any price or exchange rate - rather than silently
// saturating, because a wrong conversion is worse than a refused one.
//
// The multiplication is performed in arbitrary precision and only then reduced,
// so intermediate overflow cannot corrupt the result. Amadeus quotes exchange
// rates to sixteen decimal places, and multiplying one by a price at the
// int64 level would overflow long before the answer did.
func (a Amount) MulAmount(other Amount) (Amount, bool) {
	product := new(big.Int).Mul(big.NewInt(a.units), big.NewInt(other.units))
	scale := int(a.scale) + int(other.scale)

	// Shed excess precision from the tail, which for a rate is noise anyway.
	for scale > maxScale {
		product = roundedDiv(product, big.NewInt(10))
		scale--
	}

	if !product.IsInt64() {
		return Amount{}, false
	}
	return Amount{units: product.Int64(), scale: uint8(scale)}, true
}

// Round returns a rounded to the given number of decimal places, half away
// from zero - the convention people expect of money, and the one that does not
// systematically favour either party.
//
// Rounding to fewer places than the amount holds loses precision by design:
// currencies such as the Mongolian tögrög have no minor unit, so a converted
// price must be reduced to whole units before it is shown or charged.
func (a Amount) Round(places int) Amount {
	if places < 0 || places >= int(a.scale) {
		return a
	}

	units := big.NewInt(a.units)
	for i := int(a.scale); i > places; i-- {
		units = roundedDiv(units, big.NewInt(10))
	}
	if !units.IsInt64() {
		return a
	}
	return Amount{units: units.Int64(), scale: uint8(places)}
}

// roundedDiv divides n by d, rounding half away from zero.
func roundedDiv(n, d *big.Int) *big.Int {
	quotient, remainder := new(big.Int).QuoRem(n, d, new(big.Int))

	twice := new(big.Int).Abs(remainder)
	twice.Mul(twice, big.NewInt(2))
	if twice.CmpAbs(d) >= 0 {
		if n.Sign() < 0 {
			quotient.Sub(quotient, big.NewInt(1))
		} else {
			quotient.Add(quotient, big.NewInt(1))
		}
	}
	return quotient
}

// Mul returns a multiplied by the integer n. Saturates rather than wrapping,
// for the reason given on Add.
func (a Amount) Mul(n int) Amount {
	if n == 0 || a.units == 0 {
		return Amount{scale: a.scale}
	}

	product := a.units * int64(n)
	if product/int64(n) != a.units {
		return saturated((a.units > 0) == (n > 0))
	}
	return Amount{units: product, scale: a.scale}
}

// Compare returns -1, 0 or +1 as a is less than, equal to or greater than
// other.
func (a Amount) Compare(other Amount) int {
	x, y, _ := align(a, other)
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	default:
		return 0
	}
}

// Equal reports whether a and other represent the same value, regardless of the
// scale each is stored at.
func (a Amount) Equal(other Amount) bool { return a.Compare(other) == 0 }

// align restates both amounts at their common (larger) scale and returns the
// two unit counts. If rescaling would overflow, both are collapsed to saturated
// values at the original scale, which keeps ordering intact.
func align(a, b Amount) (x, y int64, scale uint8) {
	switch {
	case a.scale == b.scale:
		return a.units, b.units, a.scale
	case a.scale < b.scale:
		x, ok := rescale(a.units, b.scale-a.scale)
		if !ok {
			return saturatedUnits(a.units > 0), b.units, b.scale
		}
		return x, b.units, b.scale
	default:
		y, ok := rescale(b.units, a.scale-b.scale)
		if !ok {
			return a.units, saturatedUnits(b.units > 0), a.scale
		}
		return a.units, y, a.scale
	}
}

// rescale multiplies units by 10^by, reporting whether it fit in an int64.
func rescale(units int64, by uint8) (int64, bool) {
	for range by {
		if units > math.MaxInt64/10 || units < math.MinInt64/10 {
			return 0, false
		}
		units *= 10
	}
	return units, true
}

// saturated returns the largest representable amount with the given sign.
func saturated(positive bool) Amount {
	return Amount{units: saturatedUnits(positive)}
}

func saturatedUnits(positive bool) int64 {
	if positive {
		return math.MaxInt64
	}
	return math.MinInt64
}
