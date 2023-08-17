package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"process_payment_app/domain"
	"process_payment_app/repository"
	"process_payment_app/server/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/aws-sdk-go/aws"
)

const EVENT_BUS_NAME = "EVENT_BUS_NAME"

func ProcessPayment(_ context.Context, e events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	var eventItem domain.PaymentRequest
	// No es necesaria una validacion porque sabemos el valor directo a retornar.
	err := json.Unmarshal([]byte(e.Body), &eventItem)

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}), nil
	}

	if eventItem.Status != "success" {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    "Status should success",
		}), nil
	}

	handlerPayment := repository.PaymentHandler{}

	item, exist, err := handlerPayment.GetPaymentItem(eventItem.OrderId)

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	if !exist {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusNotFound,
			Message:    "order not exist",
		}), nil
	}

	err = handlerPayment.UpdatePayment(repository.ItemPaymentDB{
		OrderId: eventItem.OrderId,
		Status:  "paid",
	})

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "Couldn't update item",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	detail, err := json.Marshal(item)

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "can't create config order",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	client := cloudwatchevents.NewFromConfig(cfg)

	_, err = client.PutEvents(context.Background(),
		&cloudwatchevents.PutEventsInput{
			Entries: []types.PutEventsRequestEntry{
				{
					EventBusName: aws.String(os.Getenv(EVENT_BUS_NAME)),
					Source:       aws.String("pay_order_fn"),
					DetailType:   aws.String("order.paid"),
					Detail:       aws.String(string(detail)),
				},
			},
		})

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "can't confirm payment",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	// directamente esto no se deberia hacer, pero aguanta ser minimalista .-.
	output := domain.PaymentResponse{
		Status:  "paid",
		Message: "paid order",
	}

	out, _ := json.Marshal(output)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       string(out),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
