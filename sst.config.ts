import { SSTConfig } from "sst";
import { microservices } from "./stacks/stack";




export default {
  config(_input) {
    return {
      name: "event-driven-aws-microservices",
      region: "us-east-2",
    };
  },

  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go1.x",
    });

    
    app.stack(microservices)

  },
} satisfies SSTConfig;
