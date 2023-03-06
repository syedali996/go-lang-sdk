package models

// Response object for listing subscriptions
type ListSubscriptionsResponse struct {
    Data   *[]GetSubscriptionResponse `json:"data"`
    Paging *PagingResponse            `json:"paging"`
}
