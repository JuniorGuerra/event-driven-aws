import { StackContext, Api, Function, EventBus } from "sst/constructs";
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';


import * as cdk from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';


export function microservices({ stack }: StackContext) {

    const ordersTable = new dynamodb.Table(stack, "orders_table", {
        billingMode: dynamodb.BillingMode.PROVISIONED,
        partitionKey: {
            name: 'order_id',
            type: dynamodb.AttributeType.STRING,
        },
        removalPolicy: cdk.RemovalPolicy.DESTROY,

    });
    const paymentTable = new dynamodb.Table(stack, "payment_table", {
        billingMode: dynamodb.BillingMode.PROVISIONED,
        partitionKey: {
            name: 'order_id',
            type: dynamodb.AttributeType.STRING,
        },
        removalPolicy: cdk.RemovalPolicy.DESTROY,

    });

    ordersTable.grantReadWriteData(new iam.AccountRootPrincipal());


    const role = new iam.Role(stack, 'Role', {
        assumedBy: new iam.ServicePrincipal('dynamodb.amazonaws.com'),
    });
    

    role.addToPolicy(new iam.PolicyStatement({
        effect: iam.Effect.ALLOW,
        resources: [ordersTable.tableArn, paymentTable.tableArn],
        actions: ['dynamodb:PutItem'],
    }));


    ordersTable.metricThrottledRequestsForOperations({
        operations: [dynamodb.Operation.PUT_ITEM],
    });
    paymentTable.metricThrottledRequestsForOperations({
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
            },
            "order_paid": {
                pattern: {
                    detailType: ["order.paid"],
                },
            }
        }
    })

    const funcionProcess = new Function(stack, "process_order", {
        handler: "functions/orders-microservice/process-orders-func/cmd/main.go",
        environment: {
            DYNAMODB_TABLE_NAME: ordersTable.tableName,
        },
        permissions: [ordersTable],
    })

    const funcionCreate = new Function(stack, "process_payment", {
        handler: "functions/payment-microservice/create-payment-func/cmd/main.go",
        environment: {
            DYNAMODB_TABLE_NAME: paymentTable.tableName,
        },
        permissions: [paymentTable],
    })

    bus.addTargets(stack, "order_paid", {
        funcion: funcionProcess,
    })
    bus.addTargets(stack, "order_created", {
        funcion: funcionCreate,
    })


    const servicesEnpoints = new Api(stack, "api", {
        routes: {
            "POST /api/v1/create": "functions/orders-microservice/create-orders-func/cmd/main.go",
            "POST /api/v1/process": "functions/payment-microservice/process-payment-func/cmd/main.go",
        },
        defaults: {
            function: {
                permissions: [ordersTable, bus, paymentTable],
                environment: {
                    DYNAMODB_TABLE_NAME: ordersTable.tableName,
                    DYNAMODB_TABLE_NAME_PAYMENT: paymentTable.tableName,
                    EVENT_BUS_NAME: bus.eventBusName
                }
            },
        },
    });

    

    stack.addOutputs({
        ApiEndpoint: servicesEnpoints.url,
        OrdersTableName: ordersTable.tableName,
        PaymentTableName: paymentTable.tableName,
        EventBus: bus.eventBusName,
    });

    return {
        bus
    }
}
