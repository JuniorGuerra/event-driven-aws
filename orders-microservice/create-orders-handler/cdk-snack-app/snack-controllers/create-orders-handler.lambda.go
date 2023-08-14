package snackcontrollers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	awscdkapigw "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	awsapigwintegrations "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/jsii-runtime-go"
)

func CreateOrdersHandlerFunc(stack awscdk.Stack, dynamoDBTable awsdynamodb.Table, policeDb awsiam.Policy) awscdkapigw.HttpApi {

	createOrdersHandlerFunc := awscdklambdago.NewGoFunction(stack, jsii.String("CreateOrdersHandlerFunc"), &awscdklambdago.GoFunctionProps{
		FunctionName: jsii.String("CreateOrdersHandlerFunc"),
		Description:  jsii.String("A function that returns the IP address and user agent of the caller."),
		Entry:        jsii.String("../create-orders-handler-func/cmd"),
		Environment:  &map[string]*string{dynamoDBTableNameEnvVar: dynamoDBTable.TableName()},
	})

	createOrdersHandlerFunc.Role().AttachInlinePolicy(policeDb)

	createOrdersHandlerApi := awscdkapigw.NewHttpApi(stack, jsii.String("CreateOrdersHandlerApi"), nil)

	createOrdersHandlerApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path: jsii.String("/"),
		Methods: &[]awscdkapigw.HttpMethod{
			awscdkapigw.HttpMethod_POST,
		},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(
			jsii.String("CreateOrdersHandlerApiIntegration"),
			createOrdersHandlerFunc, nil,
		),
	})

	return createOrdersHandlerApi

}
