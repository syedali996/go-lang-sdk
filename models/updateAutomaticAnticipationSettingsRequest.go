package models

type UpdateAutomaticAnticipationSettingsRequest struct {
    Enabled          bool   `json:"enabled,omitempty"`
    Type             string `json:"type,omitempty"`
    VolumePercentage int    `json:"volume_percentage,omitempty"`
    Delay            int    `json:"delay,omitempty"`
    Days             int    `json:"days,omitempty"`
}
