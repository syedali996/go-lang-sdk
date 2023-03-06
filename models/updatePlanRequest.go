package models

// Request for updating a plan
type UpdatePlanRequest struct {
    Name                string            `json:"name"`
    Description         string            `json:"description"`
    Installments        []int             `json:"installments"`
    StatementDescriptor string            `json:"statement_descriptor"`
    Currency            string            `json:"currency"`
    Interval            string            `json:"interval"`
    IntervalCount       int               `json:"interval_count"`
    PaymentMethods      []string          `json:"payment_methods"`
    BillingType         string            `json:"billing_type"`
    Status              string            `json:"status"`
    Shippable           bool              `json:"shippable"`
    BillingDays         []int             `json:"billing_days"`
    Metadata            map[string]string `json:"metadata"`
    MinimumPrice        int               `json:"minimum_price,omitempty"`
    TrialPeriodDays     int               `json:"trial_period_days,omitempty"`
}
