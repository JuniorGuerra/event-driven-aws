package repository

type ItemOrderDB struct {
	OrderId    string `dynamodbav:"order_id"`
	UserID     string `dynamodbav:"user_id"`
	Item       string `dynamodbav:"item"`
	Quantity   int    `dynamodbav:"quantity"`
	TotalPrice int64  `dynamodbav:"total_price"`
	Status     string `dynamodbav:"status"`
}
