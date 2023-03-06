package models

// Card token data
type CreateCardTokenRequest struct {
    Number     string `json:"number"`
    HolderName string `json:"holder_name"`
    ExpMonth   int    `json:"exp_month"`
    ExpYear    int    `json:"exp_year"`
    Cvv        string `json:"cvv"`
    Brand      string `json:"brand"`
    Label      string `json:"label"`
}
