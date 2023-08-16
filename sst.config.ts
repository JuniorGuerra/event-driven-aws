import { SSTConfig } from "sst";
import { Api, EventBus } from "sst/constructs";
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as cdk from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';


export default {
  config(_input) {
    return {
      name: "orders-microservice",
      region: "us-east-1",
    };
  },



  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });



    app.stack(function Stack({ stack }) {


      const table = new dynamodb.Table(stack, "orders_table", {
        billingMode: dynamodb.BillingMode.PROVISIONED,
        partitionKey: {
          name: 'order_id',
          type: dynamodb.AttributeType.STRING,
        },
        removalPolicy: cdk.RemovalPolicy.DESTROY,
    
      });

      table.grantReadWriteData(new iam.AccountRootPrincipal());


      const role = new iam.Role(this, 'Role', {
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


      const bus = new EventBus(stack, "CreateOrdersEventBus", {
        defaults: {
          retries: 10,
        },
        rules: {
          "order_created": {
            pattern: {
              detailType: ["order.created"],
            },
            targets: {
              function: "functions/process-orders-func/cmd/main.go"
            }
          }
        }
      })

      const createOrdersApi = new Api(stack, "api", {
        routes: {
          "POST /v1/api/create": "functions/create-orders-func/cmd/main.go",
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
    });
  },
} satisfies SSTConfig;
