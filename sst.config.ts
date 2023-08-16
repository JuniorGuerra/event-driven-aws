import { SSTConfig } from "sst";
import { Api, EventBus } from "sst/constructs";
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as cdk from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';
import { ordersMicroservice } from "./stacks/orders-microservice/stack";


export default {
  config(_input) {
    return {
      name: "orders-microservice-v2",
      region: "us-east-1",
    };
  },

  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });

    app.stack(ordersMicroservice)
  },
} satisfies SSTConfig;
