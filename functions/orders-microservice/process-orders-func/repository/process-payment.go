package repository

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

func (orders PaymentHandler) UpdatePayment(payment ItemOrderDB) error {

	_, err := paymentT.DynamoDbClient.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(paymentT.TableName),

		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(payment.OrderId),
			},
		},

		ExpressionAttributeNames: map[string]*string{
			"#statusAttr": aws.String("status"),
		},

		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(payment.Status),
			},
		},

		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #statusAttr = :s"),
	})

	if err != nil {
		log.Printf("Couldn't add item to table. error - %v\n", err)
	}

	return err
}
