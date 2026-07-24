package offers_test

import (
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/v2/money"
	"github.com/techpartners-asia/amadeus-hotel-integration/v2/offers"
)

func TestPayableUsesTotalWhenNoMarkup(t *testing.T) {
	// The ordinary case, and the whole captured sandbox: SellingTotal is unset,
	// so the price to charge is Total. Reaching for SellingTotal here would
	// show the guest zero.
	price := offers.Price{
		Total: money.MustParse("600.00", "EUR"),
		// SellingTotal deliberately left zero.
	}

	payable := price.Payable()
	if payable.String() != "600 EUR" {
		t.Errorf("Payable() = %s, want 600 EUR (Total, since no markup)", payable)
	}
	if price.HasMarkup() {
		t.Error("HasMarkup() = true with no SellingTotal")
	}
}

func TestPayableUsesSellingTotalWhenMarkupApplied(t *testing.T) {
	// A travel-agency markup: SellingTotal is what the guest pays, Total is the
	// agency's cost.
	price := offers.Price{
		Total:        money.MustParse("600.00", "EUR"),
		SellingTotal: money.MustParse("660.00", "EUR"),
	}

	payable := price.Payable()
	if payable.String() != "660 EUR" {
		t.Errorf("Payable() = %s, want 660 EUR (SellingTotal, the marked-up price)", payable)
	}
	if !price.HasMarkup() {
		t.Error("HasMarkup() = false despite a higher SellingTotal")
	}
}

func TestPayableNeverReturnsZeroForAPricedOffer(t *testing.T) {
	// The trap this method exists to prevent: SellingTotal is the field with
	// the customer-sounding name, and it is empty on real offers. Payable must
	// not hand back that zero.
	for _, offer := range allOffers(t) {
		if offer.Price.Total.Amount().IsZero() {
			continue // genuinely unpriced offers are a separate case
		}
		if offer.Price.Payable().Amount().IsZero() {
			t.Errorf("offer %s has Total %s but Payable() is zero",
				offer.ID, offer.Price.Total)
		}
	}
}

func TestPayableMatchesTotalAcrossTheCapturedFixture(t *testing.T) {
	// None of the sandbox offers carry a markup, so Payable should equal Total
	// on every one - documenting that SellingTotal is simply absent here.
	markups := 0
	for _, offer := range allOffers(t) {
		if offer.Price.HasMarkup() {
			markups++
			continue
		}
		if offer.Price.Payable().String() != offer.Price.Total.String() {
			t.Errorf("offer %s: Payable() %s != Total %s with no markup",
				offer.ID, offer.Price.Payable(), offer.Price.Total)
		}
	}
	t.Logf("%d of the captured offers carry an agency markup", markups)
}

func TestHasMarkupIgnoresAnEqualSellingTotal(t *testing.T) {
	// Some sources echo Total into SellingTotal unchanged. That is not a
	// markup, and Payable returning either gives the same figure, but HasMarkup
	// must not claim one was applied.
	price := offers.Price{
		Total:        money.MustParse("600.00", "EUR"),
		SellingTotal: money.MustParse("600.00", "EUR"),
	}
	if price.HasMarkup() {
		t.Error("HasMarkup() = true when SellingTotal equals Total")
	}
	if price.Payable().String() != "600 EUR" {
		t.Errorf("Payable() = %s, want 600 EUR", price.Payable())
	}
}
