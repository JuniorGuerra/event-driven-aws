package main

import (
	"create_payment_app/server/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.CreatePayment)
}
