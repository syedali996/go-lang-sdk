package models

type CreatePhonesRequest struct {
    HomePhone   CreatePhoneRequest `json:"home_phone,omitempty"`
    MobilePhone CreatePhoneRequest `json:"mobile_phone,omitempty"`
}
