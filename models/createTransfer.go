package models

type CreateTransfer struct {
    Amount   int      `json:"amount"`
    SourceId string   `json:"source_id"`
    TargetId string   `json:"target_id"`
    Metadata []string `json:"metadata,omitempty"`
}
