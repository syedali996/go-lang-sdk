package models

// The Transaction Gateway Response
type GetGatewayResponseResponse struct {
    Code   *string                    `json:"code"`
    Errors *[]GetGatewayErrorResponse `json:"errors"`
}
