package models

type CreateAntifraudRequest struct {
    Type      string                 `json:"type"`
    Clearsale CreateClearSaleRequest `json:"clearsale"`
}
