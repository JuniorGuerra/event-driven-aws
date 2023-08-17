package repository

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const EnvironmentDBTableName = "DYNAMODB_TABLE_NAME"

type tableOrders struct {
	DynamoDbClient *dynamodb.DynamoDB
	TableName      string
}

type PaymentHandler struct{}

var paymentT tableOrders = tableOrders{}

func init() {
	newSession, err := session.NewSession()
	if err != nil {
		log.Fatalf("failed to create DynamoBD connecion error - %s\n", err.Error())
	}

	paymentT = tableOrders{
		dynamodb.New(newSession),
		os.Getenv(EnvironmentDBTableName),
	}
}

func (orders PaymentHandler) CreatePayment(payment CreatePaymentDB) error {
	item, err := dynamodbattribute.MarshalMap(payment)

	if err != nil {
		log.Printf("failed to marshal struct into dynamodb record. error - %s\n", err.Error())
		// se podria crear un evento que escuche los errores y los notifique mas completo en este lugar
		return err
	}

	_, err = paymentT.DynamoDbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(paymentT.TableName),
		Item:      item,
	})

	if err != nil {
		log.Printf("Couldn't add item to table. error - %v\n", err)
		// se podria crear un evento que escuche los errores y los notifique mas completo en este lugar
	}

	return err
}
