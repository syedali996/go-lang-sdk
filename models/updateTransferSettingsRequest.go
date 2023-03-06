package models

type UpdateTransferSettingsRequest struct {
    TransferEnabled  string `json:"transfer_enabled"`
    TransferInterval string `json:"transfer_interval"`
    TransferDay      string `json:"transfer_day"`
}
