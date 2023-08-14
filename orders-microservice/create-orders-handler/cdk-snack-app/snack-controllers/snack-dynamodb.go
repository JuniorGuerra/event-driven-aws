package snackcontrollers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

type ApigwLambdaDynamodbCdkGolangStackProps struct {
	awscdk.StackProps
}

const (
	tableName               = "orders_table"
	partitionKeyName        = "order_id"
	dynamoDBTableNameEnvVar = "DYNAMODB_TABLE_NAME"
)

func SnackDynamoDBStart(stack awscdk.Stack, props *ApigwLambdaDynamodbCdkGolangStackProps) (awsdynamodb.Table, awsiam.Policy) {

	partitionKey := &awsdynamodb.Attribute{
		Name: jsii.String(partitionKeyName),
		Type: awsdynamodb.AttributeType_STRING,
	}

	dynamoDBOrdersTable := awsdynamodb.NewTable(stack,
		jsii.String("dynamodb-table"),
		&awsdynamodb.TableProps{TableName: jsii.String(tableName),
			BillingMode:   awsdynamodb.BillingMode_PAY_PER_REQUEST,
			PartitionKey:  partitionKey,
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
	)

	dynamoDBPutItemPolicy := awsiam.NewPolicy(
		stack,
		jsii.String("policy"),
		&awsiam.PolicyProps{
			PolicyName: jsii.String("LambdaDynamoDBOrdersPutItemPolicy"),
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Effect:    awsiam.Effect_ALLOW,
					Actions:   jsii.Strings("dynamodb:PutItem"),
					Resources: jsii.Strings(*dynamoDBOrdersTable.TableArn()),
				},
				),
			},
		},
	)

	return dynamoDBOrdersTable, dynamoDBPutItemPolicy

}
