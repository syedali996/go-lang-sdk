package models

// 3D-S payment authentication response
type GetThreeDSecureResponse struct {
    Mpi           *string `json:"mpi"`
    Eci           *string `json:"eci"`
    Cavv          *string `json:"cavv"`
    TransactionId *string `json:"transaction_Id"`
    SuccessUrl    *string `json:"success_url"`
}
