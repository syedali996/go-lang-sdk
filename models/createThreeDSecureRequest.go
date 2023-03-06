package models

// Creates a 3D-S authentication payment
type CreateThreeDSecureRequest struct {
    Mpi             string `json:"mpi"`
    Cavv            string `json:"cavv,omitempty"`
    Eci             string `json:"eci,omitempty"`
    TransactionId   string `json:"transaction_id,omitempty"`
    SuccessUrl      string `json:"success_url,omitempty"`
    DsTransactionId string `json:"ds_transaction_id,omitempty"`
    Version         string `json:"version,omitempty"`
}
