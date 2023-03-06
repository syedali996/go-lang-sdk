package models

// Split
type CreateSplitRequest struct {
    Type        string                    `json:"type"`
    Amount      int                       `json:"amount"`
    RecipientId string                    `json:"recipient_id"`
    Options     CreateSplitOptionsRequest `json:"options,omitempty"`
    SplitRuleId string                    `json:"split_rule_id,omitempty"`
}
