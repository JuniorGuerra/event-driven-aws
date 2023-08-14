package handler

import (
	"app/domain"
	"app/repository"
	"app/server/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func CreateOrder(_ context.Context, e events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Println(e)
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

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusCreated,
		Body:       string(out),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
