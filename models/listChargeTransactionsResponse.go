package models

// Response object for listing charge transactions
type ListChargeTransactionsResponse struct {
    Data   *[]GetTransactionResponseInterface `json:"data"`
    Paging *PagingResponse                    `json:"paging"`
}
