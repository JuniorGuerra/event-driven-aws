package handler

import (
	"app/domain"
	"app/repository"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of repository.OrderHandler interface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateOrder(order repository.CreateOrderDB) (domain.CreateOrderEvent, error) {
	args := m.Called(order)
	return args.Get(0).(domain.CreateOrderEvent), args.Error(1)
}

func TestCreateOrder_BadRequest(t *testing.T) {
	repo := new(MockRepository)

	// Set up expectations for the repository
	order := domain.CreateOrderEvent{
		// Define fields here
		OrderID:    "1232",
		TotalPrice: 113,
	}
	repo.On("CreateOrder", mock.Anything).Return(order, nil)

	// Create your test event here, similar to events.APIGatewayV2HTTPRequest
	// For simplicity, let's assume testEvent is already defined

	resp, err := CreateOrder(context.Background(), events.APIGatewayV2HTTPRequest{})

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var responseBody domain.CreateOrderEvent
	err = json.Unmarshal([]byte(resp.Body), &responseBody)
	assert.Nil(t, err)
	// Add more assertions for the response body and headers as needed

	// repo.AssertExpectations(t)
}
func TestCreateOrder_RepositoryError(t *testing.T) {
	repo := new(MockRepository)

	order := domain.CreateOrderEvent{}
	repo.On("CreateOrder", mock.Anything).Return(order, nil)

	resp, err := CreateOrder(context.Background(), events.APIGatewayV2HTTPRequest{})

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var responseBody domain.CreateOrderEvent
	err = json.Unmarshal([]byte(resp.Body), &responseBody)
	assert.Nil(t, err)
}

func TestCreateOrder_EventError(t *testing.T) {
	repo := new(MockRepository)

	order := domain.CreateOrderEvent{}
	repo.On("CreateOrder", mock.Anything).Return(order, nil)

	resp, err := CreateOrder(context.Background(), events.APIGatewayV2HTTPRequest{
		Body: `{"user_id": "user123", "item": "item123", "quantity": 2, "total_price": 100}`,
	})

	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var responseBody domain.CreateOrderEvent
	err = json.Unmarshal([]byte(resp.Body), &responseBody)
	assert.Nil(t, err)

	repo.AssertExpectations(t)
}
