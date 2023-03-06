package models

// Request for capturing a charge
type CreateCaptureChargeRequest struct {
    Code               string               `json:"code"`
    Amount             int                  `json:"amount,omitempty"`
    Split              []CreateSplitRequest `json:"split,omitempty"`
    OperationReference string               `json:"operation_reference"`
}
