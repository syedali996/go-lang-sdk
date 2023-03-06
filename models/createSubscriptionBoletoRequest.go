package models

// Information about fines and interest on the "boleto" used from payment
type CreateSubscriptionBoletoRequest struct {
    Interest            CreateInterestRequest `json:"interest,omitempty"`
    Fine                CreateFineRequest     `json:"fine,omitempty"`
    MaxDaysToPayPastDue *int                  `json:"max_days_to_pay_past_due"`
}
