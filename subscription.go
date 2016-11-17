package braintree

import "github.com/lionelbarrow/braintree-go/nullable"

const (
	SubscriptionStatusActive       = "Active"
	SubscriptionStatusCanceled     = "Canceled"
	SubscriptionStatusExpired      = "Expired"
	SubscriptionStatusPastDue      = "Past Due"
	SubscriptionStatusPending      = "Pending"
	SubscriptionStatusUnrecognized = "Unrecognized"
)

type Subscription struct {
	XMLName                 string               `xml:"subscription" json:"-"`
	Id                      string               `xml:"id,omitempty" json:"id,omitempty"`
	Balance                 *Decimal             `xml:"balance,omitempty" json:"balance,omitempty"`
	BillingDayOfMonth       string               `xml:"billing-day-of-month,omitempty" json:"billing_day_of_month,omitempty"`
	BillingPeriodEndDate    string               `xml:"billing-period-end-date,omitempty" json:"billing_period_end_date,omitempty"`
	BillingPeriodStartDate  string               `xml:"billing-period-start-date,omitempty" json:"billing_period_start_date,omitempty"`
	CurrentBillingCycle     string               `xml:"current-billing-cycle,omitempty" json:"current_billing_cycle,omitempty"`
	DaysPastDue             string               `xml:"days-past-due,omitempty" json:"days_past_due,omitempty"`
	Discounts               []interface{}        `xml:"discounts,omitempty" json:"discounts,omitempty"`
	FailureCount            string               `xml:"failure-count,omitempty" json:"failure_count,omitempty"`
	FirstBillingDate        string               `xml:"first-billing-date,omitempty" json:"first_billing_date,omitempty"`
	MerchantAccountId       string               `xml:"merchant-account-id,omitempty" json:"merchant_account_id,omitempty"`
	NeverExpires            *nullable.NullBool   `xml:"never-expires,omitempty" json:"never_expires,omitempty"`
	NextBillAmount          *Decimal             `xml:"next-bill-amount,omitempty" json:"next_bill_amount,omitempty"`
	NextBillingPeriodAmount *Decimal             `xml:"next-billing-period-amount,omitempty" json:"next_billing_period_amount,omitempty"`
	NextBillingDate         string               `xml:"next-billing-date,omitempty" json:"next_billing_date,omitempty"`
	NumberOfBillingCycles   *nullable.NullInt64  `xml:"number-of-billing-cycles,omitempty" json:"number_of_billing_cycles,omitempty"`
	PaidThroughDate         string               `xml:"paid-through-date,omitempty" json:"paid_through_date,omitempty"`
	PaymentMethodToken      string               `xml:"payment-method-token,omitempty" json:"payment_method_token,omitempty"`
	PlanId                  string               `xml:"plan-id,omitempty" json:"plan_id,omitempty"`
	Price                   *Decimal             `xml:"price,omitempty" json:"price,omitempty"`
	Status                  string               `xml:"status,omitempty" json:"status,omitempty"`
	TrialDuration           string               `xml:"trial-duration,omitempty" json:"trial_duration,omitempty"`
	TrialDurationUnit       string               `xml:"trial-duration-unit,omitempty" json:"trial_duration_unit,omitempty"`
	TrialPeriod             *nullable.NullBool   `xml:"trial-period,omitempty" json:"trial_period,omitempty"`
	Transactions            *Transactions        `xml:"transactions,omitempty" json:"transactions,omitempty"`
	Options                 *SubscriptionOptions `xml:"options,omitempty" json:"options,omitempty"`
	AddOns                  *AddOnList           `xml:"add-ons,omitempty" json:"add_ons,omitempty"`
	// Descriptor              interface{}   `xml:"descriptor,omitempty"`   // struct with name, phone
}

type Subscriptions struct {
	Subscription []*Subscription `xml:"subscription" json:"subscription,omitempty"`
}

type SubscriptionOptions struct {
	DoNotInheritAddOnsOrDiscounts        bool `xml:"do-not-inherit-add-ons-or-discounts,omitempty" json:"do_not_inherit_add_ons_or_discounts,omitempty"`
	ProrateCharges                       bool `xml:"prorate-charges,omitempty" json:"prorate_charges,omitempty"`
	ReplaceAllAddOnsAndDiscounts         bool `xml:"replace-all-add-ons-and-discounts,omitempty" json:"replace_all_add_ons_and_discounts,omitempty"`
	RevertSubscriptionOnProrationFailure bool `xml:"revert-subscription-on-proration-failure,omitempty" json:"revert_subscription_on_proration_failure,omitempty"`
	StartImmediately                     bool `xml:"start-immediately,omitempty" json:"start_immediately,omitempty"`
}
