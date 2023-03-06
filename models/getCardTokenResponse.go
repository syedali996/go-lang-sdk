package models

// Card token data
type GetCardTokenResponse struct {
    LastFourDigits *string `json:"last_four_digits"`
    HolderName     *string `json:"holder_name"`
    HolderDocument *string `json:"holder_document"`
    ExpMonth       *string `json:"exp_month"`
    ExpYear        *string `json:"exp_year"`
    Brand          *string `json:"brand"`
    Type           *string `json:"type"`
    Label          *string `json:"label"`
}
