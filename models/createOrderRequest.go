package models

// Request for creating an order
type CreateOrderRequest struct {
    Items            []CreateOrderItemRequest `json:"items"`
    Customer         CreateCustomerRequest    `json:"customer"`
    Payments         []CreatePaymentRequest   `json:"payments"`
    Code             string                   `json:"code"`
    CustomerId       string                   `json:"customer_id"`
    Shipping         CreateShippingRequest    `json:"shipping,omitempty"`
    Metadata         map[string]string        `json:"metadata"`
    AntifraudEnabled bool                     `json:"antifraud_enabled,omitempty"`
    Ip               string                   `json:"ip,omitempty"`
    SessionId        string                   `json:"session_id,omitempty"`
    Location         CreateLocationRequest    `json:"location,omitempty"`
    Device           CreateDeviceRequest      `json:"device,omitempty"`
    Closed           bool                     `json:"closed"`
    Currency         string                   `json:"currency,omitempty"`
    Antifraud        CreateAntifraudRequest   `json:"antifraud,omitempty"`
    Submerchant      CreateSubMerchantRequest `json:"submerchant,omitempty"`
}
