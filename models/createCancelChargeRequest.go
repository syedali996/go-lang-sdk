package models

// Request for canceling a charge.
type CreateCancelChargeRequest struct {
    Amount             int                                   `json:"amount,omitempty"`
    SplitRules         []CreateCancelChargeSplitRulesRequest `json:"split_rules,omitempty"`
    Split              []CreateSplitRequest                  `json:"split,omitempty"`
    OperationReference string                                `json:"operation_reference"`
}
