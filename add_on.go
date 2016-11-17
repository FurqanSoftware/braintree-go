package braintree

type AddOnList struct {
	XMLName string        `xml:"add-ons" json:"-"`
	AddOns  []AddOn       `xml:"add-on" json:"-"`
	Add     []AddonAdd    `xml:"add,omitempty" json:"add,omitempty"`
	Remove  []string      `xml:"remove,omitempty" json:"remove,omitempty"`
	Update  []AddonUpdate `xml:"update,omitempty" json:"update,omitempty"`
}

type AddOn struct {
	XMLName string `xml:"add-on"`
	Modification
}

type AddonAdd struct {
	Amount                *Decimal `xml:"amount,omitempty" json:"amount,omitempty"`
	InheritedFromId       string   `xml:"inherited-from-id,omitempty" json:"inherited_from_id,omitempty"`
	NeverExpires          bool     `xml:"never-expires,omitempty" json:"never_expires,omitempty"`
	NumberOfBillingCycles int      `xml:"number-of-billing-cycles,omitempty" json:"number_of_billing_cycles,omitempty"`
	Quantity              int      `xml:"quantity,omitempty" json:"quantity,omitempty"`
}

type AddonUpdate struct {
	Amount                *Decimal `xml:"amount,omitempty" json:"amount,omitempty"`
	ExistingId            string   `xml:"existing-id,omitempty" json:"existing_id,omitempty"`
	NeverExpires          bool     `xml:"never-expires,omitempty" json:"never_expires,omitempty"`
	NumberOfBillingCycles int      `xml:"number-of-billing-cycles,omitempty" json:"number_of_billing_cycles,omitempty"`
	Quantity              int      `xml:"quantity,omitempty" json:"quantity,omitempty"`
}
