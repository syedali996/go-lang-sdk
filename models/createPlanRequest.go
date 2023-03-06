package models

// Request for creating a plan
type CreatePlanRequest struct {
    Name                string                     `json:"name"`
    Description         string                     `json:"description"`
    StatementDescriptor string                     `json:"statement_descriptor"`
    Items               []CreatePlanItemRequest    `json:"items"`
    Shippable           bool                       `json:"shippable"`
    PaymentMethods      []string                   `json:"payment_methods"`
    Installments        []int                      `json:"installments"`
    Currency            string                     `json:"currency"`
    Interval            string                     `json:"interval"`
    IntervalCount       int                        `json:"interval_count"`
    BillingDays         []int                      `json:"billing_days"`
    BillingType         string                     `json:"billing_type"`
    PricingScheme       CreatePricingSchemeRequest `json:"pricing_scheme"`
    Metadata            map[string]string          `json:"metadata"`
    MinimumPrice        int                        `json:"minimum_price,omitempty"`
    Cycles              int                        `json:"cycles,omitempty"`
    Quantity            int                        `json:"quantity,omitempty"`
    TrialPeriodDays     int                        `json:"trial_period_days,omitempty"`
}
