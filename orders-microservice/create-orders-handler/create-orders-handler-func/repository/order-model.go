package repository

type CreateOrderDB struct {
	OrderId    string `dynamodbav:"order_id"`
	UserID     string `dynamodbav:"user_id"`
	Item       string `dynamodbav:"item"`
	Quantity   int    `dynamodbav:"quantity"`
	TotalPrice int64  `dynamodbav:"total_price"`
}

type OrderResult struct {
	OrderId    string `json:"order_id"`
	TotalPrice int64  `json:"total_price"`
}
