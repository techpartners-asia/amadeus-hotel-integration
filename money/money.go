// Package money holds the monetary value objects shared by every bounded
// context in the SDK.
//
// Amadeus sends prices as decimal strings paired with a separate currency code
// ("total":"120.50", "currency":"EUR"). Handing that pair to callers as two
// loose strings pushes both the parsing and the currency-matching onto them,
// and the obvious parse - strconv.ParseFloat - silently loses precision on
// values a hotel bill can actually contain. Money keeps the exact decimal and
// carries the currency with it, so a total cannot drift and two currencies
// cannot be added by accident.
package money

import (
	"errors"
	"fmt"
	"strings"
)

// Currency is an ISO 4217 alphabetic currency code, e.g. "EUR".
//
// It is not validated against the ISO register: Amadeus occasionally returns
// codes outside it, and rejecting a price because its currency is unfamiliar
// would be worse than passing it through.
type Currency string

// String returns the currency code.
func (c Currency) String() string { return string(c) }

// Money is an exact monetary amount in a single currency.
//
// The zero Money is a valid zero amount with no currency, which is what an
// absent price maps to. Use IsZero to test for it.
type Money struct {
	amount   Amount
	currency Currency
}

// New returns the Money for amount in currency.
func New(amount Amount, currency Currency) Money {
	return Money{amount: amount, currency: currency}
}

// Parse builds Money from the decimal string and currency code Amadeus sends.
// An empty amount parses to a zero Money rather than an error: several Amadeus
// responses omit optional price components entirely.
func Parse(amount string, currency string) (Money, error) {
	amount = strings.TrimSpace(amount)
	if amount == "" {
		return Money{currency: Currency(currency)}, nil
	}

	parsed, err := ParseAmount(amount)
	if err != nil {
		return Money{}, fmt.Errorf("money: parsing %q %s: %w", amount, currency, err)
	}
	return Money{amount: parsed, currency: Currency(currency)}, nil
}

// MustParse is Parse for values known to be well-formed, such as test data and
// compile-time constants. It panics on a malformed amount.
func MustParse(amount string, currency string) Money {
	m, err := Parse(amount, currency)
	if err != nil {
		panic(err)
	}
	return m
}

// Amount returns the numeric part of m.
func (m Money) Amount() Amount { return m.amount }

// Currency returns the currency of m, which is "" for a zero Money.
func (m Money) Currency() Currency { return m.currency }

// IsZero reports whether m carries no amount. A zero Money is how an absent
// price is represented; it is distinct from an explicit 0.00 only in that it
// also has no currency.
func (m Money) IsZero() bool { return m.amount.IsZero() && m.currency == "" }

// String renders m as "120.50 EUR", or just the amount when no currency is set.
func (m Money) String() string {
	if m.currency == "" {
		return m.amount.String()
	}
	return m.amount.String() + " " + string(m.currency)
}

// ErrCurrencyMismatch is returned when an operation combines two different
// currencies. Amadeus can return a mix within one response - a room rate in the
// hotel's currency alongside a converted total - so this is a real condition,
// not a defensive check.
var ErrCurrencyMismatch = errors.New("money: currency mismatch")

// Add returns m+other. It fails on differing currencies; a zero Money adopts
// the other's currency, so summing a slice from a zero value works.
func (m Money) Add(other Money) (Money, error) {
	currency, err := m.combinedCurrency(other)
	if err != nil {
		return Money{}, err
	}
	return Money{amount: m.amount.Add(other.amount), currency: currency}, nil
}

// Sub returns m-other, under the same currency rules as Add.
func (m Money) Sub(other Money) (Money, error) {
	currency, err := m.combinedCurrency(other)
	if err != nil {
		return Money{}, err
	}
	return Money{amount: m.amount.Sub(other.amount), currency: currency}, nil
}

// Compare returns -1, 0 or +1 as m is less than, equal to or greater than
// other. It fails on differing currencies.
func (m Money) Compare(other Money) (int, error) {
	if _, err := m.combinedCurrency(other); err != nil {
		return 0, err
	}
	return m.amount.Compare(other.amount), nil
}

// combinedCurrency returns the currency the two operands share, treating a
// zero-value currency as "adopts the other".
func (m Money) combinedCurrency(other Money) (Currency, error) {
	switch {
	case m.currency == other.currency:
		return m.currency, nil
	case m.currency == "":
		return other.currency, nil
	case other.currency == "":
		return m.currency, nil
	default:
		return "", fmt.Errorf("%w: %s vs %s", ErrCurrencyMismatch, m.currency, other.currency)
	}
}

// Split divides m into n equal parts, returning one part and the remainder that
// would not divide evenly. part*n + remainder == m exactly.
//
// Use it to show a nightly rate from a stay total. The remainder is returned
// rather than rounded away, because deciding which night absorbs the odd cent
// is a presentation choice, not one the SDK should make silently.
func (m Money) Split(n int) (part Money, remainder Money, ok bool) {
	partAmount, remainderAmount, ok := m.amount.DivMod(n)
	if !ok {
		return Money{}, Money{}, false
	}
	return Money{amount: partAmount, currency: m.currency},
		Money{amount: remainderAmount, currency: m.currency},
		true
}

// Mul returns m multiplied by the integer n, for totalling a per-night rate
// across a stay or a per-room rate across rooms.
func (m Money) Mul(n int) Money {
	return Money{amount: m.amount.Mul(n), currency: m.currency}
}

// Sum adds every Money in values, and is the common case of totalling the price
// components of an offer. It returns a zero Money for an empty slice.
func Sum(values ...Money) (Money, error) {
	var total Money
	for _, v := range values {
		sum, err := total.Add(v)
		if err != nil {
			return Money{}, err
		}
		total = sum
	}
	return total, nil
}
