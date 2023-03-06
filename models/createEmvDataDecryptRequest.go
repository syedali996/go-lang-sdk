package models

type CreateEmvDataDecryptRequest struct {
    Cipher string                           `json:"cipher"`
    Dukpt  CreateEmvDataDukptDecryptRequest `json:"dukpt,omitempty"`
    Tags   []CreateEmvDataTlvDecryptRequest `json:"tags"`
}
