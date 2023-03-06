package models

// The Split Options Request
type CreateSplitOptionsRequest struct {
    Liable              bool `json:"liable,omitempty"`
    ChargeProcessingFee bool `json:"charge_processing_fee,omitempty"`
    ChargeRemainderFee  bool `json:"charge_remainder_fee,omitempty"`
}
