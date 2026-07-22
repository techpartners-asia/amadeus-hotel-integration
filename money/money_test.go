package money

import (
	"errors"
	"testing"
)

func TestParseFromWireFields(t *testing.T) {
	m, err := Parse("120.50", "EUR")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got := m.String(); got != "120.5 EUR" {
		t.Errorf("String() = %q, want %q", got, "120.5 EUR")
	}
	if m.Currency() != "EUR" {
		t.Errorf("Currency() = %q, want EUR", m.Currency())
	}
}

func TestParseEmptyAmountIsZeroNotError(t *testing.T) {
	// Amadeus omits optional price components rather than sending "0.00", and
	// an absent component is not a malformed one.
	m, err := Parse("", "EUR")
	if err != nil {
		t.Fatalf("Parse(\"\", \"EUR\"): %v", err)
	}
	if !m.Amount().IsZero() {
		t.Errorf("empty amount = %s, want zero", m.Amount())
	}
	if m.Currency() != "EUR" {
		t.Errorf("currency should survive an empty amount, got %q", m.Currency())
	}
}

func TestParseMalformedAmountFails(t *testing.T) {
	if _, err := Parse("not-a-price", "EUR"); err == nil {
		t.Fatal("Parse(\"not-a-price\") succeeded, want error")
	}
}

func TestIsZeroDistinguishesAbsentFromExplicitZero(t *testing.T) {
	var absent Money
	if !absent.IsZero() {
		t.Error("zero-value Money should report IsZero")
	}
	explicit := MustParse("0.00", "EUR")
	if explicit.IsZero() {
		t.Error("an explicit 0.00 EUR carries a currency and is not an absent price")
	}
}

func TestAddRejectsMixedCurrencies(t *testing.T) {
	// Amadeus can return a room rate in the hotel's currency next to a
	// converted total, so this is a real response shape, not a defensive check.
	_, err := MustParse("10", "EUR").Add(MustParse("10", "USD"))
	if !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("EUR+USD error = %v, want ErrCurrencyMismatch", err)
	}
}

func TestAddAdoptsCurrencyFromZeroValue(t *testing.T) {
	// This is what makes Sum work when it starts from a zero Money.
	var start Money
	sum, err := start.Add(MustParse("25.00", "GBP"))
	if err != nil {
		t.Fatalf("zero + GBP: %v", err)
	}
	if sum.String() != "25 GBP" {
		t.Errorf("sum = %q, want %q", sum.String(), "25 GBP")
	}
}

func TestSum(t *testing.T) {
	total, err := Sum(
		MustParse("100.00", "EUR"),
		MustParse("19.99", "EUR"),
		MustParse("0.51", "EUR"),
	)
	if err != nil {
		t.Fatalf("Sum: %v", err)
	}
	if total.String() != "120.5 EUR" {
		t.Errorf("Sum = %q, want %q", total.String(), "120.5 EUR")
	}

	empty, err := Sum()
	if err != nil || !empty.IsZero() {
		t.Errorf("Sum() = %v, %v; want zero Money and no error", empty, err)
	}
}

func TestSumPropagatesCurrencyMismatch(t *testing.T) {
	if _, err := Sum(MustParse("1", "EUR"), MustParse("1", "JPY")); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("Sum with mixed currencies: %v, want ErrCurrencyMismatch", err)
	}
}

func TestCompare(t *testing.T) {
	cheap := MustParse("99.99", "EUR")
	dear := MustParse("120.50", "EUR")

	got, err := cheap.Compare(dear)
	if err != nil || got != -1 {
		t.Errorf("Compare(99.99, 120.50) = %d, %v; want -1, nil", got, err)
	}
	if _, err := cheap.Compare(MustParse("1", "USD")); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("cross-currency Compare: %v, want ErrCurrencyMismatch", err)
	}
}

func TestSubtractingPriceComponents(t *testing.T) {
	// The shape the offers mapper needs: total minus base gives the tax portion.
	total := MustParse("130.49", "EUR")
	base := MustParse("120.50", "EUR")

	taxes, err := total.Sub(base)
	if err != nil {
		t.Fatalf("Sub: %v", err)
	}
	if taxes.String() != "9.99 EUR" {
		t.Errorf("total-base = %q, want %q", taxes.String(), "9.99 EUR")
	}
}
