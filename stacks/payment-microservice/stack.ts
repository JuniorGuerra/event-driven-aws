import { StackContext, Api, EventBus, Function, use } from "sst/constructs";
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as events from 'aws-cdk-lib/aws-events';
import * as targets from 'aws-cdk-lib/aws-events-targets'

import * as cdk from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';
import { ordersMicroservice } from "./../orders-microservice/stack";


export function paymentMicroservice({ stack }: StackContext) {

    
    const table = new dynamodb.Table(stack, "payment_table", {
        billingMode: dynamodb.BillingMode.PROVISIONED,
        partitionKey: {
            name: 'order_id',
            type: dynamodb.AttributeType.STRING,
        },
        removalPolicy: cdk.RemovalPolicy.DESTROY,

    });

    const role = new iam.Role(stack, 'Role', {
        assumedBy: new iam.ServicePrincipal('dynamodb.amazonaws.com'),
    });

    role.addToPolicy(new iam.PolicyStatement({
        effect: iam.Effect.ALLOW,
        resources: [table.tableArn],
        actions: ['dynamodb:PutItem'],
    }));


    table.metricThrottledRequestsForOperations({
        operations: [dynamodb.Operation.PUT_ITEM],
    });

    const funcion = new Function(stack, "process_payment", {
        handler: "functions/payment-microservice/create-payment-func/cmd/main.go",
        environment: {
            DYNAMODB_TABLE_NAME: table.tableName,
        },
        permissions: [table],
    })

    const { bus } = use(ordersMicroservice);

    bus.addTargets(stack, "order_created", {
        funcion: funcion,
    })

    const processPaymentApi = new Api(stack, "process-payment-api", {
        routes: {
            "POST /api/v1/process": "functions/payment-microservice/process-payment-func/cmd/main.go",
        },
        defaults: {
            function: {
                permissions: [table],
                environment: {
                    DYNAMODB_TABLE_NAME: table.tableName,
                    // EVENT_BUS_NAME: bus.eventBusName
                }
            },
        },
    });


    stack.addOutputs({
        ApiEndpoint: "POST - " + processPaymentApi.url + "/api/v1/process",
        // EventBus: bus.eventBusName,
        TableName: table.tableName
    });
}
