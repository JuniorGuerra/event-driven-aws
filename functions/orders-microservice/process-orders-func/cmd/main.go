package main

import (
	"process_orders/server/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.CreatePayment)
}

//  proximos pasos:
// crear el handler que lo cree en actualice en la base de datos
// crear otro microsevicio con dos endpoint que hagan lo mismo que este...
