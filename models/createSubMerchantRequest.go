package models

// SubMerchant
type CreateSubMerchantRequest struct {
    PaymentFacilitatorCode string               `json:"payment_facilitator_code"`
    Code                   string               `json:"code"`
    Name                   string               `json:"name"`
    MerchantCategoryCode   string               `json:"merchant_category_code"`
    Document               string               `json:"document"`
    Type                   string               `json:"type"`
    Phone                  CreatePhoneRequest   `json:"phone"`
    Address                CreateAddressRequest `json:"address"`
}
