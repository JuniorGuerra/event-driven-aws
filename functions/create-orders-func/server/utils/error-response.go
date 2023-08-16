package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorSMS struct {
	StatusCode int
	Message    string
	Data       map[string]interface{}
}

func ErrorResponse(errorsms ErrorSMS) events.APIGatewayV2HTTPResponse {

	body, err := json.Marshal(&errorsms)

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: errorsms.StatusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
