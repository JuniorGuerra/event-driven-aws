package repository

type ItemPaymentDB struct {
	OrderId    string `dynamodbav:"order_id"`
	TotalPrice int64  `dynamodbav:"total_price"`
	Status     string `dynamodbav:"status"`
}
