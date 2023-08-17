package domain

type OrderResult struct {
	OrderId    string `json:"order_id"`
	TotalPrice int64  `json:"total_price"`
	Status     string `json:"status"`
}
