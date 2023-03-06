package models

// Response object for getting a boleto
type GetSubscriptionBoletoResponse struct {
    Interest            *GetInterestResponse `json:"interest"`
    Fine                *GetFineResponse     `json:"fine"`
    MaxDaysToPayPastDue *int                 `json:"max_days_to_pay_past_due"`
}
