package models

type GetSplitOptionsResponse struct {
    Liable              *bool   `json:"liable"`
    ChargeProcessingFee *bool   `json:"charge_processing_fee"`
    ChargeRemainderFee  *string `json:"charge_remainder_fee"`
}
