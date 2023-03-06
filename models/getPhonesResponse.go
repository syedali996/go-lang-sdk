package models

type GetPhonesResponse struct {
    HomePhone   *GetPhoneResponse `json:"home_phone"`
    MobilePhone *GetPhoneResponse `json:"mobile_phone"`
}
