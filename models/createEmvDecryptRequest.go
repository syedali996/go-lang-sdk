package models

type CreateEmvDecryptRequest struct {
    IccData            string                                 `json:"icc_data"`
    CardSequenceNumber string                                 `json:"card_sequence_number"`
    Data               CreateEmvDataDecryptRequest            `json:"data"`
    Poi                CreateCardPaymentContactlessPOIRequest `json:"poi,omitempty"`
}
