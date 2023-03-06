package models

type ListWithdrawals struct {
    Data   []GetWithdrawResponse `json:"data"`
    Paging PagingResponse        `json:"paging"`
}
