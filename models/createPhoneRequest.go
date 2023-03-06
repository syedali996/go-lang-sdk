package models

type CreatePhoneRequest struct {
    CountryCode string `json:"country_code,omitempty"`
    Number      string `json:"number,omitempty"`
    AreaCode    string `json:"area_code,omitempty"`
}
