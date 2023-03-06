package models

type GetPhoneResponse struct {
    CountryCode *string `json:"country_code"`
    Number      *string `json:"number"`
    AreaCode    *string `json:"area_code"`
}
