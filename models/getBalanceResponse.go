package models

// Balance
type GetBalanceResponse struct {
    Currency           *string               `json:"currency"`
    AvailableAmount    *int64                `json:"available_amount"`
    Recipient          *GetRecipientResponse `json:"recipient"`
    TransferredAmount  *int64                `json:"transferred_amount"`
    WaitingFundsAmount *int64                `json:"waiting_funds_amount"`
}
