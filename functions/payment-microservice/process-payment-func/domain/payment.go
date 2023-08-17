package domain

type PaymentRequest struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
