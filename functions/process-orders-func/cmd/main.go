package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func readOrder(_ context.Context, event events.CloudWatchEvent) error {
	fmt.Println(event.Detail)
	fmt.Printf("received event of type %q\n", event.DetailType)
	spew.Dump(event)
	return nil
}

func main() {
	lambda.Start(readOrder)
}

//  proximos pasos:
// crear el handler que lo cree en actualice en la base de datos
// crear otro microsevicio con dos endpoint que hagan lo mismo que este...
