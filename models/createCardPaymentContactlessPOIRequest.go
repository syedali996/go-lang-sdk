package models

type CreateCardPaymentContactlessPOIRequest struct {
    SystemName    string `json:"system_name"`
    Model         string `json:"model"`
    Provider      string `json:"provider"`
    SerialNumber  string `json:"serial_number"`
    VersionNumber string `json:"version_number"`
}
