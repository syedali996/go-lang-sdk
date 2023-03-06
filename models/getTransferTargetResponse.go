package models

type GetTransferTargetResponse struct {
    TargetId *string `json:"target_id"`
    Type     *string `json:"type"`
}
