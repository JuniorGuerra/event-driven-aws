package main

import (
	snackcontrollers "cdk-snack-app/snack-controllers"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkSnackAppStackProps struct {
	awscdk.StackProps
}

func NewCdkSnackAppStack(scope constructs.Construct, id string, props *CdkSnackAppStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	dynamoDBTable, dynamoDBPutItemPolicy := snackcontrollers.SnackDynamoDBStart(stack, nil)

	createOrdersHandlerFunc := snackcontrollers.CreateOrdersHandlerFunc(stack, dynamoDBTable, dynamoDBPutItemPolicy)

	awscdk.NewCfnOutput(stack, jsii.String("CreateOrdersHandler"), &awscdk.CfnOutputProps{
		Value:       createOrdersHandlerFunc.ApiEndpoint(),
		Description: jsii.String("CreateOrdersHandler URL API Gateway"),
		ExportName:  jsii.String("CreateOrdersHandlerAPIUrl"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkSnackAppStack(app, "CdkSnackAppStack", &CdkSnackAppStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
