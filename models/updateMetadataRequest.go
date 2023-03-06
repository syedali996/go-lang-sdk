package models

// Request for updating an metadata
type UpdateMetadataRequest struct {
    Metadata map[string]string `json:"metadata"`
}
