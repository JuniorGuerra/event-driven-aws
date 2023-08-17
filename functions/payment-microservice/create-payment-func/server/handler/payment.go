package handler

import (
	"context"
	"create_payment_app/domain"
	"create_payment_app/repository"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davecgh/go-spew/spew"
)

const EVENT_BUS_NAME = "EVENT_BUS_NAME"

func CreatePayment(_ context.Context, event events.CloudWatchEvent) error {

	var eventItem domain.OrderResult
	// No es necesaria una validacion porque sabemos el valor directo a retornar.
	err := json.Unmarshal(event.Detail, &eventItem)

	if err != nil {
		// loguear error...
		spew.Dump(err.Error())
	}

	handlerPayment := repository.PaymentHandler{}

	err = handlerPayment.CreatePayment(repository.CreatePaymentDB{
		OrderId:    eventItem.OrderId,
		TotalPrice: eventItem.TotalPrice,
		Status:     "in order",
	})

	if err != nil {
		// loguear error...
		spew.Dump(err.Error())
	}

	return nil
}
