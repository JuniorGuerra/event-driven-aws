package repository

type ItemPaymentDB struct {
	OrderId    string `json:"order_id" dynamodbav:"order_id"`
	TotalPrice int64  `json:"total_price" dynamodbav:"total_price"`
	Status     string `json:"status" dynamodbav:"status"`
}
