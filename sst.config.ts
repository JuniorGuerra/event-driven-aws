import { SSTConfig } from "sst";
import { ordersMicroservice } from "./stacks/orders-microservice/stack";
import { paymentMicroservice } from "./stacks/payment-microservice/stack";


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

    app.stack(ordersMicroservice).stack(paymentMicroservice)
  },
} satisfies SSTConfig;
