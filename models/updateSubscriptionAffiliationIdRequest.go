package models

// Request for updating a Subscription Affiliation Id
type UpdateSubscriptionAffiliationIdRequest struct {
    GatewayAffiliationId string `json:"gateway_affiliation_id"`
}
