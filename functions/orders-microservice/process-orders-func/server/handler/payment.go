package handler

import (
	"context"
	"encoding/json"
	"process_orders/domain"
	"process_orders/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/davecgh/go-spew/spew"
)

func CreatePayment(_ context.Context, event events.CloudWatchEvent) error {

	var eventItem domain.CreateOrderEvent
	// No es necesaria una validacion porque sabemos el valor directo a retornar.
	err := json.Unmarshal(event.Detail, &eventItem)

	if err != nil {
		// loguear error...
		spew.Dump(err.Error())
	}

	handlerPayment := repository.PaymentHandler{}

	// actualizamos directo, no es tan necesario validar existencia.

	err = handlerPayment.UpdatePayment(repository.ItemOrderDB{
		OrderId: eventItem.OrderID,
		Status:  "ready for shipping",
	})

	if err != nil {
		spew.Dump(err.Error())
	}

	return nil
}
