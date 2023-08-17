import { SSTConfig } from "sst";
import { microservices } from "./stacks/stack";




export default {
  config(_input) {
    return {
      name: "orders-microservice-v5",
      region: "us-east-1",
    };
  },

  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });

    
    app.stack(microservices)

  },
} satisfies SSTConfig;
