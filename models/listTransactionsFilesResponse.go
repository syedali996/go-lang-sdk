package models

// Response object for listing of transactions files
type ListTransactionsFilesResponse struct {
    Data   *[]GetTransactionReportFileResponse `json:"data"`
    Paging *PagingResponse                     `json:"paging"`
}
