package handler

import (
	"app/domain"
	"app/repository"
	"app/server/utils"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/aws-sdk-go/aws"
)

const EVENT_BUS_NAME = "EVENT_BUS_NAME"

func CreateOrder(_ context.Context, e events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var createOrder domain.CreateOrderRequest

	err := json.Unmarshal([]byte(e.Body), &createOrder)

	if err != nil {

		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}), nil
	}

	if createOrder.Item == "" {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    "Item is required",
		}), nil
	}
	if createOrder.UserID == "" {

		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    "UserId is required",
		}), nil
	}
	if createOrder.TotalPrice == 0 {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusBadRequest,
			Message:    "TotalPrice is required",
		}), nil
	}

	// Aqui habria alguna logica de negocio extra antes de la db

	handlerOrder := repository.OrderHandler{}

	item, err := handlerOrder.CreateOrder(repository.CreateOrderDB{
		UserID:     createOrder.UserID,
		Item:       createOrder.Item,
		Quantity:   createOrder.Quantity,
		TotalPrice: createOrder.TotalPrice,
		Status:     "created",
	})

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "Couldn't add item to table",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	out, err := json.Marshal(item)

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
					Source:       aws.String("create_order_fn"),
					DetailType:   aws.String("order.created"),
					Detail:       aws.String(string(out)),
				},
			},
		})

	if err != nil {
		return utils.ErrorResponse(utils.ErrorSMS{
			StatusCode: http.StatusInternalServerError,
			Message:    "can't create payment order",
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		}), nil
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       string(out),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
