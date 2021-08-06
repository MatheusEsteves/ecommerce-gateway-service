package models

type CheckoutResponseData struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	PurchaseCode string `json:"purchaseCode"`
}
