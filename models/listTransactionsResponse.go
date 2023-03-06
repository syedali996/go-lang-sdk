package models

// Response object for listing transactions
type ListTransactionsResponse struct {
    Data   *[]GetTransactionResponseInterface `json:"data"`
    Paging *PagingResponse                    `json:"paging"`
}
