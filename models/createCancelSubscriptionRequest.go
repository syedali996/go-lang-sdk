package models

// Request for canceling a subscription
type CreateCancelSubscriptionRequest struct {
    CancelPendingInvoices bool `json:"cancel_pending_invoices"`
}
