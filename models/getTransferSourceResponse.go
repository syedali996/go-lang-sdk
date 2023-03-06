package models

type GetTransferSourceResponse struct {
    SourceId *string `json:"source_id"`
    Type     *string `json:"type"`
}
