package searchcriteria

// PaymentPolicy filters hotel offers by payment type, via the `paymentPolicy`
// query parameter on Hotel Search. Defaults to PaymentPolicyNone when omitted.
type PaymentPolicy string

const (
	PaymentPolicyGuarantee PaymentPolicy = "GUARANTEE"
	PaymentPolicyDeposit   PaymentPolicy = "DEPOSIT"
	// PaymentPolicyNone applies no filter and returns every payment type. It is
	// the Amadeus default; it does not mean "no payment required".
	PaymentPolicyNone PaymentPolicy = "NONE"
)

var paymentPolicyCatalog = []entry[PaymentPolicy]{
	{PaymentPolicyGuarantee, "Guarantee"},
	{PaymentPolicyDeposit, "Deposit"},
	{PaymentPolicyNone, "Any"},
}

// AllPaymentPolicies returns every payment policy Amadeus accepts.
func AllPaymentPolicies() []PaymentPolicy { return codes(paymentPolicyCatalog) }

// Label returns a human-readable name for p, or "" when p is not a known code.
func (p PaymentPolicy) Label() string { return labelOf(paymentPolicyCatalog, p) }

// IsValid reports whether p is a code Amadeus accepts.
func (p PaymentPolicy) IsValid() bool { return isValid(paymentPolicyCatalog, p) }
