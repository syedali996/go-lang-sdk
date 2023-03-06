package models

type GetAntifraudResponse struct {
    Status        *string `json:"status"`
    ReturnCode    *string `json:"return_code"`
    ReturnMessage *string `json:"return_message"`
    ProviderName  *string `json:"provider_name"`
    Score         *string `json:"score"`
}
