package models

type CreateAutomaticAnticipationSettingsRequest struct {
    Enabled          bool   `json:"enabled"`
    Type             string `json:"type"`
    VolumePercentage int    `json:"volume_percentage"`
    Delay            int    `json:"delay"`
    Days             []int  `json:"days"`
}
