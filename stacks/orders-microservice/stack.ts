import { StackContext, Api, EventBus } from "sst/constructs";
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';


import * as cdk from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';


export function ordersMicroservice({ stack }: StackContext) {

    const table = new dynamodb.Table(stack, "orders_table", {
        billingMode: dynamodb.BillingMode.PROVISIONED,
        partitionKey: {
            name: 'order_id',
            type: dynamodb.AttributeType.STRING,
        },
        removalPolicy: cdk.RemovalPolicy.DESTROY,

    });

    table.grantReadWriteData(new iam.AccountRootPrincipal());


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


    const bus = new EventBus(stack, "OrdersEventBus", {
        defaults: {
            retries: 10,
        },
        rules: {
            "order_created": {
                pattern: {
                    detailType: ["order.created"],
                },
                targets: {
                    function: "functions/orders-microservice/process-orders-func/cmd/main.go"
                }
            }
        }
    })

    const createOrdersApi = new Api(stack, "api", {
        routes: {
            "POST /api/v1/create": "functions/orders-microservice/create-orders-func/cmd/main.go",
        },
        defaults: {
            function: {
                permissions: [table, bus],
                environment: {
                    DYNAMODB_TABLE_NAME: table.tableName,
                    EVENT_BUS_NAME: bus.eventBusName
                }
            },
        },
    });

    stack.addOutputs({
        ApiEndpoint: createOrdersApi.url,
        EventBus: bus.eventBusName,
    });
}
