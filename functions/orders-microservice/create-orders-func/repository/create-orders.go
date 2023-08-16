package repository

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

const EnvironmentDBTableName = "DYNAMODB_TABLE_NAME"

type tableOrders struct {
	DynamoDbClient *dynamodb.DynamoDB
	TableName      string
}

type OrderHandler struct{}

var ordersT tableOrders = tableOrders{}

func init() {
	newSession, err := session.NewSession()
	if err != nil {
		log.Fatalf("failed to create DynamoBD connecion error - %s\n", err.Error())
	}

	ordersT = tableOrders{
		dynamodb.New(newSession),
		os.Getenv(EnvironmentDBTableName),
	}
}

func (orders OrderHandler) CreateOrder(order CreateOrderDB) (OrderResult, error) {

	id := uuid.New()
	order.OrderId = id.String()

	item, err := dynamodbattribute.MarshalMap(order)

	if err != nil {
		log.Printf("failed to marshal struct into dynamodb record. error - %s\n", err.Error())
		return OrderResult{}, err
	}

	_, err = ordersT.DynamoDbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(ordersT.TableName), Item: item,
	})

	if err != nil {
		log.Printf("Couldn't add item to table. error - %v\n", err)
	}

	// podriamos llamar a la db y retornalos de ahi, pero la logica es simple
	return OrderResult{
		OrderId:    order.OrderId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, err
}
