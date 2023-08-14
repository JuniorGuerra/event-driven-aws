package main

import (
	"app/server/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.CreateOrder)
}
